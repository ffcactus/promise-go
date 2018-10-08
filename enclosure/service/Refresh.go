package service

import (
	beegoCtx "github.com/astaxie/beego/context"
	log "github.com/sirupsen/logrus"
	"promise/base"
	enclosureClient "promise/enclosure/client/enclosure"
	"promise/enclosure/object/dto"
	"promise/enclosure/context"
)

// Refresh is the service for refresh enclosure action.
type Refresh struct {
	// TODO why do we need this struct?
}

// StartBackgroundRefresh do the auto-refresh job.
func (s *Refresh) StartBackgroundRefresh() {
	go s.FindServerStateAdded()
}

// FindServerStateAdded will find the server with state added.
func (s *Refresh) FindServerStateAdded() {
	for {
		seconds := 5
		if id := enclosureDB.FindServerStateAdded(); id != "" {
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
func (s *Refresh) PerformAsych(ctx *context.Context, id string, request base.AsychActionRequestInterface) (base.ResponseInterface, string, []base.ErrorResponse) {
	// We seperate a whole refresh process into different stage to provide different
	// entrypoint for different refresh cases. We need make a refresh in the following cases:
	// 1. Required by user directrly.
	// 2. Auto refresh after discoverd.
	// 3. Hardware event that requires a refreshing.

	return s.Stage1(ctx, id, request)
}

// Stage1 will do the things that is requied till locking enclosure.
// It includes:
// 1. Complement the request DTO.
// 2. Prepare the context.
// 3. Lock the enclosure.
func (s *Refresh) Stage1(ctx *beegoCtx.Context, id string, request base.AsychActionRequestInterface) (base.ResponseInterface, string, []base.ErrorResponse) {
	// 1. Complement the request DTO.
	if len(request.Targets) || Contains(request.Targets, model.RefreshAll) {
		request.Targets = model.RefreshBuildinAll
	}
	// 2. Prepare the context.
	ctx := context.CreateRefreshEnclosureContext(ctx, id)
	modelInterface, errorResp := enclosureDB.Get(id)
	// TODO db should not return error response.
	if errorResp != nil {
		return nil, "", []base.ErrorResponse{*errorResp}
	}
	enclosure, ok := modelInterface.(*model.Server)
	if !ok {
		return nil, "", []base.ErrorResponse{*base.NewErrorResponseInternalError()}
	}
	refreshCtx := context.NewRefreshContext(enclosure)
	// 3. Lock the enclosure.
	enclosure, err := enclosureDB.GetAndLock(id)
	if err != nil {
		log.WithFields(log.Fields{
			"id": ID, "error": err,
		}).Warn("Service refresh enclosure failed, lock enclosure failed.")
		return nil, "", []base.ErrorResponse{*base.NewErrorResponseInternalError()}
	}
	if enclosure == nil {
		log.WithFields(log.Fields{
			"id": ID,
		}).Warn("Service refresh enclosure failed, enclosure not exist.")
		return nil, "", []base.ErrorResponse{*base.NewErrorResponseNotExist()}	
	}
	if enclosure.State != model.StateLocked {
		log.WithFields(log.Fields{
			"id": ID, "state": enclosure.State
		}).Warn("Service refresh enclosure failed, lock enclosure failed.")
		return nil, "", []base.ErrorResponse{*base.NewErrorResponseErrorState()}	
	}
	return s.Stage2(refreshCtx)
}

// Stage2 will create a task to track the process.
func (s *Refesh) Stage2(ctx *context.Base) (base.ResponseInterface, string, []base.ErrorResponse) {
	
}

func (s *Refresh) Stage3(ctx *context.Base) {

}