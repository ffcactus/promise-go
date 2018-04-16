package base

// ActionServiceTemplateInterface is the interface that a concrete ActionService should have.
type ActionServiceTemplateInterface interface {
	Perform(id string, request ActionRequestInterface) (interface{}, []Message)
}

// ActionServiceInterface is the interface that ActionService have.
type ActionServiceInterface interface {
	Perform(id string, request ActionRequestInterface) (interface{}, []Message)
}

// ActionService is the implement of an action service.
type ActionService struct {
	TemplateImpl ActionServiceTemplateInterface
}

// Perform perform the action on the resource.
func (s *ActionService) Perform(id string, request ActionRequestInterface) (interface{}, []Message) {
	return s.TemplateImpl.Perform(id, request)
}