package service

import (
	"promise/base"
	"promise/pool/object/dto"
)

// Free is the service implement to allocate IPv4.
type Free struct {
}

// Perform the free IPv4 action.
func (s *Free) Perform(id string, request base.ActionRequestInterface) (base.ResponseInterface, []base.ErrorResponse) {
	var (
		response dto.GetIPv4PoolResponse
	)

	freeRequest, ok := request.(*dto.FreeIPv4Request)
	if !ok {
		return nil, []base.ErrorResponse{*base.NewErrorResponseInternalError()}
	}

	updatedPool, errorResp := ipv4PoolDB.FreeIPv4Address(id, freeRequest.Address)
	if errorResp != nil {
		return nil, []base.ErrorResponse{*errorResp}
	}
	response.Load(updatedPool)
	base.PublishUpdateMessage(&response)
	return &response, nil
}
