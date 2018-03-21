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
}

// LockServer Lock the server.
func (s *ServerStrategy) LockServer(c *context.ServerContext) error {
	success, server := c.ServerDBImplement.GetAndLockServer(c.Server.ID)
	if server == nil {
		log.WithFields(log.Fields{"id": c.Server.ID}).Info("Can not get and lock server, server not exist.")
		c.AppendMessage(commonMessage.NewResourceNotExist())
		return errors.New("failed to lock server, server not exist")
	}
	if !success {
		log.WithFields(log.Fields{"id": c.Server.ID, "state": server.State}).Info("Can not get and lock server")
		c.AppendMessage(message.NewServerLockFailed(server))
		return errors.New("failed to lock server. server can't be lock")
	}
	c.DispatchServerUpdate()
	return nil
}

// SetServerState Set server state.
func (s *ServerStrategy) SetServerState(c *context.ServerContext, state string) error {
	if c.ServerDBImplement.SetServerState(c.Server.ID, state) {
		c.DispatchServerUpdate()
		return nil
	}
	return fmt.Errorf("failed to set server %s to %s", c.Server.ID, state)
}

// SetServerHealth Set server health.
func (s *ServerStrategy) SetServerHealth(c *context.ServerContext, health string) error {
	if c.ServerDBImplement.SetServerHealth(c.Server.ID, health) {
		c.DispatchServerUpdate()
		return nil
	}
	return fmt.Errorf("failed to set server %s to %s", c.Server.ID, health)

}

// SaveServer Save the server to DB.
func (s *ServerStrategy) SaveServer(c *context.ServerContext) error {
	server, err := c.ServerDBImplement.PostServer(c.Server)
	if err != nil {
		log.WithFields(log.Fields{"hostname": c.Server.Hostname, "err": err}).Warn("Save server failed.")
		c.AppendMessage(commonMessage.NewInternalError())
		return errors.New("failed to save server")
	}
	c.Server = server
	return nil
}

// GetServerFull The server's all info.
func (s *ServerStrategy) GetServerFull(c *context.ServerContext) (*model.Server, error) {
	server := c.ServerDBImplement.GetServerFull(c.Server.ID)
	if server != nil {
		return server, nil
	}
	log.WithFields(log.Fields{"id": c.Server.ID}).Warn("Get full server failed.")
	return nil, errors.New("internel error")

}

// IndexServer Put the server into index service.
func (s *ServerStrategy) IndexServer(c *context.ServerContext) error {
	// if server, err := this.GetServerFull(context); err == nil {
	// 	serverDto := new(dto.GetServerResponse)
	// 	serverDto.Load(server)
	// 	if err := context.IndexServer(serverDto); err != nil {
	// 		log.Warn("IndexServer() failed, server id = ", server.ID, ", error = ", err)
	// 		context.AppendMessage(NewInternalError())
	// 		return errors.New("Failed to save server")
	// 	} else {
	// 		log.Debug("Index server done, server ID = ", context.Server.URI)
	// 		return nil
	// 	}
	// } else {
	// 	log.Warn("Index server failed, unable get server from DB, server ID = ", context.Server.ID)
	// 	return errors.New("Index server failed, unable get server from DB")
	// }
	return nil
}
