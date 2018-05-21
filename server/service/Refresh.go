package service

import (
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/server/context"
	"promise/server/object/model"
	"promise/server/strategy"
	"time"
)

// Refresh is the service for refresh server action.
type Refresh struct {
}

// StartBackgroundRefresh do the auto-refresh job.
func (s *Refresh) StartBackgroundRefresh() {
	go s.FindServerStateAdded()
}

// FindServerStateAdded will find the server with state added.
func (s *Refresh) FindServerStateAdded() {
	for {
		seconds := 5
		if id := serverDB.FindServerStateAdded(); id != "" {
			_, _, message := s.PerformAsych(id, nil)
			if message != nil {
				log.WithFields(log.Fields{
					"server":  id,
					"message": message[0].ID,
				}).Info("Service auto-refresh server failed.")
				seconds = 10
			} else {
				seconds = 0
			}
		} else {
			seconds = 5
		}
		time.Sleep(time.Duration(seconds) * time.Second)
	}
}

// PerformAsych will process the refresh action.
func (s *Refresh) PerformAsych(id string, request base.AsychActionRequestInterface) (base.ResponseInterface, *string, []base.Message) {
	modelInterface, message := serverDB.Get(id)
	if message != nil {
		return nil, nil, []base.Message{*message}
	}
	server, ok := modelInterface.(*model.Server)
	if !ok {
		return nil, nil, []base.Message{*base.NewMessageInternalError()}
	}
	ctx := context.CreateRefreshServerContext(server)
	st := strategy.CreateRefreshServerStrategy(server)
	task, messages := st.Execute(ctx, server)
	if messages != nil {
		return nil, nil, messages
	}
	return nil, task, nil
}
