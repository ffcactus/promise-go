package service

import (
	commonModel "promise/common/object/model"
	wsSDK "promise/sdk/ws"
	"promise/task/db"
	"promise/task/object/dto"
	"promise/task/object/model"

	"github.com/astaxie/beego"
)

// PostTask Post Task.
func PostTask(request *dto.PostTaskRequest) (*model.Task, []commonModel.Message) {
	beego.Trace("PostTask() start, task name =", request.Name)
	db := db.GetDBInstance()
	task := request.ToModel()
	createTask, err := db.PostTask(task)
	if err != nil {
		return nil, []commonModel.Message{model.NewMessageTaskSaveFailure()}
	}
	wsSDK.DispatchTaskCreate(createTask)
	return createTask, nil
}

// GetTask Get Task.
func GetTask(id string) (*model.Task, []commonModel.Message) {
	beego.Trace("GetTask() start, task ID =", id)
	db := db.GetDBInstance()
	task := db.GetTask(id)
	if task == nil {
		return nil, []commonModel.Message{model.NewMessageTaskNotExist()}
	}
	return task, nil
}

// GetTaskCollection Get task collection.
func GetTaskCollection(start int, count int) (*model.TaskCollection, []commonModel.Message) {
	beego.Trace("GetTaskCollection() start, start =", start, "count =", count)
	db := db.GetDBInstance()
	ret, err := db.GetTaskCollection(start, count)
	if err != nil {
		return nil, []commonModel.Message{model.NewMessageTaskInternalError()}
	}
	return ret, nil
}

// UpdateTask Update task.
func UpdateTask(id string, request *dto.UpdateTaskRequest) (*model.Task, []commonModel.Message) {
	beego.Trace("UpdateTask() start, task ID =", id)
	// TODO Check parameters.
	db := db.GetDBInstance()
	task := db.GetTask(id)
	if task == nil {
		return nil, []commonModel.Message{model.NewMessageTaskNotExist()}
	}
	request.UpdateModel(task)
	savedTask, err := db.UpdateTask(id, task)
	if err != nil {
		return nil, []commonModel.Message{model.NewMessageTaskSaveFailure()}
	}
	wsSDK.DispatchTaskUpdate(savedTask)
	return savedTask, nil
}

// UpdateTaskStep Update task step.
func UpdateTaskStep(id string, request *dto.UpdateTaskStepRequest) (*model.Task, []commonModel.Message) {
	beego.Trace("UpdateTaskStep() start, task ID =", id, "step name =", request.Name)
	db := db.GetDBInstance()
	task := db.GetTask(id)
	if task == nil {
		beego.Trace("UpdateTaskStep() failed, GetTask() failed, task = nil, task ID =", id)
		return nil, []commonModel.Message{model.NewMessageTaskNotExist()}
	}
	currentTime := 0
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
				beego.Trace("-------CurrentStep =", task.CurrentStep)
			}
			percentageF := (float32)(currentTime) / (float32)(task.ExpectedExecutionMs)
			task.Percentage = (int)((percentageF * 100) + 0.5)
			if task.Percentage > 100 {
				task.Percentage = 100
			}
			request.UpdateModel(task)
			savedTask, err := db.UpdateTask(id, task)
			if err != nil {
				beego.Warning("UpdateTaskStep() failed, failed to save task, step name =", request.Name)
				return nil, []commonModel.Message{model.NewMessageTaskSaveFailure()}
			}
			wsSDK.DispatchTaskUpdate(savedTask)
			return savedTask, nil
		}
	}
	beego.Warning("UpdateTaskStep() failed, can't find the step, step name =", request.Name)
	return nil, []commonModel.Message{model.NewMessageTaskNotExist()}
}
