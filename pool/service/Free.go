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
	var response dto.GetIPv4PoolResponse

	exist, updatedPool, commited, err := ipv4PoolDB.Update(id, request)
	if !exist {
		return nil, []base.Message{base.NewMessageNotExist()}
	}
	// if err != nil && err.Error() == base.ErrorUnknownPropertyValue.Error() {
	// 	return nil, []base.Message{base.NewMessageUnknownPropertyValue()}
	// }
	if err != nil || !commited {
		return nil, []base.Message{base.NewMessageTransactionError()}
	}
	response.Load(updatedPool)
	eventService.DispatchUpdateEvent(&response)
	return &response, nil
}
