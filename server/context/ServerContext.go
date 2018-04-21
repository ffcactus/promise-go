package context

import (
	serverClient "promise/server/client"
	"promise/server/db"
	"promise/server/object/model"
)

// ServerContext Server context.
type ServerContext struct {
	ErrorHandler
	CredentialHandler
	ServerClient serverClient.ServerClientInterface
	DB           db.ServerDB
}

// CreateServerContext Create server context by server.
func CreateServerContext(server *model.Server) *ServerContext {
	var context ServerContext
	context.ServerClient = serverClient.GetServerClient(server)
	return &context
}
