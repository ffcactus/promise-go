package db

import (
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	commonDB "promise/common/db"
	commonConstError "promise/common/object/consterror"
	"promise/common/util"
	"promise/task/object/dto"
	"promise/task/object/entity"
	"promise/task/object/model"
	"strings"
)

var (
	instance TaskDBImplement
)

// TaskDBImplement Task DB implement.
type TaskDBImplement struct {
}

// GetDBInstance Get DB instance.
func GetDBInstance() TaskDBInterface {
	return &instance
}

// PostTask Post Task.
func (impl *TaskDBImplement) PostTask(m *model.Task) (*model.Task, error) {
	var (
		record entity.Task
	)

	c := commonDB.GetConnection()
	record.Load(m)
	// Generate the UUID.
	record.ID = uuid.New().String()
	if err := c.Create(&record).Error; err != nil {
		log.WithFields(log.Fields{
			"name":  m.Name,
			"error": err}).
			Warn("Post task in DB failed, create resource failed.")
		return nil, err
	}
	return record.ToModel(), nil
}

// GetTask Get Task by ID.
func (impl *TaskDBImplement) GetTask(id string) *model.Task {
	var (
		record entity.Task
	)
	c := commonDB.GetConnection()

	if c.Where("\"ID\" = ?", id).First(&record).RecordNotFound() {
		log.WithFields(log.Fields{
			"id": id}).
			Warn("Get task in DB failed, resource does not exist.")
		return nil
	}
	return record.ToModel()
}

func (impl *TaskDBImplement) convertFilter(filter string) (string, error) {
	if filter == "" {
		return "", nil
	}
	cmds := strings.Split(filter, " ")
	if len(cmds) != 3 {
		return "", commonConstError.ErrorConvertFilter
	}
	switch strings.ToLower(cmds[1]) {
	case "eq":
		return "\"" + cmds[0] + "\"" + " = " + cmds[2], nil
	default:
		return "", commonConstError.ErrorConvertFilter
	}
}

// GetTaskCollection Get task collection
func (impl *TaskDBImplement) GetTaskCollection(start int64, count int64, filter string) (*model.TaskCollection, error) {
	var (
		total      int64
		selection  = []string{"\"ID\"", "\"Name\"", "\"Description\"", "\"ExecutionState\"", "\"Percentage\"", "\"ExecutionResult\""}
		collection []entity.Task
		ret        model.TaskCollection
	)

	c := commonDB.GetConnection()
	c.Table("Task").Count(&total)
	if where, err := impl.convertFilter(filter); err != nil {
		log.WithFields(log.Fields{
			"filter": filter,
			"error":  err}).
			Warn("Get task collection in DB failed, convert filter failed.")
		c.Order("\"CreatedAt\" asc").Limit(count).Offset(start).Select(selection).Find(&collection)
	} else {
		log.WithFields(log.Fields{"where": where}).Debug("Convert filter success.")
		c.Order("\"CreatedAt\" asc").Limit(count).Offset(start).Where(where).Select(selection).Find(&collection)
	}

	ret.Start = start
	ret.Count = int64(len(collection))
	ret.Total = total
	for _, v := range collection {
		m := model.TaskMember{}
		m.ID = v.ID
		m.Name = v.Name
		m.Description = v.Description
		m.CreatedAt = v.CreatedAt
		m.UpdatedAt = v.UpdatedAt
		m.CreatedByName = v.CreatedByName
		m.CreatedByURI = v.CreatedByURI
		m.TargetName = v.TargetName
		m.TargetURI = v.TargetURI
		m.CurrentStep = v.CurrentStep
		m.ExecutionState = v.ExecutionState
		m.Percentage = v.Percentage
		util.StringToStruct(v.ExecutionResult, &m.ExecutionResult)
		ret.Members = append(ret.Members, m)
	}
	return &ret, nil
}

// UpdateTask Update Task.
// It will return if the task already exist.
// It will return the updated task.
// It will return if it is committed.
// It will return error if any.
func (impl *TaskDBImplement) UpdateTask(id string, updateRequest *dto.UpdateTaskRequest) (bool, *model.Task, bool, error) {
	var (
		record entity.Task
	)

	c := commonDB.GetConnection()
	tx := c.Begin()
	if err := tx.Error; err != nil {
		log.WithFields(log.Fields{
			"id":    id,
			"error": err}).
			Warn("Update task in DB failed, start transaction failed.")
		return true, nil, false, err
	}
	if tx.Where("\"ID\" = ?", id).First(&record).RecordNotFound() {
		tx.Rollback()
		log.WithFields(log.Fields{
			"id": id}).
			Warn("Update task in DB failed, resource does not exist, transaction rollback.")
		return false, nil, false, commonConstError.ErrorResourceNotExist
	}
	m := record.ToModel()
	updateRequest.UpdateModel(m)
	record.Load(m)
	record.ID = id
	if err := tx.Save(record).Error; err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"id":    id,
			"error": err}).
			Warn("Update task in DB failed, save resource failed, transaction rollback.")
		return true, nil, false, err
	}
	return true, record.ToModel(), true, nil
}

// UpdateTaskStep Update Task step.
// It will return if the task already exist.
// It will return the updated task.
// It will return if it is committed.
// It will return error if any.
func (impl *TaskDBImplement) UpdateTaskStep(id string, updateRequest *dto.UpdateTaskStepRequest) (bool, *model.Task, bool, error) {
	var (
		record entity.Task
	)

	c := commonDB.GetConnection()
	tx := c.Begin()
	if err := tx.Error; err != nil {
		log.WithFields(log.Fields{
			"id":    id,
			"error": err}).
			Warn("Update task step in DB failed, start transaction failed.")
		return true, nil, false, err
	}
	if tx.Where("\"ID\" = ?", id).First(&record).RecordNotFound() {
		tx.Rollback()
		log.WithFields(log.Fields{
			"id": id}).
			Warn("Update task step in DB failed, resource does not exist, transaction rollback.")
		return false, nil, false, commonConstError.ErrorResourceNotExist
	}
	m := record.ToModel()
	updateRequest.UpdateModel(m)
	record.Load(m)
	record.ID = id
	if err := tx.Save(record).Error; err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"id":    id,
			"error": err}).
			Warn("Update task step in DB failed, save resource failed, transaction rollback.")
		return true, nil, false, err
	}
	return true, record.ToModel(), true, nil
}
