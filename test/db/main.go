package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"fmt"
	"time"
)

var (
	connection *gorm.DB
	done = make(chan bool)
)


// TestTable is the table in DB for this test.
type TestTable struct {
	ID        string    `gorm:"column:ID;primary_key"`
	Category  string    `gorm:"column:Category"`
	CreatedAt time.Time `gorm:"column:CreatedAt"`
	UpdatedAt time.Time `gorm:"column:UpdatedAt"`
	Value int
}

// TableName returns the table name.
func (TestTable) TableName() string {
	return "TestTable"
}

// InitConnection Init the DB connection. Each service have to call the method first.
func InitConnection() {
	if connection == nil {
		args := "host=localhost port=5432 user=postgres dbname=promise sslmode=disable password=iforgot"
		db, err := gorm.Open("postgres", args)

		if err != nil {
			fmt.Println("Open", err)
		}
		// db.LogMode(true)
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

// NewConnection creates a new connection instance.
func NewConnection() *gorm.DB {
	args := "host=localhost port=5432 user=postgres dbname=promise sslmode=disable password=iforgot"
	db, err := gorm.Open("postgres", args)
	if err != nil {
		fmt.Printf("NewConnection failed, err = %v\n", err)
	}
	// db.LogMode(true)
	return db
}

// Increase will increase the value in table.
func Increase(name string) {
	var (
		table = new(TestTable)
		// c = NewConnection()
		c = GetConnection()
	)

	tx := c.Begin()
	if err := tx.Error; err != nil {
		fmt.Printf("%s Begin() failed.\n")
	}
	// time.Sleep(time.Duration(1) * time.Second)
	if tx.Where("\"ID\" = ?", "UUID").First(table).RecordNotFound() {
		tx.Rollback()
		fmt.Printf("%s First() failed.\n", name)
		done <- false
		return
	}
	fmt.Printf("%s read value = %d\n", name, table.Value)
	table.Value++
	// time.Sleep(time.Duration(1) * time.Second)
	if err := tx.Save(table).Error; err != nil {
		tx.Rollback()
		fmt.Printf("%s Save() failed.\n", name)
		done <- false
		return
	}
	// time.Sleep(time.Duration(1) * time.Second)
	if err := tx.Commit().Error; err != nil {
		fmt.Printf("%s Commit() failed.\n", name)
		done <- false
		return
	} 
	fmt.Printf("%s Commit() successful!\n", name)
	done <-true	
}

func main() {
	// Init the connection to share with all the instances.
	InitConnection()

	c := NewConnection()
	table := new(TestTable)
	table.Value = 0

	if err := c.DropTableIfExists(table).Error; err != nil {
		fmt.Println("DropTableIfExists()", err)
	}
	if err := c.CreateTable(table).Error; err != nil {
		fmt.Println("CreateTable()", err)
	}
	table.ID = "UUID"
	if err := c.Create(table).Error; err != nil {
		fmt.Println("Create", err)
	}

	if c.First(table).RecordNotFound() {
		fmt.Printf("main First() failed\n")
	}
	fmt.Printf("--- Before Test ---\n")
	fmt.Printf("Value = %d\n", table.Value)

	fmt.Printf("--- Start Test ---\n")
	instance := 40
	for i := 0; i < instance; i++ {
		go Increase(fmt.Sprintf("Instance %d", i))
		// time.Sleep(time.Duration(1) * time.Second)
	}
	expected := 0
	for i := 0; i < instance; i++ {
		if <- done {
			expected++
		}
	}
	fmt.Println("--- Test Result ---")
	if c.First(table).RecordNotFound() {
		fmt.Printf("%s RecordNotFound(), Rollback(), err = %v\n", c.Error)
	}
	fmt.Printf("Expected = %d, Result = %d\n", expected, table.Value)
}