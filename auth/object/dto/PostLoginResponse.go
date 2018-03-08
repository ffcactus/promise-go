package dto

import (
	"promise/auth/object/model"
)

// PostLoginResponse The login response.
type PostLoginResponse struct {
	Token string `json:"Token"`
}

// Load Load from model.
func (dto *PostLoginResponse) Load(session *model.Session) {
	dto.Token = session.Token
}
