package service

import (
	// log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/director/object/dto"
)

// Node is the service
type Node struct {
	base.CRUDService
}

// Category returns the category of this service.
func (s *Node) Category() string {
	return base.CategoryNode
}

// Response creates a new response DTO.
func (s *Node) Response() base.GetResponseInterface {
	return new(dto.GetNodeResponse)
}

// DB returns the DB implementation.
// DB is not need.
func (s *Node) DB() base.DBInterface {
	return nil
}

// EventService returns the event service implementation.
func (s *Node) EventService() base.EventServiceInterface {
	return eventService
}
