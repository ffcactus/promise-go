package db

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/pool/object/entity"
	"promise/pool/object/message"
)

// IPv4Pool is the concrete DB.
type IPv4Pool struct {
	base.DB
}

// ResourceName get the resource name.
func (impl *IPv4Pool) ResourceName() string {
	return "ipv4"
}

// NewEntity return the a new entity.
func (impl *IPv4Pool) NewEntity() base.EntityInterface {
	e := new(entity.IPv4Pool)
	e.Entity.TemplateImpl = e
	return e
}

// NewEntityCollection return a collection of entity.
func (impl *IPv4Pool) NewEntityCollection() interface{} {
	return new([]entity.IPv4Pool)
}

// GetConnection return the DB connection.
func (impl *IPv4Pool) GetConnection() *gorm.DB {
	return base.GetConnection()
}

// NeedCheckDuplication return if need check duplication for entity.
func (impl *IPv4Pool) NeedCheckDuplication() bool {
	return true
}

// ConvertFindResultToCollection convert the Find() result to collection mode.
func (impl *IPv4Pool) ConvertFindResultToCollection(start int64, total int64, result interface{}) (*base.CollectionModel, *base.Message) {
	collection, ok := result.(*[]entity.IPv4Pool)
	if !ok {
		log.Error("IPv4Pool.ConvertFindResult() failed, convert data failed.")
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
func (impl *IPv4Pool) ConvertFindResultToModel(result interface{}) ([]base.ModelInterface, *base.Message) {
	collection, ok := result.(*[]entity.IPv4Pool)
	if !ok {
		log.Error("IPv4Pool.ConvertFindResult() failed, convert data failed.")
		return nil, base.NewMessageInternalError()
	}
	ret := make([]base.ModelInterface, 0)
	for _, v := range *collection {
		m := v.ToModel()
		ret = append(ret, m)
	}
	return ret, nil
}

// AllocateIPv4Address will allocate IPv4 address from IP pool.
// It will return the address if operation commited, or nil
// It will return the pool after the allocation if operation commited, or nil
// It will return message if any error.
func (impl *IPv4Pool) AllocateIPv4Address(id string, key string) (string, base.ModelInterface, *base.Message) {
	var (
		c = impl.TemplateImpl.GetConnection()
	)
	record := new(entity.IPv4Pool)
	record.Entity.TemplateImpl = record

	tx := c.Begin()
	if err := tx.Error; err != nil {
		log.WithFields(log.Fields{
			"id":    id,
			"key":   key,
			"error": err,
		}).Warn("Allocate IPv4 failed, start transaction failed.")
		return "", nil, base.NewMessageTransactionError()
	}
	exist, err := impl.GetInternal(tx, id, record)
	if !exist {
		return "", nil, base.NewMessageNotExist()
	}
	if err != nil {
		return "", nil, base.NewMessageTransactionError()
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
							"id":      record.ID,
							"key":     key,
							"address": record.Ranges[i].Addresses[j].Address,
						}).Info("Allocate IPv4, found address with key but already allocated.")
					} else {
						record.Ranges[i].Addresses[j].Allocated = true
						if record.Ranges[i].Free > 0 {
							record.Ranges[i].Free--
						}
						record.Ranges[i].Allocatable--
						if commited, err := impl.SaveAndCommit(tx, record); commited && err == nil {
							return record.Ranges[i].Addresses[j].Address, record.ToModel(), nil
						}
						return "", nil, base.NewMessageTransactionError()
					}
					// found the address with the key, but in already allocated.
					break
				}
			}
			if foundKey {
				break
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
				if record.Ranges[i].Free > 0 {
					record.Ranges[i].Free--
				}
				record.Ranges[i].Allocatable--
				commited, err := impl.SaveAndCommit(tx, record)
				if commited && err == nil {
					return record.Ranges[i].Addresses[j].Address, record.ToModel(), nil
				}
				return "", nil, base.NewMessageTransactionError()
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
				if record.Ranges[i].Free > 0 {
					record.Ranges[i].Free--
				}
				record.Ranges[i].Allocatable--
				commited, err := impl.SaveAndCommit(tx, record)
				if commited && err == nil {
					return record.Ranges[i].Addresses[j].Address, record.ToModel(), nil
				}
				return  "", nil, base.NewMessageTransactionError()
			}
		}
	}
	// So no address can allocate.
	tx.Rollback()
	log.WithFields(log.Fields{
		"id":  id,
		"key": key,
	}).Info("Allocate IPv4 failed, no allocatable address.")

	return "", nil, message.NewMessageIPv4PoolEmpty()
}

// FreeIPv4Address will free the address to pool.
// It will return message if any error.
func (impl *IPv4Pool) FreeIPv4Address(id string, address string) (base.ModelInterface, *base.Message) {
	var (
		c = impl.TemplateImpl.GetConnection()
	)
	record := new(entity.IPv4Pool)
	record.Entity.TemplateImpl = record

	tx := c.Begin()
	if err := tx.Error; err != nil {
		log.WithFields(log.Fields{
			"id":      id,
			"address": address,
			"error":   err,
		}).Warn("Free IPv4 failed, start transaction failed.")
		return nil, base.NewMessageTransactionError()
	}
	exist, err := impl.GetInternal(tx, id, record)
	if !exist {
		return nil, base.NewMessageNotExist()
	}
	if err != nil {
		return nil, base.NewMessageTransactionError()
	}
	for i := range record.Ranges {
		if !base.IPStringBetween(record.Ranges[i].Start, record.Ranges[i].End, address) {
			continue
		}
		for j := range record.Ranges[i].Addresses {
			if record.Ranges[i].Addresses[j].Address == address {
				if !record.Ranges[i].Addresses[j].Allocated {
					tx.Rollback()
					log.WithFields(log.Fields{
						"id":      id,
						"address": address,
					}).Warn("Free IPv4 failed, the address didn't allocate, transaction rollback.")
					return nil, message.NewMessageIPv4NotAllocatedError()
				}
				record.Ranges[i].Addresses[j].Allocated = false
				record.Ranges[i].Allocatable++
				if record.Ranges[i].Addresses[j].Key == "" {
					record.Ranges[i].Free++
				}
				commited, err := impl.SaveAndCommit(tx, record)
				if commited && err == nil {
					return record.ToModel(), nil
				}
				return nil, base.NewMessageTransactionError()
			}
		}
		break
	}
	// Can't find the address in pool.
	tx.Rollback()
	return nil, message.NewMessageIPv4AddressNotExistError()
}
