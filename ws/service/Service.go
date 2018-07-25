package service

import (
	"container/list"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"promise/base"
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

// StartDispatcher will recieve all the message and dispatch them as websocket.
func StartDispatcher() {
	base.InitMQService()
	defer base.StopMQService()
	base.Subscribe([]string{"*.*"}, handler)
}

// handler will handle the event.
func handler(d *amqp.Delivery) {
	count := 0
	var next *list.Element
	for each := wsConnection.Front(); each != nil; each = next {
		next = each.Next()
		if err := each.Value.(*websocket.Conn).WriteMessage(websocket.TextMessage, d.Body); err != nil {
			log.WithFields(log.Fields{
				"error":  err,
				"remain": wsConnection.Len(),
			}).Info("Send message to the listener failed, remove the listener.")
			wsConnection.Remove(each)

		} else {
			count++
		}
	}
	if count > 0 {
		log.WithFields(log.Fields{"count": count}).Info("Event dispatched.")
	}
}
