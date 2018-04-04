package db

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	commonDB "promise/common/db"
	commonConstError "promise/common/object/consterror"
	commonUtil "promise/common/util"
	"promise/pool/object/entity"
	"promise/pool/object/model"
	"strings"
)

var (
	instance PoolDBImplement
)

// PoolDBImplement is the implementation.
type PoolDBImplement struct {
}

// GetPoolDB return the singleton.
func GetPoolDB() PoolDBInterface {
	return &instance
}

// PostIPv4Pool will save the IPv4 pool if no one with the same name.
// It will return if there is one exist already with the same name.
// It will return the newly created one if commited, or nil.
// It will return if the transaction commited.
// It will return error if any.
func (impl *PoolDBImplement) PostIPv4Pool(m *model.IPv4Pool) (bool, *model.IPv4Pool, bool, error) {
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
	if err := tx.Save(&record).Error; err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"name":  m.Name,
			"error": err}).
			Warn("Post IPv4 pool in DB failed, save resource failed, transaction rollback.")
		return false, nil, false, err
	}
	if err := tx.Commit().Error; err != nil {
		log.WithFields(log.Fields{
			"name":  m.Name,
			"error": err}).
			Warn("Post IPv4 pool in DB failed, commit failed.")
		return false, nil, false, err
	}
	commonUtil.PrintJson(record.ToModel())
	return false, record.ToModel(), true, nil
}

// getOrRollback will try to get the IPv4 pool by ID, if not exist rollback.
// It will return if the resource been found.
// It will return error if any.
func (impl *PoolDBImplement) getIPv4Pool(tx *gorm.DB, id string, record *entity.IPv4Pool) (bool, error) {
	if tx.Where("\"ID\" = ?", id).First(record).RecordNotFound() {
		tx.Rollback()
		log.WithFields(log.Fields{
			"id": id}).
			Warn("Get IPv4 Pool failed, resource does not exist, transaction rollback.")
		return false, nil
	}

	if err := tx.Where("\"ID\" = ?", id).Preload("Ranges").Preload("Ranges.Addresses").First(record).Error; err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"id": id,
			"error": err}).
			Warn("Get IPv4 Pool failed, get resource failed, transaction rollback.")
		return true, err
	}
	return true, nil
}

// GetIPv4Pool get the IPv4 pool by ID.
func (impl *PoolDBImplement) GetIPv4Pool(id string) *model.IPv4Pool {
	var record entity.IPv4Pool
	c := commonDB.GetConnection()

	tx := c.Begin()
	if err := tx.Error; err != nil {
		log.WithFields(log.Fields{
			"id":  id,
			"error": err}).
			Warn("Post IPv4 pool in DB failed, start transaction failed.")
		return nil
	}

	if exist, err := impl.getIPv4Pool(tx, id, &record); exist && err == nil {
		return record.ToModel()
	}
	return nil
}

