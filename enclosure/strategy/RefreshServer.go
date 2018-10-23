package strategy

import (
	log "github.com/sirupsen/logrus"
	"promise/enclosure/context"
	"promise/enclosure/object/model"
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

// Execute performs the operation of this strategy.
func (s *RefreshServer) Execute(c context.Refresh) {
	StepStart(c, s.name)
	slots, clientError := c.GetClient().ServerSlot()
	if clientError != nil {
		// TODO we need process the alarm here.
		log.WithFields(log.Fields{
			"id": c.GetID(), "error": clientError,
		}).Warn("Strategy refresh server failed, get server slots failed.")
		StepError(c, s.name)
		return
	}
	for _, serverInModel := range c.GetEnclosure().ServerSlots {
		// If enclosure has no server URL, we check if the server exist in client's value.
		if serverInModel.ServerURL == "" {
			for _, slot := range slots {
				if slot.Index == serverInModel.Index && slot.Inserted {

				}
			}
		}
	}

	enclosure, dbError := c.GetDB().RefreshServerSlot(c.GetID(), slots)
	if dbError != nil {
		log.WithFields(log.Fields{
			"id": c.GetID(), "error": clientError,
		}).Warn("Strategy refresh server failed, DB refresh server failed.")
	}
	c.UpdateEnclosure(enclosure)
	StepFinish(c, s.name)
	log.Info("Strategy refresh server done.")
}

// prepareServer will get the server's address, username and password for adding.
func (s *RefreshServer) prepareServer(server *model.ServerSlot) (address, username, password string) {
	return "mock", "username", "password"
}

func (s *RefreshServer) addServer(server *model.ServerSlot) {

}

func (s *RefreshServer) removeServer(server *model.ServerSlot) {

}
