package dto

import (
	"net"
	"promise/base"
	"promise/pool/object/errorResp"
	"promise/pool/object/model"
)

// PostIPv4PoolRequest is the request DTO.
type PostIPv4PoolRequest struct {
	IPv4PoolResource
	Ranges []IPv4RangeRequest `json:"Ranges"`
}

// NewInstance creates a new instance.
func (PostIPv4PoolRequest) NewInstance() base.RequestInterface {
	return new(PostIPv4PoolRequest)
}

// IsValid return if the request is valid.
func (dto *PostIPv4PoolRequest) IsValid() *base.ErrorResponse {
	if dto.SubnetMask != nil && net.ParseIP(*dto.SubnetMask) == nil {
		return errorResp.NewErrorResponseIPv4FormatError()
	}
	if dto.Gateway != nil && net.ParseIP(*dto.Gateway) == nil {
		return errorResp.NewErrorResponseIPv4FormatError()
	}
	if dto.SubnetMask != nil && net.ParseIP(*dto.SubnetMask) == nil {
		return errorResp.NewErrorResponseIPv4FormatError()
	}
	if dto.DNSServers != nil {
		for _, v := range *dto.DNSServers {
			if net.ParseIP(v) == nil {
				return errorResp.NewErrorResponseIPv4FormatError()
			}
		}
	}

	if len(dto.Ranges) == 0 {
		return errorResp.NewErrorResponseIPv4RangeCountError()
	}
	for _, v := range dto.Ranges {
		start := net.ParseIP(v.Start)
		end := net.ParseIP(v.End)
		if start == nil {
			return errorResp.NewErrorResponseIPv4FormatError()
		}
		if end == nil {
			return errorResp.NewErrorResponseIPv4FormatError()
		}
		if base.IPtoInt(start) > base.IPtoInt(end) {
			return errorResp.NewErrorResponseIPv4RangeEndAddressError()
		}
		if base.IPtoInt(end)-base.IPtoInt(start)+1 > 256 {
			return errorResp.NewErrorResponseIPv4RangeSizeError()
		}
	}
	return nil
}

// String return the name for debug.
func (dto PostIPv4PoolRequest) String() string {
	return dto.Name
}

// ToModel convert the DTO to model.
func (dto *PostIPv4PoolRequest) ToModel() base.ModelInterface {
	ret := model.IPv4Pool{}
	ret.Category = base.CategoryPoolIPv4
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
		vv.Total = (uint32)(base.IPtoInt(end) - base.IPtoInt(start) + 1)
		vv.Free = vv.Total
		vv.Allocatable = vv.Total
		for i := base.IPtoInt(start); i <= base.IPtoInt(end); i++ {
			address := model.IPv4Address{}
			address.Address = base.IntToIP(i).String()
			address.Allocated = false
			vv.Addresses = append(vv.Addresses, address)
		}
		ret.Ranges = append(ret.Ranges, vv)
	}
	return &ret
}
