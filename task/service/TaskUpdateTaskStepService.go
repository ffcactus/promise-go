package service

import (
	"promise/base"
	"promise/task/object/dto"
)

// TaskUpdateTaskStepService is the impplement of ActionService for update task step.
type TaskUpdateTaskStepService struct {
}

// Perform the update task step action.
func (s *TaskUpdateTaskStepService) Perform(id string, request base.ActionRequestInterface) (interface{}, []base.Message) {
	var response dto.GetTaskResponse

	exist, updatedTask, commited, err := taskDB.UpdateTaskStep(id, request)
	if !exist {
		return nil, []base.Message{base.NewMessageNotExist()}
	}
	if err != nil && err.Error() == base.ErrorUnknownPropertyValue.Error() {
		return nil, []base.Message{base.NewMessageUnknownPropertyValue()}
	}
	if err != nil || !commited {
		return nil, []base.Message{base.NewMessageTransactionError()}
	}
	response.Load(updatedTask)
	eventService.DispatchUpdateEvent(&response)
	return &response, nil
}