package base

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/sirupsen/logrus"
	"strings"
)

// DBTemplateInterface is the interface that a concrete DB implement should have.
type DBTemplateInterface interface {
	ResourceName() string
	NeedCheckDuplication() bool
	GetConnection() *gorm.DB
	NewEntity() EntityInterface
	NewEntityCollection() interface{}
	ConvertFindResultToCollection(start int64, total int64, result interface{}) (*CollectionModel, *ErrorResponse)
	ConvertFindResultToModel(interface{}) ([]ModelInterface, *ErrorResponse)
}

// DBInterface is the interface that DB should have.
type DBInterface interface {
	Create(ModelInterface) (ModelInterface, *ErrorResponse)
	Get(id string) (ModelInterface, *ErrorResponse)
	Update(id string, request UpdateRequestInterface) (ModelInterface, *ErrorResponse)
	Delete(id string) (ModelInterface, *ErrorResponse)
	GetCollection(start int64, count int64, filter string) (*CollectionModel, *ErrorResponse)
	DeleteCollection() ([]ModelInterface, *ErrorResponse)
}

// DB is the DB implementation in Promise project.
type DB struct {
	TemplateImpl DBTemplateInterface
}

// GetInternal is part of the process to get resource in DB, since many other operation
// need this process, we seperate it out.
// It will return if the resource been found.
// It will return error if any.
func (impl *DB) GetInternal(tx *gorm.DB, id string, record EntityInterface) (bool, error) {
	var (
		name = impl.TemplateImpl.ResourceName()
	)

	preload := record.Preload()
	if tx.Where("\"ID\" = ?", id).First(record).RecordNotFound() {
		tx.Rollback()
		log.WithFields(log.Fields{
			"resource": name,
			"id":       id,
		}).Warn("DB get resource failed, resource does not exist, transaction rollback.")
		return false, ErrorResourceNotExist
	}
	if err := tx.Error; err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"resource": name,
			"id":       id,
		}).Warn("DB get resource failed, find first record failed, transaction rollback.")
		return false, ErrorResourceNotExist
	}

	tx.Where("\"ID\" = ?", id)
	for _, v := range preload {
		tx = tx.Preload(v)
	}
	if err := tx.First(record).Error; err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"resource": name,
			"id":       id,
			"error":    err,
		}).Warn("DB get resource failed, fatch failed, transaction rollback.")
		return true, err
	}
	return true, nil
}

// SaveAndCommit will save the record and do commit.
// It will return if the operation commited.
// It will return error if any.
func (impl *DB) SaveAndCommit(tx *gorm.DB, record EntityInterface) (bool, error) {
	var (
		name = impl.TemplateImpl.ResourceName()
	)

	if err := tx.Save(record).Error; err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"resource": name,
			"id":       record.GetID(),
			"error":    err,
		}).Warn("DB save and commit operation failed, save failed, transaction rollback.")
		return false, err
	}
	if err := tx.Commit().Error; err != nil {
		log.WithFields(log.Fields{
			"resource": name,
			"id":       record.GetID(),
			"error":    err,
		}).Warn("DB save and commit operation failed, commit failed.")
		return false, err
	}
	return true, nil
}

