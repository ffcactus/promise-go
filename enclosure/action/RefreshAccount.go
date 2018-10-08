package action

import (
	log "github.com/sirupsen/logrus"
	"promise/enclosure/context"
)

// RefreshAccount is the action to refresh manager.
type RefreshAccount struct {
	name                string
	messageID           string
	description         string
	expectedExecutionMs uint64
}

// NewRefreshAccount creates a new RefreshAccount action.
func NewRefreshAccount() *RefreshAccount {
	return &RefreshAccount{
		name:                "Refresh Account",
		messageID:           "Promise.Enclosure.Action.Refresh.RefreshAccount",
		description:         "Refresh enclosure manager account.",
		expectedExecutionMs: 5000,
	}
}

// Name returns the name of the action.
func (s *RefreshAccount) Name() string {
	return s.name
}

// MessageID returns the message ID of the action.
func (s *RefreshAccount) MessageID() string {
	return s.messageID
}

// Description returns the description of the action.
func (s *RefreshAccount) Description() string {
	return s.description
}

// ExpectedExecutionMs returns the expected execution time in ms of the action.
func (s *RefreshAccount) ExpectedExecutionMs() uint64 {
	return s.expectedExecutionMs
}

// Execute implements the Action interface.
func (s *RefreshAccount) Execute(c *context.Base) {
	log.Info("Action refresh account.")
}
