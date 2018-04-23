package service

import (
	"promise/base"
	"promise/sdk/event"
	"promise/server/db"
	"promise/server/object/dto"
)

var (
	serverDB = &db.Server{
		DB: base.DB{
			TemplateImpl: new(db.Server),
		},
	}

	eventService event.Service
)

// Server is the server service
type Server struct {
}

// Category returns the category of this service.
func (s *Server) Category() string {
	return base.CategoryServer
}

// Response creates a new response DTO.
func (s *Server) Response() base.GetResponseInterface {
	return new(dto.GetServerResponse)
}

// DB returns the DB implementation.
func (s *Server) DB() base.DBInterface {
	return serverDB
}

// EventService returns the event service implementation.
func (s *Server) EventService() base.EventServiceInterface {
	return eventService
}
