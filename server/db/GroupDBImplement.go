package db

import (
	"fmt"
	"github.com/google/uuid"
	"promise/common/app"
	commonDB "promise/common/db"
	"promise/server/object/entity"
	"promise/server/object/model"
)

// GroupDBImplement is the implement of GroupDBInterface.
type GroupDBImplement struct {
}

// GetGroupDB return an implement.
func GetGroupDB() GroupDBInterface {
	return new(GroupDBImplement)
}

// GetGroup will get the server group by id.
func (i *GroupDBImplement) GetGroup(id string) *model.Group {
	var sg entity.Group
	c := commonDB.GetConnection()
	if c.Where("ID = ?", id).First(&sg).RecordNotFound() {
		return nil
	}
	return sg.ToModel()
}

// GetGroupByName will get the server group by name.
func (i *GroupDBImplement) GetGroupByName(name string) *model.Group {
	var sg entity.Group
	c := commonDB.GetConnection()
	if c.Where("Name = ?", name).First(&sg).RecordNotFound() {
		return nil
	}
	return sg.ToModel()
}

// PostGroup will save the server group if no group with the same name.
func (i *GroupDBImplement) PostGroup(m *model.Group) (*model.Group, bool, error) {
	var e entity.Group

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

// GetGroupCollection Get server group collection by start and count.
func (i *GroupDBImplement) GetGroupCollection(start int, count int) (*model.GroupCollection, error) {
	var (
		total        int
		sgCollection []entity.Group
		ret          = new(model.GroupCollection)
	)

	c := commonDB.GetConnection()
	c.Table("server-group").Count(total)
	c.Order("Name asc").Limit(count).Offset(start).Select([]string{"ID", "Name"}).Find(&sgCollection)
	ret.Start = start
	ret.Count = len(sgCollection)
	ret.Total = total
	for i := range sgCollection {
		ret.Members = append(ret.Members, model.GroupMember{
			ID:   sgCollection[i].ID,
			URI:  toGroupURI(sgCollection[i].ID),
			Name: sgCollection[i].Name,
		})
	}
	return ret, nil
}

// DeleteGroup will delete server group if exist.
func (i *GroupDBImplement) DeleteGroup(id string) (bool, error) {
	var sg entity.Group

	// If I need check the existance and error at the same time, should I use transaction?
	sg.ID = id
	c := commonDB.GetConnection()
	if c.Delete(&sg).RecordNotFound() {
		return false, nil
	}
	return true, nil
}

// DeleteGroupCollection will delete all the server group except the default "all".
func (i *GroupDBImplement) DeleteGroupCollection() error {
	c := commonDB.GetConnection()
	return c.Where("name <> ?", "all").Delete(entity.Group{}).Error
}

func toGroupURI(ID string) string {
	s := app.RootURL + "/servergroup" + ID
	return s
}
