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
	log.Debug("Server management account created, server address =", c.Request.Address)
	return nil
}

// Claim the server.
func (s *MockServerPostStrategy) Claim(c *context.PostServerContext) error {
	log.Debug("Server claimed, server address =", c.Request.Address)
	return nil
}
