package base

import ()

// ActionServiceInterface is the interface that an action service should have.
type ActionServiceInterface interface {
	Perform(id string, request ActionRequestInterface) (interface{}, *string, []Message)
}

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

// Perform the update task action.
func (s *Service) Perform(id string, request ActionRequestInterface) (interface{}, *string, []Message) {
	var (
		db       = s.TemplateImpl.GetDB()
		response = s.TemplateImpl.NewResponse()
	)

	updateAction, ok := request.(UpdateActionRequestInterface)
	if !ok {
		return nil, []Message{NewMessageInternalError()}
	}

	exist, updatedTask, commited, err := db.Update(id, updateAction)
	if !exist {
		return nil, []Message{NewMessageNotExist()}
	}
	if err != nil && err.Error() == ErrorUnknownPropertyValue.Error() {
		return nil, []Message{NewMessageUnknownPropertyValue()}
	}
	if err != nil || !commited {
		return nil, []Message{NewMessageTransactionError()}
	}
	response.Load(updatedTask)
	s.TemplateImpl.GetEventService().DispatchUpdateEvent(response)
	return response, nil
}

func (s *Service) PerformSych(id string, request ActionReqeustInterface) (interface{}, []Message) {

}

// AsychServiceTemplateInterface is the interface that a concrete AsychService should have.
type AsychServiceTemplateInterface interface {
	CreateContext(AsychActionInterface) ContextInterface
	CreateStrategy(AsychActionInterface) StrategyInterface
}

// AsychServiceInterface is the interface that AsychService have.
type AsychServiceInterface interface {
	PerformAsych(id string, request ActionRequestInterface) (interface{}, *string, []Message)
}

// AsychService is the template of the service that is going to handle asychronous action.
type AsychService struct {
	TemplateImpl AsychServiceTemplateInterface
}

// PerformAsych will perform the asychronous action.
func (s *AsychService) PerformAsych(id string, request ActionRequestInterface) (interface{}, *string, []Message) {
	asychAction, ok := request.(ActionRequestInterface)
	if !ok {
		return nil, []Message{NewMessageInternalError()}
	}

	context := s.TemplateImpl.CreateContext(asychAction)
	strategy := s.TemplateImpl.CreateStrategy(asychAction)
	response, taskURI, messages := strategy.execute(context)
	if messages != nil {
		return nil, nil, messages
	}
	return response, taskURI, nil
}