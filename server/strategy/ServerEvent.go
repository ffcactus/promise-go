package strategy

import (
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/sdk/event"
	"promise/server/context"
	"promise/server/object/dto"
)

/**
 * There are many error cases have to deal with, so we put event dispatch
 * to it's own strategy.
 */

// ServerEvent is the server event strategy implementation.
type ServerEvent struct {
}

func (s *ServerEvent) dispatchServerEvent(c *context.Base, eventType string, server base.ModelInterface) {
	if server == nil {
		log.WithFields(log.Fields{
			"type":  eventType,
		}).Error("Strategy dispatch server event failed, server equals nil.")
		return	
	}
	var serverDTO = new(dto.GetServerResponse)
	if err := serverDTO.Load(server); err != nil {
		log.WithFields(log.Fields{
			"id":    server.GetID(),
			"type":  eventType,
			"error": err,
		}).Warn("Strategy dispatch server event failed, create event failed.")
		return
	}
	messages, err := event.DispatchResourceEvent(eventType, serverDTO)
	if err != nil {
		log.WithFields(log.Fields{
			"id":    server.GetID(),
			"type":  eventType,
			"error": err,
		}).Warn("Strategy dispatch server event failed, event dispatching failed.")
	}
	if messages != nil {
		log.WithFields(log.Fields{
			"id":      server.GetID(),
			"type":    eventType,
			"message": messages[0].ID,
		}).Warn("Strategy dispatch server create event failed, message returned, event dispatching failed.")
	}
}

// DispatchServerCreate will send server create event.
// Generally we don't care much about the error while sending event.
// If the server created, but event failed to dispatch, the error won't return to user.
func (s *ServerEvent) DispatchServerCreate(c *context.Base, server base.ModelInterface) {
	s.dispatchServerEvent(c, base.CreateEvent, server)
}

// DispatchServerUpdate will send server update event.
// Generally we don't care much about the error while sending event.
// If the server created, but event failed to dispatch, the error won't return to user.
func (s *ServerEvent) DispatchServerUpdate(c *context.Base, server base.ModelInterface) {
	s.dispatchServerEvent(c, base.UpdateEvent, server)
}

// DispatchServerDelete will send server delete event.
// Generally we don't care much about the error while sending event.
// If the server created, but event failed to dispatch, the error won't return to user.
func (s *ServerEvent) DispatchServerDelete(c *context.Base, server base.ModelInterface) {
	s.dispatchServerEvent(c, base.DeleteEvent, server)
}

func (s *ServerEvent) dispatchServerServerGroupEvent(c *context.Base, eventType string, ssg base.ModelInterface) {
	var ssgDTO = new(dto.GetServerServerGroupResponse)
	if err := ssgDTO.Load(ssg); err != nil {
		log.WithFields(log.Fields{
			"id":    ssg.GetID(),
			"type":  eventType,
			"error": err,
		}).Warn("Strategy dispatch server-servergroup event failed, create event failed.")
		return
	}
	messages, err := event.DispatchResourceEvent(eventType, ssgDTO)
	if err != nil {
		log.WithFields(log.Fields{
			"id":    ssg.GetID(),
			"type":  eventType,
			"error": err,
		}).Warn("Strategy dispatch server-servergroup event failed, event dispatching failed.")
	}
	if messages != nil {
		log.WithFields(log.Fields{
			"id":      ssg.GetID(),
			"type":    eventType,
			"message": messages[0].ID},
		).Warn("Strategy dispatch server-servergroup create event failed, message returned, event dispatching failed.")
	}
}

// DispatchServerServerGroupCreate will send server-servergroup create event.
// Generally we don't care much about the error while sending event.
// If the server created, but event failed to dispatch, the error won't return to user.
func (s *ServerEvent) DispatchServerServerGroupCreate(c *context.Base, ssg base.ModelInterface) {
	s.dispatchServerServerGroupEvent(c, base.CreateEvent, ssg)
}

// DispatchServerServerGroupDelete will send server-servergroup create event.
// Generally we don't care much about the error while sending event.
// If the server created, but event failed to dispatch, the error won't return to user.
func (s *ServerEvent) DispatchServerServerGroupDelete(c *context.Base, ssg base.ModelInterface) {
	s.dispatchServerServerGroupEvent(c, base.DeleteEvent, ssg)
}
