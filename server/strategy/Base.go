package strategy

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/server/context"
	"promise/server/object/message"
	"promise/server/object/model"
)

// Base is the base server strategy.
type Base struct {
	ServerEvent
}

// LockServer Lock the server.
func (s *Base) LockServer(c *context.Base, server *model.Server) error {
	success, lockedServer := c.DB.GetAndLockServer(server.ID)
	if lockedServer == nil {
		log.WithFields(log.Fields{"id": server.ID}).Info("Can not get and lock server, server not exist.")
		c.AppendMessage(*base.NewMessageNotExist())
		return errors.New("failed to lock server, server not exist")
	}
	if !success {
		log.WithFields(log.Fields{"id": server.ID, "state": server.State}).Info("Can not get and lock server.")
		c.AppendMessage(*message.NewMessageServerLockFailed(server))
		return errors.New("failed to lock server. server can't be lock")
	}
	s.DispatchServerUpdate(c, server)
	return nil
}

// SetServerState Set server state.
func (s *Base) SetServerState(c *context.Base, server *model.Server, state string) error {
	if c.DB.SetServerState(server.ID, state) {
		s.DispatchServerUpdate(c, server)
		return nil
	}
	return fmt.Errorf("failed to set server %s to %s", server.ID, state)
}

// SetServerHealth Set server health.
func (s *Base) SetServerHealth(c *context.Base, server *model.Server, health string) error {
	if c.DB.SetServerHealth(server.ID, health) {
		s.DispatchServerUpdate(c, server)
		return nil
	}
	return fmt.Errorf("failed to set server %s to %s", server.ID, health)
}
