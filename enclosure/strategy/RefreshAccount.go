package strategy

import (
	log "github.com/sirupsen/logrus"
	"promise/enclosure/context"
)

// RefreshAccount is the strategy to refresh manager.
type RefreshAccount struct {
	name                string
	messageID           string
	description         string
	expectedExecutionMs uint64
}

// NewRefreshAccount creates a new RefreshAccount strategy.
func NewRefreshAccount() *RefreshAccount {
	return &RefreshAccount{
		name:                "Refresh Account",
		messageID:           "Promise.Enclosure.Action.Refresh.RefreshAccount",
		description:         "Refresh enclosure manager account.",
		expectedExecutionMs: 5000,
	}
}

// Name returns the name of the strategy.
func (s *RefreshAccount) Name() string {
	return s.name
}

// MessageID returns the message ID of the strategy.
func (s *RefreshAccount) MessageID() string {
	return s.messageID
}

// Description returns the description of the strategy.
func (s *RefreshAccount) Description() string {
	return s.description
}

// ExpectedExecutionMs returns the expected execution time in ms of the strategy.
func (s *RefreshAccount) ExpectedExecutionMs() uint64 {
	return s.expectedExecutionMs
}

// Execute performs the operation of this strategy.
func (s *RefreshAccount) Execute(c context.Refresh) {
	log.Info("Strategy refresh account done.")
}
