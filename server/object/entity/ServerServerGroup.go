package entity

import (
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/server/object/model"
)

// ServerServerGroup is the entity of servergroup.
type ServerServerGroup struct {
	base.Entity
	ServerID      string `gorm:"column:ServerID"`
	ServerGroupID string `gorm:"column:ServerGroupID"`
}

// TableName will set the table name.
func (ServerServerGroup) TableName() string {
	return "ServerServerGroup"
}

// GetDebugName return the debug name of this entity.
func (e *ServerServerGroup) GetDebugName() string {
	return e.ServerID + " " + e.ServerGroupID
}

// GetPropertyNameForDuplicationCheck return the property name used for duplication check.
func (e *ServerServerGroup) GetPropertyNameForDuplicationCheck() string {
	return ""
}

// GetPreload return the property names that need to be preload.
func (e *ServerServerGroup) GetPreload() []string {
	return []string{}
}

// GetAssociation return all the assocations that need to delete when deleting a resource.
func (e *ServerServerGroup) GetAssociation() []interface{} {
	ret := []interface{}{}
	return ret
}

// GetTables returns the tables to delete when you want delete all the resources.
func (e *ServerServerGroup) GetTables() []interface{} {
	return []interface{}{}
}

// GetFilterNameList return all the property name that can be used in filter.
func (e *ServerServerGroup) GetFilterNameList() []string {
	return []string{}
}

// Load will load data from model. this function is used on POST.
func (e *ServerServerGroup) Load(i base.ModelInterface) error {
	m, ok := i.(*model.ServerServerGroup)
	if !ok {
		log.Error("entity.ServerServerGroup.Load() failed, convert interface failed.")
		return base.ErrorDataConvert
	}
	base.EntityLoad(&e.Entity, &m.Model)
	e.ServerID = m.ServerID
	e.ServerGroupID = m.ServerGroupID
	return nil
}

// ToModel will create a new model from entity.
func (e *ServerServerGroup) ToModel() base.ModelInterface {
	m := model.ServerServerGroup{}
	base.EntityToModel(&e.Entity, &m.Model)
	m.ServerID = e.ServerID
	m.ServerGroupID = e.ServerGroupID
	return &m
}
