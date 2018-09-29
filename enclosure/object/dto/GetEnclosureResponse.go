package dto

import (
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/enclosure/object/model"
)

// GetEnclosureResponse is the DTO.
type GetEnclosureResponse struct {
	base.GetResponse
	Name        string
	Description string
	State       string
	Health      string
	base.DeviceIdentity
	Addresses []string
}

// String return the name for debug.
func (dto GetEnclosureResponse) String() string {
	return dto.Name
}

// Load will load data from model.
func (dto *GetEnclosureResponse) Load(data base.ModelInterface) error {
	m, ok := data.(*model.Enclosure)
	if !ok {
		log.Error("GetEnclosureResponse.Load() failed, convert interface failed.")
		return base.ErrorDataConvert
	}
	dto.GetResponse.Load(&m.Model)
	dto.DeviceIdentity = m.DeviceIdentity
	dto.Name = m.Name
	dto.Description = m.Description
	dto.State = m.State
	dto.Health = m.Health
	return nil
}
