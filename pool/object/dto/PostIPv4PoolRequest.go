package dto

import (
	"net"
	"promise/common/category"
	commonDTO "promise/common/object/dto"
	"promise/pool/object/model"
	"promise/pool/util"
)

// PostIPv4PoolRequest is the request DTO.
type PostIPv4PoolRequest struct {
	commonDTO.PromiseRequest
	IPv4PoolResource
	Ranges []IPv4RangeRequest `json:"Ranges"`
}

// IsValid check if the request is valid.
func (dto *PostIPv4PoolRequest) IsValid() bool {
	if dto.SubnetMask != nil && net.ParseIP(*dto.SubnetMask) == nil {
		return false
	}
	if dto.Gateway != nil && net.ParseIP(*dto.Gateway) == nil {
		return false
	}
	if dto.SubnetMask != nil && net.ParseIP(*dto.SubnetMask) == nil {
		return false
	}
	if dto.DNSServers != nil {
		for _, v := range *dto.DNSServers {
			if net.ParseIP(v) == nil {
				return false
			}
		}
	}

	for _, v := range dto.Ranges {
		start := net.ParseIP(v.Start)
		end := net.ParseIP(v.End)
		if start == nil {
			return false
		}
		if end == nil {
			return false
		}
		if util.IPtoInt(start) > util.IPtoInt(end) {
			return false
		}
		if util.IPtoInt(end)-util.IPtoInt(start)+1 > 256 {
			return false
		}
	}
	return true
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
		start := net.ParseIP(v.Start)
		end := net.ParseIP(v.End)
		vv := model.IPv4Range{}

		vv.Start = v.Start
		vv.End = v.End
		vv.Total = (uint32)(util.IPtoInt(end) - util.IPtoInt(start) + 1)
		vv.Free = vv.Total
		vv.Allocatable = vv.Total
		for i := util.IPtoInt(start); i <= util.IPtoInt(end); i++ {
			address := model.IPv4Address{}
			address.Address = util.IntToIP(i).String()
			address.Allocated = false
			vv.Addresses = append(vv.Addresses, address)
		}
		ret.Ranges = append(ret.Ranges, vv)
	}
	return &ret
}
