package service

import (
	log "github.com/sirupsen/logrus"
	"github.com/docker/docker/client"
	"promise/base"
	"promise/director/object/dto"
)

// Node is the service
type Node struct {
	base.CRUDService
	client *client.Client
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

// Client initialize the client if hasn't yet.
func (s *Node) Client() *client.Client {
	if s.client == nil {
		var err error
		if s.client, err = client.NewClientWithOpts(client.WithVersion("1.35")); err != nil {
			log.WithFields(log.Fields{
				"err": err,
			}).Info("Service failed to create client.")
		}
	}
	return s.client
}

// GetCollection get the Node collection.
func (s *Node) GetCollection(start int64, count int64, filter string) (*base.CollectionModel, []base.Message) {
	var cli = s.Client()

	if cli == nil {
		return nil, []base.Message{*base.NewMessageInternalError()}		
	}

	return nil, nil
}
