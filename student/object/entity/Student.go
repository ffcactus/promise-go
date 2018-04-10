package entity

import (
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/student/object/model"
)

// Student is the entity of student.
type Student struct {
	base.Entity
	Name string `gorm:"column:Name"`
	Age  int    `gorm:"column:Age"`
}

// TableName will set the table name.
func (Student) TableName() string {
	return "Student"
}

// GetID return the ID.
func (e *Student) GetID() string {
	return e.ID
}

// SetID set the ID.
func (e *Student) SetID(id string) {
	e.ID = id
}

// GetDebugName return the debug name of this entity.
func (e *Student) GetDebugName() string {
	return e.Name
}

// GetPropertyNameForDuplicationCheck return the property name used for duplication check.
func (e *Student) GetPropertyNameForDuplicationCheck() string {
	return "Name"
}

// GetPreload return the property names that need to be preload.
func (e *Student) GetPreload() []string {
	return nil
}

// GetAssociation return all the association address that need to delete.
func (e *Student) GetAssociation() []interface{} {
	return nil
}

// Load will load info from model.
func (e *Student) Load(i base.ModelInterface) error {
	m, ok := i.(*model.Student)
	if !ok {
		log.Error("entity.Student.Load() failed, convert interface failed.")
		return base.ErrorDataConvert
	}
	base.EntityLoad(&e.Entity, &m.Model)
	e.Name = m.Name
	e.Age = m.Age
	return nil
}

// ToModel convert the entity to model.
func (e *Student) ToModel() base.ModelInterface {
	m := new(model.Student)
	base.EntityToModel(&e.Entity, &m.Model)
	m.Name = e.Name
	m.Age = e.Age
	return m
}

// ToMember convert the entity to member.
func (e *Student) ToMember() base.MemberInterface {
	m := new(model.StudentMember)
	base.EntityToMember(&e.Entity, &m.Member)
	m.Name = e.Name
	return m
}
