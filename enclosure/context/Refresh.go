package context

import (
	"promise/enclosure/object/dto"
)

// RefreshContext is refresh context.
type RefreshContext struct {
	Base
	Request *dto.RefreshEnclosureRequest
}

// NewRefreshContext creates a Refresh context.
func NewRefreshContext(ctx *beegoCtx.Context, id string, request *dto.RefreshEnclosureRequest) *RefreshContext {
	ret := RefreshContext{}
	ret.Context = ctx
	ret.ID = id
	ret.Request = request
	return &ret
}