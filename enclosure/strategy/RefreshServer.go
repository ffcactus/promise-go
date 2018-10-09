package strategy

import (
	log "github.com/sirupsen/logrus"
	"promise/enclosure/context"
)

// RefreshServer is the strategy to refresh manager.
type RefreshServer struct {
	name                string
	messageID           string
	description         string
	expectedExecutionMs uint64
}

// NewRefreshServer creates a new RefreshServer strategy.
func NewRefreshServer() *RefreshServer {
	return &RefreshServer{
		name:                "Refresh Server",
		messageID:           "Promise.Enclosure.Action.Refresh.RefreshServer",
		description:         "Refresh enclosure server components.",
		expectedExecutionMs: 60000,
	}
}

// Name returns the name of the strategy.
func (s *RefreshServer) Name() string {
	return s.name
}

// MessageID returns the message ID of the strategy.
func (s *RefreshServer) MessageID() string {
	return s.messageID
}

// Description returns the description of the strategy.
func (s *RefreshServer) Description() string {
	return s.description
}

// ExpectedExecutionMs returns the expected execution time in ms of the strategy.
func (s *RefreshServer) ExpectedExecutionMs() uint64 {
	return s.expectedExecutionMs
}

// Execute implements the Action interface.
func (s *RefreshServer) Execute(c *context.Base) {
	log.Info("Action refresh server.")
}
