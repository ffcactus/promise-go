package db

import (
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	commonDB "promise/common/db"
	commonConstError "promise/common/object/consterror"
	"promise/server/object/consterror"
	"promise/server/object/entity"
	"promise/server/object/model"
	"strings"
)

// DefaultServerGroupID records the ID of default servergroup. We don't have to retrieve it each time.
var (
	DefaultServerGroupID string
	sgInstance           ServerGroupDBImplement
)

// ServerGroupDBImplement is the implement of ServerGroupDBInterface.
type ServerGroupDBImplement struct {
}

// GetServerGroupDB return the singleton.
func GetServerGroupDB() ServerGroupDBInterface {
	return &sgInstance
}

// GetServerGroup will get the group by id.
func (i *ServerGroupDBImplement) GetServerGroup(id string) *model.ServerGroup {
	var record entity.ServerGroup
	c := commonDB.GetConnection()
	if c.Where("\"ID\" = ?", id).First(&record).RecordNotFound() {
		return nil
	}
	return record.ToModel()
}

// GetServerGroupByName will get the group by name.
func (i *ServerGroupDBImplement) GetServerGroupByName(name string) *model.ServerGroup {
	var sg entity.ServerGroup
	c := commonDB.GetConnection()
	if c.Where("\"Name\" = ?", name).First(&sg).RecordNotFound() {
		return nil
	}
	return sg.ToModel()
}

// PostServerGroup will save the group if no group with the same name.
func (i *ServerGroupDBImplement) PostServerGroup(m *model.ServerGroup) (*model.ServerGroup, bool, error) {
	var e entity.ServerGroup

	c := commonDB.GetConnection()
	tx := c.Begin()
	if err := tx.Error; err != nil {
		log.WithFields(log.Fields{
			"name":  m.Name,
			"error": err}).
			Warn("Post servergroup in DB failed, start transaction failed.")
		return nil, false, err
	}
	if !tx.Where("\"Name\" = ?", m.Name).First(&e).RecordNotFound() {
		tx.Rollback()
		log.WithFields(log.Fields{
			"name": m.Name}).
			Warn("Post servergroup in DB failed, duplicated resource, transaction rollback.")
		return nil, true, commonConstError.ErrorResourceExist
	}
	e.Load(m)
	e.ID = uuid.New().String()
	if err := c.Create(&e).Error; err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"name":  m.Name,
			"error": err}).
			Warn("Post servergroup in DB failed, create resource failed, transaction rollback.")
		return nil, false, err
	}
	if err := tx.Commit().Error; err != nil {
		log.WithFields(log.Fields{
			"name":  m.Name,
			"error": err}).
			Warn("Post servergroup in DB failed, commit failed.")
		return nil, false, err
	}
	return e.ToModel(), false, nil
}

func (i *ServerGroupDBImplement) convertFilter(filter string) (string, error) {
	if filter == "" {
		return "", nil
	}
	cmds := strings.Split(filter, " ")
	if len(cmds) != 3 {
		return "", commonConstError.ErrorConvertFilter
	}
	switch strings.ToLower(cmds[1]) {
	case "eq":
		return "\"" + cmds[0] + "\"" + " = " + cmds[2], nil
	default:
		return "", commonConstError.ErrorConvertFilter
	}
}

// GetServerGroupCollection Get group collection by start and count.
func (i *ServerGroupDBImplement) GetServerGroupCollection(start int, count int, filter string) (*model.ServerGroupCollection, error) {
	var (
		total        int
		sgCollection []entity.ServerGroup
		ret          = new(model.ServerGroupCollection)
	)

	c := commonDB.GetConnection()
	c.Table("ServerGroup").Count(&total)
	if where, err := i.convertFilter(filter); err != nil {
		log.WithFields(log.Fields{
			"filter": filter,
			"error":  err}).
			Warn("Get servergroup in DB failed, convert filter failed.")
		c.Order("\"Name\" asc").Limit(count).Offset(start).Select([]string{"\"ID\"", "\"Name\""}).Find(&sgCollection)
	} else {
		log.WithFields(log.Fields{"where": where}).Debug("Convert filter success.")
		c.Order("\"Name\" asc").Limit(count).Offset(start).Where(where).Select([]string{"\"ID\"", "\"Name\""}).Find(&sgCollection)
	}
	ret.Start = start
	ret.Count = len(sgCollection)
	ret.Total = total
	for i := range sgCollection {
		ret.Members = append(ret.Members, model.ServerGroupMember{
			ID:   sgCollection[i].ID,
			Name: sgCollection[i].Name,
		})
	}
	return ret, nil
}

// DeleteServerGroup will delete the servergroup and return the deleted one.
// If it's the default servergroup, previous one equals nil, error will return.
// If the servergroup doesn't exist, previous one equals nil, error equals nil.
// If any error happended in DB operation, previous one equals nil, error will return.
// In any other cases, return the previous one and error equals nil.
func (i *ServerGroupDBImplement) DeleteServerGroup(id string) (*model.ServerGroup, error) {
	var sg, previous entity.ServerGroup

	if id == DefaultServerGroupID {
		return nil, consterror.ErrorDeleteDefaultServerGroup
	}

	if id == "" {
		return nil, commonConstError.ErrorIDFormat
	}

	c := commonDB.GetConnection()

	tx := c.Begin()
	if err := tx.Error; err != nil {
		log.WithFields(log.Fields{
			"id":    id,
			"error": err}).
			Warn("Delete servergroup in DB failed, start transaction failed.")
		return nil, err
	}
	if tx.Where("\"ID\" = ?", id).First(&previous).RecordNotFound() {
		tx.Rollback()
		log.WithFields(log.Fields{
			"id": id}).
			Warn("Delete servergroup in DB failed, resource does not exist, transaction rollback.")
		return nil, nil
	}

	sg.ID = id
	if err := tx.Delete(&sg).Error; err != nil {
		log.WithFields(log.Fields{
			"id": id}).
			Warn("Delete servergroup in DB failed, delete resource failed, transaction rollback.")
		return nil, err
	}
	if err := tx.Commit().Error; err != nil {
		log.WithFields(log.Fields{
			"id":    id,
			"error": err}).
			Warn("Delete servergroup in DB failed, commit failed.")
		return nil, err
	}
	return previous.ToModel(), nil
}

// DeleteServerGroupCollection will delete all the group except the default "all".
func (i *ServerGroupDBImplement) DeleteServerGroupCollection() error {
	c := commonDB.GetConnection()
	return c.Where("\"Name\" <> ?", "all").Delete(entity.ServerGroup{}).Error
}
