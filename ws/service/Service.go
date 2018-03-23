package service

import (
	"container/list"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"promise/common/util"
	"promise/ws/object/dto"
)

var (
	// EventChannel The event channel
	EventChannel = make(chan *dto.PostEventRequest, 10)
	wsConnection = list.New()
)

// AddListener Add a listener
func AddListener(listener *websocket.Conn) {
	wsConnection.PushBack(listener)
}

// StartEventDispatcher Start the event dispater.
func StartEventDispatcher() {
	for {
		e := <-EventChannel		
		for each := wsConnection.Front(); each != nil; each = each.Next() {
			if each.Value.(*websocket.Conn).WriteMessage(websocket.TextMessage, []byte(util.StructToString(e))) != nil {
				log.Info("EventDispatcher remove listener.")
				wsConnection.Remove(each)
			}
		}
		log.WithFields(log.Fields{"type": e.Type, "category": e.Category, "resource":e.ResourceID}).Info("Event dispatched.")
	}
}

// DispatchEvent will push the event to the pipe.
func DispatchEvent(e *dto.PostEventRequest) {
	EventChannel <- e
}

/*
// DispatchServerCreate Dispatch server created.
func (h *EventHandler) DispatchServerCreate(server *model.Server) {
	serverDto := dto.GetServerResponse{}
	serverDto.Load(server)
	event := Event{
		Type:     CreateEvent,
		URI:      server.URI,
		Category: ServerCategory,
		Message:  util.StructToString(serverDto),
	}
	EventChannel <- event
}

// DispatchServerUpdate Dispatch server updated.
func (h *EventHandler) DispatchServerUpdate(server *model.Server) {
	serverDto := dto.GetServerResponse{}
	serverDto.Load(server)
	event := Event{
		Type:     UpdateEvent,
		URI:      server.URI,
		Category: ServerCategory,
		Message:  util.StructToString(serverDto),
	}
	EventChannel <- event
}

// DispatchServerDelete Dispatch server deleted.
func (h *EventHandler) DispatchServerDelete(URI string) {
	event := Event{
		Type:     DeleteEvent,
		URI:      URI,
		Category: ServerCategory,
		Message:  "",
	}
	EventChannel <- event
}
*/
