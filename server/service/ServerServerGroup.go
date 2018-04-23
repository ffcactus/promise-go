package service

import (
	"promise/base"
	"promise/server/db"
	"promise/server/object/dto"
)

var (
	serverServerGroupDB = &db.ServerServerGroup{
		DB: base.DB{
			TemplateImpl: new(db.ServerServerGroup),
		},
	}
)

// ServerServerGroup is the concrete service.
type ServerServerGroup struct {
}

// Category returns the category of this service.
func (s *ServerServerGroup) Category() string {
	return base.CategoryServerServerGroup
}

// Response creates a new response DTO.
func (s *ServerServerGroup) Response() base.GetResponseInterface {
	return new(dto.GetServerServerGroupResponse)
}

// DB returns the DB implementation.
func (s *ServerServerGroup) DB() base.DBInterface {
	return serverServerGroupDB
}

// EventService returns the event service implementation.
func (s *ServerServerGroup) EventService() base.EventServiceInterface {
	return eventService
}
