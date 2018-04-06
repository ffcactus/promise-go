package service

import (
	log "github.com/sirupsen/logrus"
	commonMessage "promise/common/object/message"
	wsSDK "promise/sdk/ws"
	"promise/task/db"
	"promise/task/object/dto"
	"promise/task/object/message"
	"promise/task/object/model"
)

// PostTask Post Task.
func PostTask(request *dto.PostTaskRequest) (*model.Task, []commonMessage.Message) {
	log.Debug("PostTask() start, task name =", request.Name)
	db := db.GetDBInstance()
	task := request.ToModel()
	createTask, err := db.PostTask(task)
	if err != nil {
		return nil, []commonMessage.Message{message.NewMessageTaskSaveFailure()}
	}
	wsSDK.DispatchTaskCreate(createTask)
	return createTask, nil
}

// GetTask Get Task.
func GetTask(id string) (*model.Task, []commonMessage.Message) {
	log.Debug("GetTask() start, task ID =", id)
	db := db.GetDBInstance()
	task := db.GetTask(id)
	if task == nil {
		return nil, []commonMessage.Message{message.NewMessageTaskNotExist()}
	}
	return task, nil
}

// GetTaskCollection Get task collection.
func GetTaskCollection(start int64, count int64, filter string) (*model.TaskCollection, []commonMessage.Message) {
	db := db.GetDBInstance()
	ret, err := db.GetTaskCollection(start, count, filter)
	if err != nil {
		return nil, []commonMessage.Message{message.NewMessageTaskInternalError()}
	}
	return ret, nil
}

// UpdateTask Update task.
func UpdateTask(id string, request *dto.UpdateTaskRequest) (*model.Task, []commonMessage.Message) {
	log.Debug("UpdateTask() start, task ID =", id)
	// TODO Check parameters.
	db := db.GetDBInstance()
	task := db.GetTask(id)
	if task == nil {
		return nil, []commonMessage.Message{message.NewMessageTaskNotExist()}
	}
	request.UpdateModel(task)
	savedTask, err := db.UpdateTask(id, task)
	if err != nil {
		return nil, []commonMessage.Message{message.NewMessageTaskSaveFailure()}
	}
	wsSDK.DispatchTaskUpdate(savedTask)
	return savedTask, nil
}

// UpdateTaskStep Update task step.
func UpdateTaskStep(id string, request *dto.UpdateTaskStepRequest) (*model.Task, []commonMessage.Message) {
	log.Debug("UpdateTaskStep() start, task ID =", id, "step name =", request.Name)
	db := db.GetDBInstance()
	task := db.GetTask(id)
	if task == nil {
		log.Debug("UpdateTaskStep() failed, GetTask() failed, task = nil, task ID =", id)
		return nil, []commonMessage.Message{message.NewMessageTaskNotExist()}
	}
	currentTime := uint64(0)
	for i := range task.TaskSteps {
		step := task.TaskSteps[i]
		currentTime += step.ExpectedExecutionMs
		if step.Name == request.Name {
			// Found the step, and update the current time.
			switch step.ExecutionState {
			case model.ExecutionStateTerminated:
			case model.ExecutionStateRunning:
			case model.ExecutionStateSuspended:
				currentTime -= step.ExpectedExecutionMs
			default:
			}
			if request.ExecutionState != nil && *request.ExecutionState == model.ExecutionStateRunning {
				task.CurrentStep = step.Name
				log.Debug("-------CurrentStep =", task.CurrentStep)
			}
			percentageF := (float32)(currentTime) / (float32)(task.ExpectedExecutionMs)
			task.Percentage = (int)((percentageF * 100) + 0.5)
			if task.Percentage > 100 {
				task.Percentage = 100
			}
			request.UpdateModel(task)
			savedTask, err := db.UpdateTask(id, task)
			if err != nil {
				log.Warn("UpdateTaskStep() failed, failed to save task, step name =", request.Name)
				return nil, []commonMessage.Message{message.NewMessageTaskSaveFailure()}
			}
			wsSDK.DispatchTaskUpdate(savedTask)
			return savedTask, nil
		}
	}
	log.Warn("UpdateTaskStep() failed, can't find the step, step name =", request.Name)
	return nil, []commonMessage.Message{message.NewMessageTaskNotExist()}
}
