package db

import (
	"github.com/jinzhu/gorm"
	"promise/apps"
	"promise/base"
	"promise/student/object/entity"
)

// StudentDB is the DB implementation for student.
type StudentDB struct {
}

// NewEntity return the a new entity.
func (impl *StudentDB) NewEntity() base.EntityInterface {
	e := new(entity.Student)
	e.Entity.TemplateImpl = e
	return e
}

// NewEntityCollection return a collection of entity.
func (impl *StudentDB) NewEntityCollection() []base.EntityInterface {
	return []entity.Student
}

// GetConnection return the DB connection.
func (impl *StudentDB) GetConnection() *gorm.DB {
	return apps.GetConnection()
}

// NeedCheckDuplication return if need check duplication for entity.
func (impl *StudentDB) NeedCheckDuplication() bool {
	return true
}
