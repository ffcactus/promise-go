package event

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
	"promise/base"
	"promise/common/app/rest"
	commonConstError "promise/common/object/consterror"
	"promise/ws/object/constvalue"
	wsDTO "promise/ws/object/dto"
	"time"
)

var (
	// WsSenderServiceURL is the service URL.
	WsSenderServiceURL = base.ProtocolScheme + base.Host + base.RootURL + base.WSSenderBaseURI
)

// Service is the implementation of EventServiceInterface
type Service struct {
}

func dispatch(event *wsDTO.PostEventRequest) ([]base.Message, error) {
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
	// TODO
	return nil, err
}

// DispatchResourceEvent can dispatch event, you have specify the event type and the DTO of GET resource response.
func DispatchResourceEvent(eventType string, dto base.ResponseInterface) ([]base.Message, error) {
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

// DispatchCreateEvent can dispatch resource create event.
func (s Service) DispatchCreateEvent(dto base.ResponseInterface) ([]base.Message, error) {
	return DispatchResourceEvent(constvalue.CreateEvent, dto)
}

// DispatchUpdateEvent can dispatch resource udpate event.
func (s Service) DispatchUpdateEvent(dto base.ResponseInterface) ([]base.Message, error) {
	return DispatchResourceEvent(constvalue.UpdateEvent, dto)
}

// DispatchDeleteEvent can dispatch resource delete event.
func (s Service) DispatchDeleteEvent(dto base.ResponseInterface) ([]base.Message, error) {
	return DispatchResourceEvent(constvalue.DeleteEvent, dto)
}

// DispatchDeleteCollectionEvent dispatch an event about resource collection deleted.
func (s Service) DispatchDeleteCollectionEvent(category string) ([]base.Message, error) {
	var (
		event wsDTO.PostEventRequest
	)
	event.CreatedAt = time.Now()
	event.Category = category
	event.Type = constvalue.DeleteCollectionEvent
	return dispatch(&event)
}
