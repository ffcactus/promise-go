package entity

import (
	"promise/auth/object/model"
)

// Session The Session object in DB.
type Session struct {
	ID        string `gorm:"primary_key"`
	AccountID string
	Token     string
	Expire    int64
}

// Load Load model to entity.
func (e *Session) Load(session *model.Session) {
	e.ID = session.ID
	e.AccountID = session.AccountID
	e.Token = session.Token
	e.Expire = session.Expire
}

// Model Convert entity to model.
func (e *Session) Model() *model.Session {
	m := new(model.Session)
	m.ID = e.ID
	m.AccountID = e.AccountID
	m.Token = e.Token
	m.Expire = e.Expire
	return m
}
