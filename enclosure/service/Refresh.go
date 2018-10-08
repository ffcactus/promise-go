package service

import (
	beegoCtx "github.com/astaxie/beego/context"
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/enclosure/action"
	"promise/enclosure/context"
	"promise/enclosure/object/dto"
	"promise/enclosure/object/model"
)

// Refresh is the service for refresh enclosure action.
type Refresh struct {
	// TODO why do we need this struct?
}

// StartBackgroundRefresh do the auto-refresh job.
// func (s *Refresh) StartBackgroundRefresh() {
// 	go s.FindServerStateAdded()
// }

// FindServerStateAdded will find the server with state added.
// func (s *Refresh) FindServerStateAdded() {
// 	for {
// 		seconds := 5
// 		if id := enclosureDB.FindServerStateAdded(); id != "" {
// 			_, _, errorResp := s.PerformAsych(id, nil)
// 			if errorResp != nil {
// 				if errorResp[0].ID == base.ErrorResponseBusy {
// 					seconds = 1
// 				} else {
// 					log.WithFields(log.Fields{
// 						"server":    id,
// 						"errorResp": errorResp[0].ID,
// 					}).Info("Service auto-refresh server failed.")
// 					seconds = 5
// 				}
// 			} else {
// 				seconds = 0
// 			}
// 		} else {
// 			seconds = 5
// 		}
// 		time.Sleep(time.Duration(seconds) * time.Second)
// 	}
// }

// PerformAsych will process the refresh action.
func (s *Refresh) PerformAsych(ctx *beegoCtx.Context, id string, request base.AsychActionRequestInterface) (base.ResponseInterface, string, []base.ErrorResponse) {
	// We seperate a whole refresh process into different stage to provide different
	// entrypoint for different refresh cases. We need make a refresh in the following cases:
	// 1. Required by user directrly.
	// 2. Auto refresh after discoverd.
	// 3. Hardware event that requires a refreshing.
	s.Stage1(ctx, id, request)
	return nil, "", nil
}

// Stage1 will do the things that is requied till locking enclosure.
// It includes:
// 1. Complement the request DTO.
// 2. Prepare the context.
// 3. Lock the enclosure.
func (s *Refresh) Stage1(ctx *beegoCtx.Context, id string, request base.AsychActionRequestInterface) {
	// 1. Complement the request DTO.
	req, _ := request.(*dto.RefreshEnclosureRequest)
	if len(req.Targets) == 0 || base.ContainsString(req.Targets, model.RefreshAll) {
		req.Targets = model.RefreshBuildinAll
	}
	// 2. Prepare the context.
	refreshCtx := context.NewRefreshContext(ctx, id, req)
	// 3. Lock the enclosure.
	modelInterface, err := enclosureDB.GetAndLock(id)
	if err != nil {
		log.WithFields(log.Fields{
			"id": id, "error": err,
		}).Warn("Service refresh enclosure failed, lock enclosure failed.")
		refreshCtx.SendResponse(nil, "", []base.ErrorResponse{*base.NewErrorResponseInternalError()})
		return
	}
	if modelInterface == nil {
		log.WithFields(log.Fields{
			"id": id,
		}).Warn("Service refresh enclosure failed, enclosure not exist.")
		refreshCtx.SendResponse(nil, "", []base.ErrorResponse{*base.NewErrorResponseNotExist()})
		return
	}
	enclosure, _ := modelInterface.(*model.Enclosure)
	if enclosure.State != model.StateLocked {
		log.WithFields(log.Fields{
			"id": id, "state": enclosure.State,
		}).Warn("Service refresh enclosure failed, lock enclosure failed.")

		refreshCtx.SendResponse(nil, "", []base.ErrorResponse{*base.NewErrorResponseErrorState()})
		return
	}
	s.Stage2(refreshCtx)
}

// Stage2 will create the action and create a task to track the process.
func (s *Refresh) Stage2(ctx *context.RefreshContext) {
	act := action.NewRefreshAction(ctx)
	act.Execute(&ctx.Base)
}

// Stage3 do the rest work of refresh.
func (s *Refresh) Stage3(ctx *context.Base) {

}
