package base

import (
	beegoCtx "github.com/astaxie/beego/context"
)

// AsychServiceTemplateInterface is the interface that a concrete AsychService should have.
type AsychServiceTemplateInterface interface {
	CreateContext(request AsychActionRequestInterface) ContextInterface
	CreateStrategy(request AsychActionRequestInterface) StrategyInterface
}

// AsychActionServiceInterface is the interface that AsychService have.
type AsychActionServiceInterface interface {
	ServiceInterface
	PerformAsych(ctx *beegoCtx.Context, id string, request AsychActionRequestInterface) (ResponseInterface, string, []ErrorResponse)
}

// AsychActionService is the service for asychronous action.
type AsychActionService struct {
	TemplateImpl AsychServiceTemplateInterface
}

// PerformAsych will perform the asychronous action.
func (s *AsychActionService) PerformAsych(ctx *beegoCtx.Context, id string, request AsychActionRequestInterface) (ResponseInterface, string, []ErrorResponse) {
	context := s.TemplateImpl.CreateContext(request)
	strategy := s.TemplateImpl.CreateStrategy(request)
	response, taskURI, errorResps := strategy.Execute(context)
	if errorResps != nil {
		return nil, "", errorResps
	}
	return response, taskURI, nil
}
