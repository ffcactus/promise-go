package entity

import (
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/enclosure/object/model"
)

// Enclosure is the entity.
type Enclosure struct {
	base.Entity
	Name        string `gorm:"column:Name"`
	Description string `gorm:"column:Description"`
	State       string `gorm:"column:State"`
	Health      string `gorm:"column:Health"`
}

// TableName will set the table name.
func (Enclosure) TableName() string {
	return "Enclosure"
}

// String return the debug name of this entity.
func (e Enclosure) String() string {
	return e.Name
}

// PropertyNameForDuplicationCheck return the property name used for duplication check.
func (e *Enclosure) PropertyNameForDuplicationCheck() string {
	return "Name"
}

// Preload return the property names that need to be preload.
func (e *Enclosure) Preload() []string {
	return []string{}
}

// Association return all the assocations that need to delete when deleting a resource.
func (e *Enclosure) Association() []interface{} {
	ret := []interface{}{}
	return ret
}

// Tables returns the tables to delete when you want delete all the resources.
func (e *Enclosure) Tables() []interface{} {
	return []interface{}{new(Enclosure)}
}

// FilterNameList return all the property name that can be used in filter.
func (e *Enclosure) FilterNameList() []string {
	return []string{"Name"}
}

// Load will load data from model. this function is used on POST.
func (e *Enclosure) Load(i base.ModelInterface) error {
	m, ok := i.(*model.Enclosure)
	if !ok {
		log.Error("entity.Enclosure.Load() failed, convert interface failed.")
		return base.ErrorDataConvert
	}
	base.EntityLoad(&e.Entity, &m.Model)
	e.Name = m.Name
	e.Description = m.Description
	e.State = m.State
	e.Health = m.Health
	return nil
}

// ToModel will create a new model from entity.
func (e *Enclosure) ToModel() base.ModelInterface {
	m := model.Enclosure{}
	base.EntityToModel(&e.Entity, &m.Model)
	m.Name = e.Name
	m.Description = e.Description
	m.State = e.State
	m.Health = e.Health
	return &m
}

// ToCollectionMember convert the entity to member.
func (e *Enclosure) ToCollectionMember() base.CollectionMemberModelInterface {
	m := new(model.EnclosureCollectionMember)
	base.EntityToMember(&e.Entity, &m.CollectionMemberModel)
	m.Name = e.Name
	m.Description = e.Description
	m.State = e.State
	m.Health = e.Health
	return m
}
