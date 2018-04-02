package dto

import (
	"promise/common/category"
	commonModel "promise/common/object/model"
	"promise/pool/object/model"
)

// PostIPv4PoolRequest is the request DTO.
type PostIPv4PoolRequest struct {
	commonModel.PromiseModel
	IPv4PoolResource
}

// ToModel will convert the DTO to model.
func (dto *PostIPv4PoolRequest) ToModel() *model.IPv4Pool {
	ret := model.IPv4Pool{}
	ret.Category = category.PoolIPv4
	ret.Name = dto.Name
	ret.Description = dto.Description
	ret.SubnetMask = dto.SubnetMask
	ret.Gateway = dto.Gateway
	ret.Domain = dto.Domain
	ret.DNSServers = dto.DNSServers
	for _, v := range dto.Ranges {
		ret.Ranges = append(ret.Ranges, model.IPv4Range{Start: v.Start, End: v.End})
	}
	return &ret
}
