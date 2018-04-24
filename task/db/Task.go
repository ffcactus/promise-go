package db

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/task/object/entity"
)

// Task is the concrete DB.
type Task struct {
	base.DB
}

// ResourceName get the resource name.
func (impl *Task) ResourceName() string {
	return "task"
}

// NewEntity return the a new entity.
func (impl *Task) NewEntity() base.EntityInterface {
	e := new(entity.Task)
	e.Entity.TemplateImpl = e
	return e
}

// NewEntityCollection return a collection of entity.
func (impl *Task) NewEntityCollection() interface{} {
	return new([]entity.Task)
}

// GetConnection return the DB connection.
func (impl *Task) GetConnection() *gorm.DB {
	return base.GetConnection()
}

// NeedCheckDuplication return if need check duplication for entity.
func (impl *Task) NeedCheckDuplication() bool {
	return false
}

// ConvertFindResultToCollection convert the Find() result to collection mode.
func (impl *Task) ConvertFindResultToCollection(start int64, total int64, result interface{}) (*base.CollectionModel, *base.Message) {
	collection, ok := result.(*[]entity.Task)
	if !ok {
		log.Error("Task.ConvertFindResult() failed.")
		return nil, base.NewMessageInternalError()
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
func (impl *Task) ConvertFindResultToModel(result interface{}) ([]base.ModelInterface, *base.Message) {
	collection, ok := result.(*[]entity.Task)
	if !ok {
		log.Error("Task.ConvertFindResult() failed.")
		return nil, base.NewMessageInternalError()
	}
	ret := make([]base.ModelInterface, 0)
	for _, v := range *collection {
		m := v.ToModel()
		ret = append(ret, m)
	}
	return ret, nil
}
