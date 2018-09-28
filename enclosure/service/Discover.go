package service

import (
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/enclosure/object/dto"
	enclosureClient "promise/enclosure/client/enclosure"
)

// Discover is the service for discover enclosure action.
type Discover struct {
}

// Perform will process the discover action.
func (s *Discover) Perform(id string, request base.ActionRequestInterface) (base.ResponseInterface, []base.ErrorResponse) {
	var (
		response dto.GetEnclosureResponse
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
	identity, clientError := client.DeviceIdentity()
	if clientError != nil {
		log.WithFields(log.Fields{"error": clientError}).Warn("Service discover enclosure failed, get device identity failed, discover abort.")
		return nil, []base.ErrorResponse{*base.NewErrorResponseFromClientError(clientError)}
	}
	enclosure.DeviceIdentity = *identity
	return &response, nil
}
