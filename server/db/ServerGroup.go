package db

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/server/object/entity"
	"promise/server/object/message"
)

// DefaultServerGroupID records the ID of default servergroup. We don't have to retrieve it each time.
var (
	DefaultServerGroupID string
	messageTransactionError = base.NewMessageTransactionError()
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
func (impl *ServerGroup) ConvertFindResultToCollection(start int64, total int64, result interface{}) (*base.CollectionModel, *base.Message) {
	collection, ok := result.(*[]entity.ServerGroup)
	if !ok {
		log.Error("ServerGroup.ConvertFindResult() failed, convert data failed.")
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
func (impl *ServerGroup) ConvertFindResultToModel(result interface{}) ([]base.ModelInterface, *base.Message) {
	collection, ok := result.(*[]entity.ServerGroup)
	if !ok {
		log.Error("ServerGroup.ConvertFindResult() failed, convert data failed.")
		return nil, base.NewMessageInternalError()
	}
	ret := make([]base.ModelInterface, 0)
	for _, v := range *collection {
		m := v.ToModel()
		ret = append(ret, m)
	}
	return ret, nil
}

// Delete will override the default Delete process.
// We should not delete the default servergroup.
func (impl *ServerGroup) Delete(id string) (base.ModelInterface, *base.Message) {
	if id == DefaultServerGroupID {
		return nil, message.NewMessageServerGroupDeleteDefault()
	}
	return impl.DB.Delete(id)
}

// DeleteCollection will override the default DeleteCollection process.
// We should not delete the default servergroup.
func (impl *ServerGroup) DeleteCollection() ([]base.ModelInterface, *base.Message) {
	var (
		name             = impl.ResourceName()
		recordCollection = impl.NewEntityCollection()
		c                = impl.GetConnection()
	)

	// We need transaction to ensure the total and the query count is consistent.
	tx := c.Begin()
	if err := tx.Error; err != nil {
		log.WithFields(log.Fields{
			"resource": name,
			"error":    err,
		}).Warn("Get collection in DB failed, start transaction failed.")
		return nil, base.NewMessageTransactionError()
	}

	if err := tx.Where("\"Name\" <> ?", "all").Find(recordCollection).Error; err != nil {
		log.WithFields(log.Fields{
			"resource": name,
			"error":    err,
		}).Warn("Delete collection in DB failed, find resource failed.")
		return nil, base.NewMessageTransactionError()
	}
	if err := tx.Where("\"Name\" <> ?", "all").Delete(entity.ServerGroup{}).Error; err != nil {
		log.WithFields(log.Fields{
			"resource": name,
			"error":    err,
		}).Warn("Delete collection in DB failed.")
		return nil, base.NewMessageTransactionError()		
	}
	ret, message := impl.TemplateImpl.ConvertFindResultToModel(recordCollection)
	if message != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"resource": name,
			"message":    message.ID,
		}).Warn("Delete collection in DB failed, convert find result failed, transaction rollback.")
		return nil, message
	}
	if err := tx.Commit().Error; err != nil {
		log.WithFields(log.Fields{
			"resource": name,
			"error":    err,
		}).Warn("Delete collection in DB failed, commit failed.")
		return nil, base.NewMessageTransactionError()
	}
	return ret, nil	
}