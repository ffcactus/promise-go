package context

import (
	beegoCtx "github.com/astaxie/beego/context"
)

// Base containes the basic context info
type Base struct {
	*beegoCtx.Context
	ID      string
}

// NewRefreshContext creates a Refresh context.
func NewRefreshContext(ctx *beegoCtx.Context, id string) *Base {
	ret := Base{}
	Base.Context = ctx
	ID = id
}