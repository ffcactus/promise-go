package db

import (
	"fmt"
	commonDB "promise/common/db"
	log "github.com/sirupsen/logrus"
	"promise/server/object/model"
	"promise/server/object/entity"
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
	if c.Where("ID = ?, ID").First(e).RecordNotFound() {
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
			"serverID": ssg.ServerID,
			"serverGroupID": ssg.ServerGroupID,
			"error": err}).
			Warn("Post server-servergroup in DB failed, start transaction failed.")
	}
	// Check if server ID exist.
	if tx.Where("ID = ?", s.ID).First(s).RecordNotFound() {
		tx.Rollback()
		log.WithFields(log.Fields{"serverID": s.ID}).
			Debug("Post server-servergroup in DB failed, server ID does not exist, transaction rollback.")
		return nil, false, fmt.Errorf("server ID does not exist")
	}
	// Check if servergroup ID exsit.
	if tx.Where("ID = ?", sg.ID).First(sg).RecordNotFound() {
		tx.Rollback()
		log.WithFields(log.Fields{"serverGroupID": sg.ID}).
			Debug("Post server-servergroup in DB failed, servergroup ID does not exist, transaction rollback.")
		return nil, false, fmt.Errorf("servergroup ID does not exist")		
	}
	// Check duplication.
	if ! tx.Where("server_id = ? AND server_group_id = ?", ssg.ServerID, ssg.ServerGroupID).First(tempSsg).RecordNotFound() {
		tx.Rollback()
		log.WithFields(log.Fields{
			"serverID": tempSsg.ServerID,
			"serverGroupID": tempSsg.ServerGroupID,
			"ID": tempSsg.ID}).
			Warn("Post server-servergroup in DB failed, duplicated resource.")		
		return tempSsg.ToModel(), true, nil
	}
	if err := tx.Create(ssg).Error; err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"serverID": ssg.ServerID,
			"serverGroupID": ssg.ServerGroupID,
			"error": err}).
			Warn("Post server-servergroup in DB failed, create resource failed, transaction rollback.")
		return nil, false, err
	}
	tx.Commit()
	return ssg.ToModel(), false, nil
}