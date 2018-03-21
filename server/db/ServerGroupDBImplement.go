package db

import (
	"fmt"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	commonDB "promise/common/db"
	"promise/server/object/entity"
	"promise/server/object/model"
	"strings"
)

// DefaultServerGroupID records the ID of default servergroup. We don't have to retrieve it each time.
var DefaultServerGroupID string

// ServerGroupDBImplement is the implement of ServerGroupDBInterface.
type ServerGroupDBImplement struct {
}

// GetServerGroupDB return an implement.
func GetServerGroupDB() ServerGroupDBInterface {
	return new(ServerGroupDBImplement)
}

// GetServerGroup will get the group by id.
func (i *ServerGroupDBImplement) GetServerGroup(id string) *model.ServerGroup {
	var sg entity.ServerGroup
	c := commonDB.GetConnection()
	if c.Where("\"ID\" = ?", id).First(&sg).RecordNotFound() {
		return nil
	}
	return sg.ToModel()
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
		return nil, true, fmt.Errorf("already exist")
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
	cmds := strings.Split(filter, " ")
	if len(cmds) != 3 {
		return "", fmt.Errorf("convert filter failed")
	}
	switch strings.ToLower(cmds[1]) {
	case "eq":
		return "\"" + cmds[0] + "\"" + " = " + cmds[2], nil
	default:
		return "", fmt.Errorf("convert filter failed")
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

// DeleteServerGroup will delete group if exist.
func (i *ServerGroupDBImplement) DeleteServerGroup(id string) (bool, error) {
	var sg entity.ServerGroup

	if id == DefaultServerGroupID {
		return true, fmt.Errorf("can not delete default servergroup")
	}

	if id == "" {
		return false, fmt.Errorf("id can not be empty")
	}
	// If I need check the existance and error at the same time, should I use transaction?
	log.Info("------ id " + id)
	sg.ID = id
	c := commonDB.GetConnection()
	if c.Delete(&sg).RecordNotFound() {
		return false, nil
	}
	return true, nil
}

// DeleteServerGroupCollection will delete all the group except the default "all".
func (i *ServerGroupDBImplement) DeleteServerGroupCollection() error {
	c := commonDB.GetConnection()
	return c.Where("\"Name\" <> ?", "all").Delete(entity.ServerGroup{}).Error
}
