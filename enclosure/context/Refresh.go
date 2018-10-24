package context

import (
	"promise/enclosure/object/dto"
	"promise/enclosure/object/model"
)

// Refresh is the context used in refresh strategy.
type Refresh interface {
	Base
	GetRequest() *dto.RefreshEnclosureRequest
	GetNextState() string
	GetNextStateReason() string
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
