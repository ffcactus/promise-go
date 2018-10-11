package context

import (
	beegoCtx "github.com/astaxie/beego/context"
	"promise/enclosure/object/dto"
	"promise/enclosure/object/model"
)

// RefreshContext is refresh context.
type RefreshContext struct {
	Base
	Request    *dto.RefreshEnclosureRequest
	NextState  string
	NextReason string
}

// NewRefreshContext creates a Refresh context.
func NewRefreshContext(ctx *beegoCtx.Context, id string, request *dto.RefreshEnclosureRequest) *RefreshContext {
	ret := RefreshContext{}
	ret.Context = ctx
	ret.ID = id
	ret.Request = request
	ret.NextState = model.StateReady
	ret.NextReason = model.StateReasonAuto
	return &ret
}
