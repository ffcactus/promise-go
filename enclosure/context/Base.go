package context

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/enclosure/client/enclosure"
	"promise/enclosure/db"
	"promise/enclosure/object/model"
)

// Base interface contains the base methods need in context operation.
type Base interface {
	UpdateEnclosure(i base.ModelInterface)
	GetID() string
	GetTaskID() string
	SetTaskID(t string)
	GetEnclosure() *model.Enclosure
	GetClient() enclosure.Client
	SetClient(t enclosure.Client)
	GetDB() *db.Enclosure
	SetDB(t *db.Enclosure)
	String() string
}

// BaseImpl implements the Base interface.
type BaseImpl struct {
	Client    enclosure.Client
	DB        *db.Enclosure
	Enclosure *model.Enclosure
	TaskID    string
	ID        string
}

// String return the debug info.
func (c BaseImpl) String() string {
	return fmt.Sprintf("(ID = %s)", c.ID)
}

// UpdateEnclosure will update the enclosure reference in context.
func (c *BaseImpl) UpdateEnclosure(i base.ModelInterface) {
	m, ok := i.(*model.Enclosure)
	if !ok {
		log.Error("Context update enclosure failed, convert to model failed.")
		return
	}
	c.Enclosure = m
}

// GetID returns enclosure ID.
func (c *BaseImpl) GetID() string {
	return c.ID
}

// GetDB returns DB implementation.
func (c *BaseImpl) GetDB() *db.Enclosure {
	return c.DB
}

// SetDB set the DB implementation in DB.
func (c *BaseImpl) SetDB(t *db.Enclosure) {
	c.DB = t
}

// GetClient returns client implementation.
func (c *BaseImpl) GetClient() enclosure.Client {
	return c.Client
}

// SetClient sets the enclosure client.
func (c *BaseImpl) SetClient(t enclosure.Client) {
	c.Client = t
}

// GetTaskID returns task ID.
func (c *BaseImpl) GetTaskID() string {
	return c.TaskID
}

// SetTaskID sets the task ID>
func (c *BaseImpl) SetTaskID(t string) {
	c.TaskID = t
}

// GetEnclosure returns current enclosure.
func (c *BaseImpl) GetEnclosure() *model.Enclosure {
	return c.Enclosure
}
