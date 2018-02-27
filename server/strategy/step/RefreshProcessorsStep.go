package step

import (
	. "promise/server/context"

	"github.com/astaxie/beego"
)

type RefeshProcessorsStep struct {
	StrategyStep
}

func (this *RefeshProcessorsStep) Execute(context *RefreshServerContext) {
	this.TaskStepStart()
	systemPageURI := *context.Server.OriginURIs.System
	processors, err := context.ServerClient.GetProcessors(systemPageURI)
	if err != nil {
		this.TaskStepError()
		beego.Warning("GetProcessors() failed, error = ", err)
	}
	if err := context.ServerDb.UpdateProcessors(context.Server.Id, processors); err != nil {
		this.TaskStepError()
		beego.Warning("UpdateProcessors() failed, error = ", err)
	}
	context.DispatchServerUpdate()
	beego.Trace("RefreshProcessors() done, server ID = ", context.Server.Id)
	this.TaskStepFinish()
}
