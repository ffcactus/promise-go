package dto

import (
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/pool/object/model"
)

// GetIPv4PoolResponse is the response DTO.
type GetIPv4PoolResponse struct {
	base.Response
	IPv4PoolResource
	Ranges      []IPv4RangeResponse `json:"Ranges"`
	Total       uint32              `json:"Total"`
	Free        uint32              `json:"Free"`
	Allocatable uint32              `json:"Allocatable"`
}

// GetDebugName return the name for debug.
func (dto *GetIPv4PoolResponse) GetDebugName() string {
	return dto.Name
}

// Load the data from model.
func (dto *GetIPv4PoolResponse) Load(data base.ModelInterface) error {
	m, ok := data.(*model.IPv4Pool)
	if !ok {
		log.Error("GetIPv4PoolResponse.Load() failed, convert interface failed.")
		return base.ErrorDataConvert
	}
	base.ResponseLoad(&dto.Response, &m.Model)
	dto.Name = m.Name
	dto.Description = m.Description
	dto.Ranges = make([]IPv4RangeResponse, 0)
	for _, v := range m.Ranges {
		vv := IPv4RangeResponse{}
		vv.Start = v.Start
		dto.Total += v.Total
		dto.Free += v.Free
		dto.Allocatable += v.Allocatable
		dto.Ranges = append(dto.Ranges, IPv4RangeResponse{
			Start:       v.Start,
			End:         v.End,
			Total:       v.Total,
			Allocatable: v.Allocatable,
			Free:        v.Free,
		})
	}
	dto.SubnetMask = m.SubnetMask
	dto.Gateway = m.Gateway
	dto.Domain = m.Domain
	dto.DNSServers = m.DNSServers
	return nil
}
