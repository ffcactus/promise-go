package entity

import (
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/server/object/model"
)

// ServerGroup is the entity of servergroup.
type ServerGroup struct {
	base.Entity
	Name        string `gorm:"column:Name"`
	Description string `gorm:"column:Description"`
}

// TableName will set the table name.
func (ServerGroup) TableName() string {
	return "ServerGroup"
}

// GetDebugName return the debug name of this entity.
func (e *ServerGroup) GetDebugName() string {
	return e.Name
}

// GetPropertyNameForDuplicationCheck return the property name used for duplication check.
func (e *ServerGroup) GetPropertyNameForDuplicationCheck() string {
	return "Name"
}

// GetPreload return the property names that need to be preload.
func (e *ServerGroup) GetPreload() []string {
	return []string{}
}

// GetAssociation return all the assocations that need to delete when deleting a resource.
func (e *ServerGroup) GetAssociation() []interface{} {
	ret := []interface{}{}
	return ret
}

// GetTables returns the tables to delete when you want delete all the resources.
func (e *ServerGroup) GetTables() []interface{} {
	return []interface{}{}
}

// GetFilterNameList return all the property name that can be used in filter.
func (e *ServerGroup) GetFilterNameList() []string {
	return []string{"Name"}
}

// Load will load data from model. this function is used on POST.
func (e *ServerGroup) Load(i base.ModelInterface) error {
	m, ok := i.(*model.ServerGroup)
	if !ok {
		log.Error("entity.ServerGroup.Load() failed, convert interface failed.")
		return base.ErrorDataConvert
	}
	base.EntityLoad(&e.Entity, &m.Model)
	e.Name = m.Name
	e.Description = m.Description
	return nil
}

// ToModel will create a new model from entity.
func (e *ServerGroup) ToModel() base.ModelInterface {
	m := model.ServerGroup{}
	base.EntityToModel(&e.Entity, &m.Model)
	m.Name = e.Name
	m.Description = e.Description
	return &m
}