// Create is the default implement to post resource in DB.
// It will return the newly created one if commited, or nil.
// It will return errorResp if any error.
func (impl *DB) Create(m ModelInterface) (ModelInterface, *ErrorResponse) {
	var (
		name   = impl.TemplateImpl.ResourceName()
		record = impl.TemplateImpl.NewEntity()
		c      = impl.TemplateImpl.GetConnection()
	)
	tx := c.Begin()
	if err := tx.Error; err != nil {
		log.WithFields(log.Fields{
			"resource": name,
			"error":    err,
		}).Warn("DB create resource failed, start transaction failed.")
		return nil, NewErrorResponseTransactionError()
	}
	if impl.TemplateImpl.NeedCheckDuplication() {
		where := "\"" + record.PropertyNameForDuplicationCheck() + "\" = ?"
		if !tx.Where(where, m.ValueForDuplicationCheck()).First(record).RecordNotFound() {
			tx.Rollback()
			log.WithFields(log.Fields{
				"resource": name,
				"existed":  record.GetID(),
				"name":     record,
			}).Warn("DB create resource failed, duplicated resource, transaction rollback.")
			return nil, NewErrorResponseDuplicate()
		}
	}

	record.Load(m)
	record.SetID(uuid.New().String())
	if err := tx.Create(record).Error; err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"resource": name,
			"name":     m,
			"error":    err,
		}).Warn("DB create resource failed, create resource failed, transaction rollback.")
		return nil, NewErrorResponseTransactionError()
	}
	if err := tx.Commit().Error; err != nil {
		log.WithFields(log.Fields{
			"resource": name,
			"name":     m,
			"error":    err,
		}).Warn("DB create resource failed, commit failed.")
		return nil, NewErrorResponseTransactionError()
	}
	return record.ToModel(), nil
}

// Get is the default implement to get resource in DB.
// If the resource does not exist in the DB return nil.
func (impl *DB) Get(id string) (ModelInterface, *ErrorResponse) {
	var (
		name   = impl.TemplateImpl.ResourceName()
		record = impl.TemplateImpl.NewEntity()
		c      = impl.TemplateImpl.GetConnection()
	)
	tx := c.Begin()
	if err := tx.Error; err != nil {
		log.WithFields(log.Fields{
			"resource": name,
			"id":       id,
			"error":    err,
		}).Warn("DB get resource failed, start transaction failed.")
		return nil, NewErrorResponseTransactionError()
	}
	exist, err := impl.GetInternal(tx, id, record)
	if !exist {
		return nil, NewErrorResponseNotExist()
	}
	if err != nil {
		return nil, NewErrorResponseTransactionError()
	}
	if err := tx.Commit().Error; err != nil {
		log.WithFields(log.Fields{
			"resource": name,
			"id":       id,
			"error":    err,
		}).Warn("DB get resource failed, commit transaction failed.")
		return nil, NewErrorResponseTransactionError()
	}
	return record.ToModel(), nil
}

// Update is the default implement to update resource in DB.
// It will return the updated resource.
// It will return errorResp if any error.
func (impl *DB) Update(id string, request UpdateRequestInterface) (ModelInterface, *ErrorResponse) {
	var (
		name   = impl.TemplateImpl.ResourceName()
		record = impl.TemplateImpl.NewEntity()
		c      = impl.TemplateImpl.GetConnection()
	)
	tx := c.Begin()
	if err := tx.Error; err != nil {
		log.WithFields(log.Fields{
			"resource": name,
			"id":       id,
			"error":    err,
		}).Warn("Update resource in DB failed, start transaction failed.")
		return nil, NewErrorResponseTransactionError()
	}
	exist, err := impl.GetInternal(tx, id, record)
	if !exist {
		return nil, NewErrorResponseNotExist()
	}
	if err != nil {
		return nil, NewErrorResponseTransactionError()
	}

	m := record.ToModel()
	if err := request.UpdateModel(m); err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"resource": name,
			"id":       id,
			"error":    err,
		}).Warn("Update resource in DB failed, update model failed, transaction rollback.")
		return nil, NewErrorResponseUnknownPropertyValue()
	}
	record.Load(m)
	record.SetID(id)
	commited, err := impl.SaveAndCommit(tx, record)
	if err != nil || !commited {
		return nil, NewErrorResponseTransactionError()
	}
	return record.ToModel(), nil
}

