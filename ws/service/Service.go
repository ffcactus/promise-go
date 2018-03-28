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
func AddListener(listener *websocket.Conn) int {
	wsConnection.PushBack(listener)
	return wsConnection.Len()
}

// StartEventDispatcher Start the event dispater.
func StartEventDispatcher() {
	for {
		e := <-EventChannel
		count := 0
		var next *list.Element
		for each := wsConnection.Front(); each != nil; each = next {
			next = each.Next()
			if err := each.Value.(*websocket.Conn).WriteMessage(websocket.TextMessage, []byte(util.StructToString(e))); err != nil {
				log.WithFields(log.Fields{"error": err, "remain": wsConnection.Len()}).Info("Send message to the listener failed, remove the listener.")				
				wsConnection.Remove(each)
				
			} else {
				count++
			}
		}
		if count > 0 {
			log.WithFields(log.Fields{"count": count, "type": e.Type, "category": e.Category, "resource": e.ResourceID}).Info("Event dispatched.")
		}
	}
}

// DispatchEvent will push the event to the pipe.
func DispatchEvent(e *dto.PostEventRequest) {
	EventChannel <- e
}
