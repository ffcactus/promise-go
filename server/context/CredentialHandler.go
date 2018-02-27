package context

import (
	"promise/server/object/model"
)

// CredentialHandlerInterface The interface of credential handler.
type CredentialHandlerInterface interface {
	GetCredential(server model.Server) (username, password string)
}

// CredentialHandler The implementation of credential handler.
type CredentialHandler struct {
}

// GetCredential Get server credential.
func (h *CredentialHandler) GetCredential(server model.Server) (username, password string) {
	return "", ""
}
