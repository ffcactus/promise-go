package db

import (
	"fmt"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	commonDB "promise/common/db"
	"promise/server/object/entity"
	"promise/server/object/model"
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
	if c.Where("ID = ?", id).First(&sg).RecordNotFound() {
		return nil
	}
	return sg.ToModel()
}

// GetServerGroupByName will get the group by name.
func (i *ServerGroupDBImplement) GetServerGroupByName(name string) *model.ServerGroup {
	var sg entity.ServerGroup
	c := commonDB.GetConnection()
	if c.Where("Name = ?", name).First(&sg).RecordNotFound() {
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
			"name": m.Name,
			"error": err}).
			Warn("Post servergroup in DB failed, start transaction failed.")		
		return nil, false, err
	}
	if !tx.Where("Name = ?", m.Name).First(&e).RecordNotFound() {
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
			"name": m.Name,
			"error": err}).
			Warn("Post servergroup in DB failed, create resource failed, transaction rollback.")
		return nil, false, err
	}
	if err := tx.Commit().Error; err != nil {
		log.WithFields(log.Fields{
			"name": m.Name,
			"error": err}).
			Warn("Post servergroup in DB failed, commit failed.")			
		return nil, false, err
	}
	return e.ToModel(), false, nil	

}

// GetServerGroupCollection Get group collection by start and count.
func (i *ServerGroupDBImplement) GetServerGroupCollection(start int, count int) (*model.ServerGroupCollection, error) {
	var (
		total        int
		sgCollection []entity.ServerGroup
		ret          = new(model.ServerGroupCollection)
	)

	c := commonDB.GetConnection()
	c.Table("server-group").Count(total)
	c.Order("Name asc").Limit(count).Offset(start).Select([]string{"ID", "Name"}).Find(&sgCollection)
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

	// If I need check the existance and error at the same time, should I use transaction?
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
	return c.Where("name <> ?", "all").Delete(entity.ServerGroup{}).Error
}
