package strategy

import (
	log "github.com/sirupsen/logrus"
	"promise/server/context"
	"promise/server/object/model"
)

// DiscoverMockServerStrategy is the strategy for mock server.
type DiscoverMockServerStrategy struct {
	DiscoverRackServerStrategy
}

// CreateManagementAccount Create the management account.
func (s *DiscoverMockServerStrategy) CreateManagementAccount(c *context.DiscoverServer, server *model.Server) error {
	log.WithFields(log.Fields{"address": c.Request().Hostname}).Info("Create management account.")
	return nil
}

// Claim the server.
func (s *DiscoverMockServerStrategy) Claim(c *context.DiscoverServer, server *model.Server) error {
	log.WithFields(log.Fields{"address": c.Request().Hostname}).Info("Claim server.")
	return nil
}
