package db

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/enclosure/object/entity"	
)

// Enclosure is the DB implementation for enclosure.
type Enclosure struct {
	base.DB
}

// GetConnection return the DB connection.
func (impl *Enclosure) GetConnection() *gorm.DB {
	return base.GetConnection()
}

// ResourceName get the resource name.
func (impl *Enclosure) ResourceName() string {
	return "enclosure"
}

// NewEntity return the a new entity.
func (impl *Enclosure) NewEntity() base.EntityInterface {
	return new(entity.Enclosure)
}

// NewEntityCollection return a collection of entity.
func (impl *Enclosure) NewEntityCollection() interface{} {
	return new([]entity.Enclosure)
}

// NeedCheckDuplication return if need check duplication for entity.
func (impl *Enclosure) NeedCheckDuplication() bool {
	return true
}

// ConvertFindResultToCollection convert the Find() result to collection mode.
func (impl *Enclosure) ConvertFindResultToCollection(start int64, total int64, result interface{}) (*base.CollectionModel, *base.ErrorResponse) {
	collection, ok := result.(*[]entity.Enclosure)
	if !ok {
		log.Error("Enclosure.ConvertFindResult() failed, convert data failed.")
		return nil, base.NewErrorResponseInternalError()
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
func (impl *Enclosure) ConvertFindResultToModel(result interface{}) ([]base.ModelInterface, *base.ErrorResponse) {
	collection, ok := result.(*[]entity.Enclosure)
	if !ok {
		log.Error("Enclosure.ConvertFindResult() failed, convert data failed.")
		return nil, base.NewErrorResponseInternalError()
	}
	ret := make([]base.ModelInterface, 0)
	for _, v := range *collection {
		m := v.ToModel()
		ret = append(ret, m)
	}
	return ret, nil
}
