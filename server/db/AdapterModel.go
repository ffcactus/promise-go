package db

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/server/object/entity"
)

// AdapterModel is the concrete DB.
type AdapterModel struct {
	base.DB
}

// ResourceName get the resource name.
func (impl *AdapterModel) ResourceName() string {
	return "adapterconfig"
}

// NewEntity return the a new entity.
func (impl *AdapterModel) NewEntity() base.EntityInterface {
	return new(entity.AdapterModel)
}

// NewEntityCollection return a collection of entity.
func (impl *AdapterModel) NewEntityCollection() interface{} {
	return new([]entity.AdapterModel)
}

// GetConnection return the DB connection.
func (impl *AdapterModel) GetConnection() *gorm.DB {
	return base.GetConnection()
}

// NeedCheckDuplication return if need check duplication for entity.
func (impl *AdapterModel) NeedCheckDuplication() bool {
	return false
}

// ConvertFindResultToCollection convert the Find() result to collection mode.
func (impl *AdapterModel) ConvertFindResultToCollection(start int64, total int64, result interface{}) (*base.CollectionModel, *base.Message) {
	collection, ok := result.(*[]entity.AdapterModel)
	if !ok {
		log.Error("AdapterModel.ConvertFindResult() failed, convert data failed.")
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
func (impl *AdapterModel) ConvertFindResultToModel(result interface{}) ([]base.ModelInterface, *base.Message) {
	collection, ok := result.(*[]entity.AdapterModel)
	if !ok {
		log.Error("AdapterModel.ConvertFindResult() failed, convert data failed.")
		return nil, base.NewMessageInternalError()
	}
	ret := make([]base.ModelInterface, 0)
	for _, v := range *collection {
		m := v.ToModel()
		ret = append(ret, m)
	}
	return ret, nil
}
