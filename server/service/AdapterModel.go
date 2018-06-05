package service

import (
	"promise/base"
	"promise/server/db"
	"promise/server/object/dto"
)

var (
	adapterModelDB = &db.AdapterModel{
		DB: base.DB{
			TemplateImpl: new(db.AdapterModel),
		},
	}
)

// AdapterModel is the servergroup service.
type AdapterModel struct {
}

// Category returns the category of this service.
func (s *AdapterModel) Category() string {
	return base.CategoryAdapterModel
}

// Response creates a new response DTO.
func (s *AdapterModel) Response() base.GetResponseInterface {
	return new(dto.GetAdapterModelResponse)
}

// DB returns the DB implementation.
func (s *AdapterModel) DB() base.DBInterface {
	return adapterModelDB
}

// EventService returns the event service implementation.
func (s *AdapterModel) EventService() base.EventServiceInterface {
	return eventService
}
