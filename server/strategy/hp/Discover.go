package hp

import (
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/server/context"
	"promise/server/object/constvalue"
	"promise/server/object/model"
	"promise/server/strategy/common"
)

// Discover strategy for HP servers.
type Discover struct {
	common.Common
}

// Get the management account.
func (s *Discover) getManagementAccount(c *context.DiscoverServer) (string, string) {
	// Should ask the auth service to provider the management account.
	return "Director", "Huawei12#$"
}

// CreateManagementAccount will create a management account.
func (s *Discover) CreateManagementAccount(c *context.DiscoverServer, server *model.Server) error {
	// For simplicity we just use the orginal credential.
	credential := *server.OriginUsername + " " + *server.OriginPassword
	server.Credential = credential
	log.WithFields(log.Fields{"hostname": c.Request().Hostname}).Info("Create management account.")
	return nil
}

// Claim the server.
func (s *Discover) Claim(c *context.DiscoverServer, server *model.Server) error {
	log.WithFields(log.Fields{"address": c.Request().Hostname}).Info("Claim server.")
	return nil
}

// Execute will execute all the steps.
func (s *Discover) Execute(c *context.DiscoverServer, tempServer *model.Server) (base.ModelInterface, error) {
	if err := s.CreateManagementAccount(c, tempServer); err != nil {
		return nil, err
	}
	if err := s.Claim(c, tempServer); err != nil {
		return nil, err
	}
	// Set the servers init state and health.
	tempServer.State = constvalue.ServerStateAdded
	tempServer.Health = constvalue.ServerHealthOK
	server, ssg, err := c.DB.CreateServer(tempServer)
	if err != nil {
		c.AppendErrorResponse(*base.NewErrorResponseTransactionError())
		return nil, err
	}
	// tempServer = server
	// Dispatch event.
	s.DispatchServerCreate(&c.Base, server)
	s.DispatchServerServerGroupCreate(&c.Base, ssg)

	return server, nil
}
