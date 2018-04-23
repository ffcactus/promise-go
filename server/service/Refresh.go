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

// FindServerStateAdded will find the server with state added.
func (s *Refresh) FindServerStateAdded() {
	for {
		seconds := 5
		if id := serverDB.FindServerStateAdded(); id != "" {
			s.Perform(id, nil)
			seconds = 0
		} else {
			seconds = 5
		}
		time.Sleep(time.Duration(seconds) * time.Second)
	}
}

// Perform will process the refresh action.
func (s *Refresh) Perform(id string, request base.ActionRequestInterface) (base.ResponseInterface, []base.Message) {
	modelInterface := serverDB.Get(id)
	if modelInterface == nil {
		return nil, []base.Message{base.NewMessageNotExist()}
	}
	server, ok := modelInterface.(*model.Server)
	if !ok {
		return nil, []base.Message{base.NewMessageInternalError()}
	}
	ctx := context.CreateRefreshServerContext(server)
	st := strategy.CreateRefreshServerStrategy(server)
	if err := st.Execute(ctx, server); err != nil {
		return nil, ctx.Messages()
	}
	return nil, nil
}
