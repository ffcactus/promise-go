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

// DebugInfo return the debug name of this entity.
func (e *ServerServerGroup) DebugInfo() string {
	return e.ServerID + " " + e.ServerGroupID
}

// PropertyNameForDuplicationCheck return the property name used for duplication check.
func (e *ServerServerGroup) PropertyNameForDuplicationCheck() string {
	return ""
}

// Preload return the property names that need to be preload.
func (e *ServerServerGroup) Preload() []string {
	return []string{}
}

// Association return all the assocations that need to delete when deleting a resource.
func (e *ServerServerGroup) Association() []interface{} {
	return []interface{}{}
}

// Tables returns the tables to delete when you want delete all the resources.
func (e *ServerServerGroup) Tables() []interface{} {
	return []interface{}{new(ServerServerGroup)}
}

// FilterNameList return all the property name that can be used in filter.
func (e *ServerServerGroup) FilterNameList() []string {
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

// ToCollectionMember convert the entity to member.
func (e *ServerServerGroup) ToCollectionMember() base.CollectionMemberModelInterface {
	m := new(model.ServerServerGroupCollectionMember)
	base.EntityToMember(&e.Entity, &m.CollectionMemberModel)
	m.ServerID = e.ServerID
	m.ServerGroupID = e.ServerGroupID
	return m
}
