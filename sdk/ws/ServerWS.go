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
	"promise/ws/object/constValue"
	wsDTO "promise/ws/object/dto"
	"time"
)

var (
	// WsSenderServiceURL is the service URL.
	WsSenderServiceURL = app.ProtocolScheme + app.Host + app.RootURL + commonConstValue.WSSenderBaseURI
)

func dispatch(event *wsDTO.PostEventRequest) ([]commonDTO.Message, error) {
	messages, err := rest.Do(
		http.MethodPost,
		WsSenderServiceURL,
		event,
		nil,
		[]int{http.StatusCreated})
	if err != nil {
		log.WithFields(log.Fields{
			"category": event.Category,
			"type":     event.Type,
			"resource": event.ResourceID,
			"message":  messages[0].ID,
			"error":    err}).Warn("Dispatch event failed.")
	}
	return messages, err
}

func dispatchResourceCreateOrUpdate(eventType string, dto commonDTO.PromiseResponseInterface) ([]commonDTO.Message, error) {
	var (
		event wsDTO.PostEventRequest
	)
	event.CreatedAt = time.Now()
	event.Category = dto.GetCategory()
	event.Type = eventType
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
	return dispatch(&event)
}

// DispatchResourceCreate dispatch an event about resource created.
func DispatchResourceCreate(dto commonDTO.PromiseResponseInterface) ([]commonDTO.Message, error) {
	return dispatchResourceCreateOrUpdate(constValue.CreateEvent, dto)
}

// DispatchResourceUpdate dispatch an event about resource updated.
func DispatchResourceUpdate(dto commonDTO.PromiseResponseInterface) ([]commonDTO.Message, error) {
	return dispatchResourceCreateOrUpdate(constValue.UpdateEvent, dto)
}

// DispatchResourceDelete dispatch an event about resource deleted.
func DispatchResourceDelete(category string, id string) ([]commonDTO.Message, error) {
	var (
		event wsDTO.PostEventRequest
	)
	event.CreatedAt = time.Now()
	event.Category = category
	event.Type = constValue.DeleteEvent
	event.ResourceID = id
	return dispatch(&event)
}

// DispatchResourceCollectionDelete dispatch an event about resource collection deleted.
func DispatchResourceCollectionDelete(category string) ([]commonDTO.Message, error) {
	var (
		event wsDTO.PostEventRequest
	)
	event.CreatedAt = time.Now()
	event.Category = category
	event.Type = constValue.DeleteCollectionEvent
	return dispatch(&event)
}
