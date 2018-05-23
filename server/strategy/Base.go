package strategy

import (
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/server/context"
	"promise/server/object/model"
)

// Base is the base server strategy.
type Base struct {
	ServerEvent
}

// LockServer Lock the server.
func (s *Base) LockServer(c *context.Base, server *model.Server) *base.Message {
	success, lockedServer := c.DB.GetAndLockServer(server.ID)
	if lockedServer == nil {
		log.WithFields(log.Fields{"id": server.ID}).Info("Can not get and lock server, server not exist.")
		return base.NewMessageNotExist()
	}
	if !success {
		log.WithFields(log.Fields{"id": server.ID, "state": server.State}).Info("Can not get and lock server.")
		return base.NewMessageErrorState()
	}
	s.DispatchServerUpdate(c, server)
	return nil
}

// SetServerState Set server state.
func (s *Base) SetServerState(c *context.Base, server *model.Server, state string) error {
	updatedServer, err := c.DB.SetServerState(server.ID, state)
	if err != nil {
		return base.ErrorTransaction
	}
	s.DispatchServerUpdate(c, updatedServer)
	return nil
}

// SetServerHealth Set server health.
func (s *Base) SetServerHealth(c *context.Base, server *model.Server, health string) error {
	updatedServer, err := c.DB.SetServerHealth(server.ID, health)
	if err != nil {
		return base.ErrorTransaction
	}
	s.DispatchServerUpdate(c, updatedServer)
	return nil
}