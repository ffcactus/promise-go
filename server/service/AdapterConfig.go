package service

import (
	"promise/base"
	"promise/server/db"
	"promise/server/object/dto"
)

var (
	adapterConfigDB = &db.AdapterConfig{
		DB: base.DB{
			TemplateImpl: new(db.AdapterConfig),
		},
	}
)

// AdapterConfig is the servergroup service.
type AdapterConfig struct {
}

// Category returns the category of this service.
func (s *AdapterConfig) Category() string {
	return base.CategoryAdapterConfig
}

// Response creates a new response DTO.
func (s *AdapterConfig) Response() base.GetResponseInterface {
	return new(dto.GetAdapterConfigResponse)
}

// DB returns the DB implementation.
func (s *AdapterConfig) DB() base.DBInterface {
	return adapterConfigDB
}
