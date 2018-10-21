package strategy

import (
	log "github.com/sirupsen/logrus"
	"promise/enclosure/context"
)

// RefreshPower is the strategy to refresh power.
type RefreshPower struct {
	name                string
	messageID           string
	description         string
	expectedExecutionMs uint64
}

// NewRefreshPower creates a new RefreshPower strategy.
func NewRefreshPower() *RefreshPower {
	return &RefreshPower{
		name:                "Refresh Power",
		messageID:           "Promise.Enclosure.Action.Refresh.RefreshPower",
		description:         "Refresh enclosure power components and their settings.",
		expectedExecutionMs: 5000,
	}
}

// Name returns the name of the strategy.
func (s *RefreshPower) Name() string {
	return s.name
}

// MessageID returns the message ID of the strategy.
func (s *RefreshPower) MessageID() string {
	return s.messageID
}

// Description returns the description of the strategy.
func (s *RefreshPower) Description() string {
	return s.description
}

// ExpectedExecutionMs returns the expected execution time in ms of the strategy.
func (s *RefreshPower) ExpectedExecutionMs() uint64 {
	return s.expectedExecutionMs
}

// Execute implements the Action interface.
func (s *RefreshPower) Execute(c *context.Base) {
	log.Info("Action refresh power.")
	StepStart(c, s.name)
	slots, clientError := c.Client.PowerSlot()
	if clientError != nil {
		// TODO we need process the alarm here.
		log.WithFields(log.Fields{
			"id": c.ID, "error": clientError,
		}).Warn("Strategy refresh power failed, get power slots failed.")
		StepError(c, s.name)
		return
	}
	enclosure, dbError := c.DB.RefreshPowerSlot(c.ID, slots)
	if dbError != nil {
		log.WithFields(log.Fields{
			"id": c.ID, "error": clientError,
		}).Warn("Strategy refresh power failed, DB refresh power failed.")
	}
	c.UpdateEnclosure(enclosure)
	StepFinish(c, s.name)
}
