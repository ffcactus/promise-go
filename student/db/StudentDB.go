package db

import (
	"github.com/jinzhu/gorm"
	"promise/base"
	"promise/apps"
	"promise/student/object/entity"
)

// StudentDB is the DB implementation for student.
type StudentDB struct {

}

// NewEntity return the a new entity.
func (impl *StudentDB) NewEntity() base.EntityInterface {
	// return new(entity.Student)
	return &base.Entity{
		TemplateImpl: new(entity.Student),
	}
}

// GetConnection return the DB connection.
func (impl *StudentDB) GetConnection() *gorm.DB {
	return apps.GetConnection()
}

// NeedCheckDuplication return if need check duplication for entity.
func (impl *StudentDB) NeedCheckDuplication() bool {
	return true
}