// Delete is the default implement to delete resource from DB.
// It will return the deleted one if commited.
// It will return errorResp if any error.
func (impl *DB) Delete(id string) (ModelInterface, *ErrorResponse) {
	var (
		name     = impl.TemplateImpl.ResourceName()
		record   = impl.TemplateImpl.NewEntity()
		previous = impl.TemplateImpl.NewEntity()
		c        = impl.TemplateImpl.GetConnection()
	)

	tx := c.Begin()
	if err := tx.Error; err != nil {
		log.WithFields(log.Fields{
			"resource": name,
			"id":       id,
			"error":    err,
		}).Warn("DB delete resource failed, start transaction failed.")
		return nil, NewErrorResponseTransactionError()
	}
	exist, err := impl.GetInternal(tx, id, previous)
	// Rollback in GetInternal.
	if !exist {
		return nil, NewErrorResponseNotExist()
	}
	if err != nil {
		return nil, NewErrorResponseTransactionError()
	}

	record.SetID(id)
	for _, v := range record.Association() {
		if err := tx.Delete(v).Error; err != nil {
			tx.Rollback()
			log.WithFields(log.Fields{
				"resource": name,
				"id":       id,
				"error":    err,
			}).Warn("DB delete resource failed, delete association failed, transaction rollback.")
			return nil, NewErrorResponseTransactionError()
		}
	}
	if err := tx.Delete(record).Error; err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"resource": name,
			"id":       id,
			"error":    err,
		}).Warn("DB delete resource failed, delete resource failed, transaction rollback.")
		return nil, NewErrorResponseTransactionError()
	}
	if err := tx.Commit().Error; err != nil {
		log.WithFields(log.Fields{
			"resource": name,
			"id":       id,
			"error":    err,
		}).Warn("DB delete resource failed, commit failed.")
		return nil, NewErrorResponseTransactionError()
	}
	return previous.ToModel(), nil
}

func (impl *DB) convertFilter(filter string, filterNames []string) (string, error) {
	if filter == "" {
		return "", nil
	}
	cmds := strings.Split(filter, " ")
	if len(cmds) != 3 {
		return "", ErrorConvertFilter
	}
	for _, v := range filterNames {
		if cmds[0] == v {
			switch strings.ToLower(cmds[1]) {
			case "eq":
				return "\"" + cmds[0] + "\"" + " = " + cmds[2], nil
			default:
				return "", ErrorConvertFilter
			}
		}
	}
	return "", ErrorUnknownFilterName

}

// GetCollection get the collection in DB.
// It returns nil if any error.
func (impl *DB) GetCollection(start int64, count int64, filter string) (*CollectionModel, *ErrorResponse) {
	var (
		name             = impl.TemplateImpl.ResourceName()
		total            int64
		record           = impl.TemplateImpl.NewEntity()
		recordCollection = impl.TemplateImpl.NewEntityCollection()
		c                = impl.TemplateImpl.GetConnection()
	)

	// Check filter first to avoid starting a transaction.
	where, err := impl.convertFilter(filter, record.FilterNameList())
	if err != nil {
		log.WithFields(log.Fields{
			"resource": name,
			"filter":   filter,
			"error":    err,
		}).Warn("DB get collection failed, convert filter failed.")
		return nil, NewErrorResponseUnknownFilterName()
	}
	log.WithFields(log.Fields{
		"resource": name,
		"where":    where,
	}).Debug("DB convert filter success.")

	// We need transaction to ensure the total and the query count is consistent.
	tx := c.Begin()
	if err := tx.Error; err != nil {
		log.WithFields(log.Fields{
			"resource": name,
			"error":    err,
		}).Warn("DB get collection failed, start transaction failed.")
		return nil, NewErrorResponseTransactionError()
	}

	// Get total count.
	if err := tx.Table(record.TableName()).Count(&total).Error; err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"resource": name,
			"error":    err,
		}).Warn("DB get collection failed, get count failed.")
		return nil, NewErrorResponseTransactionError()
	}

	// Find all the matches.
	if err := tx.Limit(count).Offset(start).Where(where).Find(recordCollection).Error; err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"resource": name,
			"error":    err,
		}).Warn("DB get collection failed, find resource failed.")
		return nil, NewErrorResponseTransactionError()
	}
	if err := tx.Commit().Error; err != nil {
		log.WithFields(log.Fields{
			"resource": name,
			"error":    err,
		}).Warn("DB get collection failed, commit failed.")
		return nil, NewErrorResponseTransactionError()
	}
	return impl.TemplateImpl.ConvertFindResultToCollection(start, total, recordCollection)
}

