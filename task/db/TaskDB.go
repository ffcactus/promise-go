package db

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/task/object/entity"
)

// TaskDB is the DB implementation for task.
type TaskDB struct {
	base.DB
}

// GetResourceName get the resource name.
func (impl *TaskDB) GetResourceName() string {
	return "task"
}

// NewEntity return the a new entity.
func (impl *TaskDB) NewEntity() base.EntityInterface {
	e := new(entity.Task)
	e.Entity.TemplateImpl = e
	return e
}

// NewEntityCollection return a collection of entity.
func (impl *TaskDB) NewEntityCollection() interface{} {
	return new([]entity.Task)
}

// GetConnection return the DB connection.
func (impl *TaskDB) GetConnection() *gorm.DB {
	return base.GetConnection()
}

// NeedCheckDuplication return if need check duplication for entity.
func (impl *TaskDB) NeedCheckDuplication() bool {
	return false
}

// ConvertFindResultToCollection convert the Find() result to collection mode.
func (impl *TaskDB) ConvertFindResultToCollection(start int64, total int64, result interface{}) (*base.CollectionModel, error) {
	collection, ok := result.(*[]entity.Task)
	if !ok {
		log.Error("TaskDB.ConvertFindResult() failed.")
		return nil, base.ErrorDataConvert
	}
	ret := base.CollectionModel{}
	ret.Start = start
	ret.Count = int64(len(*collection))
	ret.Total = total
	for _, v := range *collection {
		ret.Members = append(ret.Members, v.ToCollectionMember())
	}
	return &ret, nil
}

// ConvertFindResultToModel convert the Find() result to model slice
func (impl *TaskDB) ConvertFindResultToModel(result interface{}) ([]base.ModelInterface, error) {
	collection, ok := result.(*[]entity.Task)
	if !ok {
		log.Error("TaskDB.ConvertFindResult() failed.")
		return nil, base.ErrorDataConvert
	}
	ret := make([]base.ModelInterface, 0)
	for _, v := range *collection {
		m := v.ToModel()
		ret = append(ret, m)
	}
	return ret, nil
}

// UpdateTaskStep update the task step.
func (impl *TaskDB) UpdateTaskStep(id string, request base.ActionRequestInterface) (bool, base.ModelInterface, bool, error) {
	var (
		record = impl.NewEntity()
		c      = impl.GetConnection()
	)
	tx := c.Begin()
	if err := tx.Error; err != nil {
		log.WithFields(log.Fields{
			"id":    id,
			"error": err,
		}).Warn("Update task step in DB failed, start transaction failed.")
		return true, nil, false, err
	}
	if exist, err := impl.GetInternal(tx, id, record); !exist || err != nil {
		return false, nil, false, err
	}
	m := record.ToModel()
	if err := request.UpdateModel(m); err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"id":    id,
			"error": err}).
			Warn("Update task step in DB failed, update model failed, transaction rollback.")
		return true, nil, false, err
	}
	record.Load(m)
	record.SetID(id)
	if err := tx.Save(record).Error; err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"id":    id,
			"error": err}).
			Warn("Update task step in DB failed, save resource failed, transaction rollback.")
		return true, nil, false, err
	}
	if err := tx.Commit().Error; err != nil {
		log.WithFields(log.Fields{
			"id":    id,
			"error": err,
		}).Warn("Update task step in DB failed, commit failed.")
		return true, nil, false, err
	}
	return true, record.ToModel(), true, nil
}

// UpdateTask update the task.
// func (impl *TaskDB) UpdateTask(id string, request base.ActionRequestInterface) (bool, base.ModelInterface, bool, error) {
// 	var (
// 		record = impl.NewEntity()
// 		c      = impl.GetConnection()
// 	)
// 	tx := c.Begin()
// 	if err := tx.Error; err != nil {
// 		log.WithFields(log.Fields{
// 			"id":       id,
// 			"error":    err,
// 		}).Warn("Update task in DB failed, start transaction failed.")
// 		return true, nil, false, err
// 	}
// 	if exist, err := impl.GetInternal(tx, id, record); !exist || err != nil {
// 		return false, nil, false, err
// 	}
// 	m := record.ToModel()
// 	if err := request.UpdateModel(m); err != nil {
// 		tx.Rollback()
// 		log.WithFields(log.Fields{
// 			"id":    id,
// 			"error": err}).
// 			Warn("Update task in DB failed, update model failed, transaction rollback.")
// 		return true, nil, false, err
// 	}
// 	record.Load(m)
// 	record.SetID(id)
// 	if err := tx.Save(record).Error; err != nil {
// 		tx.Rollback()
// 		log.WithFields(log.Fields{
// 			"id":    id,
// 			"error": err}).
// 			Warn("Update task in DB failed, save resource failed, transaction rollback.")
// 		return true, nil, false, err
// 	}
// 	if err := tx.Commit().Error; err != nil {
// 		log.WithFields(log.Fields{
// 			"id":       id,
// 			"error":    err,
// 		}).Warn("Update task in DB failed, commit failed.")
// 		return true, nil, false, err
// 	}
// 	return true, record.ToModel(), true, nil
// }
