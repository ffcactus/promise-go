package dto

import (
    log "github.com/sirupsen/logrus"
	"promise/common/object/constError"
	"promise/pool/object/model"
	commonDTO "promise/common/object/dto"
)

// GetIPv4PoolResponse is the response DTO.
type GetIPv4PoolResponse struct {
	commonDTO.PromiseResponse
	IPv4PoolResource
}

// Load the data from model.
func (dto *GetIPv4PoolResponse) Load(data interface{}) error {
	m, ok := data.(*model.IPv4Pool)
	if !ok {
		log.Warn("GetIPv4PoolResponse load data from model failed.")
		return constError.ErrorDataConvert
	}
	dto.PromiseResponse.Load(&m.PromiseModel)
	dto.Name = m.Name
	dto.Description = m.Description
	dto.Ranges = make([]IPv4Range, 0)
	for _, v := range m.Ranges {
		dto.Ranges = append(dto.Ranges, IPv4Range{Start: v.Start, End: v.End})
	}
	dto.SubnetMask = m.SubnetMask
	dto.Gateway = m.Gateway
	dto.Domain = m.Domain
	dto.DNSServers = m.DNSServers
	return nil
}