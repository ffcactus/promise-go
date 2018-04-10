package base

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

// DBInterface is the interface that DB should have.
type DBInterface interface {
	Post(ModelInterface) (bool, ModelInterface, bool, error)
	Get(id string) ModelInterface
	Delete(id string) (bool, ModelInterface, bool, error)
}

// DBTemplateInterface is the interface that a concrete DB implement should have.
type DBTemplateInterface interface {
	NeedCheckDuplication() bool
	GetConnection() *gorm.DB
	NewEntity() EntityInterface
}

// DB is the DB implementation in Promise project.
type DB struct {
	TemplateImpl DBTemplateInterface
}

// Post is the default implement to post resource in DB.
// It will return if there is one exist already with the same name.
// It will return the newly created one if commited, or nil.
// It will return if the transaction commited.
// It will return error if any.
func (i *DB) Post(m ModelInterface) (bool, ModelInterface, bool, error) {
	var (
		record = i.TemplateImpl.NewEntity()
		c = i.TemplateImpl.GetConnection()
	)

	tx := c.Begin()
	if err := tx.Error; err != nil {
		log.WithFields(log.Fields{
			"resource": m.GetDebugName(),
			"error":    err,
		}).Warn("Post resource in DB failed, start transaction failed.")
		return false, nil, false, err
	}
	if i.TemplateImpl.NeedCheckDuplication() {
		where := "\"" + record.GetPropertyNameForDuplicationCheck() + "\" = ?"
		if !tx.Where(where, m.GetValueForDuplicationCheck()).First(record).RecordNotFound() {
			tx.Rollback()
			log.WithFields(log.Fields{
				"id":   record.GetID(),
				"name": record.GetDebugName(),
			}).Warn("Post resource in DB failed, duplicated resource, transaction rollback.")
			return true, nil, false, ErrorResourceNotExist
		}
	}

	record.Load(m)
	record.SetID(uuid.New().String())
	PrintJSON(record)
	if err := c.Create(&record).Error; err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"name":  m.GetDebugName(),
			"error": err,
		}).Warn("Post resource in DB failed, create resource failed, transaction rollback.")
		return false, nil, false, err
	}
	if err := tx.Save(&record).Error; err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"name":  m.GetDebugName(),
			"error": err,
		}).Warn("Post resource in DB failed, save resource failed, transaction rollback.")
		return false, nil, false, err
	}
	if err := tx.Commit().Error; err != nil {
		log.WithFields(log.Fields{
			"name":  m.GetDebugName(),
			"error": err,
		}).Warn("Post resource in DB failed, commit failed.")
		return false, nil, false, err
	}
	return false, record.ToModel(), true, nil
}

// get is part of the process to get resource in DB, since many other operation 
// need this process, we seperate it out.
// It will return if the resource been found.
// It will return error if any.
func (i *DB) get(tx *gorm.DB, id string, record EntityInterface) (bool, error) {
	preload := record.GetPreload()
	if tx.Where("\"ID\" = ?", id).First(record).RecordNotFound() {
		tx.Rollback()
		log.WithFields(log.Fields{
			"id": id,
		}).Warn("Get resource in DB failed, resource does not exist, transaction rollback.")
		return false, ErrorResourceNotExist
	}

	tx.Where("\"ID\" = ?", id)
	for _, v := range preload {
		tx = tx.Preload(v)
	}
	if err := tx.First(record).Error; err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"id":    id,
			"error": err,
		}).Warn("Get resource in DB failed, fatch failed, transaction rollback.")
		return true, err
	}
	return true, nil
}

// Get is the default implement to get resource in DB.
// If the resource does not exist in the DB return nil.
func (i *DB) Get(id string) ModelInterface {
	var (
		record = i.TemplateImpl.NewEntity()
		c = i.TemplateImpl.GetConnection()
	)
	tx := c.Begin()
	if err := tx.Error; err != nil {
		log.WithFields(log.Fields{
			"id":    id,
			"error": err,
		}).Warn("Get resource in DB failed, start transaction failed.")
		return nil
	}
	if exist, err := i.get(tx, id, record); exist && err == nil {
		return record.ToModel()
	}
	return nil
}

// Delete is the default implement to delete resource from DB.
// It will return if the one exist.
// It will return the deleted one if commited.
// It will return wether the operation commited.
// It will return error if any.
func (i *DB) Delete(id string) (bool, ModelInterface, bool, error) {
	var (
		record = i.TemplateImpl.NewEntity()
		previous = i.TemplateImpl.NewEntity()
		c = i.TemplateImpl.GetConnection()
	)

	if id == "" {
		return true, nil, false, ErrorIDFormat
	}
	tx := c.Begin()
	if err := tx.Error; err != nil {
		log.WithFields(log.Fields{
			"id":    id,
			"error": err,
		}).Warn("Delete resource from DB failed, start transaction failed.")
		return true, nil, false, err
	}
	if exist, err := i.get(tx, id, previous); err != nil || !exist {
		return false, nil, false, err
	}
	record.SetID(id)
	for _, v := range record.GetAssociation() {
		if err := tx.Delete(v).Error; err != nil {
			tx.Rollback()
			log.WithFields(log.Fields{
				"id": id,
			}).Warn("Delete resource from DB failed, delete association failed, transaction rollback.")
			return true, nil, false, err
		}
	}
	if err := tx.Commit().Error; err != nil {
		log.WithFields(log.Fields{
			"id":    id,
			"error": err,
		}).Warn("Delete resource from DB failed, commit failed.")
		return true, nil, false, err	
	}
	return true, previous.ToModel(), true, nil
}