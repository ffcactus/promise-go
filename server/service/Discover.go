package service

import (
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/server/context"
	"promise/server/object/dto"
	"promise/server/object/message"
	"promise/server/strategy"
)

// Discover is the service for discover server action.
type Discover struct {
}

// Perform will process the discover action.
func (s *Discover) Perform(id string, request base.ActionRequestInterface) (base.ResponseInterface, []base.Message) {
	discoverRequest, ok := request.(*dto.DiscoverServerRequest)
	if !ok {
		log.Error("Perform discover server failed, convert request failed.")
		return nil, []base.Message{*base.NewMessageInternalError()}
	}
	serverBasicInfo, err := Probe(discoverRequest)
	server := serverBasicInfo.CreateServer()
	ctx := context.CreateDiscoverServerContext(server, discoverRequest)
	st := strategy.CreateDiscoverServerStrategy(server)
	response, err := st.Execute(ctx, server)
	if err != nil {
		return nil, []base.Message{*message.NewMessageServerDiscoverFailed()}
	}
	getResponse, ok := response.(base.GetResponseInterface)
	if !ok {
		log.Error("Perform discover server failed, convert response failed.")
		return nil, []base.Message{*base.NewMessageInternalError()}
	}
	eventService.DispatchCreateEvent(getResponse)
	return getResponse, nil
}
