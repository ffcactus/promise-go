package strategy

import (
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/sdk/event"
	"promise/server/context"
	"promise/server/object/dto"
	"promise/server/object/model"
)

/**
 * There are many error cases have to deal with, so we put event dispatch
 * to it's own strategy.
 */

// ServerEventStrategy is the server event strategy implementation.
type ServerEventStrategy struct {
}

func (s *ServerEventStrategy) dispatchServerEvent(c *context.ServerContext, eventType string, server *model.Server) {
	var serverDTO = new(dto.GetServerResponse)
	if err := serverDTO.Load(server); err != nil {
		log.WithFields(log.Fields{
			"id":    server.ID,
			"type":  eventType,
			"error": err}).Warn("Dispatch server event failed, create event failed.")
		return
	}
	messages, err := event.DispatchResourceEvent(eventType, serverDTO)
	if err != nil {
		log.WithFields(log.Fields{
			"id":    server.ID,
			"type":  eventType,
			"error": err}).
			Warn("Dispatch server event failed, event dispatching failed.")
	}
	if messages != nil {
		log.WithFields(log.Fields{
			"id":      server.ID,
			"type":    eventType,
			"message": messages[0].ID}).
			Warn("Dispatch server create event failed, message returned, event dispatching failed.")
	}
}

// DispatchServerCreate will send server create event.
// Generally we don't care much about the error while sending event.
// If the server created, but event failed to dispatch, the error won't return to user.
func (s *ServerEventStrategy) DispatchServerCreate(c *context.ServerContext, server *model.Server) {
	s.dispatchServerEvent(c, base.CreateEvent, server)
}

// DispatchServerUpdate will send server update event.
// Generally we don't care much about the error while sending event.
// If the server created, but event failed to dispatch, the error won't return to user.
func (s *ServerEventStrategy) DispatchServerUpdate(c *context.ServerContext, server *model.Server) {
	s.dispatchServerEvent(c, base.UpdateEvent, server)
}

// DispatchServerDelete will send server delete event.
// Generally we don't care much about the error while sending event.
// If the server created, but event failed to dispatch, the error won't return to user.
func (s *ServerEventStrategy) DispatchServerDelete(c *context.ServerContext, server *model.Server) {
	s.dispatchServerEvent(c, base.DeleteEvent, server)
}

func (s *ServerEventStrategy) dispatchServerServerGroupEvent(c *context.ServerContext, eventType string, ssg *model.ServerServerGroup) {
	var ssgDTO = new(dto.GetServerServerGroupResponse)
	if err := ssgDTO.Load(ssg); err != nil {
		log.WithFields(log.Fields{
			"id":          ssg.ID,
			"type":        eventType,
			"server":      ssg.ServerID,
			"servergroup": ssg.ServerGroupID,
			"error":       err,
		}).Warn("Dispatch server-servergroup event failed, create event failed.")
		return
	}
	messages, err := event.DispatchResourceEvent(eventType, ssgDTO)
	if err != nil {
		log.WithFields(log.Fields{
			"id":          ssg.ID,
			"type":        eventType,
			"server":      ssg.ServerID,
			"servergroup": ssg.ServerGroupID,
			"error":       err,
		}).Warn("Dispatch server-servergroup event failed, event dispatching failed.")
	}
	if messages != nil {
		log.WithFields(log.Fields{
			"id":          ssg.ID,
			"type":        eventType,
			"server":      ssg.ServerID,
			"servergroup": ssg.ServerGroupID,
			"message":     messages[0].ID},
		).Warn("Dispatch server-servergroup create event failed, message returned, event dispatching failed.")
	}
}

// DispatchServerServerGroupCreate will send server-servergroup create event.
// Generally we don't care much about the error while sending event.
// If the server created, but event failed to dispatch, the error won't return to user.
func (s *ServerEventStrategy) DispatchServerServerGroupCreate(c *context.ServerContext, ssg *model.ServerServerGroup) {
	s.dispatchServerServerGroupEvent(c, base.CreateEvent, ssg)
}

// DispatchServerServerGroupDelete will send server-servergroup create event.
// Generally we don't care much about the error while sending event.
// If the server created, but event failed to dispatch, the error won't return to user.
func (s *ServerEventStrategy) DispatchServerServerGroupDelete(c *context.ServerContext, ssg *model.ServerServerGroup) {
	s.dispatchServerServerGroupEvent(c, base.DeleteEvent, ssg)
}
