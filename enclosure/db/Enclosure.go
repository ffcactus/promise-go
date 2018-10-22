package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/enclosure/object/entity"
	"promise/enclosure/object/model"
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

// Exist returns if the enclosure already exist in the DB.
// If it exists, return it.
func (impl *Enclosure) Exist(e *model.Enclosure) (bool, base.ModelInterface) {
	return false, nil
}

// GetAndLock will try to lock the enclosure by ID. Please note that it is not enought to just check the error,
// you need check the state of the returned enclosure to see if it has been locked.
// to see if it is locked successfully.
// 
// The first return value is the enclosure when everything works fine or nil if failed to get and lock enclosure.
// The second and third value are the previous state and statereason respectively.
// The fourh return value indicates if any error happened.
// If the enclosure does not exist, return (nil, "", "", nil).
// If the enclosure can't be locked, return the (enclosure, state, reason, nil).
// Note that in this case the return value is the same to return of successfully locked.
// For any DB operation error, return (nil, "", "", error).
func (impl *Enclosure) GetAndLock(ID string) (base.ModelInterface, string, string, error) {
	var (
		c         = impl.TemplateImpl.GetConnection()
		enclosure = new(entity.Enclosure)
		rollback  = false
		state, reason string
	)

	// Transaction start.
	tx := c.Begin()
	if err := tx.Error; err != nil {
		log.WithFields(log.Fields{
			"id":    ID,
			"error": err}).
			Warn("DB get and lock enclosure failed, start transaction failed.")
		return nil, "", "", err
	}

	defer func() {
		if rollback {
			tx.Rollback()
			log.WithFields(log.Fields{
				"id": ID,
			}).Warn("DB get and lock enclosure failed, transaction roll back.")
		}
	}()

	if err := tx.Exec("SET TRANSACTION ISOLATION LEVEL SERIALIZABLE;").Error; err != nil {
		rollback = true
		log.WithFields(log.Fields{
			"id":    ID,
			"error": err,
		}).Warn("DB get and lock enclosure failed, set transaction isolation level to serializable failed.")
		return nil, "", "", err
	}
	if tx.Where("\"ID\" = ?", ID).First(enclosure).RecordNotFound() {
		rollback = true
		log.WithFields(log.Fields{
			"id": ID,
		}).Warn("DB get and lock enclosure failed, enclosure does not exist.")
		return nil, "", "", nil
	}
	state = enclosure.State
	reason = enclosure.StateReason
	if !model.EnclosureLockable(enclosure.State) {
		// Server not ready, rollback.
		rollback = true
		log.WithFields(log.Fields{
			"id":    ID,
			"state": enclosure.State,
		}).Warn("DB get and lock enclosure failed, enclosure not lockable.")
		return enclosure.ToModel(), state, reason, fmt.Errorf("unlockable state")
	}
	// Change the state.
	if err := tx.Model(enclosure).UpdateColumn("State", model.StateLocked).Error; err != nil {
		rollback = true
		log.WithFields(log.Fields{
			"id":    ID,
			"state": enclosure.State,
		}).Warn("DB get and lock enclosure failed, update state failed.")
		return nil, "", "", err
	}
	found, err := impl.GetInternal(tx, ID, enclosure)
	if !found || err != nil {
		rollback = true
		log.WithFields(log.Fields{
			"id":    ID,
			"error": err,
		}).Warn("DB get and lock enclosure failed, load enclosure failed.")
		return nil, "", "", err
	}
	// Commit.
	if err := tx.Commit().Error; err != nil {
		log.WithFields(log.Fields{
			"id":    ID,
			"error": err,
		}).Warn("DB get and lock enclosure failed, commit failed.")
		return nil, "", "", err
	}
	log.WithFields(log.Fields{
		"id":    ID,
		"state": enclosure.State,
	}).Info("DB get and lock enclosure success.")
	return enclosure.ToModel(), state, reason, nil
}

