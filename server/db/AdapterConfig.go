package db

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/server/object/entity"
)

// AdapterConfig is the DB implementation for adapter config.
type AdapterConfig struct {
	base.DB
}

// GetConnection return the DB connection.
func (impl *AdapterConfig) GetConnection() *gorm.DB {
	return base.GetConnection()
}

// ResourceName get the resource name.
func (impl *AdapterConfig) ResourceName() string {
	return "adapterconfig"
}

// NewEntity return the a new entity.
func (impl *AdapterConfig) NewEntity() base.EntityInterface {
	return new(entity.AdapterConfig)
}

// NewEntityCollection return a collection of entity.
func (impl *AdapterConfig) NewEntityCollection() interface{} {
	return new([]entity.AdapterConfig)
}

// NeedCheckDuplication return if need check duplication for entity.
func (impl *AdapterConfig) NeedCheckDuplication() bool {
	return false
}

// ConvertFindResultToCollection convert the Find() result to collection mode.
func (impl *AdapterConfig) ConvertFindResultToCollection(start int64, total int64, result interface{}) (*base.CollectionModel, *base.Message) {
	collection, ok := result.(*[]entity.AdapterConfig)
	if !ok {
		log.Error("AdapterConfig.ConvertFindResult() failed, convert data failed.")
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
func (impl *AdapterConfig) ConvertFindResultToModel(result interface{}) ([]base.ModelInterface, *base.Message) {
	collection, ok := result.(*[]entity.AdapterConfig)
	if !ok {
		log.Error("AdapterConfig.ConvertFindResult() failed, convert data failed.")
		return nil, base.NewMessageInternalError()
	}
	ret := make([]base.ModelInterface, 0)
	for _, v := range *collection {
		m := v.ToModel()
		ret = append(ret, m)
	}
	return ret, nil
}
