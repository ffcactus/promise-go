package strategy

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	commonMessage "promise/common/object/message"
	"promise/server/context"
	"promise/server/object/message"
	"promise/server/object/model"
)

// ServerStrategy Common server strategy.
type ServerStrategy struct {
	ServerEventStrategy
}

// LockServer Lock the server.
func (s *ServerStrategy) LockServer(c *context.ServerContext, server *model.Server) error {
	success, lockedServer := c.DB.GetAndLockServer(server.ID)
	if lockedServer == nil {
		log.WithFields(log.Fields{"id": server.ID}).Info("Can not get and lock server, server not exist.")
		c.AppendMessage(commonMessage.NewResourceNotExist())
		return errors.New("failed to lock server, server not exist")
	}
	if !success {
		log.WithFields(log.Fields{"id": server.ID, "state": server.State}).Info("Can not get and lock server.")
		c.AppendMessage(message.NewServerLockFailed(server))
		return errors.New("failed to lock server. server can't be lock")
	}
	s.DispatchServerUpdate(c, server)
	return nil
}

// SetServerState Set server state.
func (s *ServerStrategy) SetServerState(c *context.ServerContext, server *model.Server, state string) error {
	if c.DB.SetServerState(server.ID, state) {
		s.DispatchServerUpdate(c, server)
		return nil
	}
	return fmt.Errorf("failed to set server %s to %s", server.ID, state)
}

// SetServerHealth Set server health.
func (s *ServerStrategy) SetServerHealth(c *context.ServerContext, server *model.Server, health string) error {
	if c.DB.SetServerHealth(server.ID, health) {
		s.DispatchServerUpdate(c, server)
		return nil
	}
	return fmt.Errorf("failed to set server %s to %s", server.ID, health)
}