// SetState sets the state and state reason to the enclosure specified by ID.
// On success, return the enclosure with the new state and state reason.
// If the enclosure not exist, return nil and nil.
// For other DB operation error , return nil and error.
func (impl *Enclosure) SetState(ID, state, reason string) (base.ModelInterface, error) {
	var (
		c         = impl.TemplateImpl.GetConnection()
		enclosure = new(entity.Enclosure)
		rollback  = false
	)
	// Use transaction for the enclosure may be removed before update the state.
	tx := c.Begin()
	if err := tx.Error; err != nil {
		log.WithFields(log.Fields{
			"id":    ID,
			"op":    "SetState",
			"error": err,
		}).Warn("DB operation failed, start transaction failed.")
		return nil, base.ErrorTransaction
	}

	defer func() {
		if rollback {
			tx.Rollback()
		}
	}()

	found, err := impl.GetInternal(tx, ID, enclosure)
	if !found || err != nil {
		rollback = true
		log.WithFields(log.Fields{
			"id": ID,
			"op": "SetState",
		}).Warn("DB operation failed, load enclosure failed.")
		return nil, base.ErrorTransaction
	}
	if err := tx.Model(enclosure).UpdateColumn(entity.Enclosure{State: state, StateReason: reason}).Error; err != nil {
		rollback = true
		log.WithFields(log.Fields{
			"id":    ID,
			"op":    "SetState",
			"error": err,
		}).Warn("DB opertion failed, update enclosure failed, transaction rollback.")
		return nil, base.ErrorTransaction
	}
	if err := tx.Commit().Error; err != nil {
		log.WithFields(log.Fields{
			"id":    ID,
			"op":    "SetState",
			"error": err,
		}).Warn("DB opertion failed, commit failed.")
		return nil, base.ErrorTransaction
	}
	return enclosure.ToModel(), nil
}

// RefreshManagerSlot refreshes the manager slots to the enclosure given by id in the DB.
func (impl *Enclosure) RefreshManagerSlot(id string, slots []model.ManagerSlot) (base.ModelInterface, error) {
	var (
		c         = impl.TemplateImpl.GetConnection()
		enclosure entity.Enclosure
		tx        *gorm.DB
		rollback  = false
	)
	tx = c.Begin()
	if err := tx.Error; err != nil {
		log.WithFields(log.Fields{
			"id":    id,
			"op":    "RefreshManagerSlot",
			"error": err,
		}).Warn("DB operation failed, start transaction failed.")
		return nil, err
	}

	defer func() {
		if rollback {
			tx.Rollback()
		}
	}()

	found, err := impl.GetInternal(tx, id, &enclosure)
	if !found || err != nil {
		rollback = true
		log.WithFields(log.Fields{
			"id": id,
			"op": "RefreshManagerSlot",
		}).Warn("DB operation failed, load enclosure failed.")
		// TODO we can't make sure this is a transaction error.
		// Should GetInternal just return error?
		return nil, base.ErrorTransaction
	}

	if err := impl.deleteManagerSlot(tx, &enclosure); err != nil {
		rollback = true
		log.WithFields(log.Fields{
			"id":    id,
			"op":    "RefreshManagerSlot",
			"error": err,
		}).Warn("DB operation failed, delete association failed, transaction rollback.")
		return nil, err
	}
	enclosure.ManagerSlots = []entity.ManagerSlot{}
	for _, v := range slots {
		each := entity.ManagerSlot{}
		each.Load(&v)
		enclosure.ManagerSlots = append(enclosure.ManagerSlots, each)
	}
	if err := tx.Save(&enclosure).Error; err != nil {
		rollback = true
		log.WithFields(log.Fields{
			"id":    id,
			"op":    "RefreshManagerSlot",
			"error": err,
		}).Warn("DB opertion failed, save enclosure failed, transaction rollback.")
		return nil, err
	}
	if err := tx.Commit().Error; err != nil {
		log.WithFields(log.Fields{
			"id":    id,
			"op":    "RefreshManagerSlot",
			"error": err,
		}).Warn("DB opertion failed, commit failed.")
		return nil, err
	}
	return enclosure.ToModel(), nil
}

func (impl *Enclosure) deleteManagerSlot(c *gorm.DB, enclosure *entity.Enclosure) error {
	for i := range enclosure.ManagerSlots {
		if err := c.Delete(enclosure.ManagerSlots[i]).Error; err != nil {
			return err
		}
	}
	return nil
}

