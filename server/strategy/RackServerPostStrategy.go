package strategy

import (
	log "github.com/sirupsen/logrus"
	wsSDK "promise/sdk/ws"
	"promise/server/context"
	"promise/server/object/constvalue"
	"promise/server/object/message"
)

// RackServerPostStrategy is the strategy for post rack server.
type RackServerPostStrategy struct {
	ServerStrategy
}

// Get the management account.
func (s *RackServerPostStrategy) getManagementAccount(c *context.PostServerContext) (string, string) {
	// Should ask the auth service to provider the management account.
	return "Director", "Huawei12#$"
}

// CreateManagementAccount will create a management account.
func (s *RackServerPostStrategy) CreateManagementAccount(c *context.PostServerContext) error {
	username, password := s.getManagementAccount(c)
	if err := c.ServerClient.CreateManagementAccount(username, password); err != nil {
		c.AppendMessage(message.NewServerAccountExist(c.Server))
		return err
	}
	// After the management account created on the server, update it in the context.
	// TODO
	credential := username + " " + password
	c.Server.Credential = credential
	log.WithFields(log.Fields{"hostname": c.Request.Hostname}).Info("Create management account.")
	return nil
}

// Claim the server.
func (s *RackServerPostStrategy) Claim(c *context.PostServerContext) error {
	log.WithFields(log.Fields{"hostname": c.Request.Hostname}).Info("Claim server.")
	return nil
}

// Execute will execute all the steps.
func (s *RackServerPostStrategy) Execute(c *context.PostServerContext) error {
	if err := s.CreateManagementAccount(c); err != nil {
		return err
	}
	if err := s.Claim(c); err != nil {
		return err
	}
	// Set the servers init state and health.
	c.Server.State = constvalue.ServerStateAdded
	c.Server.Health = constvalue.ServerHealthOK
	if err := s.SaveServer(&c.ServerContext); err != nil {
		return err
	}
	// At this point, the server in the context is what we can use..
	wsSDK.DispatchServerCreate(c.Server)
	return nil
}