// DeleteCollection delete the collection in DB.
// It returns all the deleted resources,
// It returnes if committed.
// It returens error if any.
func (impl *DB) DeleteCollection() ([]ModelInterface, *ErrorResponse) {
	var (
		name             = impl.TemplateImpl.ResourceName()
		recordCollection = impl.TemplateImpl.NewEntityCollection()
		c                = impl.TemplateImpl.GetConnection()
		tables           = impl.TemplateImpl.NewEntity().Tables()
	)

	// We need transaction to ensure the total and the query count is consistent.
	tx := c.Begin()
	if err := tx.Error; err != nil {
		log.WithFields(log.Fields{
			"resource": name,
			"error":    err,
		}).Warn("DB delete collection failed, start transaction failed.")
		return nil, NewErrorResponseTransactionError()
	}

	if err := tx.Find(recordCollection).Error; err != nil {
		log.WithFields(log.Fields{
			"resource": name,
			"error":    err,
		}).Warn("DB delete collection failed, find resource failed.")
		return nil, NewErrorResponseTransactionError()
	}
	for _, v := range tables {
		if err := tx.Delete(v).Error; err != nil {
			tx.Rollback()
			log.WithFields(log.Fields{
				"resource": name,
				"error":    err,
			}).Warn("DB delete collection failed, delete resources failed, transaction rollback.")
			return nil, NewErrorResponseTransactionError()
		}
	}
	ret, errorResp := impl.TemplateImpl.ConvertFindResultToModel(recordCollection)
	if errorResp != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"resource": name,
			"errorResp":  errorResp.ID,
		}).Warn("DB delete collection failed, convert find result failed, transaction rollback.")
		return nil, errorResp
	}
	if err := tx.Commit().Error; err != nil {
		log.WithFields(log.Fields{
			"resource": name,
			"error":    err,
		}).Warn("DB delete collection failed, commit failed.")
		return nil, NewErrorResponseTransactionError()
	}
	return ret, nil
}

// TableInfo The tables in DB.
type TableInfo struct {
	Name string
	Info interface{}
}

var connection *gorm.DB

// InitConnection Init the DB connection. Each service have to call the method first.
func InitConnection() error {
	if connection == nil {
		log.Info("Init DB connection.")
		args := "host=db port=5432 user=postgres dbname=promise sslmode=disable password=iforgot"
		db, err := gorm.Open("postgres", args)
		// args := "host=100.100.194.103 port=5432 user=gaussdba dbname=NETADAPTOR sslmode=disable password=Huawei12#$"
		// db, err := gorm.Open("gauss", args)
		if err != nil {
			log.WithFields(log.Fields{
				"error": err,
			}).Error("DB open failed.")
			return err
		}
		// db.LogMode(true)
		db.SingularTable(true)
		connection = db
	} else {
		log.Info("DB connection exist.")
	}
	return nil
}

// GetConnection Get the DB connection.
func GetConnection() *gorm.DB {
	return connection
}

// CreateTables Create all the tables.
func CreateTables(tables []TableInfo) bool {
	c := GetConnection()
	success := true
	for i := range tables {
		if err := c.CreateTable(tables[i].Info).Error; err != nil {
			success = false
			log.WithFields(log.Fields{
				"Table": tables[i].Name,
				"error": err,
			}).Error("DB create table failed.")
		}
	}
	return success
}

// RemoveTables Remove all the tables.
func RemoveTables(tables []TableInfo) bool {
	c := GetConnection()
	success := true
	for i := range tables {
		if err := c.DropTableIfExists(tables[i].Info).Error; err != nil {
			success = false
			log.WithFields(log.Fields{
				"Table": tables[i].Name,
				"error": err,
			}).Error("DB remove table failed.")
		}
	}
	return success
}
