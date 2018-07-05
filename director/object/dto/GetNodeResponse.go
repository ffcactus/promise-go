package dto

import (
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/director/object/model"
)

// GetNodeResponse is DTO.
type GetNodeResponse struct {
	base.GetResponse
	Hostname string `json:"Hostname"`
	Status string `json:"Status"`
	Availibility string `json:"Availibility"`
	ManagerStatus string `json:"ManagerStatus"`
}

// DebugInfo return the name for debug.
func (dto *GetNodeResponse) DebugInfo() string {
	return dto.Hostname
}

// Load the data from model.
func (dto *GetNodeResponse) Load(data base.ModelInterface) error {
	m, ok := data.(*model.Node)
	if !ok {
		log.Error("GetNodeResponse.Load() failed, convert interface failed.")
		return base.ErrorDataConvert
	}
	dto.GetResponse.Load(&m.Model)
	dto.Hostname = m.Hostname
	dto.Status = m.Status
	dto.Availibility = m.Availibility
	dto.ManagerStatus = m.ManagerStatus	
	return nil
}