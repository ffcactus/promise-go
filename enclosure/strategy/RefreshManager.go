package strategy

import (
	log "github.com/sirupsen/logrus"
	"promise/enclosure/context"
	taskModel "promise/task/object/model"
)

// RefreshManager is the strategy to refresh manager.
type RefreshManager struct {
	Task
	name                string
	messageID           string
	description         string
	expectedExecutionMs uint64
}

// NewRefreshManager creates a new RefreshManager strategy.
func NewRefreshManager() *RefreshManager {
	return &RefreshManager{
		name:                "Refresh Manager",
		messageID:           "Promise.Enclosure.Action.Refresh.RefreshManager",
		description:         "Refresh enclosure manager components and their settings.",
		expectedExecutionMs: 5000,
	}
}

// Name returns the name of the strategy.
func (s *RefreshManager) Name() string {
	return s.name
}

// MessageID returns the message ID of the strategy.
func (s *RefreshManager) MessageID() string {
	return s.messageID
}

// Description returns the description of the strategy.
func (s *RefreshManager) Description() string {
	return s.description
}

// ExpectedExecutionMs returns the expected execution time in ms of the strategy.
func (s *RefreshManager) ExpectedExecutionMs() uint64 {
	return s.expectedExecutionMs
}

// Execute implements the Action interface.
func (s *RefreshManager) Execute(c *context.Base) {
	log.Info("Action refresh manager.")
	s.StepStart(c, s.Name)
	slots, clientError := c.Client.ManagerSlot()
	if clientError != nil {
		// TODO we need process the alarm here.
		log.WithFields(log.Fields{
			"id": c.ID, "error": clientError,
		}).Warn("Strategy refresh manager failed, get manager slots failed.")
		s.UpdateStepExecutionState(c, s.Name, taskModel.ExecutionStateRunning)
		return
	}
	enclosure, dbError := c.DB.RefreshManagerSlot(c.ID, slots)
	if dbError != nil {
		log.WithFields(log.Fields{
			"id": c.ID, "error": clientError,
		}).Warn("Strategy refresh manager failed, DB refresh manager failed.")
	}
	c.Enclosure = enclosure
}
