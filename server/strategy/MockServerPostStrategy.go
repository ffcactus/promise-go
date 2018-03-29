package strategy

import (
	log "github.com/sirupsen/logrus"
	"promise/server/context"
	"promise/server/object/model"
)

// MockServerPostStrategy is the strategy for mock server.
type MockServerPostStrategy struct {
	RackServerPostStrategy
}

// CreateManagementAccount Create the management account.
func (s *MockServerPostStrategy) CreateManagementAccount(c *context.PostServerContext, server *model.Server) error {
	log.WithFields(log.Fields{"address": c.Request.Hostname}).Info("Create management account.")
	return nil
}

// Claim the server.
func (s *MockServerPostStrategy) Claim(c *context.PostServerContext, server *model.Server) error {
	log.WithFields(log.Fields{"address": c.Request.Hostname}).Info("Claim server.")
	return nil
}
