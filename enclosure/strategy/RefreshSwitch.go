package strategy

import (
	log "github.com/sirupsen/logrus"
	"promise/enclosure/context"
)

// RefreshSwitch is the strategy to refresh manager.
type RefreshSwitch struct {
	name                string
	messageID           string
	description         string
	expectedExecutionMs uint64
}

// NewRefreshSwitch creates a new RefreshSwitch strategy.
func NewRefreshSwitch() *RefreshSwitch {
	return &RefreshSwitch{
		name:                "Refresh Switch",
		messageID:           "Promise.Enclosure.Action.Refresh.RefreshSwitch",
		description:         "Refresh enclosure server components.",
		expectedExecutionMs: 60000,
	}
}

// Name returns the name of the strategy.
func (s *RefreshSwitch) Name() string {
	return s.name
}

// MessageID returns the message ID of the strategy.
func (s *RefreshSwitch) MessageID() string {
	return s.messageID
}

// Description returns the description of the strategy.
func (s *RefreshSwitch) Description() string {
	return s.description
}

// ExpectedExecutionMs returns the expected execution time in ms of the strategy.
func (s *RefreshSwitch) ExpectedExecutionMs() uint64 {
	return s.expectedExecutionMs
}

// Execute performs the operation of this strategy.
func (s *RefreshSwitch) Execute(c context.Refresh) {
	StepStart(c, s.name)
	slots, clientError := c.GetClient().SwitchSlot()
	if clientError != nil {
		// TODO we need process the alarm here.
		log.WithFields(log.Fields{
			"id": c.GetID(), "error": clientError,
		}).Warn("Strategy refresh switch failed, get switch slots failed.")
		StepError(c, s.name)
		return
	}
	enclosure, dbError := c.GetDB().RefreshSwitchSlot(c.GetID(), slots)
	if dbError != nil {
		log.WithFields(log.Fields{
			"id": c.GetID(), "error": clientError,
		}).Warn("Strategy refresh switch failed, DB refresh switch failed.")
	}
	c.UpdateEnclosure(enclosure)
	c.DispatchUpdateEvent()
	StepFinish(c, s.name)
	log.Info("Strategy refresh switch done.")
}
