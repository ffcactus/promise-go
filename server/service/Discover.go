package service

import (
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/server/context"
	"promise/server/object/dto"
	"promise/server/object/errorResp"
	"promise/server/strategy"
)

// Discover is the service for discover server action.
type Discover struct {
}

// Perform will process the discover action.
func (s *Discover) Perform(id string, request base.ActionRequestInterface) (base.ResponseInterface, []base.ErrorResponse) {
	var (
		response dto.GetServerResponse
	)
	discoverRequest, ok := request.(*dto.DiscoverServerRequest)
	if !ok {
		log.Error("Service perform discover server failed, convert request failed.")
		return nil, []base.ErrorResponse{*base.NewErrorResponseInternalError()}
	}
	serverBasicInfo, err := Probe(discoverRequest)
	server := serverBasicInfo.CreateServer()
	ctx := context.CreateDiscoverServerContext(server, discoverRequest)
	st := strategy.CreateDiscoverServerStrategy(server)
	model, err := st.Execute(ctx, server)
	if err != nil {
		return nil, []base.ErrorResponse{*errorResp.NewErrorResponseServerDiscoverFailed()}
	}
	if err := response.Load(model); err != nil {
		return nil, []base.ErrorResponse{*base.NewErrorResponseInternalError()}
	}
	return &response, nil
}
