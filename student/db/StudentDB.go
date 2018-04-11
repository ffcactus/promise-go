package db

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"promise/apps"
	"promise/base"
	"promise/student/object/entity"
)

// StudentDB is the DB implementation for student.
type StudentDB struct {
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
	return apps.GetConnection()
}

// NeedCheckDuplication return if need check duplication for entity.
func (impl *StudentDB) NeedCheckDuplication() bool {
	return true
}

// ConvertFindResult convert the Find() result to model.
func (impl *StudentDB) ConvertFindResult(start int64, total int64, result interface{}) (*base.CollectionModel, error) {
	collection, ok := result.(*[]entity.Student)
	if !ok {
		log.Error("StudentDB.ConvertFindResult() failed.")
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
