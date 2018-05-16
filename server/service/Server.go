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
	base.CRUDService
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

// Delete will override the default process.
// We not only remove the server but the server-servergroup.
func (s *Server) Delete(id string) []base.Message {
	var (
		response = s.Response()
	)

	server, ssg, message := serverDB.DeleteServer(id)
	if message != nil {
		return []base.Message{*message}
	}
	response.Load(server)
	for _, v := range ssg {
		ssgResponse := dto.GetServerServerGroupResponse{}
		ssgResponse.Load(v)
		s.EventService().DispatchDeleteEvent(&ssgResponse)
	}
	s.EventService().DispatchDeleteEvent(response)
	return nil
}

// DeleteCollection will override the default process.
// We not only remove the server but the server-servergroup.
func (s *Server) DeleteCollection() []base.Message {
	records, ssg, message := serverDB.DeleteServerCollection()
	if message != nil {
		return []base.Message{*message}
	}
	for _, v := range records {
		response := s.TemplateImpl.Response()
		response.Load(v)
		s.TemplateImpl.EventService().DispatchDeleteEvent(response)
	}
	for _, v := range ssg {
		ssgResponse := dto.GetServerServerGroupResponse{}
		ssgResponse.Load(v)
		s.EventService().DispatchDeleteEvent(&ssgResponse)
	}
	s.TemplateImpl.EventService().DispatchDeleteCollectionEvent(s.TemplateImpl.Category())
	return nil
}
