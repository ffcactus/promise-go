package strategy

import (
	log "github.com/sirupsen/logrus"
	"promise/server/context"
)

// MockServerPostStrategy is the strategy for mock server.
type MockServerPostStrategy struct {
	RackServerPostStrategy
}

// CreateManagementAccount Create the management account.
func (s *MockServerPostStrategy) CreateManagementAccount(c *context.PostServerContext) error {
	log.WithFields(log.Fields{"address": c.Request.Address}).Info("Create management account.")
	return nil
}

// Claim the server.
func (s *MockServerPostStrategy) Claim(c *context.PostServerContext) error {
	log.WithFields(log.Fields{"address": c.Request.Address}).Info("Claim server.")
	return nil
}
