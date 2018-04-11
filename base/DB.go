package base

import (
	"reflect"
	"strings"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

// DBTemplateInterface is the interface that a concrete DB implement should have.
type DBTemplateInterface interface {
	NeedCheckDuplication() bool
	GetConnection() *gorm.DB
	NewEntity() EntityInterface
	// NewEntityCollection() []base.EntityInterface
}

// DBInterface is the interface that DB should have.
type DBInterface interface {
	Post(ModelInterface) (bool, ModelInterface, bool, error)
	Get(id string) ModelInterface
	Delete(id string) (bool, ModelInterface, bool, error)
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
func (impl *DB) Post(m ModelInterface) (bool, ModelInterface, bool, error) {
	var (
		record = impl.TemplateImpl.NewEntity()
		c      = impl.TemplateImpl.GetConnection()
	)
	tx := c.Begin()
	if err := tx.Error; err != nil {
		log.WithFields(log.Fields{
			"resource": m.GetDebugName(),
			"error":    err,
		}).Warn("Post resource in DB failed, start transaction failed.")
		return false, nil, false, err
	}
	if impl.TemplateImpl.NeedCheckDuplication() {
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
	if err := c.Create(record).Error; err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"name":  m.GetDebugName(),
			"error": err,
		}).Warn("Post resource in DB failed, create resource failed, transaction rollback.")
		return false, nil, false, err
	}
	if err := tx.Save(record).Error; err != nil {
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
func (impl *DB) get(tx *gorm.DB, id string, record EntityInterface) (bool, error) {
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
func (impl *DB) Get(id string) ModelInterface {
	var (
		record = impl.TemplateImpl.NewEntity()
		c      = impl.TemplateImpl.GetConnection()
	)
	tx := c.Begin()
	if err := tx.Error; err != nil {
		log.WithFields(log.Fields{
			"id":    id,
			"error": err,
		}).Warn("Get resource in DB failed, start transaction failed.")
		return nil
	}
	if exist, err := impl.get(tx, id, record); exist && err == nil {
		return record.ToModel()
	}
	return nil
}

// Delete is the default implement to delete resource from DB.
// It will return if the one exist.
// It will return the deleted one if commited.
// It will return wether the operation commited.
// It will return error if any.
func (impl *DB) Delete(id string) (bool, ModelInterface, bool, error) {
	var (
		record   = impl.TemplateImpl.NewEntity()
		previous = impl.TemplateImpl.NewEntity()
		c        = impl.TemplateImpl.GetConnection()
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
	if exist, err := impl.get(tx, id, previous); err != nil || !exist {
		return false, nil, false, err
	}
	record.SetID(id)
	for _, v := range record.GetAssociation() {
		if err := tx.Delete(v).Error; err != nil {
			tx.Rollback()
			log.WithFields(log.Fields{
				"id":    id,
				"error": err,
			}).Warn("Delete resource from DB failed, delete association failed, transaction rollback.")
			return true, nil, false, err
		}
	}
	if err := tx.Delete(record).Error; err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"id":    id,
			"error": err,
		}).Warn("Delete resource from DB failed, delete resource failed, transaction rollback.")
		return true, nil, false, err
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

func (impl *DB) convertFilter(filter string) (string, error) {
	if filter == "" {
		return "", nil
	}
	cmds := strings.Split(filter, " ")
	if len(cmds) != 3 {
		return "", ErrorConvertFilter
	}
	switch strings.ToLower(cmds[1]) {
	case "eq":
		return "\"" + cmds[0] + "\"" + " = " + cmds[2], nil
	default:
		return "", ErrorConvertFilter
	}
}

// GetCollection get the resource collection in DB.
// It returns nil if any error.
func (impl *DB) GetCollection(start int64, count int64, filter string) (*CollectionModel, error) {
	var (
		total int64
		record = impl.TemplateImpl.NewEntity()
		// recordCollection = impl.TemplateImpl.NewEntityCollection()
		recordCollection = reflect.SliceOf(reflect.TypeOf(record))
		ret = new(CollectionModel)
		c        = impl.TemplateImpl.GetConnection()
	)

	if err := c.Table(record.TableName()).Count(&total).Error; err != nil {
		log.WithFields(log.Fields{
			"error":  err,
		}).Warn("Get resource collection in DB failed, get count failed.")
		return nil, err		
	}
	where, err := impl.convertFilter(filter)
	if err != nil {
		log.WithFields(log.Fields{
			"filter": filter,
			"error":  err}).
			Warn("Get resource collection in DB failed, convert filter failed.")
		return nil, err
	}
	log.WithFields(log.Fields{
		"where": where,
	}).Debug("Convert filter success.")
	c.Limit(count).Offset(start).Where(where).Find(recordCollection)
	
	ret.Start = start
	_rc := reflect.ValueOf(recordCollection)
	_count := _rc.Len()
	ret.Count = int64(_count)
	ret.Total = total
	for i := 0; i < _count; i++ {
		_interface, ok := _rc.Index(i).Interface().(EntityInterface)
		if !ok {
			log.Error("Get resource collection in DB failed, convert data failed.")
		}
		ret.Members = append(ret.Members, _interface.ToMember())
	}
	log.Info("--- collection DB returned ---")
	PrintJSON(ret)
	return ret, nil
}