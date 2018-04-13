package dto

import (
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/student/object/model"
)

// GetStudentResponse is the DTO of get student response.
type GetStudentResponse struct {
	base.Response
	Name   string
	Age    int
	Phones []PhoneResponse
}

// GetDebugName return the name for debug.
func (dto *GetStudentResponse) GetDebugName() string {
	return dto.Name
}

// Load will load info from model.
func (dto *GetStudentResponse) Load(i base.ModelInterface) error {
	m, ok := i.(*model.Student)
	if !ok {
		log.Error("GetStudentResponse.Load() failed, convert interface failed.")
		return base.ErrorDataConvert
	}
	base.ResponseLoad(&dto.Response, &m.Model)
	dto.Name = m.Name
	dto.Age = m.Age
	for _, v := range m.Phones {
		phone := PhoneResponse{}
		phone.Load(&v)
		dto.Phones = append(dto.Phones, phone)
	}
	return nil
}
