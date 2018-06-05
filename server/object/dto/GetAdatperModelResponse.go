package dto

import (
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/server/object/model"
)

// GetAdapterModelResponse is DTO.
type GetAdapterModelResponse struct {
	base.GetResponse
	Name string `json:"Name"`
}

// DebugInfo return the name for debug.
func (dto *GetAdapterModelResponse) DebugInfo() string {
	return dto.Name
}

// Load will load data from model.
func (dto *GetAdapterModelResponse) Load(data base.ModelInterface) error {
	m, ok := data.(*model.AdapterModel)
	if !ok {
		log.Error("GetAdapterModelResponse.Load() failed, convert interface failed.")
		return base.ErrorDataConvert
	}
	dto.GetResponse.Load(&m.Model)
	dto.Name = m.Name
	return nil
}
