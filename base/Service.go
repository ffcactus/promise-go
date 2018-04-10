package base

import ()

// ServiceInterface is the interface that a service should have.
type ServiceInterface interface {
	Post(RequestInterface) (ModelInterface, []MessageInterface)
	Get(id string) (ModelInterface, []MessageInterface)
	Delete(id string) []MessageInterface
}

// ServiceTemplateInterface is the interface that a concrete service should have.
type ServiceTemplateInterface interface {
	GetDB() DBInterface
	NewResponse() ResponseInterface
	GetEventService() EventServiceInterface
}

// Service is the service in Promise project.
type Service struct {
	TemplateImpl ServiceTemplateInterface
}

// Post is the default post resource implement in service.
func (s *Service) Post(request RequestInterface) (ModelInterface, []MessageInterface) {
	var (
		db       = s.TemplateImpl.GetDB()
		response = s.TemplateImpl.NewResponse()
		model    = request.ToModel()
	)
	exist, posted, commited, err := db.Post(model)
	if exist {
		return nil, []MessageInterface{NewMessageResourceDuplicate()}
	}
	if err != nil || !commited {
		return nil, []MessageInterface{NewMessageTransactionError()}
	}
	response.Load(posted)
	s.TemplateImpl.GetEventService().DispatchCreateEvent(response)
	return posted, nil
}

// Get is the default get resource implement in service.
func (s *Service) Get(id string) (ModelInterface, []MessageInterface) {
	var (
		db = s.TemplateImpl.GetDB()
	)
	model := db.Get(id)
	if model == nil {
		return nil, []MessageInterface{NewMessageNotExist()}
	}
	return model, nil
}

// Delete is the default delete resource implement in service.
func (s *Service) Delete(id string) []MessageInterface {
	var (
		db       = s.TemplateImpl.GetDB()
		response = s.TemplateImpl.NewResponse()
	)

	exist, model, commited, err := db.Delete(id)
	if !exist {
		return []MessageInterface{NewMessageNotExist()}
	}
	if err != nil || !commited {
		return []MessageInterface{NewMessageTransactionError()}
	}
	response.Load(model)
	s.TemplateImpl.GetEventService().DispatchDeleteEvent(response)
	return nil
}
