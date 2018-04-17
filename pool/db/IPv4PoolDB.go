package db

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/pool/object/entity"
)

// IPv4PoolDB is the concrete DB.
type IPv4PoolDB struct {
	base.DB
}

// GetResourceName get the resource name.
func (impl *IPv4PoolDB) GetResourceName() string {
	return "ipv4"
}

// NewEntity return the a new entity.
func (impl *IPv4PoolDB) NewEntity() base.EntityInterface {
	e := new(entity.IPv4Pool)
	e.Entity.TemplateImpl = e
	return e
}

// NewEntityCollection return a collection of entity.
func (impl *IPv4PoolDB) NewEntityCollection() interface{} {
	return new([]entity.IPv4Pool)
}

// GetConnection return the DB connection.
func (impl *IPv4PoolDB) GetConnection() *gorm.DB {
	return base.GetConnection()
}

// NeedCheckDuplication return if need check duplication for entity.
func (impl *IPv4PoolDB) NeedCheckDuplication() bool {
	return true
}

// ConvertFindResultToCollection convert the Find() result to collection mode.
func (impl *IPv4PoolDB) ConvertFindResultToCollection(start int64, total int64, result interface{}) (*base.CollectionModel, error) {
	collection, ok := result.(*[]entity.IPv4Pool)
	if !ok {
		log.Error("IPv4PoolDB.ConvertFindResult() failed.")
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
func (impl *IPv4PoolDB) ConvertFindResultToModel(result interface{}) ([]base.ModelInterface, error) {
	collection, ok := result.(*[]entity.IPv4Pool)
	if !ok {
		log.Error("IPv4PoolDB.ConvertFindResult() failed.")
		return nil, base.ErrorDataConvert
	}
	ret := make([]base.ModelInterface, 0)
	for _, v := range *collection {
		m := v.ToModel()
		ret = append(ret, m)
	}
	return ret, nil
}
