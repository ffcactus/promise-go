package service

import (
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/server/context"
	"promise/server/object/dto"
	"promise/server/strategy"
)

// Discover is the service for discover server action.
type Discover struct {
}

// Perform will process the action.
func (s *Discover) Perform(id string, request ActionRequestInterface) (base.ResponseInterface, []base.Message) {
	discoverRequest, ok := request.(*dto.DiscoverServerRequest)
	if !ok {
		log.Error("Perform discover server failed, convert request failed.")
		return nil, []base.Message{base.NewMessageInternalError()}
	}
	serverBasicInfo, err := Probe(request)
	server := serverBaiscInfo.CreateServer()
	ctx := context.CreatePostServerContext(server, discoverRequest)
	st := strategy.CreatePostServerStrategy(server)
	response, _, messages := st.Execute(ctx)
	getResponse, ok := response.(base.GetResponseInterfaced)
	if !ok {
		log.Error("Perform discover server failed, convert response failed.")
		return nil, []base.Message{base.NewMessageInternalError()}
	}
	eventService.DispatchCreateEvent(getResponse)
	return getResponse, messages
}

