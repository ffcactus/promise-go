package context

import (
	beegoCtx "github.com/astaxie/beego/context"
	"promise/enclosure/object/dto"
	"promise/enclosure/object/model"
)

// RefreshContext is refresh context.
type RefreshContext struct {
	Base
	Request   *dto.RefreshEnclosureRequest
	Enclosure *model.Enclosure
}

// NewRefreshContext creates a Refresh context.
func NewRefreshContext(ctx *beegoCtx.Context, id string, request *dto.RefreshEnclosureRequest) *RefreshContext {
	ret := RefreshContext{}
	ret.Context = ctx
	ret.ID = id
	ret.Request = request
	return &ret
}