func (impl *PoolDBImplement) convertFilter(filter string) (string, error) {
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
func (impl *PoolDBImplement) GetIPv4PoolCollection(start int, count int, filter string) (*model.IPv4PoolCollection, error) {
	var (
		total      int
		collection []entity.IPv4Pool
		ret        = new(model.IPv4PoolCollection)
	)

	c := commonDB.GetConnection()
	c.Table("IPv4Pool").Count(&total)
	if where, err := impl.convertFilter(filter); err != nil {
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
func (impl *PoolDBImplement) DeleteIPv4Pool(id string) (bool, *model.IPv4Pool, bool, error) {
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

	if exist, err := impl.getIPv4Pool(tx, id, &record); err != nil || !exist {
		return exist, nil, false, err
	}

	for _, i := range record.Ranges {
		for _, j := range i.Addresses {
			if err := tx.Delete(&j).Error; err != nil {
				tx.Rollback()
				log.WithFields(log.Fields{
					"id": id}).
					Warn("Delete IPv4 pool in DB failed, delete Addresses failed, transaction rollback.")
			}
		}
		if err := tx.Delete(&i).Error; err != nil {
			tx.Rollback()
			log.WithFields(log.Fields{
				"id": id}).
				Warn("Delete IPv4 pool in DB failed, delete Ranges failed, transaction rollback.")
		}
	}
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
func (impl *PoolDBImplement) DeleteIPv4PoolCollection() ([]model.IPv4Pool, bool, error) {
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

// saveAndCommit will save the record and do commit.
func (impl *PoolDBImplement) saveAndCommit(tx *gorm.DB, record *entity.IPv4Pool) (bool, error) {
	if err := tx.Save(record).Error; err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"id":  record.ID,
			"error": err}).
			Warn("Save and commit operation failed, save failed, transaction rollback.")
		return false,err
	}
	if err := tx.Commit().Error; err != nil {
		log.WithFields(log.Fields{
			"id":  record.ID,
			"error": err}).
			Warn("Save and commit operation failed, commit failed.")
		return false, err
	}
	return true, nil
}

// AllocateIPv4Address will allocate IPv4 address from IP pool.
// It will return if the pool exist.
// It will return the address if operation commited, or nil
// It will return the pool after the allocation if operation commited, or nil
// It will return if the operation commited.
// It will return error if any, or nil.
func (impl *PoolDBImplement) AllocateIPv4Address(id string, key string) (bool, string, *model.IPv4Pool, bool, error) {
	var record entity.IPv4Pool
	if id == "" {
		return false, "", nil, false, commonConstError.ErrorIDFormat
	}
	c := commonDB.GetConnection()

	tx := c.Begin()
	if err := tx.Error; err != nil {
		log.WithFields(log.Fields{
			"id":    id,
			"key": key,
			"error": err}).
			Warn("Allocate IPv4 failed, start transaction failed.")
		return true, "", nil, false, err
	}
	exist, err := impl.getIPv4Pool(tx, id, &record)
	if !exist {
		return false, "", nil, false, commonConstError.ErrorResourceNotExist
	}
	if err != nil {
		return true, "", nil, false, err
	}

	foundKey := false
	// If try to find the address with the same key.
	if key != "" {		
		for i := range record.Ranges {
			for j := range record.Ranges[i].Addresses {
				if record.Ranges[i].Addresses[j].Key == key {
					foundKey = true
					if record.Ranges[i].Addresses[j].Allocated == true {
						log.WithFields(log.Fields{
							"id":  record.ID,
							"key": key,
							"address":record.Ranges[i].Addresses[j].Address}).
							Info("Allocate IPv4, found address with key but already allocated.")					
					} else {
						record.Ranges[i].Addresses[j].Allocated = true
						record.Ranges[i].Free--
						record.Ranges[i].Allocatable--
						if commited, err := impl.saveAndCommit(tx, &record); commited && err == nil {
							return true, record.Ranges[i].Addresses[j].Address, record.ToModel(), true, nil
						}
						return true, "", nil, false, err							
					}
					// found the address with the key, but in already allocated.
					break;
				}
			}
			if foundKey {
				break;
			}
		}
	}

	// if the key == nil, we don't have to find the address with the key.
	for i := range record.Ranges {
		if record.Ranges[i].Free > 0 {
			for j := range record.Ranges[i].Addresses {
				if record.Ranges[i].Addresses[j].Key != "" || record.Ranges[i].Addresses[j].Allocated == true {
					continue
				}
				record.Ranges[i].Addresses[j].Allocated = true
				record.Ranges[i].Addresses[j].Key = key
				record.Ranges[i].Free--
				record.Ranges[i].Allocatable--
				commited, err := impl.saveAndCommit(tx, &record)
				if commited && err == nil {
					return true, record.Ranges[i].Addresses[j].Address, record.ToModel(), true, nil	
				}
				return true, "", nil, commited, err										
			}
		}
	}
	// So no free address, try to use the allocatable address.
	for i := range record.Ranges {
		if record.Ranges[i].Allocatable > 0 {
			for j := range record.Ranges[i].Addresses {
				if record.Ranges[i].Addresses[j].Allocated == true {
					continue
				}
				record.Ranges[i].Addresses[j].Allocated = true
				record.Ranges[i].Addresses[j].Key = key
				record.Ranges[i].Free--
				record.Ranges[i].Allocatable--
				commited, err := impl.saveAndCommit(tx, &record)
				if commited && err == nil {
					return true, record.Ranges[i].Addresses[j].Address, record.ToModel(), true, nil	
				}
				return true, "", nil, commited, err				
			}
		}
	}
	// So no address can allocate.
	tx.Rollback()
	log.WithFields(log.Fields{
		"id":  id,
		"key": key}).
		Info("Allocate IPv4 failed, no allocatable address.")					

	return true, "", nil, false, nil
}