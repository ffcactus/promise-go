package base

import (
	"promise/base"
	wsSDK "promise/sdk/ws"
)

// ServiceInterface is the interface that a service should have.
type ServiceInterface interface {
	Post(RequestInterface) (ModelInterface, error)
	GetDB() base.DBInterface
	NewResponse() base.ResposneInterface
}

// Service is the service in Promise project.
type Service struct {
	Interface ServiceInterface
}

// Post is the default method to do post in service.
func (s *Service) Post(request RequestInterface) (ModelInterface, error) {
	var (
		db = s.Interface.GetDB()
		responseDTO = s.Interface.NewResponse()
	)

	exist, posted, commited, err := db.Post(request.ToModel())
	if exist {
		return nil, []base.Message{base.NewMessageResourceDuplicate()}
	}
	if err != nil || !commited {
		return nil, []base.Message{base.NewMessageTransactionError()}
	}
	responseDTO.Load(posted)
	wsSDK.DispatchResourceCreateEvent(responseDTO)
	return posted, nil
}