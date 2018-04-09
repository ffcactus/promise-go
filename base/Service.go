package base

import (
// wsSDK "promise/sdk/event"
)

// ServiceInterface is the interface that a service should have.
type ServiceInterface interface {
	// Post(RequestInterface) (ModelInterface, []MessageInterface)
	GetDB() DBInterface
	NewResponse() ResponseInterface
	GetEventService() EventServiceInterface
}

// Service is the service in Promise project.
type Service struct {
	Interface ServiceInterface
}

// Post is the default method to do post in service.
func (s *Service) Post(request RequestInterface) (ModelInterface, []MessageInterface) {
	var (
		db       = s.Interface.GetDB()
		response = s.Interface.NewResponse()
	)

	exist, posted, commited, err := db.Post(request.ToModel())
	if exist {
		return nil, []MessageInterface{NewMessageResourceDuplicate()}
	}
	if err != nil || !commited {
		return nil, []MessageInterface{NewMessageTransactionError()}
	}
	response.Load(posted)
	s.Interface.GetEventService().DispatchCreateEvent(response)
	return posted, nil
}
