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
	return s.Prepare(ctx, id, request)
}

// Prepare do the things that is requied before returning the response to client.
// It includes:
// 1. Complement the request DTO.
// 2. Prepare the context.
// 3. Lock the enclosure.
// 4. Create the task.
func (s *Refresh) Prepare(ctx *beegoCtx.Context, id string, request base.AsychActionRequestInterface) (base.ResponseInterface, string, []base.ErrorResponse) {
	var (
		response dto.GetEnclosureResponse
		needRestoreState = true
	)

	// 1. Complement the request DTO.
	req, _ := request.(*dto.RefreshEnclosureRequest)
	if len(req.Targets) == 0 || base.ContainsString(req.Targets, model.RefreshAll) {
		req.Targets = model.RefreshBuildinAll
	}
	// 2. Prepare the context.
	refreshCtx := context.NewRefreshContext(ctx, id, req)
	refreshCtx.DB = enclosureDB
	// 3. Lock the enclosure.
	modelInterface, preState, preReason, err := enclosureDB.GetAndLock(id)
	defer func() {
		if needRestoreState {
			if _, err := refreshCtx.DB.SetState(refreshCtx.ID, preState, preReason); err != nil {
				log.WithFields(log.Fields{"id": id}).Error("Service refresh enclosure failed, prepare refresh failed, restore enclosure state on error failed.")
			} else {
				log.WithFields(log.Fields{
					"id": id, "state": preState, "reason": preReason,
				}).Warn("Service refresh enclosure failed, prepare refresh failed, restore enclosure state.")
			}
		}
	}()
	
	needRestoreState = false
	// TODO: Maybe we need define DB error to indicate this is transaction error.
	if modelInterface == nil && err != nil {
		log.WithFields(log.Fields{
			"id": id, "error": err,
		}).Warn("Service refresh enclosure failed, DB lock enclosure failed.")
		return nil, "", []base.ErrorResponse{*base.NewErrorResponseTransactionError()}
	}
	if modelInterface != nil && err == nil {
		enclosure, _ := modelInterface.(*model.Enclosure)
		response.Load(enclosure)
		if enclosure.State != model.StateLocked {
			log.WithFields(log.Fields{
				"id": id, "state": enclosure.State,
			}).Warn("Service refresh enclosure failed, enclosure can't be locked.")
			return response, "", []base.ErrorResponse{*base.NewErrorResponseBusy()}
		}
	}
	if modelInterface == nil {
		log.WithFields(log.Fields{
			"id": id,
		}).Warn("Service refresh enclosure failed, enclosure not exist.")
		return nil, "", []base.ErrorResponse{*base.NewErrorResponseNotExist()}
	}
	log.WithFields(log.Fields{
		"id": id,
	}).Info("Service refresh enclosure, lock enclosure success.")
	needRestoreState= true
	enclosure, _ := modelInterface.(*model.Enclosure)
	refreshCtx.Enclosure = enclosure
	refreshCtx.Client = enclosureClient.NewClient(enclosure)
	log.WithFields(log.Fields{
		"id": id, "client": refreshCtx.Client,
	}).Info("Service refresh enclosure, enclosure client created.")
	// 4. Create the task.
	act := strategy.NewRefresh(refreshCtx)
	createTaskRequest := act.Task()
	createTaskResponse, err := taskSDK.CreateTask(createTaskRequest)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Warn("Service refresh failed, create task failed.")
		return nil, "", []base.ErrorResponse{*base.NewErrorResponseInternalError()}
	}
	log.WithFields(log.Fields{"task": createTaskResponse.GetID()}).Info("Service refresh, create task.")
	refreshCtx.TaskID = createTaskResponse.ID
	response.Load(refreshCtx.Enclosure)
	log.WithFields(log.Fields{"ctx": refreshCtx}).Info("Service response to client before execute strategy.")
	go act.Execute(&refreshCtx.Base)
	// Let the goroutine to set the state accordingly.
	needRestoreState = false
	return response, createTaskResponse.URI, nil
}