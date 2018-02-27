package service

import (
	. "promise/common/object/model"
	. "promise/server/context"
	"promise/server/db"
	"promise/server/object/dto"
	. "promise/server/object/model"
	. "promise/server/strategy"
	"time"

	"github.com/astaxie/beego"
)

// Post server.
// Since the server is not known yet. we have to create the context step by step.
func PostServer(request *dto.PostServerRequest) (*Server, []Message) {
	dbImpl := db.GetDBInstance()

	// Try to get the basic info.
	serverBasicInfo, err := Probe(request)
	if err != nil {
		beego.Info("PostServer() failed, GetBasicInfo() failed, request address = ", request.Address, ", error = ", err)
		return nil, []Message{NewPostFailed()}
	}

	server := serverBasicInfo.CreateServer()
	if exist, existedServer := dbImpl.IsServerExist(server); exist {
		return nil, []Message{NewServerExist(existedServer)}
	}

	// Before save the server to the DB. We need configure the server first.
	context := CreatePostServerContext(server, request)
	strategy := CreatePostServerStrategy(server)
	if err := strategy.Execute(context); err != nil {
		return nil, context.Messages()
	}
	return context.Server, nil
}

// Get server by server ID.
func GetServer(id string) (*Server, []Message) {
	dbImpl := db.GetDBInstance()
	server := dbImpl.GetServerFull(id)
	// util.PrintJson(server)
	if server == nil {
		return nil, []Message{NewServerNotExist()}
	}
	return server, nil
}

func GetServerCollection(start int, count int) (*ServerCollection, []Message) {
	dbImpl := db.GetDBInstance()
	if ret, err := dbImpl.GetServerCollection(start, count); err != nil {
		return nil, []Message{NewInternalError()}
	} else {
		return ret, nil
	}
}

func RefreshServer(id string) (*dto.RefreshServerResponse, []Message) {
	dbImpl := db.GetDBInstance()
	server := dbImpl.GetServer(id)
	if server == nil {
		return nil, []Message{NewServerNotExist()}
	}
	context := CreateRefreshServerContext(server)
	strategy := CreateRefreshServerStrategy(server)
	if err := strategy.Execute(context); err != nil {
		return nil, context.Messages()
	}
	return nil, nil
}

// This is the channel used for auto-refresh server with state "Added"
//var AutoRefreshChannel chan(string) = make(chan string)

// This is the producer of auto-refreshing server with state "Added"
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
