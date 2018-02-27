package db

import (
	"promise/auth/object/model"
)

// AuthDBInterface Auth DB interface.
type AuthDBInterface interface {
	GetAccountByName(username string) *model.Account
	PostAccount(account *model.Account) *model.Account
	PostSession(account *model.Session) *model.Session
	GetSessionByToken(token string) *model.Session
}
