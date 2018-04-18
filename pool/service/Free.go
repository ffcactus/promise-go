package service

import (
	"promise/base"
	"promise/pool/object/dto"
)

// Free is the service implement to allocate IPv4.
type Free struct {
}

// Perform the free IPv4 action.
func (s *Free) Perform(id string, request base.ActionRequestInterface) (interface{}, []base.Message) {
	var (
		response dto.GetIPv4PoolResponse
	)

	freeRequest, ok := request.(*dto.FreeIPv4Request)
	if !ok {
		return nil, []base.Message{base.NewMessageInternalError()}
	}

	exist, updatedPool, commited, err := ipv4PoolDB.FreeIPv4Address(id, freeRequest.Address)
	if !exist {
		return nil, []base.Message{base.NewMessageNotExist()}
	}
	if err != nil && err.Error() == base.ErrorIPv4NotInPool.Error() {
		return nil, []base.Message{base.NewMessageIPv4AddressNotExistError()}
	}
	if err != nil && err.Error() == base.ErrorIPv4NotAllocated.Error() {
		return nil, []base.Message{base.NewMessageIPv4NotAllocatedError()}
	}
	if err != nil || !commited {
		return nil, []base.Message{base.NewMessageTransactionError()}
	}
	response.Load(updatedPool)
	eventService.DispatchUpdateEvent(&response)
	return &response, nil
}
