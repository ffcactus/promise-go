package service

import (
	beegoCtx "github.com/astaxie/beego/context"
	log "github.com/sirupsen/logrus"
	"promise/base"
	enclosureClient "promise/enclosure/client/enclosure"
	"promise/enclosure/context"
	"promise/enclosure/object/dto"
	"promise/enclosure/object/model"
	"promise/enclosure/strategy"
	taskSDK "promise/sdk/task"
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
	refreshCtx.DB = enclosureDB
	// 3. Lock the enclosure.
	modelInterface, err := enclosureDB.GetAndLock(id)
	defer func() {
		if modelInterface, err := refreshCtx.DB.SetState(refreshCtx.ID, refreshCtx.NextState, refreshCtx.NextReason); err != nil {
			log.WithFields(log.Fields{"id": id}).Error("Service refresh enclosure failed, unlock enclosure failed.")
		} else {
			enclosure, _ := modelInterface.(*model.Enclosure)
			log.WithFields(log.Fields{
				"id": id, "state": enclosure.State, "reason": enclosure.StateReason,
			}).Info("Service refresh enclosure done, set enclosure state.")
		}
	}()
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
	refreshCtx.Enclosure = enclosure
	refreshCtx.Client = enclosureClient.NewClient(enclosure)
	log.WithFields(log.Fields{
		"id": id,
	}).Info("Service refresh enclosure, lock enclosure success.")
	s.Stage2(refreshCtx)
}

// Stage2 will create the action and create a task to track the process.
func (s *Refresh) Stage2(ctx *context.RefreshContext) {
	var (
		response dto.GetEnclosureResponse
	)
	act := strategy.NewRefresh(ctx)
	createTaskRequest := act.Task()
	createTaskResponse, err := taskSDK.CreateTask(createTaskRequest)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Warn("Service refresh failed, create task failed.")
		ctx.SendResponse(nil, "", []base.ErrorResponse{*base.NewErrorResponseInternalError()})
	}
	log.WithFields(log.Fields{"task": createTaskResponse.GetID()}).Info("Service refresh, create task.")
	ctx.TaskID = createTaskResponse.ID
	response.Load(ctx.Enclosure)
	// Send response to client.
	ctx.SendResponse(response, ctx.TaskID, nil)
	log.WithFields(log.Fields{"ctx": ctx}).Info("Service response to client before execute strategy.")
	act.Execute(&ctx.Base)
}

// Stage3 do the rest work of refresh.
func (s *Refresh) Stage3(ctx *context.Base) {

}
