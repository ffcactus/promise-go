package ws

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
	"promise/common/app"
	"promise/common/app/rest"
	commonConstError "promise/common/object/constError"
	commonConstValue "promise/common/object/constValue"
	commonDTO "promise/common/object/dto"
	"promise/common/util"
	serverDTO "promise/server/object/dto"
	"promise/server/object/model"
	"promise/ws/object/constValue"
	wsDTO "promise/ws/object/dto"
	"time"
)

var (
	// WsServerRoot The root of the service.
	WsServerRoot = app.ProtocolScheme + app.Host + app.RootURL + commonConstValue.WSSenderBaseURI
)

func dispatchServerCreateOrUpdate(server *model.Server, eventType string) ([]commonDTO.Message, error) {
	var (
		s     serverDTO.GetServerResponse
		event wsDTO.PostEventRequest
	)
	event.CreatedAt = time.Now()
	event.Category = commonConstValue.CategoryServer
	event.Type = eventType // constValue.CreateEvent
	event.ResourceID = server.ID
	s.Load(server)
	b, err := json.Marshal(s)
	if err != nil {
		log.WithFields(log.Fields{
			"category": event.Category,
			"type":     event.Type,
			"resource": event.ResourceID,
			"error":    err}).Warn("Dispatch server event failed, failed to unmarshal resource.")
		return nil, commonConstError.ErrorDataConvert
	}
	event.Data = json.RawMessage(b)

	messages, err := rest.Do(
		http.MethodPost,
		WsServerRoot,
		event,
		nil,
		[]int{http.StatusCreated})
	return messages, err
}

// DispatchServerCreate Dispatch server created.
func DispatchServerCreate(server *model.Server) ([]commonDTO.Message, error) {
	return dispatchServerCreateOrUpdate(server, constValue.CreateEvent)
}

// DispatchServerUpdate Dispatch server updated.
func DispatchServerUpdate(server *model.Server) ([]commonDTO.Message, error) {
	return dispatchServerCreateOrUpdate(server, constValue.UpdateEvent)
}

// DispatchServerDelete Dispatch server deleted.
func DispatchServerDelete(id string) ([]commonDTO.Message, error) {
	var (
		event wsDTO.PostEventRequest
	)
	event.CreatedAt = time.Now()
	event.Category = commonConstValue.CategoryServer
	event.Type = constValue.DeleteEvent
	event.ResourceID = id
	messages, err := rest.Do(
		http.MethodPost,
		WsServerRoot,
		event,
		nil,
		[]int{http.StatusCreated})
	return messages, err
}

func dispatchResourceCreateOrUpdate(dto commonDTO.PromiseResponseInterface, eventType string) ([]commonDTO.Message, error) {
	var (
		event wsDTO.PostEventRequest
	)
	event.CreatedAt = time.Now()
	event.Category = dto.GetCategory()
	event.Type = eventType // constValue.CreateEvent
	event.ResourceID = dto.GetID()
	b, err := json.Marshal(dto)
	if err != nil {
		log.WithFields(log.Fields{
			"category": event.Category,
			"type":     event.Type,
			"resource": event.ResourceID,
			"error":    err}).Warn("Dispatch event failed, failed to unmarshal resource.")
		return nil, commonConstError.ErrorDataConvert
	}
	event.Data = json.RawMessage(b)

	messages, err := rest.Do(
		http.MethodPost,
		WsServerRoot,
		event,
		nil,
		[]int{http.StatusCreated})
	return messages, err
}

// DispatchServerGroupCreate Dispatch servergroup created.
func DispatchServerGroupCreate(sg *model.ServerGroup) ([]commonDTO.Message, error) {
	var dto serverDTO.GetServerGroupResponse
	dto.Load(sg)
	util.PrintJson(dto)
	return dispatchResourceCreateOrUpdate(&dto, constValue.CreateEvent)
}

// DispatchServerGroupUpdate Dispatch servergroup updated.
func DispatchServerGroupUpdate(sg *model.ServerGroup) ([]commonDTO.Message, error) {
	var dto serverDTO.GetServerGroupResponse
	dto.Load(sg)
	util.PrintJson(dto)
	return dispatchResourceCreateOrUpdate(&dto, constValue.UpdateEvent)
}

// DispatchServerGroupDelete Dispatch servergroup deleted.
func DispatchServerGroupDelete(id string) ([]commonDTO.Message, error) {
	var (
		event wsDTO.PostEventRequest
	)
	event.CreatedAt = time.Now()
	event.Category = commonConstValue.CategoryServerGroup
	event.Type = constValue.DeleteEvent
	event.ResourceID = id
	messages, err := rest.Do(
		http.MethodPost,
		WsServerRoot,
		event,
		nil,
		[]int{http.StatusCreated})
	return messages, err
}
