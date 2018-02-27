package db

import (
	"promise/auth/object/entity"
	"promise/auth/object/model"
	commonDB "promise/common/db"

	"github.com/google/uuid"
)

// AuthDBImplement Auth DB implementation.
type AuthDBImplement struct {
}

// GetDBInstance Get DB instance.
func GetDBInstance() AuthDBInterface {
	return &AuthDBImplement{}
}

// GetAccountByName Get user by name.
// In any case that the user can't be found, return nil.
func (impl *AuthDBImplement) GetAccountByName(username string) *model.Account {
	c := commonDB.GetConnection()
	account := new(entity.Account)
	if c.Where("Name = ?", username).First(account).RecordNotFound() {
		return nil
	}
	return account.Model()
}

// PostAccount Post a user in DB.
func (impl *AuthDBImplement) PostAccount(account *model.Account) *model.Account {
	c := commonDB.GetConnection()
	e := new(entity.Account)
	e.Load(account)
	e.ID = uuid.New().String()
	c.Create(e)
	return e.Model()

}

// PostSession Post a session in DB.
func (impl *AuthDBImplement) PostSession(session *model.Session) *model.Session {
	c := commonDB.GetConnection()
	e := new(entity.Session)
	e.Load(session)
	c.Create(e)
	return e.Model()
}

// GetSessionByToken Get session by token.
func (impl *AuthDBImplement) GetSessionByToken(token string) *model.Session {
	c := commonDB.GetConnection()
	session := new(entity.Session)
	if c.Where("Token = ?", token).First(session).RecordNotFound() {
		return nil
	}
	return session.Model()
}
