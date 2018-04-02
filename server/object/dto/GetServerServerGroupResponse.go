package dto

import (
	log "github.com/sirupsen/logrus"
	"promise/common/object/consterror"
	"promise/common/object/constvalue"
	commonDTO "promise/common/object/dto"
	"promise/server/object/model"
)

// GetServerServerGroupResponse is the DTO.
type GetServerServerGroupResponse struct {
	commonDTO.PromiseResponse
	ServerID       string `json:"ServerID"`
	ServerURI      string `json:"ServerURI"`
	ServerGroupID  string `json:"ServerGroupID"`
	ServerGroupURI string `json:"ServerGroupURI"`
}

// Load the data from model.
func (dto *GetServerServerGroupResponse) Load(data interface{}) error {
	m, ok := data.(*model.ServerServerGroup)
	if !ok {
		log.Warn("GetServerServerGroupResponse load data from model failed.")
		return consterror.ErrorDataConvert
	}
	dto.PromiseResponse.Load(&m.PromiseModel)
	dto.ServerID = m.ServerID
	dto.ServerURI = constvalue.ToServerURI(m.ServerID)
	dto.ServerGroupID = m.ServerGroupID
	dto.ServerGroupURI = constvalue.ToServerGroupURI(m.ServerGroupID)
	return nil
}
