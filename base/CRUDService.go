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
	Create(request PostRequestInterface) (ModelInterface, []ErrorResponse)
	Update(id string, request UpdateRequestInterface) (GetResponseInterface, []ErrorResponse)
	Get(id string) (ModelInterface, []ErrorResponse)
	Delete(id string) []ErrorResponse
	GetCollection(start int64, count int64, filter string) (*CollectionModel, []ErrorResponse)
	DeleteCollection() []ErrorResponse
}

// CRUDService is the service for CRUD operations.
type CRUDService struct {
	TemplateImpl CRUDServiceTemplateInterface
}

// Create is the default process to post resource in DB.
func (s *CRUDService) Create(request PostRequestInterface) (ModelInterface, []ErrorResponse) {
	var (
		db       = s.TemplateImpl.DB()
		response = s.TemplateImpl.Response()
		model    = request.ToModel()
	)
	posted, errorResp := db.Create(model)
	if errorResp != nil {
		return nil, []ErrorResponse{*errorResp}
	}
	response.Load(posted)
	PublishCreateMessage(response)
	return posted, nil
}

// Get is the default process to get resource in DB.
func (s *CRUDService) Get(id string) (ModelInterface, []ErrorResponse) {
	var (
		db = s.TemplateImpl.DB()
	)
	model, errorResp := db.Get(id)
	if errorResp != nil {
		return nil, []ErrorResponse{*errorResp}
	}
	return model, nil
}

// Delete is the default process to delete resource in DB.
func (s *CRUDService) Delete(id string) []ErrorResponse {
	var (
		db       = s.TemplateImpl.DB()
		response = s.TemplateImpl.Response()
	)

	model, errorResp := db.Delete(id)
	if errorResp != nil {
		return []ErrorResponse{*errorResp}
	}
	response.Load(model)
	PublishDeleteMessage(response)
	return nil
}

// GetCollection is the default process to get collection in DB.
func (s *CRUDService) GetCollection(start int64, count int64, filter string) (*CollectionModel, []ErrorResponse) {
	var (
		db = s.TemplateImpl.DB()
	)
	collection, errorResp := db.GetCollection(start, count, filter)
	if errorResp != nil {
		return nil, []ErrorResponse{*errorResp}
	}
	return collection, nil
}

// DeleteCollection is the default process to  delete collection in DB.
func (s *CRUDService) DeleteCollection() []ErrorResponse {
	var (
		db = s.TemplateImpl.DB()
	)
	records, errorResp := db.DeleteCollection()
	if errorResp != nil {
		return []ErrorResponse{*errorResp}
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
func (s *CRUDService) Update(id string, request UpdateRequestInterface) (GetResponseInterface, []ErrorResponse) {
	var (
		db       = s.TemplateImpl.DB()
		response = s.TemplateImpl.Response()
	)

	updateAction, ok := request.(UpdateRequestInterface)
	if !ok {
		return nil, []ErrorResponse{*NewErrorResponseInternalError()}
	}

	updatedTask, errorResp := db.Update(id, updateAction)
	if errorResp != nil {
		return nil, []ErrorResponse{*errorResp}
	}
	response.Load(updatedTask)
	PublishUpdateMessage(response)
	return response, nil
}
