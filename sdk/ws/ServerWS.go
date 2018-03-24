package ws

import (
	"net/http"
	"time"
	"encoding/json"
	"promise/server/object/model"
	"promise/common/app"
	commonConstError "promise/common/object/constError"
	commonConstValue "promise/common/object/constValue"
	commonDTO "promise/common/object/dto"
	"promise/common/app/rest"
	"promise/ws/object/constValue"
	serverDTO "promise/server/object/dto"
	wsDTO "promise/ws/object/dto"
	log "github.com/sirupsen/logrus"
)

var (
	// WsServerRoot The root of the service.
	WsServerRoot = app.ProtocolScheme + app.Host + app.RootURL + commonConstValue.WSSenderBaseURI
)

func dispatchServerCreateOrUpdate(server *model.Server, eventType string) ([]commonDTO.Message, error) {
	var (
		s serverDTO.GetServerResponse
		event wsDTO.PostEventRequest
	)
	event.CreatedAt = time.Now()
	event.Category = commonConstValue.CategoryServer
	event.Type = eventType // constValue.CreateEvent
	event.ResourceID = server.ID
	s.Load(server)
	b, err := json.Marshal(s); 
	if err != nil {
		log.WithFields(log.Fields{
			"category":event.Category, 
			"type":event.Type, 
			"resource": event.ResourceID, 
			"error":err}).Warn("Dispatch server event failed, failed to unmarshal resource.")
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
