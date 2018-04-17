package service

import (
	"promise/base"
	"promise/pool/db"
	"promise/pool/object/dto"
	"promise/sdk/event"
)

var (
	ipv4PoolDB = &db.IPv4PoolDB{
		DB: base.DB{
			TemplateImpl: new(db.IPv4PoolDB),
		},
	}

	eventService event.Service
)

// IPv4Pool is the concrete service.
type IPv4Pool struct {
}

// GetCategory returns the category of this service.
func (s *IPv4Pool) GetCategory() string {
	return base.CategoryPoolIPv4
}

// NewResponse creates a new response DTO.
func (s *IPv4Pool) NewResponse() base.ResponseInterface {
	return new(dto.GetIPv4PoolResponse)
}

// GetDB returns the DB implementation.
func (s *IPv4Pool) GetDB() base.DBInterface {
	return ipv4PoolDB
}

// GetEventService returns the event service implementation.
func (s *IPv4Pool) GetEventService() base.EventServiceInterface {
	return eventService
}
