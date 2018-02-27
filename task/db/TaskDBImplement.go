package db

import (
	commonDB "promise/common/db"
	"promise/common/util"
	"promise/task/object/entity"
	"promise/task/object/model"

	"github.com/google/uuid"
)

// TaskDBImplement Task DB implement.
type TaskDBImplement struct {
}

// GetDBInstance Get DB instance.
func GetDBInstance() TaskDBInterface {
	return &TaskDBImplement{}
}

// PostTask Post Task.
func (i *TaskDBImplement) PostTask(task *model.Task) (*model.Task, error) {
	c := commonDB.GetConnection()

	var e = createTaskEntity(task)
	// Generate the UUID.
	e.ID = uuid.New().String()
	c.Create(e)
	return createTaskModel(e), nil
}

// GetTask Get Task by ID.
func (i *TaskDBImplement) GetTask(ID string) *model.Task {
	c := commonDB.GetConnection()
	var task = new(entity.Task)
	c.Where("ID = ?", ID).First(task)
	if task.ID != ID {
		return nil
	}
	return createTaskModel(task)
}

// GetTaskCollection Get task collection
func (i *TaskDBImplement) GetTaskCollection(start int, count int) (*model.TaskCollection, error) {
	var (
		total int
		task  []entity.Task
		ret   model.TaskCollection
	)

	c := commonDB.GetConnection()
	c.Table("task").Count(&total)
	c.Order("Created_At asc").Limit(count).Offset(start).Select([]string{"Id", "Name", "Description", "Execution_State", "Percentage", "Execution_Result"}).Find(&task)
	ret.Start = start
	ret.Count = len(task)
	ret.Total = total
	for i := range task {
		m := new(model.TaskMember)
		m.URI = "/promise/v1/task/" + task[i].ID
		m.Name = task[i].Name
		m.Description = task[i].Description
		m.CreatedAt = task[i].CreatedAt
		m.UpdatedAt = task[i].UpdatedAt
		m.CreatedByName = task[i].CreatedByName
		m.CreatedByURI = task[i].CreatedByURI
		m.TargetName = task[i].TargetName
		m.TargetURI = task[i].TargetURI
		m.CurrentStep = task[i].CurrentStep
		m.ExecutionState = task[i].ExecutionState
		m.Percentage = task[i].Percentage
		util.StringToStruct(task[i].ExecutionResult, &m.ExecutionResult)

		ret.Members = append(ret.Members, *m)
	}
	return &ret, nil
}

// UpdateTask Update Task.
// Note: Assume the task exists already.
func (i *TaskDBImplement) UpdateTask(id string, task *model.Task) (*model.Task, error) {
	c := commonDB.GetConnection()
	var e = createTaskEntity(task)
	e.ID = id
	c.Save(e)
	return createTaskModel(e), nil
}

func createTaskEntity(m *model.Task) *entity.Task {
	e := new(entity.Task)
	e.ID = uuid.New().String()
	e.MessageID = m.MessageID
	e.Name = m.Name
	e.Description = m.Description
	e.ExecutionState = m.ExecutionState
	e.CreatedByName = m.CreatedByName
	e.CreatedByURI = m.CreatedByURI
	e.TargetName = m.TargetName
	e.TargetURI = m.TargetURI
	e.ExpectedExecutionMs = m.ExpectedExecutionMs
	e.Percentage = m.Percentage
	e.CreatedAt = m.CreatedAt
	e.UpdatedAt = m.UpdatedAt
	e.CurrentStep = m.CurrentStep
	e.TaskSteps = util.StructToString(m.TaskSteps)
	e.ExecutionResult = util.StructToString(m.ExecutionResult)
	return e
}

func createTaskModel(e *entity.Task) *model.Task {
	m := new(model.Task)
	m.ID = e.ID
	m.URI = "/promise/v1/task/" + e.ID
	m.MessageID = e.MessageID
	m.Name = e.Name
	m.Description = e.Description
	m.ExecutionState = e.ExecutionState
	m.CreatedByName = e.CreatedByName
	m.CreatedByURI = e.CreatedByURI
	m.TargetName = e.TargetName
	m.TargetURI = e.TargetURI
	m.ExpectedExecutionMs = e.ExpectedExecutionMs
	m.Percentage = e.Percentage
	m.CreatedAt = e.CreatedAt
	m.UpdatedAt = e.UpdatedAt
	m.CurrentStep = e.CurrentStep
	util.StringToStruct(e.ExecutionResult, &m.ExecutionResult)
	util.StringToStruct(e.TaskSteps, &m.TaskSteps)
	for i := range e.SubTasks {
		m.SubTasks = append(m.SubTasks, *createTaskModel(&e.SubTasks[i]))
	}
	return m
}
