package strategy

import (
	log "github.com/sirupsen/logrus"
	wsSDK "promise/sdk/ws"
	"promise/server/object/dto"
	"promise/server/context"
	"promise/server/object/model"
	"promise/common/category"
)

/**
 * There are many error cases have to deal with, so we put event dispatch
 * to it's own strategy.
 */

// ServerEventStrategy is the server event strategy implementation.
type ServerEventStrategy struct {

}

// DispatchServerCreate will send server create event.
// Generally we don't care much about the error while sending event.
// If the server created, but event failed to dispatch, the error won't return to user.
func (s *ServerEventStrategy) DispatchServerCreate(c *context.ServerContext, server *model.Server) {
	var serverDTO = new(dto.GetServerResponse)
	if err := serverDTO.Load(server); err != nil {
		log.WithFields(log.Fields{"id": server.ID, "error": err}).Warn("Dispatch server create event failed, create event failed.")
		return
	}
	messages, err := wsSDK.DispatchResourceCreate(serverDTO)
	if err != nil {
		log.WithFields(log.Fields{
			"id": server.ID, 
			"error": err}).
			Warn("Dispatch server create event failed, event dispatching failed.")
	}
	if messages != nil {
		log.WithFields(log.Fields{
			"id": server.ID, 
			"message": messages[0].ID}).
			Warn("Dispatch server create event failed, event dispatching failed.")
	}
}

// DispatchServerUpdate will send server update event.
// Generally we don't care much about the error while sending event.
// If the server created, but event failed to dispatch, the error won't return to user.
func (s *ServerEventStrategy) DispatchServerUpdate(c *context.ServerContext, server *model.Server) {
	var serverDTO = new(dto.GetServerResponse)
	if err := serverDTO.Load(server); err != nil {
		log.WithFields(log.Fields{"id": server.ID, "error": err}).Warn("Dispatch server update event failed, create event failed.")
		return
	}
	messages, err := wsSDK.DispatchResourceUpdate(serverDTO)
	if err != nil {
		log.WithFields(log.Fields{
			"id": server.ID, 
			"error": err}).
			Warn("Dispatch server update event failed, event dispatching failed.")
	}
	if messages != nil {
		log.WithFields(log.Fields{
			"id": server.ID, 
			"message": messages[0].ID}).
			Warn("Dispatch server update event failed, event dispatching failed.")
	}
}

// DispatchServerDelete will send server delete event.
// Generally we don't care much about the error while sending event.
// If the server created, but event failed to dispatch, the error won't return to user.
func (s *ServerEventStrategy) DispatchServerDelete(c *context.ServerContext, id string) {
	messages, err := wsSDK.DispatchResourceDelete(category.Server, id)
	if err != nil {
		log.WithFields(log.Fields{
			"id": id, 
			"error": err}).
			Warn("Dispatch server delete event failed, event dispatching failed.")
	}
	if messages != nil {
		log.WithFields(log.Fields{
			"id": id, 
			"message": messages[0].ID}).
			Warn("Dispatch server delete event failed, event dispatching failed.")
	}
}

// DispatchServerServerGroupCreate will send server-servergroup create event.
// Generally we don't care much about the error while sending event.
// If the server created, but event failed to dispatch, the error won't return to user.
func (s *ServerEventStrategy) DispatchServerServerGroupCreate(c *context.ServerContext, ssg *model.ServerServerGroup) {
	var ssgDTO = new(dto.GetServerServerGroupResponse)
	if err := ssgDTO.Load(ssg); err != nil {
		log.WithFields(log.Fields{
			"id": ssg.ID, 
			"server": ssg.ServerID,
			"servergroup": ssg.ServerGroupID,
			"error": err}).Warn("Dispatch server-servergroup create event failed, create event failed.")
		return
	}
	messages, err := wsSDK.DispatchResourceCreate(ssgDTO)
	if err != nil {
		log.WithFields(log.Fields{
			"id": ssg.ID, 
			"server": ssg.ServerID,
			"servergroup": ssg.ServerGroupID,
			"error": err}).
			Warn("Dispatch server-servergroup create event failed, event dispatching failed.")
	}
	if messages != nil {
		log.WithFields(log.Fields{
			"id": ssg.ID, 
			"server": ssg.ServerID,
			"servergroup": ssg.ServerGroupID,
			"message": messages[0].ID}).
			Warn("Dispatch server-servergroup create event failed, event dispatching failed.")
	}
}