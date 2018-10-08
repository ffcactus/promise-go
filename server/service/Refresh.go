package service

import (
	beegoCtx "github.com/astaxie/beego/context"
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
			_, _, errorResp := s.PerformAsych(id, nil)
			if errorResp != nil {
				if errorResp[0].ID == base.ErrorResponseBusy {
					seconds = 1
				} else {
					log.WithFields(log.Fields{
						"server":    id,
						"errorResp": errorResp[0].ID,
					}).Info("Service auto-refresh server failed.")
					seconds = 5
				}
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
func (s *Refresh) PerformAsych(ctx *beegoCtx.Context, id string, request base.AsychActionRequestInterface) (base.ResponseInterface, string, []base.ErrorResponse) {
	modelInterface, errorResp := serverDB.Get(id)
	if errorResp != nil {
		return nil, nil, []base.ErrorResponse{*errorResp}
	}
	server, ok := modelInterface.(*model.Server)
	if !ok {
		return nil, nil, []base.ErrorResponse{*base.NewErrorResponseInternalError()}
	}
	ctx := context.CreateRefreshServerContext(server)
	st := strategy.CreateRefreshServerStrategy(server)
	task, errorResps := st.Execute(ctx, server)
	if errorResps != nil {
		return nil, nil, errorResps
	}
	return nil, task, nil
}
