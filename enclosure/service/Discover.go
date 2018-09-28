package service

import (
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/enclosure/object/dto"
)

// Discover is the service for discover enclosure action.
type Discover struct {
}

// Perform will process the discover action.
func (s *Discover) Perform(id string, request base.ActionRequestInterface) (base.ResponseInterface, []base.ErrorResponse) {
	var (
		response dto.GetEnclosureResponse
	)

	_, ok := request.(*dto.DiscoverEnclosureRequest)
	if !ok {
		log.Error("Service perform discover enclosure failed, convert request failed.")
		return nil, []base.ErrorResponse{*base.NewErrorResponseInternalError()}
	}
	return &response, nil
}
