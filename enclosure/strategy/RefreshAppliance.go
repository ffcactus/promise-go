package strategy

import (
	log "github.com/sirupsen/logrus"
	"promise/enclosure/context"
)

// RefreshAppliance is the strategy to refresh appliance.
type RefreshAppliance struct {
	name                string
	messageID           string
	description         string
	expectedExecutionMs uint64
}

// NewRefreshAppliance creates a new RefreshAppliance strategy.
func NewRefreshAppliance() *RefreshAppliance {
	return &RefreshAppliance{
		name:                "Refresh Appliance",
		messageID:           "Promise.Enclosure.Action.Refresh.RefreshAppliance",
		description:         "Refresh enclosure appliance components and their settings.",
		expectedExecutionMs: 5000,
	}
}

// Name returns the name of the strategy.
func (s *RefreshAppliance) Name() string {
	return s.name
}

// MessageID returns the message ID of the strategy.
func (s *RefreshAppliance) MessageID() string {
	return s.messageID
}

// Description returns the description of the strategy.
func (s *RefreshAppliance) Description() string {
	return s.description
}

// ExpectedExecutionMs returns the expected execution time in ms of the strategy.
func (s *RefreshAppliance) ExpectedExecutionMs() uint64 {
	return s.expectedExecutionMs
}

// Execute performs the operation of this strategy.
func (s *RefreshAppliance) Execute(c context.Refresh) {
	log.Info("Strategy refresh appliance.")
	StepStart(c, s.name)
	slots, clientError := c.GetClient().ApplianceSlot()
	if clientError != nil {
		// TODO we need process the alarm here.
		log.WithFields(log.Fields{
			"id": c.GetID(), "error": clientError,
		}).Warn("Strategy refresh appliance failed, get appliance slots failed.")
		StepError(c, s.name)
		return
	}
	enclosure, dbError := c.GetDB().RefreshApplianceSlot(c.GetID(), slots)
	if dbError != nil {
		log.WithFields(log.Fields{
			"id": c.GetID(), "error": clientError,
		}).Warn("Strategy refresh appliance failed, DB refresh appliance failed.")
	}
	c.UpdateEnclosure(enclosure)
	StepFinish(c, s.name)
}
