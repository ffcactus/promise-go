package controller

import (
	"github.com/astaxie/beego"
	log "github.com/sirupsen/logrus"
	"net/http"
	commonDto "promise/common/object/dto"
	dto "promise/server/object/dto"
	"promise/server/service"
)

// ServerIDController Server ID controller.
type ServerIDController struct {
}

// ResourceName returns the name this controller handle of.
func (c *IPv4IDController) ResourceName() string {
	return "server"
}

// Response creates a new response DTO.
func (c *IPv4IDController) Response() base.GetResponseInterface {
	return new(dto.GetServerResponse)
}

// Service returns the service.
func (c *IPv4IDController) Service() base.CRUDServiceInterface {
	return serverService
}
