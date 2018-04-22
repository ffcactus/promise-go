package base

// CRUDServiceTemplateInterface is the interface that a concrete CURD Service should have.
type CRUDServiceTemplateInterface interface {
	DB() DBInterface
	Response() GetResponseInterface
	Category() string
	EventService() EventServiceInterface
}

// CRUDServiceInterface is the interface that a CRUD Service have.
type CRUDServiceInterface interface {
	ServiceInterface
	Create(request PostRequestInterface) (ModelInterface, []Message)
	Update(id string, request UpdateRequestInterface) (GetResponseInterface, []Message)
	Get(id string) (ModelInterface, []Message)
	Delete(id string) []Message
	GetCollection(start int64, count int64, filter string) (*CollectionModel, []Message)
	DeleteCollection() []Message
}

// CRUDService is the service for CRUD operations.
type CRUDService struct {
	TemplateImpl CRUDServiceTemplateInterface
}

// Create is the default process to post resource in DB.
func (s *CRUDService) Create(request PostRequestInterface) (ModelInterface, []Message) {
	var (
		db       = s.TemplateImpl.DB()
		response = s.TemplateImpl.Response()
		model    = request.ToModel()
	)
	exist, posted, commited, err := db.Create(model)
	if exist {
		return nil, []Message{NewMessageDuplicate()}
	}
	if err != nil || !commited {
		return nil, []Message{NewMessageTransactionError()}
	}
	response.Load(posted)
	s.TemplateImpl.EventService().DispatchCreateEvent(response)
	return posted, nil
}

// Get is the default process to get resource in DB.
func (s *CRUDService) Get(id string) (ModelInterface, []Message) {
	var (
		db = s.TemplateImpl.DB()
	)
	model := db.Get(id)
	if model == nil {
		return nil, []Message{NewMessageNotExist()}
	}
	return model, nil
}

// Delete is the default process to delete resource in DB.
func (s *CRUDService) Delete(id string) []Message {
	var (
		db       = s.TemplateImpl.DB()
		response = s.TemplateImpl.Response()
	)

	exist, model, commited, err := db.Delete(id)
	if !exist {
		return []Message{NewMessageNotExist()}
	}
	if err != nil || !commited {
		return []Message{NewMessageTransactionError()}
	}
	response.Load(model)
	s.TemplateImpl.EventService().DispatchDeleteEvent(response)
	return nil
}

// GetCollection is the default process to get collection in DB.
func (s *CRUDService) GetCollection(start int64, count int64, filter string) (*CollectionModel, []Message) {
	var (
		db = s.TemplateImpl.DB()
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

// DeleteCollection is the default process to  delete collection in DB.
func (s *CRUDService) DeleteCollection() []Message {
	var (
		db = s.TemplateImpl.DB()
	)
	records, commited, err := db.DeleteCollection()
	if err != nil || !commited {
		return []Message{NewMessageTransactionError()}
	}
	for _, v := range records {
		response := s.TemplateImpl.Response()
		response.Load(v)
		s.TemplateImpl.EventService().DispatchDeleteEvent(response)
	}
	s.TemplateImpl.EventService().DispatchDeleteCollectionEvent(s.TemplateImpl.Category())
	return nil
}

// Update is the default process to update resource in DB.
func (s *CRUDService) Update(id string, request UpdateRequestInterface) (GetResponseInterface, []Message) {
	var (
		db       = s.TemplateImpl.DB()
		response = s.TemplateImpl.Response()
	)

	updateAction, ok := request.(UpdateRequestInterface)
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
	s.TemplateImpl.EventService().DispatchUpdateEvent(response)
	return response, nil
}
