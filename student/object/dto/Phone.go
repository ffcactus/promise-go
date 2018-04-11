package dto

import (
	"promise/student/object/model"
)

// The Phone DTO.
type Phone struct {
	Number string `json:"Number`
}

// ToModel convert the DTO to model.
func (dto *Phone) ToModel() *model.Phone {
	m := model.Phone{}
	m.Number = dto.Number
	return &m
}

// PhoneResponse is the phone in response.
type PhoneResponse struct {
	ID uint64 `json:"ID"`
	Phone
}

// Load will load data from model.
func (dto *PhoneResponse) Load(m *model.Phone) {
	dto.ID = m.ID
	dto.Number = m.Number
}