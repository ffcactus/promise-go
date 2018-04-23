package db

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/server/object/entity"
)

// DefaultServerGroupID records the ID of default servergroup. We don't have to retrieve it each time.
var (
	DefaultServerGroupID string
)

// ServerGroup is the concrete DB.
type ServerGroup struct {
	base.DB
}

// ResourceName get the resource name.
func (impl *ServerGroup) ResourceName() string {
	return "servergroup"
}

// NewEntity return the a new entity.
func (impl *ServerGroup) NewEntity() base.EntityInterface {
	e := new(entity.ServerGroup)
	e.Entity.TemplateImpl = e
	return e
}

// NewEntityCollection return a collection of entity.
func (impl *ServerGroup) NewEntityCollection() interface{} {
	return new([]entity.ServerGroup)
}

// GetConnection return the DB connection.
func (impl *ServerGroup) GetConnection() *gorm.DB {
	return base.GetConnection()
}

// NeedCheckDuplication return if need check duplication for entity.
func (impl *ServerGroup) NeedCheckDuplication() bool {
	return true
}

// ConvertFindResultToCollection convert the Find() result to collection mode.
func (impl *ServerGroup) ConvertFindResultToCollection(start int64, total int64, result interface{}) (*base.CollectionModel, error) {
	collection, ok := result.(*[]entity.ServerGroup)
	if !ok {
		log.Error("ServerGroup.ConvertFindResult() failed, convert data failed.")
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
func (impl *ServerGroup) ConvertFindResultToModel(result interface{}) ([]base.ModelInterface, error) {
	collection, ok := result.(*[]entity.ServerGroup)
	if !ok {
		log.Error("ServerGroup.ConvertFindResult() failed, convert data failed.")
		return nil, base.ErrorDataConvert
	}
	ret := make([]base.ModelInterface, 0)
	for _, v := range *collection {
		m := v.ToModel()
		ret = append(ret, m)
	}
	return ret, nil
}
