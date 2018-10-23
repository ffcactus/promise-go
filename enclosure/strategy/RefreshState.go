package strategy

import (
	log "github.com/sirupsen/logrus"
	"promise/enclosure/context"
)

// RefreshState is the strategy to refresh the state of the enclosure after the refresh is done.
type RefreshState struct {
	name                string
	messageID           string
	description         string
	expectedExecutionMs uint64
}

// NewRefreshState creates a new RefreshState strategy.
func NewRefreshState() *RefreshState {
	return &RefreshState{
		name:                "Refresh State",
		messageID:           "Promise.Enclosure.Action.Refresh.RefreshState",
		description:         "Refresh enclosure state.",
		expectedExecutionMs: 1000,
	}
}

// Name returns the name of the strategy.
func (s *RefreshState) Name() string {
	return s.name
}

// MessageID returns the message ID of the strategy.
func (s *RefreshState) MessageID() string {
	return s.messageID
}

// Description returns the description of the strategy.
func (s *RefreshState) Description() string {
	return s.description
}

// ExpectedExecutionMs returns the expected execution time in ms of the strategy.
func (s *RefreshState) ExpectedExecutionMs() uint64 {
	return s.expectedExecutionMs
}

// Execute performs the operation of this strategy.
func (s *RefreshState) Execute(c context.Refresh) {
	StepStart(c, s.name)
	state := c.GetNextState()
	reason := c.GetNextStateReason()
	enclosure, err := c.GetDB().SetState(c.GetID(), state, reason)
	if err != nil {
		log.WithFields(log.Fields{
			"id":     c.GetID(),
			"state":  state,
			"reason": reason,
			"error":  err,
		}).Error("Strategy refresh state failed, update state in DB failed.")
		StepError(c, s.name)
		return
	}
	c.UpdateEnclosure(enclosure)
	StepFinish(c, s.name)
	log.WithFields(log.Fields{
		"id":     c.GetID(),
		"state":  state,
		"reason": reason,
	}).Info("Strategy refresh state done.")
}
