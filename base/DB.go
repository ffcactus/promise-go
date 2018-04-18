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
	GetResourceName() string
	NeedCheckDuplication() bool
	GetConnection() *gorm.DB
	NewEntity() EntityInterface
	NewEntityCollection() interface{}
	ConvertFindResultToCollection(start int64, total int64, result interface{}) (*CollectionModel, error)
	ConvertFindResultToModel(interface{}) ([]ModelInterface, error)
}

// DBInterface is the interface that DB should have.
type DBInterface interface {
	Post(ModelInterface) (bool, ModelInterface, bool, error)
	Get(id string) ModelInterface
	Update(id string, request UpdateActionRequestInterface) (bool, ModelInterface, bool, error)
	Delete(id string) (bool, ModelInterface, bool, error)
	GetCollection(start int64, count int64, filter string) (*CollectionModel, error)
	DeleteCollection() ([]ModelInterface, bool, error)
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
		name   = impl.TemplateImpl.GetResourceName()
		record = impl.TemplateImpl.NewEntity()
		c      = impl.TemplateImpl.GetConnection()
	)
	tx := c.Begin()
	if err := tx.Error; err != nil {
		log.WithFields(log.Fields{
			"resource": name,
			"error":    err,
		}).Warn("Post resource in DB failed, start transaction failed.")
		return false, nil, false, err
	}
	if impl.TemplateImpl.NeedCheckDuplication() {
		where := "\"" + record.GetPropertyNameForDuplicationCheck() + "\" = ?"
		if !tx.Where(where, m.GetValueForDuplicationCheck()).First(record).RecordNotFound() {
			tx.Rollback()
			log.WithFields(log.Fields{
				"resource": name,
				"id":       record.GetID(),
				"name":     record.GetDebugName(),
			}).Warn("Post resource in DB failed, duplicated resource, transaction rollback.")
			return true, nil, false, ErrorResourceNotExist
		}
	}

	record.Load(m)
	record.SetID(uuid.New().String())
	if err := c.Create(record).Error; err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"resource": name,
			"name":     m.GetDebugName(),
			"error":    err,
		}).Warn("Post resource in DB failed, create resource failed, transaction rollback.")
		return false, nil, false, err
	}
	if err := tx.Save(record).Error; err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"resource": name,
			"name":     m.GetDebugName(),
			"error":    err,
		}).Warn("Post resource in DB failed, save resource failed, transaction rollback.")
		return false, nil, false, err
	}
	if err := tx.Commit().Error; err != nil {
		log.WithFields(log.Fields{
			"resource": name,
			"name":     m.GetDebugName(),
			"error":    err,
		}).Warn("Post resource in DB failed, commit failed.")
		return false, nil, false, err
	}
	return false, record.ToModel(), true, nil
}

// GetInternal is part of the process to get resource in DB, since many other operation
// need this process, we seperate it out.
// It will return if the resource been found.
// It will return error if any.
func (impl *DB) GetInternal(tx *gorm.DB, id string, record EntityInterface) (bool, error) {
	var (
		name = impl.TemplateImpl.GetResourceName()
	)

	preload := record.GetPreload()
	if tx.Where("\"ID\" = ?", id).First(record).RecordNotFound() {
		tx.Rollback()
		log.WithFields(log.Fields{
			"resource": name,
			"id":       id,
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
			"resource": name,
			"id":       id,
			"error":    err,
		}).Warn("Get resource in DB failed, fatch failed, transaction rollback.")
		return true, err
	}
	return true, nil
}

// Get is the default implement to get resource in DB.
// If the resource does not exist in the DB return nil.
func (impl *DB) Get(id string) ModelInterface {
	var (
		name   = impl.TemplateImpl.GetResourceName()
		record = impl.TemplateImpl.NewEntity()
		c      = impl.TemplateImpl.GetConnection()
	)
	tx := c.Begin()
	if err := tx.Error; err != nil {
		log.WithFields(log.Fields{
			"resource": name,
			"id":       id,
			"error":    err,
		}).Warn("Get resource in DB failed, start transaction failed.")
		return nil
	}
	if exist, err := impl.GetInternal(tx, id, record); exist && err == nil {
		return record.ToModel()
	}
	return nil
}

// Update is the default implement to update resource in DB.
// It will return if the one exist.
// It will return the updated resource.
// It will return wether the operation commited.
// It will return error if any.
func (impl *DB) Update(id string, request UpdateActionRequestInterface) (bool, ModelInterface, bool, error) {
	var (
		name   = impl.TemplateImpl.GetResourceName()
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
		return true, nil, false, err
	}
	if exist, err := impl.GetInternal(tx, id, record); !exist || err != nil {
		return false, nil, false, err
	}
	m := record.ToModel()
	if err := request.UpdateModel(m); err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"resource": name,
			"id":       id,
			"error":    err}).
			Warn("Update resource in DB failed, update model failed, transaction rollback.")
		return true, nil, false, err
	}
	record.Load(m)
	record.SetID(id)
	if err := tx.Save(record).Error; err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"resource": name,
			"id":       id,
			"error":    err}).
			Warn("Update resource in DB failed, save resource failed, transaction rollback.")
		return true, nil, false, err
	}
	if err := tx.Commit().Error; err != nil {
		log.WithFields(log.Fields{
			"resource": name,
			"id":       id,
			"error":    err,
		}).Warn("Update resource in DB failed, commit failed.")
		return true, nil, false, err
	}
	return true, record.ToModel(), true, nil
}

