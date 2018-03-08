package db

import (
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// TableInfo The tables in DB.
type TableInfo struct {
	Name string
	Info interface{}
}

var connection *gorm.DB

// InitConnection Init the DB connection. Each service have to call the method first.
func InitConnection() error {
	if connection == nil {
		beego.Info("Init DB connection.")
		args := "host=localhost port=5432 user=postgres dbname=promise sslmode=disable password=iforgot"
		db, err := gorm.Open("postgres", args)
		if err != nil {
			beego.Info("gorm.Open() failed, error = ", err)
			return err
		}
		db.LogMode(false)
		db.SingularTable(true)
		connection = db
	} else {
		beego.Info("DB connection exist.")
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
			beego.Error("Failed to create table", tables[i].Name, err)
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
			beego.Error("Failed to remove table", tables[i].Name, err)
		}
	}
	return success
}
