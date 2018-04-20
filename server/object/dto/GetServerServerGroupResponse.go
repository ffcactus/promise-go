package dto

import (
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/server/object/model"
)

// GetServerServerGroupResponse is the response DTO.
type GetServerServerGroupResponse struct {
	base.Response
	ServerID       string `json:"ServerID"`
	ServerURI      string `json:"ServerURI"`
	ServerGroupID  string `json:"ServerGroupID"`
	ServerGroupURI string `json:"ServerGroupURI"`
}

// GetDebugName return the name for debug.
func (dto *GetServerServerGroupResponse) GetDebugName() string {
	return dto.ServerID + " " + dto.ServerGroupID
}

// Load the data from model.
func (dto *GetServerServerGroupResponse) Load(data base.ModelInterface) error {
	m, ok := data.(*model.ServerServerGroup)
	if !ok {
		log.Error("GetIPv4PoolResponse.Load() failed, convert interface failed.")
		return base.ErrorDataConvert
	}
	base.ResponseLoad(&dto.Response, &m.Model)
	dto.ServerID = m.ServerID
	dto.ServerURI = base.ToServerURI(m.ServerID)
	dto.ServerGroupID = m.ServerGroupID
	dto.ServerGroupURI = base.ToServerGroupURI(m.ServerGroupID)
	return nil
}
