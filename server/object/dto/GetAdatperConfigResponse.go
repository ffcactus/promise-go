package dto

import (
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/server/object/model"
)

// GetAdapterConfigResponse is DTO.
type GetAdapterConfigResponse struct {
	base.GetResponse
	Name string `json:"Name"`
}

// DebugInfo return the name for debug.
func (dto *GetAdapterConfigResponse) DebugInfo() string {
	return dto.Name
}

// Load will load data from model.
func (dto *GetAdapterConfigResponse) Load(data base.ModelInterface) error {
	m, ok := data.(*model.AdapterConfig)
	if !ok {
		log.Error("GetAdapterConfigResponse.Load() failed, convert interface failed.")
		return base.ErrorDataConvert
	}
	dto.GetResponse.Load(&m.Model)
	dto.Name = m.Name
	return nil
}
