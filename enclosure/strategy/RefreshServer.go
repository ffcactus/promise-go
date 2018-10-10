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
	slots, clientError := c.Client.ServerSlot()
	if clientError != nil {
		// TODO we need process the alarm here.
		log.WithFields(log.Fields{
			"id": c.ID, "error": clientError,
		}).Warn("Strategy refresh server failed, get server slots failed.")
	}
	enclosure, dbError := c.DB.RefreshServerSlot(c.ID, slots)
	if dbError != nil {
		log.WithFields(log.Fields{
			"id": c.ID, "error": clientError,
		}).Warn("Strategy refresh server failed, DB refresh server failed.")
	}
	c.Enclosure = enclosure
}
