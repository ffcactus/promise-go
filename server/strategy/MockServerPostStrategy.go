package strategy

import (
	"github.com/astaxie/beego"
	"promise/server/context"
)

// MockServerPostStrategy is the strategy for mock server.
type MockServerPostStrategy struct {
	RackServerPostStrategy
}

// CreateManagementAccount Create the management account.
func (s *MockServerPostStrategy) CreateManagementAccount(c *context.PostServerContext) error {
	beego.Trace("Server management account created, server address =", c.Request.Address)
	return nil
}

// Claim the server.
func (s *MockServerPostStrategy) Claim(c *context.PostServerContext) error {
	beego.Trace("Server claimed, server address =", c.Request.Address)
	return nil
}
