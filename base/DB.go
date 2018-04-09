package base

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

// DBInterface is the interface that DB should have.
type DBInterface interface {
	NeedCheckDuplication() bool
	GetConnection() *gorm.DB
	NewEntity() EntityInterface
	Post(ModelInterface) (bool, ModelInterface, bool, error)
}

// DB is the DB implementation in Promise project.
type DB struct {
	Interface DBInterface
}

// Post will save model to the DB.
// It will return if there is one exist already with the same name.
// It will return the newly created one if commited, or nil.
// It will return if the transaction commited.
// It will return error if any.
func (i *DB) Post(m ModelInterface) (bool, ModelInterface, bool, error) {
	var (
		record = i.Interface.NewEntity()
	)
	c := i.Interface.GetConnection()

	tx := c.Begin()
	if err := tx.Error; err != nil {
		log.WithFields(log.Fields{
			"resource": m.GetDebugName(),
			"error":    err,
		}).Warn("Post resource in DB failed, start transaction failed.")
		return false, nil, false, err
	}
	if i.Interface.NeedCheckDuplication() {
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
