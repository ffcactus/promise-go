package step

import (
	. "promise/server/context"
	taskModel "promise/task/object/model"
)

type StrategyStep struct {
	MessageId           *string
	Name                string
	Description         string
	ExpectedExecutionMs int
}

func (this *StrategyStep) ExecuteResultState() taskModel.ExecutionResultState {
	return taskModel.ExecutionStateReady
}

func (this *StrategyStep) Execute(context ServerContext) {

}

func (this *StrategyStep) TaskStepStart() {

}

func (this *StrategyStep) TaskStepFinish() {

}

func (this *StrategyStep) TaskStepError() {

}
