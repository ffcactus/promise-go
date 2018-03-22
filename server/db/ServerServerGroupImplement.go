package db

import (
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	commonDB "promise/common/db"
	"promise/server/object/entity"
	"promise/server/object/constError"
	commonConstError "promise/common/object/constError"
	"promise/server/object/model"
	"strings"
)

// ServerServerGroupImplement is the SQL implement.
type ServerServerGroupImplement struct {
}

// GetServerServerGroupInstance will return the server-servergroup DB impl.
func GetServerServerGroupInstance() ServerServerGroupInterface {
	return new(ServerServerGroupImplement)
}

// GetServerServerGroup will get the resource by ID.
func (i *ServerServerGroupImplement) GetServerServerGroup(id string) *model.ServerServerGroup {
	var e = new(entity.ServerServerGroup)

	c := commonDB.GetConnection()
	if c.Where("\"ID\" = ?", id).First(e).RecordNotFound() {
		return nil
	}
	return e.ToModel()
}

// PostServerServerGroup will post the server-servergroup.
// We need check the duplication that if the pair of server ID and servergroup ID exist, iti's a duplicated one.
// We need make sure both server ID and servergroup ID exist.
// So we need transaction here.
func (i *ServerServerGroupImplement) PostServerServerGroup(m *model.ServerServerGroup) (*model.ServerServerGroup, bool, error) {
	var tempSsg = new(entity.ServerServerGroup)
	var ssg = new(entity.ServerServerGroup)
	var s = new(entity.Server)
	var sg = new(entity.ServerGroup)

	ssg.Load(m)
	s.ID = m.ServerID
	sg.ID = m.ServerGroupID

	c := commonDB.GetConnection()
	tx := c.Begin()
	if err := tx.Error; err != nil {
		log.WithFields(log.Fields{
			"serverID":      ssg.ServerID,
			"serverGroupID": ssg.ServerGroupID,
			"error":         err}).
			Warn("Post server-servergroup in DB failed, start transaction failed.")
	}
	// Check if server ID exist.
	if tx.Where("\"id\" = ?", s.ID).First(s).RecordNotFound() {
		tx.Rollback()
		log.WithFields(log.Fields{"serverID": s.ID}).
			Warn("Post server-servergroup in DB failed, server ID does not exist, transaction rollback.")
		return nil, false, commonConstError.ErrorResourceNotExist
	}
	// Check if servergroup ID exsit.
	if tx.Where("\"ID\" = ?", sg.ID).First(sg).RecordNotFound() {
		tx.Rollback()
		log.WithFields(log.Fields{"serverGroupID": sg.ID}).
			Warn("Post server-servergroup in DB failed, servergroup ID does not exist, transaction rollback.")
		return nil, false, commonConstError.ErrorResourceNotExist
	}
	// Check duplication.
	if !tx.Where("\"ServerID\" = ? AND \"ServerGroupID\" = ?", ssg.ServerID, ssg.ServerGroupID).First(tempSsg).RecordNotFound() {
		tx.Rollback()
		log.WithFields(log.Fields{
			"serverID":      tempSsg.ServerID,
			"serverGroupID": tempSsg.ServerGroupID,
			"ID":            tempSsg.ID}).
			Warn("Post server-servergroup in DB failed, duplicated resource.")
		return tempSsg.ToModel(), true, nil
	}
	ssg.ID = uuid.New().String()
	if err := tx.Create(ssg).Error; err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"serverID":      ssg.ServerID,
			"serverGroupID": ssg.ServerGroupID,
			"error":         err}).
			Warn("Post server-servergroup in DB failed, create resource failed, transaction rollback.")
		return nil, false, err
	}
	if err := tx.Commit().Error; err != nil {
		log.WithFields(log.Fields{
			"serverID":      ssg.ServerID,
			"serverGroupID": ssg.ServerGroupID,
			"error":         err}).
			Warn("Post server-servergroup in DB failed, commit failed.")
		return nil, false, err
	}
	return ssg.ToModel(), false, nil
}

func (i *ServerServerGroupImplement) convertFilter(filter string) (string, error) {
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

// GetServerServerGroupCollection Get group collection by start and count.
func (i *ServerServerGroupImplement) GetServerServerGroupCollection(start int, count int, filter string) (*model.ServerServerGroupCollection, error) {
	var (
		total      int
		collection []entity.ServerServerGroup
		ret        = new(model.ServerServerGroupCollection)
	)

	c := commonDB.GetConnection()
	c.Table("ServerServerGroup").Count(&total)
	if where, err := i.convertFilter(filter); err != nil {
		log.WithFields(log.Fields{
			"filter": filter,
			"error":  err}).
			Warn("Get server-servergroup in DB failed, convert filter failed.")
		c.Limit(count).Offset(start).Find(&collection)
	} else {
		log.WithFields(log.Fields{"where": where}).Debug("Convert filter success.")
		c.Limit(count).Offset(start).Where(where).Find(&collection)
	}
	ret.Start = start
	ret.Count = len(collection)
	ret.Total = total
	for i := range collection {
		ret.Members = append(ret.Members, model.ServerServerGroupMember{
			ID:            collection[i].ID,
			ServerID:      collection[i].ServerID,
			ServerGroupID: collection[i].ServerGroupID,
		})
	}
	return ret, nil
}

// DeleteServerServerGroup will delete the servergroup and return the deleted one.
// If it's the default servergroup, previous one equals nil, error will return.
// If the server-servergroup doesn't exist, previous one equals nil, error equals nil.
// If any error happended in DB operation, previous one equals nil, error will return.
// In any other cases, return the previous one and error equals nil.
func (i *ServerServerGroupImplement) DeleteServerServerGroup(id string) (*model.ServerServerGroup, error) {
	var ssg, previous entity.ServerServerGroup

	if id == "" {
		return nil, commonConstError.ErrorIDFormat
	}

	c := commonDB.GetConnection()

	tx := c.Begin()
	if err := tx.Error; err != nil {
		log.WithFields(log.Fields{
			"id":  id,
			"error": err}).
			Warn("Delete server-servergroup in DB failed, start transaction failed.")
		return nil, err
	}
	if tx.Where("\"ID\" = ?", id).First(&previous).RecordNotFound() {
		tx.Rollback()
		log.WithFields(log.Fields{
			"id": id}).
			Warn("Delete server-servergroup in DB failed, resource does not exist, transaction rollback.")
		return nil, nil
	}
	if previous.ServerGroupID == DefaultServerGroupID {
		return nil, constError.ErrorDeleteDefaultServerServerGroup
	}
	ssg.ID = id	
	if err := tx.Delete(&ssg).Error; err!= nil {
		return nil, err
	}
	if err := tx.Commit().Error; err != nil {
		log.WithFields(log.Fields{
			"id":  id,
			"error": err}).
			Warn("Delete server-servergroup in DB failed, commit failed.")
		return nil, err
	}
	return previous.ToModel(), nil
}

// DeleteServerServerGroupCollection will remove all server-servergroup
// except the default association to default server group.
func (i *ServerServerGroupImplement) DeleteServerServerGroupCollection() error {
	c := commonDB.GetConnection()
	return c.Where("\"ServerGroupID\" <> ?", DefaultServerGroupID).Delete(entity.ServerServerGroup{}).Error
}
