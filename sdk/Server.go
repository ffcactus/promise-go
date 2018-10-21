package sdk

import (
	"promise/base"
)

// Server represents the server SDK.
type Server struct {
}

// Discover perform server discover operation.
func (o *Server) Discover(name, description, address, username, password string) base.ClientError {

}
