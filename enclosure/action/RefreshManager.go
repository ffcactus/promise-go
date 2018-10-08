package action

import (
	log "github.com/sirupsen/logrus"
	"promise/enclosure/context"
)

// RefreshManager is the action to refresh manager.
type RefreshManager struct {
	name                string
	messageID           string
	description         string
	expectedExecutionMs uint64
}

// NewRefreshManager creates a new RefreshManager action.
func NewRefreshManager() *RefreshManager {
	return &RefreshManager{
		name:                "Refresh Manager",
		messageID:           "Promise.Enclosure.Action.Refresh.RefreshManager",
		description:         "Refresh enclosure manager components and their settings.",
		expectedExecutionMs: 5000,
	}
}

// Name returns the name of the action.
func (s *RefreshManager) Name() string {
	return s.name
}

// MessageID returns the message ID of the action.
func (s *RefreshManager) MessageID() string {
	return s.messageID
}

// Description returns the description of the action.
func (s *RefreshManager) Description() string {
	return s.description
}

// ExpectedExecutionMs returns the expected execution time in ms of the action.
func (s *RefreshManager) ExpectedExecutionMs() uint64 {
	return s.expectedExecutionMs
}

// Execute implements the Action interface.
func (s *RefreshManager) Execute(c *context.Base) {
	log.Info("Action refresh manager.")
}
