package service

import (
	"promise/base"
	"promise/pool/object/dto"
)

// Allocate is the service implement to allocate IPv4.
type Allocate struct {
}

// Perform the allocate IPv4 action.
func (s *Allocate) Perform(id string, request base.ActionRequestInterface) (interface{}, []base.Message) {
	var (
		response dto.AllocateIPv4Response
	)

	allocateRequest, ok := request.(*dto.AllocateIPv4Request)
	if !ok {
		return nil, []base.Message{base.NewMessageInternalError()}
	}
	key := ""
	if allocateRequest.Key != nil {
		key = *allocateRequest.Key
	}
	exist, address, updatedPool, commited, err := ipv4PoolDB.AllocateIPv4Address(id, key)
	if !exist {
		return nil, []base.Message{base.NewMessageNotExist()}
	}
	if exist && address == "" && !commited && err == nil {
		return nil, []base.Message{base.NewMessageIPv4PoolEmpty()}
	}
	if err != nil || !commited {
		return nil, []base.Message{base.NewMessageTransactionError()}
	}
	response.Address = address
	if err := response.Pool.Load(updatedPool); err != nil {
		return nil, []base.Message{base.NewMessageInternalError()}
	}
	eventService.DispatchUpdateEvent(&response.Pool)
	return &response, nil
}
