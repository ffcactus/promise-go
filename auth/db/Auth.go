package db

import (
	"github.com/google/uuid"
	"promise/auth/object/entity"
	"promise/auth/object/model"
	"promise/base"
)

// Auth Auth DB implementation.
type Auth struct {
}

// GetDBInstance Get DB instance.
func GetDBInstance() *Auth {
	return &Auth{}
}

// GetAccountByName Get user by name.
// In any case that the user can't be found, return nil.
func (impl *Auth) GetAccountByName(username string) *model.Account {
	c := base.GetConnection()
	account := new(entity.Account)
	if c.Where("Name = ?", username).First(account).RecordNotFound() {
		return nil
	}
	return account.Model()
}

// PostAccount Post a user in DB.
func (impl *Auth) PostAccount(account *model.Account) *model.Account {
	c := base.GetConnection()
	e := new(entity.Account)
	e.Load(account)
	e.ID = uuid.New().String()
	c.Create(e)
	return e.Model()

}

// PostSession Post a session in DB.
func (impl *Auth) PostSession(session *model.Session) *model.Session {
	c := base.GetConnection()
	e := new(entity.Session)
	e.Load(session)
	c.Create(e)
	return e.Model()
}

// GetSessionByToken Get session by token.
func (impl *Auth) GetSessionByToken(token string) *model.Session {
	c := base.GetConnection()
	session := new(entity.Session)
	if c.Where("Token = ?", token).First(session).RecordNotFound() {
		return nil
	}
	return session.Model()
}
