package db

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/server/object/entity"
)

// ServerServerGroup is the concrete DB.
type ServerServerGroup struct {
	base.DB
}

// ResourceName get the resource name.
func (impl *ServerServerGroup) ResourceName() string {
	return "servergroup"
}

// NewEntity return the a new entity.
func (impl *ServerServerGroup) NewEntity() base.EntityInterface {
	e := new(entity.ServerServerGroup)
	e.Entity.TemplateImpl = e
	return e
}

// NewEntityCollection return a collection of entity.
func (impl *ServerServerGroup) NewEntityCollection() interface{} {
	return new([]entity.ServerServerGroup)
}

// GetConnection return the DB connection.
func (impl *ServerServerGroup) GetConnection() *gorm.DB {
	return base.GetConnection()
}

// NeedCheckDuplication return if need check duplication for entity.
func (impl *ServerServerGroup) NeedCheckDuplication() bool {
	return true
}

// ConvertFindResultToCollection convert the Find() result to collection mode.
func (impl *ServerServerGroup) ConvertFindResultToCollection(start int64, total int64, result interface{}) (*base.CollectionModel, *base.Message) {
	collection, ok := result.(*[]entity.ServerServerGroup)
	if !ok {
		log.Error("ServerServerGroup.ConvertFindResult() failed, convert data failed.")
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
func (impl *ServerServerGroup) ConvertFindResultToModel(result interface{}) ([]base.ModelInterface, *base.Message) {
	collection, ok := result.(*[]entity.ServerServerGroup)
	if !ok {
		log.Error("ServerServerGroup.ConvertFindResult() failed, convert data failed.")
		return nil, base.NewMessageInternalError()
	}
	ret := make([]base.ModelInterface, 0)
	for _, v := range *collection {
		m := v.ToModel()
		ret = append(ret, m)
	}
	return ret, nil
}
