package db

import (
	"fmt"
	"github.com/google/uuid"
	commonDB "promise/common/db"
	"promise/server/object/entity"
	"promise/server/object/model"
)

// ServerGroupDBImplement is the implement of ServerGroupDBInterface.
type ServerGroupDBImplement struct {
}

// GetServerGroupDB return an implement.
func GetServerGroupDB() ServerGroupDBInterface {
	return new(ServerGroupDBImplement)
}

// GetServerGroupByName will get the server group by name.
func (i *ServerGroupDBImplement) GetServerGroupByName(name string) *model.ServerGroup {
	var sg entity.ServerGroup
	c := commonDB.GetConnection()
	if c.Where("Name = ?", name).First(&sg).RecordNotFound() {
		return nil
	}
	return sg.ToModel()
}

// PostServerGroup will save the server group if no group with the same name.
func (i *ServerGroupDBImplement) PostServerGroup(m *model.ServerGroup) (*model.ServerGroup, bool, error) {
	var e entity.ServerGroup

	c := commonDB.GetConnection()
	tx := c.Begin()
	if tx.Where("Name = ?", m.Name).First(&e).RecordNotFound() {
		e.Load(m)
		e.ID = uuid.New().String()
		if err := c.Create(&e).Error; err != nil {
			tx.Rollback()
			return nil, false, err
		} else {
			tx.Commit()
			return e.ToModel(), false, nil
		}
	} else {
		tx.Rollback()
		return nil, true, fmt.Errorf("already exist.")
	}
}
