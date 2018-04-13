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
		return nil, []commonMessage.Message{commonMessage.NewDuplicate()}
	}

	// Before save the server to the DB. We need configure the server first.
	ctx := context.CreatePostServerContext(server, request)
	st := strategy.CreatePostServerStrategy(server)
	postServer, err := st.Execute(ctx, server)
	if err != nil {
		return nil, ctx.Messages()
	}
	return postServer, nil
}

// GetServer will get server by server ID.
func GetServer(id string) (*model.Server, []commonMessage.Message) {
	dbImpl := db.GetDBInstance()
	server := dbImpl.GetServerFull(id)
	// util.PrintJson(server)
	if server == nil {
		return nil, []commonMessage.Message{commonMessage.NewNotExist()}
	}
	return server, nil
}

// GetServerCollection will get server collection.
func GetServerCollection(start int64, count int64, filter string) (*model.ServerCollection, []commonMessage.Message) {
	dbImpl := db.GetDBInstance()
	ret, err := dbImpl.GetServerCollection(start, count, filter)
	if err != nil {
		return nil, []commonMessage.Message{commonMessage.NewTransactionError()}
	}
	return ret, nil
}

// RefreshServer will refresh server.
func RefreshServer(id string) (*dto.RefreshServerResponse, []commonMessage.Message) {
	dbImpl := db.GetDBInstance()
	server := dbImpl.GetServer(id)
	if server == nil {
		return nil, []commonMessage.Message{commonMessage.NewNotExist()}
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
	exist, server, ssg, commited, err := dbImpl.DeleteServer(id)
	if !exist {
		return []commonMessage.Message{commonMessage.NewNotExist()}
	}
	if err != nil || !commited {
		return []commonMessage.Message{commonMessage.NewTransactionError()}
	}
	var eventStrategy strategy.ServerEventStrategy
	eventStrategy.DispatchServerDelete(nil, server)
	for _, each := range ssg {
		eventStrategy.DispatchServerServerGroupDelete(nil, &each)
	}
	return nil
}

// DeleteServerCollection will delete all the servers.
func DeleteServerCollection() []commonMessage.Message {
	// TODO right now, we just remove servers from DB.
	dbImpl := db.GetDBInstance()
	err := dbImpl.DeleteServerCollection()
	if err != nil {
		return []commonMessage.Message{commonMessage.NewTransactionError()}
	}
	return nil
}
