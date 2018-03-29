package strategy

import (
	log "github.com/sirupsen/logrus"
	"promise/server/context"
	"promise/server/object/constValue"
	commonMessage "promise/common/object/message"
	"promise/server/object/message"
	"promise/server/object/model"
)

// RackServerPostStrategy is the strategy for post rack server.
type RackServerPostStrategy struct {
	ServerStrategy
	ServerEventStrategy
}

// Get the management account.
func (s *RackServerPostStrategy) getManagementAccount(c *context.PostServerContext) (string, string) {
	// Should ask the auth service to provider the management account.
	return "Director", "Huawei12#$"
}

// CreateManagementAccount will create a management account.
func (s *RackServerPostStrategy) CreateManagementAccount(c *context.PostServerContext, server *model.Server) error {
	username, password := s.getManagementAccount(c)
	if err := c.ServerClient.CreateManagementAccount(username, password); err != nil {
		c.AppendMessage(message.NewServerAccountExist(server))
		return err
	}
	// After the management account created on the server, update it in the context.
	// TODO
	credential := username + " " + password
	server.Credential = credential
	log.WithFields(log.Fields{"hostname": c.Request.Hostname}).Info("Create management account.")
	return nil
}

// Claim the server.
func (s *RackServerPostStrategy) Claim(c *context.PostServerContext, server *model.Server) error {
	log.WithFields(log.Fields{"hostname": c.Request.Hostname}).Info("Claim server.")
	return nil
}

// Execute will execute all the steps.
func (s *RackServerPostStrategy) Execute(c *context.PostServerContext, tempServer *model.Server) error {
	if err := s.CreateManagementAccount(c, tempServer); err != nil {
		return err
	}
	if err := s.Claim(c, tempServer); err != nil {
		return err
	}
	// Set the servers init state and health.
	tempServer.State = constValue.ServerStateAdded
	tempServer.Health = constValue.ServerHealthOK
	server, ssg, err := c.DB.PostServer(tempServer)
	if err != nil {
		c.AppendMessage(commonMessage.NewTransactionError())
		return err
	}
	// Dispatch event.
	s.DispatchServerCreate(&c.ServerContext, server)
	s.DispatchServerServerGroupCreate(&c.ServerContext, ssg)
	
	return nil
}
