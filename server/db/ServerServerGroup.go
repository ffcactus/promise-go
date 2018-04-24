package db

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/server/object/entity"
	"promise/server/object/message"
)

// ServerServerGroup is the concrete DB.
type ServerServerGroup struct {
	base.DB
}

// ResourceName get the resource name.
func (impl *ServerServerGroup) ResourceName() string {
	return "server-servergroup"
}

// NewEntity return the a new entity.
func (impl *ServerServerGroup) NewEntity() base.EntityInterface {
	return new(entity.ServerServerGroup)
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

// Delete will override the default process.
// We should not delete the default server-servergroup relationship, since a server
// should always in the default servergroup.
func (impl *ServerServerGroup) Delete(id string) (base.ModelInterface, *base.Message) {
	var (
		ssg, previous entity.ServerServerGroup
		c                = impl.GetConnection()
	)

	tx := c.Begin()
	if err := tx.Error; err != nil {
		log.WithFields(log.Fields{
			"id":    id,
			"error": err}).
			Warn("Delete server-servergroup in DB failed, start transaction failed.")
		return nil, base.NewMessageTransactionError()
	}
	if tx.Where("\"ID\" = ?", id).First(&previous).RecordNotFound() {
		tx.Rollback()
		log.WithFields(log.Fields{
			"id": id}).
			Warn("Delete server-servergroup in DB failed, resource does not exist, transaction rollback.")
		return nil, base.NewMessageNotExist()
	}
	if previous.ServerGroupID == DefaultServerGroupID {
		tx.Rollback()
		log.WithFields(log.Fields{
			"id": id}).
			Warn("Delete server-servergroup in DB failed, delete default server-servergroup, transaction rollback.")
		return nil, message.NewMessageServerServerGroupDeleteDefault()
	}
	ssg.ID = id
	if err := tx.Delete(&ssg).Error; err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"id": id}).
			Warn("Delete server-servergroup in DB failed, delete resource failed, transaction rollback.")		
		return nil, base.NewMessageTransactionError()
	}
	if err := tx.Commit().Error; err != nil {
		log.WithFields(log.Fields{
			"id":    id,
			"error": err}).
			Warn("Delete server-servergroup in DB failed, commit failed.")
		return nil, base.NewMessageTransactionError()
	}
	return previous.ToModel(), nil
}

// DeleteCollection will override the default process.
// We should not delete the default server-servergroup relationship, since a server
// should always in the default servergroup.
func (impl *ServerServerGroup) DeleteCollection() ([]base.ModelInterface, *base.Message) {
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
		}).Warn("Delete collection in DB failed, start transaction failed.")
		return nil, base.NewMessageTransactionError()
	}

	if err := tx.Where("\"ServerGroupID\" <> ?", DefaultServerGroupID).Find(recordCollection).Error; err != nil {
		log.WithFields(log.Fields{
			"resource": name,
			"error":    err,
		}).Warn("Delete collection in DB failed, find resource failed.")
		return nil, base.NewMessageTransactionError()
	}

	if err := tx.Where("\"ServerGroupID\" <> ?", DefaultServerGroupID).Delete(entity.ServerServerGroup{}).Error; err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"resource": name,
			"error":    err,
		}).Warn("Delete collection in DB failed, delete resources failed, transaction rollback.")
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