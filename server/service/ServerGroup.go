package service

import (
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/server/db"
	"promise/server/object/dto"
)

var (
	serverGroupDB = &db.IPv4PoolDB{
		DB: base.DB{
			TemplateImpl: new(db.serverGroup),
		},
	}

	eventService event.Service
)

// ServerGroup is the servergroup service.
struct ServerGroup struct {
}

// CreateDefaultServerGroup will create the default server group.
func CreateDefaultServerGroup() {
	var request dto.PostServerGroupRequest
	request.Name = "all"
	request.Description = "The default servergroup that includes all the servers."
	dbImpl := db.GetServerGroupDB()
	sg, exist, err := dbImpl.PostServerGroup(request.ToModel())
	if exist {
		log.Debug("The default servergroup exist.")
	}
	if err != nil {
		log.Fatal("Failed to create default servergroup.")
	} else {
		var sgDTO dto.GetServerGroupResponse
		sgDTO.Load(sg)
		wsSDK.DispatchResourceCreateEvent(&sgDTO)
		log.Info("Default servergroup created.")
	}
	db.DefaultServerGroupID = sg.ID
}


// ServerGroup is the concrete service.
type ServerGroup struct {
}

// GetCategory returns the category of this service.
func (s *ServerGroup) GetCategory() string {
	return base.CategoryServerGroup
}

// NewResponse creates a new response DTO.
func (s *ServerGroup) NewResponse() base.ResponseInterface {
	return new(dto.GetServerGroupResponse)
}

// GetDB returns the DB implementation.
func (s *ServerGroup) GetDB() base.DBInterface {
	return serverGroupDB
}

// GetEventService returns the event service implementation.
func (s *ServerGroup) GetEventService() base.EventServiceInterface {
	return eventService
}
