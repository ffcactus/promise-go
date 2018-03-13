package service

import (
	"container/list"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"promise/common/util"
)

var (
	// CreateEvent create event.
	CreateEvent = "Create"
	// UpdateEvent update event.
	UpdateEvent = "Update"
	// DeleteEvent delete event.
	DeleteEvent = "Delete"

	// ServerCategory server category
	ServerCategory = "Server"
)

// Event The event object.
type Event struct {
	Type     string
	URI      string
	Category string
	Message  string
}

var (
	// EventChannel The event channel
	EventChannel = make(chan *Event, 10)
	wsConnection = list.New()
)

// AddListener Add a listener
func AddListener(listener *websocket.Conn) {
	wsConnection.PushBack(listener)
	log.Info("EventDispatcher add listener.")
}

// StartEventDispatcher Start the event dispater.
func StartEventDispatcher() {
	for {
		e := <-EventChannel
		log.Debug("StartEventDispatcher(), event type =", e.Type, "Category =", e.Category, "URI =", e.URI)
		for each := wsConnection.Front(); each != nil; each = each.Next() {
			if each.Value.(*websocket.Conn).WriteMessage(websocket.TextMessage, []byte(util.StructToString(e))) != nil {
				log.Info("EventDispatcher remove listener.")
				wsConnection.Remove(each)
			}
		}
	}
}

// DispatchEvent will push the event to the pipe.
func DispatchEvent(e *Event) {
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
