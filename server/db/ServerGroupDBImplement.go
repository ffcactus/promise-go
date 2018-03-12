package db

import (
	"fmt"
	"github.com/google/uuid"
	"promise/common/app"
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

// GetServerGroup will get the server group by id.
func (i *ServerGroupDBImplement) GetServerGroup(id string) *model.ServerGroup {
	var sg entity.ServerGroup
	c := commonDB.GetConnection()
	if c.Where("ID = ?", id).First(&sg).RecordNotFound() {
		return nil
	}
	return sg.ToModel()
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
		}
		tx.Commit()
		return e.ToModel(), false, nil
	}
	tx.Rollback()
	return nil, true, fmt.Errorf("already exist")
}

// GetServerGroupCollection Get server group collection by start and count.
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
			URI:  toServerGroupURI(sgCollection[i].ID),
			Name: sgCollection[i].Name,
		})
	}
	return ret, nil
}

// DeleteServerGroup will delete server group if exist.
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

// DeleteServerGroupCollection will delete all the server group except the default "all".
func (i *ServerGroupDBImplement) DeleteServerGroupCollection() error {
	c := commonDB.GetConnection()
	return c.Where("name <> ?", "all").Delete(entity.ServerGroup{}).Error
}

func toServerGroupURI(ID string) string {
	s := app.RootURL + "/servergroup" + ID
	return s
}
