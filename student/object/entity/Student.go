package entity

import (
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/student/object/model"
)

// Student is the entity of student.
type Student struct {
	base.Entity
	Name   string  `gorm:"column:Name"`
	Age    int     `gorm:"column:Age"`
	Phones []Phone `gorm:"column:Phones;ForeignKey:StudentRef"`
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
	return []string{"Phones"}
}

// GetFilterNameList return all the property name that can be used in filter.
func (e *Student) GetFilterNameList() []string {
	return []string{"Name", "Age"}
}

// GetAssociation return all the assocations that need to delete when deleting a resource.
func (e *Student) GetAssociation() []interface{} {
	ret := []interface{}{}
	for _, v := range e.Phones {
		ret = append(ret, v) 
	}
	ret = append(ret, e)
	return ret
}

// GetTables returns the tables to delete when you want delete all the resources.
func (e *Student) GetTables() []interface{} {
	return []interface{}{Phone{}, Student{}}
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
	for _, v := range m.Phones {
		phone := Phone{}
		phone.Load(&v)
		e.Phones = append(e.Phones, phone)
	}
	return nil
}

// ToModel converts the entity to model.
func (e *Student) ToModel() base.ModelInterface {
	m := new(model.Student)
	base.EntityToModel(&e.Entity, &m.Model)
	m.Name = e.Name
	m.Age = e.Age
	for _, v := range e.Phones {
		phone := v.ToModel()
		m.Phones = append(m.Phones, *phone)
	}
	return m
}

// ToCollectionMember convert the entity to member.
func (e *Student) ToCollectionMember() base.CollectionMemberModelInterface {
	m := new(model.StudentCollectionMember)
	base.EntityToMember(&e.Entity, &m.CollectionMemberModel)
	m.Name = e.Name
	return m
}
