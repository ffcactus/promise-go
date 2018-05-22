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

// SubTestTable is the sub tables in TestTable.
type SubTestTable struct {
	ID uint64 `gorm:"column:ID;primary_key"`
	MainValueRef string `gorm:"column:MainValueRef"`
	Value int `gorm:"column:Value"`
}

// TestTable is the table in DB for this test.
type TestTable struct {
	ID        string    `gorm:"column:ID;primary_key"`
	Category  string    `gorm:"column:Category"`
	CreatedAt time.Time `gorm:"column:CreatedAt"`
	UpdatedAt time.Time `gorm:"column:UpdatedAt"`
	Value int           `gorm:"column:Value"`
	SubValue []SubTestTable `gorm:"column:SubValue;ForeignKey:MainValueRef"`
}

// TableName returns the table name.
func (TestTable) TableName() string {
	return "TestTable"
}

// TableName returns the table name.
func (SubTestTable) TableName() string {
	return "SubTestTable"
}

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

// NewConnection creates a new connection instance.
func NewConnection() *gorm.DB {
	args := "host=localhost port=5432 user=postgres dbname=promise sslmode=disable password=iforgot"
	db, err := gorm.Open("postgres", args)
	if err != nil {
		fmt.Printf("NewConnection failed, err = %v\n", err)
	}
	db.LogMode(true)
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

	subValue := make([]SubTestTable, 0)
	subValue = append(subValue, SubTestTable{Value: 1})
	subValue = append(subValue, SubTestTable{Value: 2})
	table.SubValue = subValue

	if err := c.DropTableIfExists(new(TestTable)).Error; err != nil {
		fmt.Println("DropTableIfExists()", err)
	}
	if err := c.DropTableIfExists(new(SubTestTable)).Error; err != nil {
		fmt.Println("DropTableIfExists()", err)
	}
	if err := c.CreateTable(new(TestTable)).Error; err != nil {
		fmt.Println("CreateTable()", err)
	}
	if err := c.CreateTable(new(SubTestTable)).Error; err != nil {
		fmt.Println("CreateTable()", err)
	}

	// Start transaction to save main value.
	table.ID = "UUID"
	tx := c.Begin()
	if err := tx.Error; err != nil {
		fmt.Printf("Begin() failed.\n")
		return
	}
	if !tx.Where("\"Value\" = ?", 100).First(table).RecordNotFound() {
		fmt.Printf("Duplicated.\n")
		tx.Rollback()
		return
	}
	if err := tx.Create(table).Error; err != nil {
		fmt.Printf("Create() failed.\n")
		tx.Rollback()
		return
	}
	if err := tx.Save(table).Error; err != nil {
		fmt.Printf("Save() failed\n")
		tx.Rollback()
		return
	}

	if err := tx.Commit().Error; err != nil {
		fmt.Printf("Commit() failed.\n")
		return
	}

	if c.First(table).RecordNotFound() {
		fmt.Printf("main First() failed\n")
	}
	fmt.Printf("--- Before Test ---\n")
	fmt.Printf("Value = %d\n", table.Value)

	fmt.Printf("--- Start Test ---\n")
	instance := 10
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