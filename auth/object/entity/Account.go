package entity

import (
	"promise/auth/object/model"
)

// The Account object in DB.
type Account struct {
	ID           string `gorm:"column:ID;primary_key"`
	Name         string `gorm:"column:Name"`
	PasswordHash string `gorm:"column:PasswordHash"`
}

// TableName will set the table name.
func (Account) TableName() string {
	return "Account"
}

// Load Load model to entity.
func (e *Account) Load(user *model.Account) {
	e.Name = user.Name
	e.PasswordHash = user.PasswordHash
}

// Model Convert entity to model.
func (e *Account) Model() *model.Account {
	m := new(model.Account)
	m.ID = e.ID
	m.Name = e.Name
	return m
}
