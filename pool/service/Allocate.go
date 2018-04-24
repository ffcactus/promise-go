package service

import (
	"promise/base"
	"promise/pool/object/dto"
)

// Allocate is the service implement to allocate IPv4.
type Allocate struct {
}

// Perform the allocate IPv4 action.
func (s *Allocate) Perform(id string, request base.ActionRequestInterface) (base.ResponseInterface, []base.Message) {
	var (
		response dto.AllocateIPv4Response
	)

	allocateRequest, ok := request.(*dto.AllocateIPv4Request)
	if !ok {
		return nil, []base.Message{*base.NewMessageInternalError()}
	}
	key := ""
	if allocateRequest.Key != nil {
		key = *allocateRequest.Key
	}
	address, updatedPool, message := ipv4PoolDB.AllocateIPv4Address(id, key)
	if message != nil {
		return nil, []base.Message{*message}
	}
	response.Address = address
	if err := response.Pool.Load(updatedPool); err != nil {
		return nil, []base.Message{*base.NewMessageInternalError()}
	}
	eventService.DispatchUpdateEvent(&response.Pool)
	return &response, nil
}
