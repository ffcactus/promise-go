package service

import (
	"promise/base"
	"promise/pool/object/dto"
)

// Allocate is the service implement to allocate IPv4.
type Allocate struct {
}

// Perform the allocate IPv4 action.
func (s *Allocate) Perform(id string, request base.ActionRequestInterface) (base.ResponseInterface, []base.ErrorResponse) {
	var (
		response dto.AllocateIPv4Response
	)

	allocateRequest, ok := request.(*dto.AllocateIPv4Request)
	if !ok {
		return nil, []base.ErrorResponse{*base.NewErrorResponseInternalError()}
	}
	key := ""
	if allocateRequest.Key != nil {
		key = *allocateRequest.Key
	}
	address, updatedPool, errorResp := ipv4PoolDB.AllocateIPv4Address(id, key)
	if errorResp != nil {
		return nil, []base.ErrorResponse{*errorResp}
	}
	response.Address = address
	if err := response.Pool.Load(updatedPool); err != nil {
		return nil, []base.ErrorResponse{*base.NewErrorResponseInternalError()}
	}
	base.PublishUpdateMessage(&response.Pool)
	return &response, nil
}
