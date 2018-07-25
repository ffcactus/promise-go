package base

// CRUDServiceTemplateInterface is the interface that a concrete CURD Service should have.
type CRUDServiceTemplateInterface interface {
	DB() DBInterface
	Response() GetResponseInterface
	Category() string
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
	posted, message := db.Create(model)
	if message != nil {
		return nil, []Message{*message}
	}
	response.Load(posted)
	PublishCreateMessage(response)
	return posted, nil
}

// Get is the default process to get resource in DB.
func (s *CRUDService) Get(id string) (ModelInterface, []Message) {
	var (
		db = s.TemplateImpl.DB()
	)
	model, message := db.Get(id)
	if message != nil {
		return nil, []Message{*message}
	}
	return model, nil
}

// Delete is the default process to delete resource in DB.
func (s *CRUDService) Delete(id string) []Message {
	var (
		db       = s.TemplateImpl.DB()
		response = s.TemplateImpl.Response()
	)

	model, message := db.Delete(id)
	if message != nil {
		return []Message{*message}
	}
	response.Load(model)
	PublishDeleteMessage(response)
	return nil
}

// GetCollection is the default process to get collection in DB.
func (s *CRUDService) GetCollection(start int64, count int64, filter string) (*CollectionModel, []Message) {
	var (
		db = s.TemplateImpl.DB()
	)
	collection, message := db.GetCollection(start, count, filter)
	if message != nil {
		return nil, []Message{*message}
	}
	return collection, nil
}

// DeleteCollection is the default process to  delete collection in DB.
func (s *CRUDService) DeleteCollection() []Message {
	var (
		db = s.TemplateImpl.DB()
	)
	records, message := db.DeleteCollection()
	if message != nil {
		return []Message{*message}
	}
	for _, v := range records {
		response := s.TemplateImpl.Response()
		response.Load(v)
		PublishDeleteMessage(response)
	}
	PublishDeleteCollectionMessage(s.TemplateImpl.Category())
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
		return nil, []Message{*NewMessageInternalError()}
	}

	updatedTask, message := db.Update(id, updateAction)
	if message != nil {
		return nil, []Message{*message}
	}
	response.Load(updatedTask)
	PublishUpdateMessage(response)
	return response, nil
}
