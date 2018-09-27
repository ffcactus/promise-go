package service

import (
	"promise/base"
	"promise/enclosure/db"
	"promise/enclosure/object/dto"
)

var (
	enclosureDB = &db.Enclosure{
		DB: base.DB{
			TemplateImpl: new(db.Enclosure),
		},
	}
)

// Enclosure is the enclosure service.
type Enclosure struct {
	base.CRUDService
}

// Category returns the category of this service.
func (s *Enclosure) Category() string {
	return base.CategoryEnclosure
}

// Response creates a new response DTO.
func (s *Enclosure) Response() base.GetResponseInterface {
	return new(dto.GetEnclosureResponse)
}

// DB returns the DB implementation.
func (s *Enclosure) DB() base.DBInterface {
	return enclosureDB
}
