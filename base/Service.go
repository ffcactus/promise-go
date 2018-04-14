package base

import ()

// ServiceInterface is the interface that a service should have.
type ServiceInterface interface {
	Post(RequestInterface) (ModelInterface, []Message)
	Get(id string) (ModelInterface, []Message)
	Delete(id string) []Message
	GetCollection(start int64, count int64, filter string) (*CollectionModel, []Message)
	DeleteCollection() []Message
}

// ServiceTemplateInterface is the interface that a concrete service should have.
type ServiceTemplateInterface interface {
	GetDB() DBInterface
	NewResponse() ResponseInterface
	GetCategory() string
	GetEventService() EventServiceInterface
}

// Service is the service in Promise project.
type Service struct {
	TemplateImpl ServiceTemplateInterface
}

// Post is the default post resource implement in service.
func (s *Service) Post(request RequestInterface) (ModelInterface, []Message) {
	var (
		db       = s.TemplateImpl.GetDB()
		response = s.TemplateImpl.NewResponse()
		model    = request.ToModel()
	)
	exist, posted, commited, err := db.Post(model)
	if exist {
		return nil, []Message{NewMessageDuplicate()}
	}
	if err != nil || !commited {
		return nil, []Message{NewMessageTransactionError()}
	}
	response.Load(posted)
	s.TemplateImpl.GetEventService().DispatchCreateEvent(response)
	return posted, nil
}

// Get is the default get resource implement in service.
func (s *Service) Get(id string) (ModelInterface, []Message) {
	var (
		db = s.TemplateImpl.GetDB()
	)
	model := db.Get(id)
	if model == nil {
		return nil, []Message{NewMessageNotExist()}
	}
	return model, nil
}

// Delete is the default delete resource implement in service.
func (s *Service) Delete(id string) []Message {
	var (
		db       = s.TemplateImpl.GetDB()
		response = s.TemplateImpl.NewResponse()
	)

	exist, model, commited, err := db.Delete(id)
	if !exist {
		return []Message{NewMessageNotExist()}
	}
	if err != nil || !commited {
		return []Message{NewMessageTransactionError()}
	}
	response.Load(model)
	s.TemplateImpl.GetEventService().DispatchDeleteEvent(response)
	return nil
}

// GetCollection is the default get collection implement in service.
func (s *Service) GetCollection(start int64, count int64, filter string) (*CollectionModel, []Message) {
	var (
		db = s.TemplateImpl.GetDB()
	)
	collection, err := db.GetCollection(start, count, filter)
	if err != nil && err.Error() == ErrorUnknownFilterName.Error() {
		return nil, []Message{NewMessageUnknownFilterName()}
	}
	if err != nil {
		return nil, []Message{NewMessageTransactionError()}
	}
	return collection, nil
}

// DeleteCollection is the default delete collection implement in service.
func (s *Service) DeleteCollection() []Message {
	var (
		db = s.TemplateImpl.GetDB()
	)
	records, commited, err := db.DeleteCollection()
	if err != nil || !commited {
		return []Message{NewMessageTransactionError()}
	}
	for _, v := range records {
		response := s.TemplateImpl.NewResponse()
		response.Load(v)
		s.TemplateImpl.GetEventService().DispatchDeleteEvent(response)
	}
	s.TemplateImpl.GetEventService().DispatchDeleteCollectionEvent(s.TemplateImpl.GetCategory())
	return nil
}
