package db

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/server/object/entity"
)

// ServerServerGroupDB is the concrete DB.
type ServerServerGroupDB struct {
	base.DB
}


// GetResourceName get the resource name.
func (impl *ServerServerGroupDB) GetResourceName() string {
	return "servergroup"
}

// NewEntity return the a new entity.
func (impl *ServerServerGroupDB) NewEntity() base.EntityInterface {
	e := new(entity.ServerServerGroup)
	e.Entity.TemplateImpl = e
	return e
}

// NewEntityCollection return a collection of entity.
func (impl *ServerServerGroupDB) NewEntityCollection() interface{} {
	return new([]entity.ServerServerGroup)
}

// GetConnection return the DB connection.
func (impl *ServerServerGroupDB) GetConnection() *gorm.DB {
	return base.GetConnection()
}

// NeedCheckDuplication return if need check duplication for entity.
func (impl *ServerServerGroupDB) NeedCheckDuplication() bool {
	return true
}

// ConvertFindResultToCollection convert the Find() result to collection mode.
func (impl *ServerServerGroupDB) ConvertFindResultToCollection(start int64, total int64, result interface{}) (*base.CollectionModel, error) {
	collection, ok := result.(*[]entity.ServerServerGroup)
	if !ok {
		log.Error("ServerServerGroupDB.ConvertFindResult() failed, convert data failed.")
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
func (impl *ServerServerGroupDB) ConvertFindResultToModel(result interface{}) ([]base.ModelInterface, error) {
	collection, ok := result.(*[]entity.ServerServerGroup)
	if !ok {
		log.Error("ServerServerGroupDB.ConvertFindResult() failed, convert data failed.")
		return nil, base.ErrorDataConvert
	}
	ret := make([]base.ModelInterface, 0)
	for _, v := range *collection {
		m := v.ToModel()
		ret = append(ret, m)
	}
	return ret, nil
}
