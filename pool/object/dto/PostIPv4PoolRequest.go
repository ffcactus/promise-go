package dto

import (
	"net"
	"promise/base"
	"promise/pool/object/model"
)

// PostIPv4PoolRequest is the request DTO.
type PostIPv4PoolRequest struct {
	IPv4PoolResource
	Ranges []IPv4RangeRequest `json:"Ranges"`
}

// NewInstance creates a new instance.
func (dto *PostIPv4PoolRequest) NewInstance() base.RequestInterface {
	return new(PostIPv4PoolRequest)
}

// IsValid return if the request is valid.
func (dto *PostIPv4PoolRequest) IsValid() *base.Message {
	if dto.SubnetMask != nil && net.ParseIP(*dto.SubnetMask) == nil {
		message := base.NewMessageIPv4FormatError()
		return &message
	}
	if dto.Gateway != nil && net.ParseIP(*dto.Gateway) == nil {
		message := base.NewMessageIPv4FormatError()
		return &message
	}
	if dto.SubnetMask != nil && net.ParseIP(*dto.SubnetMask) == nil {
		message := base.NewMessageIPv4FormatError()
		return &message
	}
	if dto.DNSServers != nil {
		for _, v := range *dto.DNSServers {
			if net.ParseIP(v) == nil {
				message := base.NewMessageIPv4FormatError()
				return &message
			}
		}
	}

	if len(dto.Ranges) == 0 {
		message := base.NewMessageIPv4RangeCountError()
		return &message
	}
	for _, v := range dto.Ranges {
		start := net.ParseIP(v.Start)
		end := net.ParseIP(v.End)
		if start == nil {
			message := base.NewMessageIPv4FormatError()
			return &message
		}
		if end == nil {
			message := base.NewMessageIPv4FormatError()
			return &message
		}
		if base.IPtoInt(start) > base.IPtoInt(end) {
			message := base.NewMessageIPv4RangeEndAddressError()
			return &message
		}
		if base.IPtoInt(end)-base.IPtoInt(start)+1 > 256 {
			message := base.NewMessageIPv4RangeSizeError()
			return &message
		}
	}
	return nil
}

// DebugInfo return the name for debug.
func (dto *PostIPv4PoolRequest) DebugInfo() string {
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
