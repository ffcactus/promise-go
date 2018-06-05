package entity

import (
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/server/object/model"
)

// ServerGroup is the entity.
type AdapterConfig struct {
	base.Entity
	Name        string `gorm:"column:Name"`
}

// TableName will set the table name.
func (AdapterConfig) TableName() string {
	return "AdapterConfig"
}

// DebugInfo return the debug name of this entity.
func (e *AdapterConfig) DebugInfo() string {
	return e.Name
}

// PropertyNameForDuplicationCheck return the property name used for duplication check.
func (e *AdapterConfig) PropertyNameForDuplicationCheck() string {
	return "Name"
}

// Preload return the property names that need to be preload.
func (e *AdapterConfig) Preload() []string {
	return []string{}
}

// Association return all the assocations that need to delete when deleting a resource.
func (e *AdapterConfig) Association() []interface{} {
	ret := []interface{}{}
	return ret
}

// Tables returns the tables to delete when you want delete all the resources.
func (e *AdapterConfig) Tables() []interface{} {
	return []interface{}{new(AdapterConfig)}
}

// FilterNameList return all the property name that can be used in filter.
func (e *AdapterConfig) FilterNameList() []string {
	return []string{"Name"}
}

// Load will load data from model. this function is used on POST.
func (e *AdapterConfig) Load(i base.ModelInterface) error {
	m, ok := i.(*model.AdapterConfig)
	if !ok {
		log.Error("entity.AdapterConfig.Load() failed, convert interface failed.")
		return base.ErrorDataConvert
	}
	base.EntityLoad(&e.Entity, &m.Model)
	e.Name = m.Name
	return nil
}

// ToModel will create a new model from entity.
func (e *AdapterConfig) ToModel() base.ModelInterface {
	m := model.AdapterConfig{}
	base.EntityToModel(&e.Entity, &m.Model)
	m.Name = e.Name
	return &m
}

// ToCollectionMember convert the entity to member.
func (e *AdapterConfig) ToCollectionMember() base.CollectionMemberModelInterface {
	m := new(model.AdapterConfigCollectionMember)
	base.EntityToMember(&e.Entity, &m.CollectionMemberModel)
	m.Name = e.Name
	return m
}
