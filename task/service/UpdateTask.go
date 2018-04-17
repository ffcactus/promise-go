package service

import (
	"promise/base"
	"promise/task/object/dto"
)

// UpdateTask is the service implement to update task.
type UpdateTask struct {
}

// Perform the update task action.
func (s *UpdateTask) Perform(id string, request base.ActionRequestInterface) (interface{}, []base.Message) {
	var response dto.GetTaskResponse

	exist, updatedTask, commited, err := taskDB.Update(id, request)
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
