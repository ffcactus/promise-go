package base

// AsychServiceTemplateInterface is the interface that a concrete AsychService should have.
type AsychServiceTemplateInterface interface {
	CreateContext(request AsychActionRequestInterface) ContextInterface
	CreateStrategy(request AsychActionRequestInterface) StrategyInterface
}

// AsychActionServiceInterface is the interface that AsychService have.
type AsychActionServiceInterface interface {
	ServiceInterface
	PerformAsych(id string, request AsychActionRequestInterface) (ResponseInterface, *string, []Message)
}

// AsychActionService is the service for asychronous action.
type AsychActionService struct {
	TemplateImpl AsychServiceTemplateInterface
}

// PerformAsych will perform the asychronous action.
func (s *AsychActionService) PerformAsych(id string, request AsychActionRequestInterface) (ResponseInterface, *string, []Message) {
	context := s.TemplateImpl.CreateContext(request)
	strategy := s.TemplateImpl.CreateStrategy(request)
	response, taskURI, messages := strategy.Execute(context)
	if messages != nil {
		return nil, nil, messages
	}
	return response, taskURI, nil
}
