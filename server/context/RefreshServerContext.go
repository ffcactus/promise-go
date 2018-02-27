package context

import (
	m "promise/server/object/model"
)

// RefreshServerContext Refresh server context.
type RefreshServerContext struct {
	ServerContext
}

// CreateRefreshServerContext Create a new instance by server.
func CreateRefreshServerContext(server *m.Server) *RefreshServerContext {
	var context RefreshServerContext
	context.ServerContext = *CreateServerContext(server)
	return &context
}