// RefreshServerSlot refreshes the server slots to the enclosure given by id in the DB.
func (impl *Enclosure) RefreshServerSlot(id string, slots []model.ServerSlot) (base.ModelInterface, error) {
	var (
		c         = impl.TemplateImpl.GetConnection()
		enclosure entity.Enclosure
		tx        *gorm.DB
		rollback  = false
	)
	tx = c.Begin()
	if err := tx.Error; err != nil {
		log.WithFields(log.Fields{
			"id":    id,
			"op":    "RefreshServerSlot",
			"error": err,
		}).Warn("DB operation failed, start transaction failed.")
		return nil, err
	}

	defer func() {
		if rollback {
			tx.Rollback()
		}
	}()

	found, err := impl.GetInternal(tx, id, &enclosure)
	if !found || err != nil {
		rollback = true
		log.WithFields(log.Fields{
			"id": id,
			"op": "RefreshServerSlot",
		}).Warn("DB operation failed, load enclosure failed.")
		// TODO we can't make sure this is a transaction error.
		// Should GetInternal just return error?
		return nil, base.ErrorTransaction
	}

	if err := impl.deleteServerSlot(tx, &enclosure); err != nil {
		rollback = true
		log.WithFields(log.Fields{
			"id":    id,
			"op":    "RefreshServerSlot",
			"error": err,
		}).Warn("DB operation failed, delete association failed, transaction rollback.")
		return nil, err
	}
	enclosure.ServerSlots = []entity.ServerSlot{}
	for _, v := range slots {
		each := entity.ServerSlot{}
		each.Load(&v)
		enclosure.ServerSlots = append(enclosure.ServerSlots, each)
	}
	if err := tx.Save(&enclosure).Error; err != nil {
		rollback = true
		log.WithFields(log.Fields{
			"id":    id,
			"op":    "RefreshServerSlot",
			"error": err,
		}).Warn("DB opertion failed, save enclosure failed, transaction rollback.")
		return nil, err
	}
	if err := tx.Commit().Error; err != nil {
		log.WithFields(log.Fields{
			"id":    id,
			"op":    "RefreshServerSlot",
			"error": err,
		}).Warn("DB opertion failed, commit failed.")
		return nil, err
	}
	return enclosure.ToModel(), nil
}

func (impl *Enclosure) deleteServerSlot(c *gorm.DB, enclosure *entity.Enclosure) error {
	for i := range enclosure.ServerSlots {
		if err := c.Delete(enclosure.ServerSlots[i]).Error; err != nil {
			return err
		}
	}
	return nil
}

// RefreshSwitchSlot refreshes the switch slots to the enclosure given by id in the DB.
func (impl *Enclosure) RefreshSwitchSlot(id string, slots []model.SwitchSlot) (base.ModelInterface, error) {
	var (
		c         = impl.TemplateImpl.GetConnection()
		enclosure entity.Enclosure
		tx        *gorm.DB
		rollback  = false
	)
	tx = c.Begin()
	if err := tx.Error; err != nil {
		log.WithFields(log.Fields{
			"id":    id,
			"op":    "RefreshSwitchSlot",
			"error": err,
		}).Warn("DB operation failed, start transaction failed.")
		return nil, err
	}

	defer func() {
		if rollback {
			tx.Rollback()
		}
	}()

	found, err := impl.GetInternal(tx, id, &enclosure)
	if !found || err != nil {
		rollback = true
		log.WithFields(log.Fields{
			"id": id,
			"op": "RefreshSwitchSlot",
		}).Warn("DB operation failed, load enclosure failed.")
		// TODO we can't make sure this is a transaction error.
		// Should GetInternal just return error?
		return nil, base.ErrorTransaction
	}

	if err := impl.deleteSwitchSlot(tx, &enclosure); err != nil {
		rollback = true
		log.WithFields(log.Fields{
			"id":    id,
			"op":    "RefreshSwitchSlot",
			"error": err,
		}).Warn("DB operation failed, delete association failed, transaction rollback.")
		return nil, err
	}
	enclosure.SwitchSlots = []entity.SwitchSlot{}
	for _, v := range slots {
		each := entity.SwitchSlot{}
		each.Load(&v)
		enclosure.SwitchSlots = append(enclosure.SwitchSlots, each)
	}
	if err := tx.Save(&enclosure).Error; err != nil {
		rollback = true
		log.WithFields(log.Fields{
			"id":    id,
			"op":    "RefreshSwitchSlot",
			"error": err,
		}).Warn("DB opertion failed, save enclosure failed, transaction rollback.")
		return nil, err
	}
	if err := tx.Commit().Error; err != nil {
		log.WithFields(log.Fields{
			"id":    id,
			"op":    "RefreshSwitchSlot",
			"error": err,
		}).Warn("DB opertion failed, commit failed.")
		return nil, err
	}
	return enclosure.ToModel(), nil
}

