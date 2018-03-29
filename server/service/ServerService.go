package service

import (
	commonMessage "promise/common/object/message"
	"promise/server/context"
	"promise/server/db"
	"promise/server/object/dto"
	"promise/server/object/message"
	"promise/server/object/model"
	"promise/server/strategy"
	"time"
)

// PostServer since the server is not known yet. we have to create the context step by step.
func PostServer(request *dto.PostServerRequest) (*model.Server, []commonMessage.Message) {
	dbImpl := db.GetDBInstance()

	// Try to get the basic info.
	serverBasicInfo, err := Probe(request)
	if err != nil {
		return nil, []commonMessage.Message{message.NewServerPostFailed()}
	}

	server := serverBasicInfo.CreateServer()
	if exist, _ := dbImpl.IsServerExist(server); exist {
		return nil, []commonMessage.Message{commonMessage.NewResourceDuplicate()}
	}

	// Before save the server to the DB. We need configure the server first.
	ctx := context.CreatePostServerContext(server, request)
	st := strategy.CreatePostServerStrategy(server)
	if err := st.Execute(ctx, server); err != nil {
		return nil, ctx.Messages()
	}
	return server, nil
}

// GetServer will get server by server ID.
func GetServer(id string) (*model.Server, []commonMessage.Message) {
	dbImpl := db.GetDBInstance()
	server := dbImpl.GetServerFull(id)
	// util.PrintJson(server)
	if server == nil {
		return nil, []commonMessage.Message{commonMessage.NewResourceNotExist()}
	}
	return server, nil
}

// GetServerCollection will get server collection.
func GetServerCollection(start int, count int) (*model.ServerCollection, []commonMessage.Message) {
	dbImpl := db.GetDBInstance()
	ret, err := dbImpl.GetServerCollection(start, count)
	if err != nil {
		return nil, []commonMessage.Message{commonMessage.NewInternalError()}
	}
	return ret, nil
}

// RefreshServer will refresh server.
func RefreshServer(id string) (*dto.RefreshServerResponse, []commonMessage.Message) {
	dbImpl := db.GetDBInstance()
	server := dbImpl.GetServer(id)
	if server == nil {
		return nil, []commonMessage.Message{commonMessage.NewResourceNotExist()}
	}
	ctx := context.CreateRefreshServerContext(server)
	st := strategy.CreateRefreshServerStrategy(server)
	if err := st.Execute(ctx, server); err != nil {
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

// DeleteServer will delete server group by ID.
func DeleteServer(id string) []commonMessage.Message {
	// TODO right now, we just remove server from DB.
	dbImpl := db.GetDBInstance()
	exist, err := dbImpl.DeleteServer(id)
	if err != nil {
		return []commonMessage.Message{commonMessage.NewInternalError()}
	}
	if !exist {
		return []commonMessage.Message{commonMessage.NewResourceNotExist()}
	}
	return nil
}

// DeleteServerCollection will delete all the servers.
func DeleteServerCollection() []commonMessage.Message {
	// TODO right now, we just remove servers from DB.
	dbImpl := db.GetDBInstance()
	err := dbImpl.DeleteServerCollection()
	if err != nil {
		return []commonMessage.Message{commonMessage.NewInternalError()}
	}
	return nil
}
