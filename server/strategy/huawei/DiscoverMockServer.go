package huawei

import (
	log "github.com/sirupsen/logrus"
	"promise/server/context"
	"promise/server/object/model"
)

// DiscoverMockServer is the strategy for mock server.
type DiscoverMockServer struct {
	DiscoverRackServer
}

// CreateManagementAccount Create the management account.
func (s *DiscoverMockServer) CreateManagementAccount(c *context.DiscoverServer, server *model.Server) error {
	log.WithFields(log.Fields{"address": c.Request().Hostname}).Info("Create management account.")
	return nil
}

// Claim the server.
func (s *DiscoverMockServer) Claim(c *context.DiscoverServer, server *model.Server) error {
	log.WithFields(log.Fields{"address": c.Request().Hostname}).Info("Claim server.")
	return nil
}
