package service

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"promise/server/client"
	"promise/server/object/dto"
	"promise/server/object/model"
)

// Probe will try to probe the server.
func Probe(request *dto.PostServerRequest) (*model.ServerBasicInfo, error) {
	c := client.FindBestClient(request.Hostname, request.Username, request.Password)
	if c == nil {
		log.WithFields(log.Fields{"hostname": request.Hostname}).Warn("Probe server failed, can not find client.")
		return nil, fmt.Errorf("failed to get server client")
	}

	serverBasicInfo, err := c.GetBasicInfo()
	if err != nil {
		log.WithFields(log.Fields{"hostname": request.Hostname, "err": err}).Warn("Probe server failed, can not get basic info.")
	}

	serverBasicInfo.Hostname = request.Hostname
	serverBasicInfo.OriginUsername = &request.Username
	serverBasicInfo.OriginPassword = &request.Password
	serverBasicInfo.Name = "Name???"
	serverBasicInfo.Description = "Description???"
	return serverBasicInfo, nil
}
