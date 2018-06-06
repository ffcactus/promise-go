package service

import (
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/server/db"
	"promise/server/object/dto"
)

var (
	serverGroupDB = &db.ServerGroup{
		DB: base.DB{
			TemplateImpl: new(db.ServerGroup),
		},
	}
)

// ServerGroup is the servergroup service.
type ServerGroup struct {
}

// CreateDefaultServerGroup will create the default server group.
func CreateDefaultServerGroup() {
	var (
		request  dto.PostServerGroupRequest
		response dto.GetServerGroupResponse
	)
	request.Name = "all"
	request.Description = "The default servergroup that includes all the servers."
	sg, message := serverGroupDB.Create(request.ToModel())
	if message == nil {
		response.Load(sg)
		eventService.DispatchCreateEvent(&response)
		log.WithFields(log.Fields{
			"id": sg.GetID(),
		}).Info("Service create the default servergroup created.")
		db.DefaultServerGroupID = sg.GetID()
	} else if message.ID == base.MessageDuplicate {
		log.Debug("Service found the default servergroup exist.")
	} else {
		log.WithFields(log.Fields{
			"message": message.ID,
		}).Error("Service failed to create default servergroup.")
	}
}

// Category returns the category of this service.
func (s *ServerGroup) Category() string {
	return base.CategoryServerGroup
}

// Response creates a new response DTO.
func (s *ServerGroup) Response() base.GetResponseInterface {
	return new(dto.GetServerGroupResponse)
}

// DB returns the DB implementation.
func (s *ServerGroup) DB() base.DBInterface {
	return serverGroupDB
}

// EventService returns the event service implementation.
func (s *ServerGroup) EventService() base.EventServiceInterface {
	return eventService
}
