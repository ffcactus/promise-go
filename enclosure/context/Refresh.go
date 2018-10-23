package context

import (
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/enclosure/object/dto"
	"promise/enclosure/object/model"
)

// Refresh is the context used in refresh strategy.
type Refresh interface {
	Base
	GetRequest() *dto.RefreshEnclosureRequest
	GetNextState() string
	GetNextStateReason() string
	DispatchUpdateEvent()
}

// RefreshImpl implements Refresh context.
type RefreshImpl struct {
	BaseImpl
	Request         *dto.RefreshEnclosureRequest
	NextState       string
	NextStateReason string
}

// NewRefresh creates a Refresh context.
func NewRefresh(id string, request *dto.RefreshEnclosureRequest) Refresh {
	ret := RefreshImpl{}
	ret.ID = id
	ret.Request = request
	ret.NextState = model.StateReady
	ret.NextStateReason = model.StateReasonAuto
	return &ret
}

// GetRequest returns request.
func (c *RefreshImpl) GetRequest() *dto.RefreshEnclosureRequest {
	return c.Request
}

// GetNextState returns next state.
func (c *RefreshImpl) GetNextState() string {
	return c.NextState
}

// GetNextStateReason returns next state reason.
func (c *RefreshImpl) GetNextStateReason() string {
	return c.NextStateReason
}

// DispatchUpdateEvent will send an update event using the enclosure in the context.
func (c *RefreshImpl) DispatchUpdateEvent() {
	response := dto.GetEnclosureResponse{}
	if err := response.Load(c.GetEnclosure()); err != nil {
		log.WithFields(log.Fields{
			"id":    c.GetID(),
			"error": err,
		}).Warn("Context dispatch update event failed, create event failed.")
		return
	}
	err := base.PublishResourceMessage(base.UpdateOperation, &response)
	if err != nil {
		log.WithFields(log.Fields{
			"id":    c.GetID(),
			"error": err,
		}).Warn("Context dispatch update event failed, event dispatching failed.")
	}
}
