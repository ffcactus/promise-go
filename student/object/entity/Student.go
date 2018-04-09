package entity

import (
	"promise/base"
)

// Student is the entity of student.
type Student struct {
	base.Entity
	Name        string `gorm:"column:Name"`
	Age int `gorm:"column:Age"`
}

// TableName will set the table name.
func (Student) TableName() string {
	return "Student"
}