package service

import (
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
			s.PerformAsych(id, nil)
			seconds = 0
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
	task, err := st.Execute(ctx, server)
	if err != nil {
		return nil, nil, ctx.Messages()
	}
	return nil, task, nil
}
