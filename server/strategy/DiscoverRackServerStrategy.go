package strategy

import (
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/server/context"
	"promise/server/object/constvalue"
	"promise/server/object/message"
	"promise/server/object/model"
)

// DiscoverRackServerStrategy is the strategy for post rack server.
type DiscoverRackServerStrategy struct {
	ServerStrategy
	ServerEventStrategy
}

// Get the management account.
func (s *DiscoverRackServerStrategy) getManagementAccount(c *context.DiscoverServer) (string, string) {
	// Should ask the auth service to provider the management account.
	return "Director", "Huawei12#$"
}

// CreateManagementAccount will create a management account.
func (s *DiscoverRackServerStrategy) CreateManagementAccount(c *context.DiscoverServer, server *model.Server) error {
	username, password := s.getManagementAccount(c)
	if err := c.ServerClient.CreateManagementAccount(username, password); err != nil {
		c.AppendMessage(message.NewMessageServerAccountExist(server))
		return err
	}
	// After the management account created on the server, update it in the context.
	// TODO
	credential := username + " " + password
	server.Credential = credential
	log.WithFields(log.Fields{"hostname": c.Request().Hostname}).Info("Create management account.")
	return nil
}

// Claim the server.
func (s *DiscoverRackServerStrategy) Claim(c *context.DiscoverServer, server *model.Server) error {
	log.WithFields(log.Fields{"hostname": c.Request().Hostname}).Info("Claim server.")
	return nil
}

// Execute will execute all the steps.
func (s *DiscoverRackServerStrategy) Execute(c *context.DiscoverServer, tempServer *model.Server) (*model.Server, error) {
	if err := s.CreateManagementAccount(c, tempServer); err != nil {
		return nil, err
	}
	if err := s.Claim(c, tempServer); err != nil {
		return nil, err
	}
	// Set the servers init state and health.
	tempServer.State = constvalue.ServerStateAdded
	tempServer.Health = constvalue.ServerHealthOK
	server, ssg, err := c.DB.Create(tempServer)
	if err != nil {
		c.AppendMessage(base.NewMessageTransactionError())
		return nil, err
	}
	tempServer = server
	// Dispatch event.
	s.DispatchServerCreate(&c.ServerContext, server)
	s.DispatchServerServerGroupCreate(&c.ServerContext, ssg)

	return server, nil
}
