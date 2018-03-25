package dto

import (
	log "github.com/sirupsen/logrus"
	"promise/common/object/constError"
	"promise/common/object/constValue"
	commonDTO "promise/common/object/dto"
	"promise/server/object/model"
)

// GetServerGroupResponse is the DTO.
type GetServerGroupResponse struct {
	commonDTO.PromiseResponse
	Name        string `json:"Name"`
	Description string `json:"Description"`
}

// Load the data from model.
func (dto *GetServerGroupResponse) Load(data interface{}) error {
	m, ok := data.(*model.ServerGroup)
	if !ok {
		log.Warn("GetServerGroupResponse load data from model failed.")
		return constError.ErrorDataConvert
	}
	dto.PromiseResponse.Load(&m.PromiseModel)
	dto.URI = constValue.ToServerGroupURI(m.ID)
	dto.Name = m.Name
	dto.Description = m.Description
	return nil
}
