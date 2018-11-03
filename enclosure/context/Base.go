package context

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/enclosure/client/enclosure"
	"promise/enclosure/db"
	"promise/enclosure/object/model"
	"promise/enclosure/object/dto"
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
	DispatchCreateEvent()
	DispatchUpdateEvent()
	DispatchDeleteEvent()			
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

// GetEnclosureResponse translate the enclosure into response.
func (c *BaseImpl) GetEnclosureResponse() *dto.GetEnclosureResponse {
	response := dto.GetEnclosureResponse{}
	if err := response.Load(c.GetEnclosure()); err != nil {
		log.WithFields(log.Fields{
			"id":    c.GetID(),
			"error": err,
		}).Warn("Context get response failed, create response failed.")
		return nil
	}
	return &response
}

// DispatchCreateEvent will send an create event using the enclosure in the context.
func (c *BaseImpl) DispatchCreateEvent() {
	response := c.GetEnclosureResponse()
	if response == nil {
		return
	}
	err := base.PublishResourceMessage(base.CreateOperation, response)
	if err != nil {
		log.WithFields(log.Fields{
			"id":    c.GetID(),
			"type": base.CreateOperation,
			"error": err,
		}).Warn("Context dispatch event failed, event dispatching failed.")
	}
}

// DispatchUpdateEvent will send an update event using the enclosure in the context.
func (c *BaseImpl) DispatchUpdateEvent() {
	response := c.GetEnclosureResponse()
	if response == nil {
		return
	}
	err := base.PublishResourceMessage(base.UpdateOperation, response)
	if err != nil {
		log.WithFields(log.Fields{
			"id":    c.GetID(),
			"type": base.UpdateOperation,
			"error": err,
		}).Warn("Context dispatch event failed, event dispatching failed.")
	}
}

// DispatchDeleteEvent will send an delete event using the enclosure in the context.
func (c *BaseImpl) DispatchDeleteEvent() {
	response := c.GetEnclosureResponse()
	if response == nil {
		return
	}
	err := base.PublishResourceMessage(base.DeleteOperation, response)
	if err != nil {
		log.WithFields(log.Fields{
			"id":    c.GetID(),
			"type": base.DeleteOperation,
			"error": err,
		}).Warn("Context dispatch event failed, event dispatching failed.")
	}
}