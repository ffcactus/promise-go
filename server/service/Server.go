package service

import (
	"promise/base"
	"promise/sdk/event"
	"promise/server/context"
	"promise/server/db"
	"promise/server/object/dto"
	"promise/server/object/message"
	"promise/server/object/model"
	"promise/server/strategy"
	"time"
)

var (
	serverDB = &db.Server{
		DB: base.DB{
			TemplateImpl: new(db.Server)
		}
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
	return new(dto.GetIPv4PoolResponse)
}

// DB returns the DB implementation.
func (s *Server) DB() base.DBInterface {
	return serverDB
}

// EventService returns the event service implementation.
func (s *Server) EventService() base.EventServiceInterface {
	return eventService
}


// FindServerStateAdded will find the server with state added.
func FindServerStateAdded() {
	for {
		seconds := 5
		if id := serverDB.FindServerStateAdded(); id != "" {
			RefreshServer(id)
			seconds = 0
		} else {
			seconds = 5
		}
		time.Sleep(time.Duration(seconds) * time.Second)
	}
}