func (impl *Enclosure) deleteSwitchSlot(c *gorm.DB, enclosure *entity.Enclosure) error {
	for i := range enclosure.SwitchSlots {
		if err := c.Delete(enclosure.SwitchSlots[i]).Error; err != nil {
			return err
		}
	}
	return nil
}

// RefreshApplianceSlot refreshes the appliance slots to the enclosure given by id in the DB.
func (impl *Enclosure) RefreshApplianceSlot(id string, slots []model.ApplianceSlot) (base.ModelInterface, error) {
	var (
		c         = impl.TemplateImpl.GetConnection()
		enclosure entity.Enclosure
		tx        *gorm.DB
		rollback  = false
	)
	tx = c.Begin()
	if err := tx.Error; err != nil {
		log.WithFields(log.Fields{
			"id":    id,
			"op":    "RefreshApplianceSlot",
			"error": err,
		}).Warn("DB operation failed, start transaction failed.")
		return nil, err
	}

	defer func() {
		if rollback {
			tx.Rollback()
		}
	}()

	found, err := impl.GetInternal(tx, id, &enclosure)
	if !found || err != nil {
		rollback = true
		log.WithFields(log.Fields{
			"id": id,
			"op": "RefreshApplianceSlot",
		}).Warn("DB operation failed, load enclosure failed.")
		// TODO we can't make sure this is a transaction error.
		// Should GetInternal just return error?
		return nil, base.ErrorTransaction
	}

	if err := impl.deleteApplianceSlot(tx, &enclosure); err != nil {
		rollback = true
		log.WithFields(log.Fields{
			"id":    id,
			"op":    "RefreshApplianceSlot",
			"error": err,
		}).Warn("DB operation failed, delete association failed, transaction rollback.")
		return nil, err
	}
	enclosure.ApplianceSlots = []entity.ApplianceSlot{}
	for _, v := range slots {
		each := entity.ApplianceSlot{}
		each.Load(&v)
		enclosure.ApplianceSlots = append(enclosure.ApplianceSlots, each)
	}
	if err := tx.Save(&enclosure).Error; err != nil {
		rollback = true
		log.WithFields(log.Fields{
			"id":    id,
			"op":    "RefreshApplianceSlot",
			"error": err,
		}).Warn("DB opertion failed, save enclosure failed, transaction rollback.")
		return nil, err
	}
	if err := tx.Commit().Error; err != nil {
		log.WithFields(log.Fields{
			"id":    id,
			"op":    "RefreshApplianceSlot",
			"error": err,
		}).Warn("DB opertion failed, commit failed.")
		return nil, err
	}
	return enclosure.ToModel(), nil
}

func (impl *Enclosure) deleteApplianceSlot(c *gorm.DB, enclosure *entity.Enclosure) error {
	for i := range enclosure.ApplianceSlots {
		if err := c.Delete(enclosure.ApplianceSlots[i]).Error; err != nil {
			return err
		}
	}
	return nil
}

