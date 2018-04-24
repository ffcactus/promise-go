package db

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/student/object/entity"
)

// StudentDB is the DB implementation for student.
type StudentDB struct {
}

// ResourceName get the resource name.
func (impl *StudentDB) ResourceName() string {
	return "student"
}

// NewEntity return the a new entity.
func (impl *StudentDB) NewEntity() base.EntityInterface {
	e := new(entity.Student)
	e.Entity.TemplateImpl = e
	return e
}

// NewEntityCollection return a collection of entity.
func (impl *StudentDB) NewEntityCollection() interface{} {
	return new([]entity.Student)
}

// GetConnection return the DB connection.
func (impl *StudentDB) GetConnection() *gorm.DB {
	return base.GetConnection()
}

// NeedCheckDuplication return if need check duplication for entity.
func (impl *StudentDB) NeedCheckDuplication() bool {
	return true
}

// ConvertFindResultToCollection convert the Find() result to collection mode.
func (impl *StudentDB) ConvertFindResultToCollection(start int64, total int64, result interface{}) (*base.CollectionModel, *base.Message) {
	collection, ok := result.(*[]entity.Student)
	if !ok {
		log.Error("StudentDB.ConvertFindResult() failed.")
		message := base.NewMessageInternalError()
		return nil, &message
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
func (impl *StudentDB) ConvertFindResultToModel(result interface{}) ([]base.ModelInterface, *base.Message) {
	collection, ok := result.(*[]entity.Student)
	if !ok {
		log.Error("StudentDB.ConvertFindResult() failed.")
		message := base.NewMessageInternalError()
		return nil, &message
	}
	ret := make([]base.ModelInterface, 0)
	for _, v := range *collection {
		m := v.ToModel()
		ret = append(ret, m)
	}
	return ret, nil
}
