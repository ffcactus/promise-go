package dto

import (
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/server/object/model"
)

// GetServerGroupResponse is the DTO.
type GetServerGroupResponse struct {
	base.Response
	Name        string `json:"Name"`
	Description string `json:"Description"`
}

// GetDebugName return the name for debug.
func (dto *GetServerGroupResponse) GetDebugName() string {
	return dto.Name
}

// Load the data from model.
func (dto *GetServerGroupResponse) Load(data base.ModelInterface) error {
	m, ok := data.(*model.ServerGroup)
	if !ok {
		log.Warn("GetServerGroupResponse.Load() failed, convert interface failed.")
		return base.ErrorDataConvert
	}
	base.ResponseLoad(&dto.Response, &m.Model)
	dto.Name = m.Name
	dto.Description = m.Description
	return nil
}