// RefreshPowerSlot refreshes the power slots to the enclosure given by id in the DB.
func (impl *Enclosure) RefreshPowerSlot(id string, slots []model.PowerSlot) (base.ModelInterface, error) {
	var (
		c         = impl.TemplateImpl.GetConnection()
		enclosure entity.Enclosure
		tx        *gorm.DB
		rollback  = false
	)
	tx = c.Begin()
	if err := tx.Error; err != nil {
		log.WithFields(log.Fields{
			"id":    id,
			"op":    "RefreshPowerSlot",
			"error": err,
		}).Warn("DB operation failed, start transaction failed.")
		return nil, err
	}

	defer func() {
		if rollback {
			tx.Rollback()
		}
	}()

	found, err := impl.GetInternal(tx, id, &enclosure)
	if !found || err != nil {
		rollback = true
		log.WithFields(log.Fields{
			"id": id,
			"op": "RefreshPowerSlot",
		}).Warn("DB operation failed, load enclosure failed.")
		// TODO we can't make sure this is a transaction error.
		// Should GetInternal just return error?
		return nil, base.ErrorTransaction
	}

	if err := impl.deletePowerSlot(tx, &enclosure); err != nil {
		rollback = true
		log.WithFields(log.Fields{
			"id":    id,
			"op":    "RefreshPowerSlot",
			"error": err,
		}).Warn("DB operation failed, delete association failed, transaction rollback.")
		return nil, err
	}
	enclosure.PowerSlots = []entity.PowerSlot{}
	for _, v := range slots {
		each := entity.PowerSlot{}
		each.Load(&v)
		enclosure.PowerSlots = append(enclosure.PowerSlots, each)
	}
	if err := tx.Save(&enclosure).Error; err != nil {
		rollback = true
		log.WithFields(log.Fields{
			"id":    id,
			"op":    "RefreshPowerSlot",
			"error": err,
		}).Warn("DB opertion failed, save enclosure failed, transaction rollback.")
		return nil, err
	}
	if err := tx.Commit().Error; err != nil {
		log.WithFields(log.Fields{
			"id":    id,
			"op":    "RefreshPowerSlot",
			"error": err,
		}).Warn("DB opertion failed, commit failed.")
		return nil, err
	}
	return enclosure.ToModel(), nil
}

func (impl *Enclosure) deletePowerSlot(c *gorm.DB, enclosure *entity.Enclosure) error {
	for i := range enclosure.PowerSlots {
		if err := c.Delete(enclosure.PowerSlots[i]).Error; err != nil {
			return err
		}
	}
	return nil
}

// RefreshFanSlot refreshes the fan slots to the enclosure given by id in the DB.
func (impl *Enclosure) RefreshFanSlot(id string, slots []model.FanSlot) (base.ModelInterface, error) {
	var (
		c         = impl.TemplateImpl.GetConnection()
		enclosure entity.Enclosure
		tx        *gorm.DB
		rollback  = false
	)
	tx = c.Begin()
	if err := tx.Error; err != nil {
		log.WithFields(log.Fields{
			"id":    id,
			"op":    "RefreshFanSlot",
			"error": err,
		}).Warn("DB operation failed, start transaction failed.")
		return nil, err
	}

	defer func() {
		if rollback {
			tx.Rollback()
		}
	}()

	found, err := impl.GetInternal(tx, id, &enclosure)
	if !found || err != nil {
		rollback = true
		log.WithFields(log.Fields{
			"id": id,
			"op": "RefreshFanSlot",
		}).Warn("DB operation failed, load enclosure failed.")
		// TODO we can't make sure this is a transaction error.
		// Should GetInternal just return error?
		return nil, base.ErrorTransaction
	}

	if err := impl.deleteFanSlot(tx, &enclosure); err != nil {
		rollback = true
		log.WithFields(log.Fields{
			"id":    id,
			"op":    "RefreshFanSlot",
			"error": err,
		}).Warn("DB operation failed, delete association failed, transaction rollback.")
		return nil, err
	}
	enclosure.FanSlots = []entity.FanSlot{}
	for _, v := range slots {
		each := entity.FanSlot{}
		each.Load(&v)
		enclosure.FanSlots = append(enclosure.FanSlots, each)
	}
	if err := tx.Save(&enclosure).Error; err != nil {
		rollback = true
		log.WithFields(log.Fields{
			"id":    id,
			"op":    "RefreshFanSlot",
			"error": err,
		}).Warn("DB opertion failed, save enclosure failed, transaction rollback.")
		return nil, err
	}
	if err := tx.Commit().Error; err != nil {
		log.WithFields(log.Fields{
			"id":    id,
			"op":    "RefreshFanSlot",
			"error": err,
		}).Warn("DB opertion failed, commit failed.")
		return nil, err
	}
	return enclosure.ToModel(), nil
}

func (impl *Enclosure) deleteFanSlot(c *gorm.DB, enclosure *entity.Enclosure) error {
	for i := range enclosure.FanSlots {
		if err := c.Delete(enclosure.FanSlots[i]).Error; err != nil {
			return err
		}
	}
	return nil
}
