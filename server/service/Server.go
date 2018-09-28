package service

import (
	"promise/base"
	"promise/server/db"
	"promise/server/object/dto"
)

var (
	serverDB = &db.Server{
		DB: base.DB{
			TemplateImpl: new(db.Server),
		},
	}
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

// Delete will override the default process.
// We not only remove the server but the server-servergroup.
func (s *Server) Delete(id string) []base.ErrorResponse {
	var (
		response = s.Response()
	)

	server, ssg, errorResp := serverDB.DeleteServer(id)
	if errorResp != nil {
		return []base.ErrorResponse{*errorResp}
	}
	response.Load(server)
	for _, v := range ssg {
		ssgResponse := dto.GetServerServerGroupResponse{}
		ssgResponse.Load(v)
		base.PublishDeleteMessage(&ssgResponse)
	}
	base.PublishDeleteMessage(response)
	return nil
}

// DeleteCollection will override the default process.
// We not only remove the server but the server-servergroup.
func (s *Server) DeleteCollection() []base.ErrorResponse {
	records, ssg, errorResp := serverDB.DeleteServerCollection()
	if errorResp != nil {
		return []base.ErrorResponse{*errorResp}
	}
	for _, v := range records {
		response := s.TemplateImpl.Response()
		response.Load(v)
		base.PublishDeleteMessage(response)
	}
	for _, v := range ssg {
		ssgResponse := dto.GetServerServerGroupResponse{}
		ssgResponse.Load(v)
		base.PublishDeleteMessage(&ssgResponse)
	}
	base.PublishDeleteCollectionMessage(s.TemplateImpl.Category())
	return nil
}
