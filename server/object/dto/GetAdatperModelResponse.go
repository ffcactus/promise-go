package dto

import (
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/server/object/model"
)

// RAIDCapability describe the capability of an FCoE adapter.
type RAIDCapability struct {
	Version int `json:"Version"`
}

// Load will load data from model.
func (dto *RAIDCapability) Load(m model.RAIDCapability) {
	dto.Version = m.Version
}

// EthernetCapability describe the capability of an FCoE adapter.
type EthernetCapability struct {
	Version int `json:"Version"`
}

// Load will load data from model.
func (dto *EthernetCapability) Load(m model.EthernetCapability) {
	dto.Version = m.Version
}

// FCoECapability describe the capability of an FCoE adapter.
type FCoECapability struct {
	Version int `json:"Version"`
}

// Load will load data from model.
func (dto *FCoECapability) Load(m model.FCoECapability) {
	dto.Version = m.Version
}

// AdapterCapability describe the capability of an adapter.
type AdapterCapability struct {
	Version  int                 `json:"Version"`
	RAID     *RAIDCapability     `json:"RAID"`
	Ethernet *EthernetCapability `json:"Ethernet"`
	FCoE     *FCoECapability     `json:"FCoE"`
}

// Load will load data from model.
func (dto *AdapterCapability) Load(m model.AdapterCapability) {
	dto.Version = m.Version
	if m.RAID != nil {
		dto.RAID = new(RAIDCapability)
		dto.RAID.Load(*m.RAID)
	}
	if m.Ethernet != nil {
		dto.Ethernet = new(EthernetCapability)
		dto.Ethernet.Load(*m.Ethernet)
	}
	if m.FCoE != nil {
		dto.FCoE = new(FCoECapability)
		dto.FCoE.Load(*m.FCoE)
	}
}

// GetAdapterModelResponse is DTO.
type GetAdapterModelResponse struct {
	base.GetResponse
	Name       string            `json:"Name"`
	Type       string            `json:"Type"`
	Capability AdapterCapability `json:"Capability"`
}

// String return the name for debug.
func (dto GetAdapterModelResponse) String() string {
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
	dto.Type = m.Type
	dto.Capability.Load(m.Capability)
	return nil
}
