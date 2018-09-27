package dto

import (
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/server/object/model"
)

// GetServerGroupResponse is the DTO.
type GetServerGroupResponse struct {
	base.GetResponse
	Name        string `json:"Name"`
	Description string `json:"Description"`
}

// String return the name for debug.
func (dto GetServerGroupResponse) String() string {
	return dto.Name
}

// Load the data from model.
func (dto *GetServerGroupResponse) Load(data base.ModelInterface) error {
	m, ok := data.(*model.ServerGroup)
	if !ok {
		log.Error("GetServerGroupResponse.Load() failed, convert interface failed.")
		return base.ErrorDataConvert
	}
	dto.GetResponse.Load(&m.Model)
	dto.Name = m.Name
	dto.Description = m.Description
	return nil
}
