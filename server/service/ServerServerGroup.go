package service

import (
	"promise/base"
	"promise/sdk/event"
	"promise/server/db"
	"promise/server/object/dto"
)

var (
	serverServerGroupDB = &db.ServerServerGroupDB{
		DB: base.DB{
			TemplateImpl: new(db.ServerServerGroupDB),
		},
	}
)

// ServerServerGroup is the concrete service.
type ServerServerGroup struct {
}

// GetCategory returns the category of this service.
func (s *ServerServerGroup) GetCategory() string {
	return base.CategoryServerServerGroup
}

// NewResponse creates a new response DTO.
func (s *ServerServerGroup) NewResponse() base.ResponseInterface {
	return new(dto.GetServerServerGroupResponse)
}

// GetDB returns the DB implementation.
func (s *ServerServerGroup) GetDB() base.DBInterface {
	return serverServerGroupDB
}

// GetEventService returns the event service implementation.
func (s *ServerServerGroup) GetEventService() base.EventServiceInterface {
	return eventService
}
