package service

import (
	"promise/auth/db"
	"promise/auth/object/dto"
	"promise/auth/object/model"
	commonModel "promise/common/object/model"

	log "github.com/sirupsen/logrus"
)

// Login On success return the session.
func Login(request *dto.PostLoginRequest) (*model.Session, []commonModel.Message) {
	dbInstance := db.GetDBInstance()
	account := dbInstance.GetAccountByName(request.Name)
	if account == nil {
		return nil, []commonModel.Message{model.NewMessageAuthIncorrectCredential()}
	}
	// We should valid the password here.
	session := CreateSession(account)
	savedSession := dbInstance.PostSession(session)
	if savedSession == nil {
		log.Warn("Failed to save session in DB.")
		return nil, []commonModel.Message{model.NewMessageAuthInternalError()}
	}
	log.Info("User", request.Name, "login.")
	return savedSession, nil
}

// GetSession Get session by token
func GetSession(token string) (*model.Session, []commonModel.Message) {
	dbInstance := db.GetDBInstance()
	session := dbInstance.GetSessionByToken(token)
	if session == nil {
		return nil, []commonModel.Message{model.NewMessageAuthNotFoundSession()}
	}
	return session, nil
}

// CreateDefaultAdmin Create the default admin account if it's not exist.
func CreateDefaultAdmin() error {
	dbInstance := db.GetDBInstance()
	account := dbInstance.GetAccountByName("admin")
	if account == nil {
		defaultAdmin := new(model.Account)
		defaultAdmin.Name = "admin"
		defaultAdmin.PasswordHash = "password_hash"
		dbInstance.PostAccount(defaultAdmin)
		log.Info("Default admin account is created.")
	}
	return nil
}

// CreateSession Create a Session.
func CreateSession(account *model.Account) *model.Session {
	session := new(model.Session)
	session.AccountID = account.ID
	session.Token = "new-session-token"
	return session
}
