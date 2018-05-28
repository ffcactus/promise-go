package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

type Table struct {
	ID uint64 `gorm:"column:ID;primary_key"`
	State string `gorm:"column:State"`
	Value int `gorm:"column:Value"`
}

func (Table) TableName() string {
	return "Table"
}

var (
	instance = 2
	connection *gorm.DB
	done       = make(chan bool)
)

// InitConnection Init the DB connection. Each service have to call the method first.
func InitConnection() {
	if connection == nil {
		args := "host=localhost port=5432 user=postgres dbname=promise sslmode=disable password=iforgot"
		db, err := gorm.Open("postgres", args)

		if err != nil {
			fmt.Println("Open", err)
		}
		db.LogMode(true)
		db.SingularTable(true)
		connection = db
	} else {
		fmt.Println("DB connection exist.")
	}
}

// GetConnection Get the DB connection.
func GetConnection() *gorm.DB {
	return connection
}

// GetAndLock will do get and lock in transaction.
func GetAndLock(instance int) {
	c := GetConnection()
	// c.Exec("SET TRANSACTION ISOLATION LEVEL SERIALIZABLE;")
	// defer c.Exec("SET TRANSACTION ISOLATION LEVEL Read committed;")
	tx := c.Begin()
	table := new(Table)
	if err := tx.Error; err != nil {
		fmt.Printf("%d Begin() failed.\n", instance)
		return
	}
	tx.Exec("SET TRANSACTION ISOLATION LEVEL SERIALIZABLE;")

	time.Sleep(time.Duration(1) * time.Second)
	if tx.Where("\"State\" = ?", "Added").First(table).RecordNotFound() {
		fmt.Printf("%d No record.\n", instance)
		tx.Rollback()
		done <- false
		return
	}
	fmt.Printf("%d find the Added.\n", instance)
	time.Sleep(time.Duration(1) * time.Second)
	table.State = "Locked"
	table.Value = instance
	if err := tx.Save(table).Error; err != nil {
		fmt.Printf("%d Save() failed.\n", instance)
		tx.Rollback()	
		done <- false
		return
	}
	fmt.Printf("%d save as locked.\n", instance)
	time.Sleep(time.Duration(1) * time.Second)
	if err := tx.Commit().Error; err != nil {
		fmt.Printf("%d Commit() failed.\n", instance)
		done <- false
		return
	}
	fmt.Printf("%d Commit() successful!\n", instance)
	done <- true
}

func main() {
	var (
		table = new(Table)
	)

	InitConnection()
	c := GetConnection()
	if err := c.DropTableIfExists(table).Error; err != nil {
		fmt.Println("DropTableIfExists()", err)
	}
	if err := c.CreateTable(table).Error; err != nil {
		fmt.Println("CreateTable()", err)
	}
	table.ID = 1
	table.State = "Added"
	table.Value = 0

	if err := c.Create(table).Error; err != nil {
		fmt.Printf("Save() failed\n")
		return
	}

	for i := 0; i < instance; i++ {
		go GetAndLock(i + 1)
	}

	for i := 0; i < instance; i++ {
		<-done
		fmt.Printf("instance done.\n")
	}
}