// Delete is the default implement to delete resource from DB.
// It will return if the one exist.
// It will return the deleted one if commited.
// It will return wether the operation commited.
// It will return error if any.
func (impl *DB) Delete(id string) (bool, ModelInterface, bool, error) {
	var (
		name     = impl.TemplateImpl.GetResourceName()
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
			"resource": name,
			"id":       id,
			"error":    err,
		}).Warn("Delete resource from DB failed, start transaction failed.")
		return true, nil, false, err
	}
	if exist, err := impl.GetInternal(tx, id, previous); err != nil || !exist {
		return false, nil, false, err
	}
	record.SetID(id)
	for _, v := range record.GetAssociation() {
		if err := tx.Delete(v).Error; err != nil {
			tx.Rollback()
			log.WithFields(log.Fields{
				"resource": name,
				"id":       id,
				"error":    err,
			}).Warn("Delete resource from DB failed, delete association failed, transaction rollback.")
			return true, nil, false, err
		}
	}
	if err := tx.Delete(record).Error; err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"resource": name,
			"id":       id,
			"error":    err,
		}).Warn("Delete resource from DB failed, delete resource failed, transaction rollback.")
		return true, nil, false, err
	}
	if err := tx.Commit().Error; err != nil {
		log.WithFields(log.Fields{
			"resource": name,
			"id":       id,
			"error":    err,
		}).Warn("Delete resource from DB failed, commit failed.")
		return true, nil, false, err
	}
	return true, previous.ToModel(), true, nil
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
func (impl *DB) GetCollection(start int64, count int64, filter string) (*CollectionModel, error) {
	var (
		name             = impl.TemplateImpl.GetResourceName()
		total            int64
		record           = impl.TemplateImpl.NewEntity()
		recordCollection = impl.TemplateImpl.NewEntityCollection()
		c                = impl.TemplateImpl.GetConnection()
	)

	// Check filter first to avoid starting a transaction.
	where, err := impl.convertFilter(filter, record.GetFilterNameList())
	if err != nil {
		log.WithFields(log.Fields{
			"resource": name,
			"filter":   filter,
			"error":    err}).
			Warn("Get collection in DB failed, convert filter failed.")
		return nil, err
	}
	log.WithFields(log.Fields{
		"resource": name,
		"where":    where,
	}).Debug("Convert filter success.")

	// We need transaction to ensure the total and the query count is consistent.
	tx := c.Begin()
	if err := tx.Error; err != nil {
		log.WithFields(log.Fields{
			"resource": name,
			"error":    err}).
			Warn("Get collection in DB failed, start transaction failed.")
		return nil, err
	}

	// Get total count.
	if err := tx.Table(record.TableName()).Count(&total).Error; err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"resource": name,
			"error":    err,
		}).Warn("Get collection in DB failed, get count failed.")
		return nil, err
	}

	// Find all the matches.
	if err := tx.Limit(count).Offset(start).Where(where).Find(recordCollection).Error; err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"resource": name,
			"error":    err}).
			Warn("Get collection in DB failed, find resource failed.")
		return nil, err
	}
	if err := tx.Commit().Error; err != nil {
		log.WithFields(log.Fields{
			"resource": name,
			"error":    err}).
			Warn("Get collection in DB failed, commit failed.")
		return nil, err
	}
	return impl.TemplateImpl.ConvertFindResultToCollection(start, total, recordCollection)
}

// DeleteCollection delete the collection in DB.
// It returns all the deleted resources,
// It returnes if committed.
// It returens error if any.
func (impl *DB) DeleteCollection() ([]ModelInterface, bool, error) {
	var (
		name             = impl.TemplateImpl.GetResourceName()
		recordCollection = impl.TemplateImpl.NewEntityCollection()
		c                = impl.TemplateImpl.GetConnection()
		tables           = impl.TemplateImpl.NewEntity().GetTables()
	)

	// We need transaction to ensure the total and the query count is consistent.
	tx := c.Begin()
	if err := tx.Error; err != nil {
		log.WithFields(log.Fields{
			"resource": name,
			"error":    err}).
			Warn("Get collection in DB failed, start transaction failed.")
		return nil, false, err
	}

	if err := tx.Find(recordCollection).Error; err != nil {
		log.WithFields(log.Fields{
			"resource": name,
			"error":    err}).
			Warn("Delete collection in DB failed, find resource failed.")
		return nil, false, err
	}
	for _, v := range tables {
		if err := tx.Delete(v).Error; err != nil {
			tx.Rollback()
			log.WithFields(log.Fields{
				"resource": name,
				"error":    err}).
				Warn("Delete collection in DB failed, delete resources failed, transaction rollback.")
			return nil, false, err
		}
	}
	ret, err := impl.TemplateImpl.ConvertFindResultToModel(recordCollection)
	if err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"resource": name,
			"error":    err}).
			Warn("Delete collection in DB failed, convert find result failed, transaction rollback.")
		return nil, false, err
	}
	if err := tx.Commit().Error; err != nil {
		log.WithFields(log.Fields{
			"resource": name,
			"error":    err}).
			Warn("Delete collection in DB failed, commit failed.")
		return nil, false, err
	}
	return ret, true, nil
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
		args := "host=localhost port=5432 user=postgres dbname=promise sslmode=disable password=iforgot"
		db, err := gorm.Open("postgres", args)
		// args := "host=100.100.194.103 port=5432 user=gaussdba dbname=NETADAPTOR sslmode=disable password=Huawei12#$"
		// db, err := gorm.Open("gauss", args)
		if err != nil {
			log.Info("gorm.Open() failed, error = ", err)
			return err
		}
		db.LogMode(true)
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
			log.Error("Failed to create table", tables[i].Name, err)
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
			log.Error("Failed to remove table", tables[i].Name, err)
		}
	}
	return success
}