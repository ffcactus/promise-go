package context

import (
	m "promise/server/object/model"
)

// RefreshServer Refresh server context.
type RefreshServer struct {
	Base
}

// CreateRefreshServerContext Create a new instance by server.
func CreateRefreshServerContext(server *m.Server) *RefreshServer {
	var context RefreshServer
	context.Base = *CreateServerContext(server)
	return &context
}
