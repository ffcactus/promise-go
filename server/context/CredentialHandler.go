package context

import (
	"promise/server/object/model"
	"strings"
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
	splite := strings.Split(server.Credential, " ")
	return splite[0], splite[1]
}
