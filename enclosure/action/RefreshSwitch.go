package action

import (
	log "github.com/sirupsen/logrus"
	"promise/enclosure/context"
)

// RefreshSwitch is the action to refresh manager.
type RefreshSwitch struct {
	name                string
	messageID           string
	description         string
	expectedExecutionMs uint64
}

// NewRefreshSwitch creates a new RefreshSwitch action.
func NewRefreshSwitch() *RefreshSwitch {
	return &RefreshSwitch{
		name:                "Refresh Switch",
		messageID:           "Promise.Enclosure.Action.Refresh.RefreshSwitch",
		description:         "Refresh enclosure server components.",
		expectedExecutionMs: 60000,
	}
}

// Name returns the name of the action.
func (s *RefreshSwitch) Name() string {
	return s.name
}

// MessageID returns the message ID of the action.
func (s *RefreshSwitch) MessageID() string {
	return s.messageID
}

// Description returns the description of the action.
func (s *RefreshSwitch) Description() string {
	return s.description
}

// ExpectedExecutionMs returns the expected execution time in ms of the action.
func (s *RefreshSwitch) ExpectedExecutionMs() uint64 {
	return s.expectedExecutionMs
}

// Execute implements the Action interface.
func (s *RefreshSwitch) Execute(c *context.Base) {
	log.Info("Action refresh switch.")
}
