package service

import (
	log "github.com/sirupsen/logrus"
	"promise/base"
	enclosureClient "promise/enclosure/client/enclosure"
	"promise/enclosure/object/dto"
)

// Discover is the service for discover enclosure action.
type Discover struct {
}

// Perform will process the discover action.
func (s *Discover) Perform(id string, request base.ActionRequestInterface) (base.ResponseInterface, []base.ErrorResponse) {
	var (
		response    dto.GetEnclosureResponse
		clientError base.ClientError
		identity    *base.DeviceIdentity
	)

	discoverRequest, ok := request.(*dto.DiscoverEnclosureRequest)
	if !ok {
		log.Error("Service perform discover enclosure failed, convert request failed.")
		return nil, []base.ErrorResponse{*base.NewErrorResponseInternalError()}
	}

	enclosure := discoverRequest.NewEnclosure()
	client := enclosureClient.NewClient(enclosure)
	if client == nil {
		log.Warn("Service discover enclosure failed, create client failed, discover abort.")
		return nil, []base.ErrorResponse{*base.NewErrorResponseInternalError()}
	}
	// get identity and check existance.
	identity, clientError = client.DeviceIdentity()
	if clientError != nil {
		log.WithFields(log.Fields{"error": clientError}).Warn("Service discover enclosure failed, get device identity failed, discover abort.")
		return nil, []base.ErrorResponse{*base.NewErrorResponseFromClientError(clientError)}
	}
	enclosure.DeviceIdentity = *identity
	if exist, _ := enclosureDB.Exist(enclosure); exist {
		return nil, []base.ErrorResponse{*base.NewErrorResponseDuplicate()}
	}
	// save
	created, errResp := enclosureDB.Create(enclosure)
	if errResp != nil {
		return nil, []base.ErrorResponse{*errResp}
	}
	if err := response.Load(created); err != nil {
		return nil, []base.ErrorResponse{*base.NewErrorResponseInternalError()}
	}
	return &response, nil
}
