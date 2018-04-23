package context

import (
	"promise/server/object/dto"
	"promise/server/object/model"
)

// DiscoverServer The context for post server.
type DiscoverServer struct {
	Base
	DiscoverRequest *dto.DiscoverServerRequest
}

// CreateDiscoverServerContext Create the context based on server and request.
func CreateDiscoverServerContext(server *model.Server, request *dto.DiscoverServerRequest) *DiscoverServer {
	var context DiscoverServer
	context.Base = *CreateServerContext(server)
	context.DiscoverRequest = request
	return &context
}

// Request Get request from the context.
func (c *DiscoverServer) Request() *dto.DiscoverServerRequest {
	return c.DiscoverRequest
}

// GetCredential Get the server credential.
// For POST server, the credential is provided directly in the request.
func (c *DiscoverServer) GetCredential() (username, password string) {
	return c.DiscoverRequest.Username, c.DiscoverRequest.Password
}
