package ws

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
	"promise/base"
	"promise/common/app"
	"promise/common/app/rest"
	commonConstError "promise/common/object/consterror"
	commonConstValue "promise/common/object/constvalue"
	commonDTO "promise/common/object/dto"
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
		m := ""
		if messages != nil && len(messages) > 1 {
			m = messages[0].ID
		}
		log.WithFields(log.Fields{
			"category": event.Category,
			"type":     event.Type,
			"resource": event.ResourceID,
			"message":  m,
			"error":    err}).Warn("Dispatch event failed.")
	}
	return messages, err
}

// DispatchResourceEvent can dispatch event, you have specify the event type and the DTO of GET resource response.
func DispatchResourceEvent(eventType string, dto commonDTO.PromiseResponseInterface) ([]commonDTO.Message, error) {
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

// DispatchResourceCreateEvent can dispatch resource create event.
func DispatchResourceCreateEvent(dto commonDTO.PromiseResponseInterface) ([]commonDTO.Message, error) {
	return DispatchResourceEvent(base.CreateEvent, dto)
}

// DispatchResourceUpdateEvent can dispatch resource udpate event.
func DispatchResourceUpdateEvent(dto commonDTO.PromiseResponseInterface) ([]commonDTO.Message, error) {
	return DispatchResourceEvent(base.UpdateEvent, dto)
}

// DispatchResourceDeleteEvent can dispatch resource delete event.
func DispatchResourceDeleteEvent(dto commonDTO.PromiseResponseInterface) ([]commonDTO.Message, error) {
	return DispatchResourceEvent(base.DeleteEvent, dto)
}

// DispatchResourceCollectionDeleteEvent dispatch an event about resource collection deleted.
func DispatchResourceCollectionDeleteEvent(category string) ([]commonDTO.Message, error) {
	var (
		event wsDTO.PostEventRequest
	)
	event.CreatedAt = time.Now()
	event.Category = category
	event.Type = base.DeleteCollectionEvent
	return dispatch(&event)
}
