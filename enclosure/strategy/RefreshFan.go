package strategy

import (
	log "github.com/sirupsen/logrus"
	"promise/enclosure/context"
)

// RefreshFan is the strategy to refresh fan.
type RefreshFan struct {
	name                string
	messageID           string
	description         string
	expectedExecutionMs uint64
}

// NewRefreshFan creates a new RefreshFan strategy.
func NewRefreshFan() *RefreshFan {
	return &RefreshFan{
		name:                "Refresh Fan",
		messageID:           "Promise.Enclosure.Action.Refresh.RefreshFan",
		description:         "Refresh enclosure fan components and their settings.",
		expectedExecutionMs: 5000,
	}
}

// Name returns the name of the strategy.
func (s *RefreshFan) Name() string {
	return s.name
}

// MessageID returns the message ID of the strategy.
func (s *RefreshFan) MessageID() string {
	return s.messageID
}

// Description returns the description of the strategy.
func (s *RefreshFan) Description() string {
	return s.description
}

// ExpectedExecutionMs returns the expected execution time in ms of the strategy.
func (s *RefreshFan) ExpectedExecutionMs() uint64 {
	return s.expectedExecutionMs
}

// Execute implements the Action interface.
func (s *RefreshFan) Execute(c *context.Base) {
	log.Info("Action refresh fan.")
	slots, clientError := c.Client.FanSlot()
	if clientError != nil {
		// TODO we need process the alarm here.
		log.WithFields(log.Fields{
			"id": c.ID, "error": clientError,
		}).Warn("Strategy refresh fan failed, get fan slots failed.")
	}
	enclosure, dbError := c.DB.RefreshFanSlot(c.ID, slots)
	if dbError != nil {
		log.WithFields(log.Fields{
			"id": c.ID, "error": clientError,
		}).Warn("Strategy refresh fan failed, DB refresh fan failed.")
	}
	c.Enclosure = enclosure
}
