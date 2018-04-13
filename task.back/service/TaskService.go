package service

import (
	commonMessage "promise/common/object/message"
	wsSDK "promise/sdk/ws"
	"promise/task/db"
	"promise/task/object/dto"
	"promise/task/object/message"
	"promise/task/object/model"
)

// PostTask Post Task.
func PostTask(request *dto.PostTaskRequest) (*model.Task, []commonMessage.Message) {
	var (
		dbImpl  = db.GetDBInstance()
		taskDTO dto.GetTaskResponse
	)

	posted, err := dbImpl.PostTask(request.ToModel())
	if err != nil {
		return nil, []commonMessage.Message{commonMessage.NewTransactionError()}
	}
	taskDTO.Load(posted)
	wsSDK.DispatchResourceCreateEvent(&taskDTO)
	return posted, nil
}

// GetTask Get Task.
func GetTask(id string) (*model.Task, []commonMessage.Message) {
	var (
		dbImpl = db.GetDBInstance()
	)

	task := dbImpl.GetTask(id)
	if task == nil {
		return nil, []commonMessage.Message{message.NewMessageTaskNotExist()}
	}
	return task, nil
}

// GetTaskCollection Get task collection.
func GetTaskCollection(start int64, count int64, filter string) (*model.TaskCollection, []commonMessage.Message) {
	var (
		dbImpl = db.GetDBInstance()
	)

	ret, err := dbImpl.GetTaskCollection(start, count, filter)
	if err != nil {
		return nil, []commonMessage.Message{message.NewMessageTaskInternalError()}
	}
	return ret, nil
}

// UpdateTask Update task.
func UpdateTask(id string, updateRequest *dto.UpdateTaskRequest) (*model.Task, []commonMessage.Message) {
	var (
		dbImpl = db.GetDBInstance()
	)

	// We need pass the DTO to DB implementation, the updated
	// should side in a transaction, you can't check the existence here and do the
	// update somewhere else.
	exist, updatedTask, commited, err := dbImpl.UpdateTask(id, updateRequest)
	if !exist {
		return nil, []commonMessage.Message{commonMessage.NewResourceNotExist()}
	}
	if err != nil || !commited {
		return nil, []commonMessage.Message{commonMessage.NewTransactionError()}
	}
	var taskDTO dto.GetTaskResponse
	taskDTO.Load(updatedTask)
	wsSDK.DispatchResourceUpdateEvent(&taskDTO)
	return updatedTask, nil
}

// UpdateTaskStep Update task step.
func UpdateTaskStep(id string, request *dto.UpdateTaskStepRequest) (*model.Task, []commonMessage.Message) {
	var (
		dbImpl = db.GetDBInstance()
	)

	// We need pass the DTO to DB implementation, the updated
	// should side in a transaction, you can't check the existence here and do the
	// update somewhere else.
	exist, updatedTask, commited, err := dbImpl.UpdateTaskStep(id, request)
	if !exist {
		return nil, []commonMessage.Message{commonMessage.NewResourceNotExist()}
	}
	if err != nil || !commited {
		return nil, []commonMessage.Message{commonMessage.NewTransactionError()}
	}
	var taskDTO dto.GetTaskResponse
	taskDTO.Load(updatedTask)
	wsSDK.DispatchResourceUpdateEvent(&taskDTO)
	return updatedTask, nil
}
