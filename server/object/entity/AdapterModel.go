package entity

import (
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/server/object/model"
)

// AdapterModel is the entity.
type AdapterModel struct {
	base.Entity
	Name string `gorm:"column:Name"`
}

// TableName will set the table name.
func (AdapterModel) TableName() string {
	return "AdapterModel"
}

// DebugInfo return the debug name of this entity.
func (e *AdapterModel) DebugInfo() string {
	return e.Name
}

// PropertyNameForDuplicationCheck return the property name used for duplication check.
func (e *AdapterModel) PropertyNameForDuplicationCheck() string {
	return "Name"
}

// Preload return the property names that need to be preload.
func (e *AdapterModel) Preload() []string {
	return []string{}
}

// Association return all the assocations that need to delete when deleting a resource.
func (e *AdapterModel) Association() []interface{} {
	ret := []interface{}{}
	return ret
}

// Tables returns the tables to delete when you want delete all the resources.
func (e *AdapterModel) Tables() []interface{} {
	return []interface{}{new(AdapterModel)}
}

// FilterNameList return all the property name that can be used in filter.
func (e *AdapterModel) FilterNameList() []string {
	return []string{"Name"}
}

// Load will load data from model. this function is used on POST.
func (e *AdapterModel) Load(i base.ModelInterface) error {
	m, ok := i.(*model.AdapterModel)
	if !ok {
		log.Error("entity.AdapterModel.Load() failed, convert interface failed.")
		return base.ErrorDataConvert
	}
	base.EntityLoad(&e.Entity, &m.Model)
	e.Name = m.Name
	return nil
}

// ToModel will create a new model from entity.
func (e *AdapterModel) ToModel() base.ModelInterface {
	m := model.AdapterModel{}
	base.EntityToModel(&e.Entity, &m.Model)
	m.Name = e.Name
	return &m
}

// ToCollectionMember convert the entity to member.
func (e *AdapterModel) ToCollectionMember() base.CollectionMemberModelInterface {
	m := new(model.AdapterModelCollectionMember)
	base.EntityToMember(&e.Entity, &m.CollectionMemberModel)
	m.Name = e.Name
	return m
}
