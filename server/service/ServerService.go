package service

import (
	commonM "promise/common/object/model"
	"promise/server/context"
	"promise/server/db"
	"promise/server/object/dto"
	"promise/server/object/message"
	"promise/server/object/model"
	"promise/server/strategy"
	"time"
)

// PostServer since the server is not known yet. we have to create the context step by step.
func PostServer(request *dto.PostServerRequest) (*model.Server, []commonM.Message) {
	dbImpl := db.GetDBInstance()

	// Try to get the basic info.
	serverBasicInfo, err := Probe(request)
	if err != nil {
		return nil, []commonM.Message{message.NewServerPostFailed()}
	}

	server := serverBasicInfo.CreateServer()
	if exist, existedServer := dbImpl.IsServerExist(server); exist {
		return nil, []commonM.Message{message.NewServerExist(existedServer)}
	}

	// Before save the server to the DB. We need configure the server first.
	ctx := context.CreatePostServerContext(server, request)
	st := strategy.CreatePostServerStrategy(server)
	if err := st.Execute(ctx); err != nil {
		return nil, ctx.Messages()
	}
	return ctx.Server, nil
}

// GetServer will get server by server ID.
func GetServer(id string) (*model.Server, []commonM.Message) {
	dbImpl := db.GetDBInstance()
	server := dbImpl.GetServerFull(id)
	// util.PrintJson(server)
	if server == nil {
		return nil, []commonM.Message{message.NewServerNotExist()}
	}
	return server, nil
}

// GetServerCollection will get server collection.
func GetServerCollection(start int, count int) (*model.ServerCollection, []commonM.Message) {
	dbImpl := db.GetDBInstance()
	ret, err := dbImpl.GetServerCollection(start, count)
	if err != nil {
		return nil, []commonM.Message{message.NewServerInternalError()}
	}
	return ret, nil
}

// RefreshServer will refresh server.
func RefreshServer(id string) (*dto.RefreshServerResponse, []commonM.Message) {
	dbImpl := db.GetDBInstance()
	server := dbImpl.GetServer(id)
	if server == nil {
		return nil, []commonM.Message{message.NewServerNotExist()}
	}
	ctx := context.CreateRefreshServerContext(server)
	st := strategy.CreateRefreshServerStrategy(server)
	if err := st.Execute(ctx); err != nil {
		return nil, ctx.Messages()
	}
	return nil, nil
}

// FindServerStateAdded will find the server with state added.
func FindServerStateAdded() {
	dbImpl := db.GetDBInstance()

	for {
		seconds := 5
		if id := dbImpl.FindServerStateAdded(); id != "" {
			RefreshServer(id)
			seconds = 0
		} else {
			seconds = 5
		}
		time.Sleep(time.Duration(seconds) * time.Second)
	}
}
