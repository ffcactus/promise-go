package db

import (
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	commonDB "promise/common/db"
	commonConstError "promise/common/object/constError"
	"promise/pool/object/entity"
	"promise/pool/object/model"
	"strings"
)

var (
	poolDBInstance PoolDBImplement
)

// PoolDBImplement is the implementation.
type PoolDBImplement struct {
}

// GetPoolDB return the singleton.
func GetPoolDB() PoolDBInterface {
	return &poolDBInstance
}

// PostIPv4Pool will save the IPv4 pool if no one with the same name.
// It will return if there is one exist already with the same name.
// It will return the newly created one if commited, or nil.
// It will return if the transaction commited.
// It will return error if any.
func (i *PoolDBImplement) PostIPv4Pool(m *model.IPv4Pool) (bool, *model.IPv4Pool, bool, error) {
	var record entity.IPv4Pool

	c := commonDB.GetConnection()
	tx := c.Begin()
	if err := tx.Error; err != nil {
		log.WithFields(log.Fields{
			"name":  m.Name,
			"error": err}).
			Warn("Post IPv4 pool in DB failed, start transaction failed.")
		return false, nil, false, err
	}
	if !tx.Where("\"Name\" = ?", m.Name).First(&record).RecordNotFound() {
		tx.Rollback()
		log.WithFields(log.Fields{
			"id":   record.ID,
			"name": m.Name}).
			Warn("Post IPv4 pool in DB failed, duplicated resource, transaction rollback.")
		return true, nil, false, commonConstError.ErrorResourceNotExist
	}
	record.Load(m)
	record.ID = uuid.New().String()
	if err := c.Create(&record).Error; err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"name":  m.Name,
			"error": err}).
			Warn("Post IPv4 pool in DB failed, create resource failed, transaction rollback.")
		return false, nil, false, err
	}
	if err := tx.Commit().Error; err != nil {
		log.WithFields(log.Fields{
			"name":  m.Name,
			"error": err}).
			Warn("Post IPv4 pool in DB failed, commit failed.")
		return false, nil, false, err
	}
	return false, record.ToModel(), true, nil
}

// GetIPv4Pool get the IPv4 pool by ID.
func (i *PoolDBImplement) GetIPv4Pool(id string) *model.IPv4Pool {
	var record entity.IPv4Pool
	c := commonDB.GetConnection()
	if c.Where("\"ID\" = ?", id).First(&record).RecordNotFound() {
		return nil
	}
	return record.ToModel()
}

func (i *PoolDBImplement) convertFilter(filter string) (string, error) {
	if filter == "" {
		return "", nil
	}
	cmds := strings.Split(filter, " ")
	if len(cmds) != 3 {
		return "", commonConstError.ErrorConvertFilter
	}
	switch strings.ToLower(cmds[1]) {
	case "eq":
		return "\"" + cmds[0] + "\"" + " = " + cmds[2], nil
	default:
		return "", commonConstError.ErrorConvertFilter
	}
}

// GetIPv4PoolCollection Get IPv4 pool collection by start and count.
func (i *PoolDBImplement) GetIPv4PoolCollection(start int, count int, filter string) (*model.IPv4PoolCollection, error) {
	var (
		total      int
		collection []entity.IPv4Pool
		ret        = new(model.IPv4PoolCollection)
	)

	c := commonDB.GetConnection()
	c.Table("IPv4Pool").Count(&total)
	if where, err := i.convertFilter(filter); err != nil {
		log.WithFields(log.Fields{
			"filter": filter,
			"error":  err}).
			Warn("Get IPv4 pool in DB failed, convert filter failed.")
		c.Order("\"Name\" asc").Limit(count).Offset(start).Select([]string{"\"ID\"", "\"Name\""}).Find(&collection)
	} else {
		log.WithFields(log.Fields{"where": where}).Debug("Convert filter success.")
		c.Order("\"Name\" asc").Limit(count).Offset(start).Where(where).Select([]string{"\"ID\"", "\"Name\""}).Find(&collection)
	}
	ret.Start = start
	ret.Count = len(collection)
	ret.Total = total
	for _, v := range collection {
		ret.Members = append(ret.Members, model.IPv4PoolMember{
			ID:   v.ID,
			Name: v.Name,
		})
	}
	return ret, nil
}

// DeleteIPv4Pool delete the IPv4 pool by ID.
// It will return if the one exist.
// It will return the deleted one if commited.
// It will return wether the operation commited.
// It will return error if any.
func (i *PoolDBImplement) DeleteIPv4Pool(id string) (bool, *model.IPv4Pool, bool, error) {
	var record, previous entity.IPv4Pool

	if id == "" {
		return true, nil, false, commonConstError.ErrorIDFormat
	}
	c := commonDB.GetConnection()

	tx := c.Begin()
	if err := tx.Error; err != nil {
		log.WithFields(log.Fields{
			"id":    id,
			"error": err}).
			Warn("Delete IPv4 pool in DB failed, start transaction failed.")
		return true, nil, false, err
	}
	if tx.Where("\"ID\" = ?", id).First(&previous).RecordNotFound() {
		tx.Rollback()
		log.WithFields(log.Fields{
			"id": id}).
			Warn("Delete IPv4 pool in DB failed, resource does not exist, transaction rollback.")
		return false, nil, false, commonConstError.ErrorResourceNotExist
	}

	record.ID = id
	if err := tx.Delete(&record).Error; err != nil {
		log.WithFields(log.Fields{
			"id": id}).
			Warn("Delete IPv4 pool in DB failed, delete resource failed, transaction rollback.")
		return true, nil, false, err
	}
	if err := tx.Commit().Error; err != nil {
		log.WithFields(log.Fields{
			"id":    id,
			"error": err}).
			Warn("Delete IPv4 pool in DB failed, commit failed.")
		return true, nil, false, err
	}
	return true, previous.ToModel(), true, nil
}

// DeleteIPv4PoolCollection delete the IPv4 pool by ID.
// It will return the deleted resources.
// It will return wether the commit success.
// It will return error if any.
func (i *PoolDBImplement) DeleteIPv4PoolCollection() ([]model.IPv4Pool, bool, error) {
	var records []entity.IPv4Pool
	c := commonDB.GetConnection()

	tx := c.Begin()
	if err := tx.Error; err != nil {
		log.WithFields(log.Fields{
			"error": err}).
			Warn("Delete IPv4 pool collection in DB failed, start transaction failed.")
		return nil, false, err
	}
	if err := tx.Find(&records).Error; err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"error": err}).
			Warn("Delete IPv4 pool collection in DB failed, get the collection failed, transaction rollback.")
		return nil, false, err
	}
	if err := tx.Delete(entity.IPv4Pool{}).Error; err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"error": err}).
			Warn("Delete IPv4 pool collection in DB failed, delete collection failed, transaction rollback.")
		return nil, false, err
	}
	if err := tx.Commit().Error; err != nil {
		log.WithFields(log.Fields{
			"error": err}).
			Warn("Delete IPv4 pool in DB failed, commit failed.")
		return nil, false, err
	}
	var deleted = make([]model.IPv4Pool, 0)
	for _, v := range records {
		deleted = append(deleted, *v.ToModel())
	}
	return deleted, true, nil
}
