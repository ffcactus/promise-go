package context

import (
	"promise/server/object/dto"
	"promise/server/object/model"
)

// PostServerContext The context for post server.
type PostServerContext struct {
	ServerContext
	Request *dto.PostServerRequest // POST server request
}

// CreatePostServerContext Create the context based on server and request.
func CreatePostServerContext(server *model.Server, request *dto.PostServerRequest) *PostServerContext {
	var context PostServerContext
	context.ServerContext = *CreateServerContext(server)
	context.Request = request
	return &context
}

// GetRequest Get request from the context.
func (c *PostServerContext) GetRequest() interface{} {
	return c.Request
}

// GetCredential Get the server credential.
// For POST server, the credential is provided directly in the request.
func (c *PostServerContext) GetCredential() (username, password string) {
	return c.Request.Username, c.Request.Password
}
