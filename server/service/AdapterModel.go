package service

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"promise/base"
	"promise/server/db"
	"promise/server/object/constvalue"
	"promise/server/object/dto"
	"promise/server/object/model"
)

var (
	adapterModelDB = &db.AdapterModel{
		DB: base.DB{
			TemplateImpl: new(db.AdapterModel),
		},
	}
	models = make(map[string]string)
)

// AdapterModel is the servergroup service.
type AdapterModel struct {
}

// Category returns the category of this service.
func (s *AdapterModel) Category() string {
	return base.CategoryAdapterModel
}

// Response creates a new response DTO.
func (s *AdapterModel) Response() base.GetResponseInterface {
	return new(dto.GetAdapterModelResponse)
}

// DB returns the DB implementation.
func (s *AdapterModel) DB() base.DBInterface {
	return adapterModelDB
}

// EventService returns the event service implementation.
func (s *AdapterModel) EventService() base.EventServiceInterface {
	return eventService
}

// LoadModel will load the model from local files.
func LoadModel() {
	var (
		m        model.AdapterModel
		response dto.GetAdapterModelResponse
	)
	files, err := ioutil.ReadDir(constvalue.AdapterModelPath)
	if err != nil {
		log.WithFields(log.Fields{
			"path": constvalue.AdapterModelPath,
		}).Error("Service failed to read adapter model directory.")
	}
	for _, file := range files {
		content, err := ioutil.ReadFile(constvalue.AdapterModelPath + "/" + file.Name())
		if err != nil {
			log.WithFields(log.Fields{
				"name": file.Name(),
			}).Error("Service failed to read adapter model file.")
		}
		if err := json.Unmarshal(content, &m); err != nil {
			log.WithFields(log.Fields{
				"name": file.Name(),
			}).Error("Service failed to convert adapter file to model.")
		}
		// Don't forget to set the category here.
		m.Category = base.CategoryAdapterModel
		save, message := adapterModelDB.Create(&m)
		if message == nil {
			response.Load(&m)
			eventService.DispatchCreateEvent(&response)
			log.WithFields(log.Fields{
				"id": save.GetID(),
			}).Info("Service save adapter model to DB.")
		} else if message.ID == base.MessageDuplicate {
			log.WithFields(log.Fields{
				"name": file.Name(),
			}).Info("Service found the adapter model exist.")
		} else {
			log.WithFields(log.Fields{
				"message": message.ID,
				"name":    file.Name(),
			}).Info("Service failed to create adapter model to DB.")
		}
	}
}
