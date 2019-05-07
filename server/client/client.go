package client

import (
	log "github.com/sirupsen/logrus"
	"promise/server/client/dell"
	"promise/server/client/hp"
	"promise/server/client/huawei"
	"promise/server/client/mock"
	"promise/server/object/model"
	"strings"
)

// ServerClientInterface is the client interface for server device.
type ServerClientInterface interface {
	String() string
	Support() bool
	GetProtocol() string
	GetBasicInfo() (*model.ServerBasicInfo, error)
	CreateManagementAccount(username string, password string) error
	GetProcessors(systemID string) ([]model.Processor, error)
	GetMemory(systemID string) ([]model.Memory, error)
	GetEthernetInterfaces(systemID string) ([]model.EthernetInterface, error)
	GetNetworkInterfaces(systemID string) ([]model.NetworkInterface, error)
	GetStorages(systemID string) ([]model.Storage, error)
	// For chassis info
	GetBoards(chassisID string) ([]model.Board, error)
	GetPower(chassisID string) (*model.Power, error)
	GetThermal(chassisID string) (*model.Thermal, error)
	GetNetworkAdapters(chassisID string) ([]model.NetworkAdapter, error)
	GetDrives(chassisID string) ([]model.Drive, error)
	GetPCIeDevices(chassisID string) ([]model.PCIeDevice, error)
}

// FindBestClient will find the best client based on the server credential, which means it will detect the server real time.
func FindBestClient(vender *string, hostname string, username string, password string) ServerClientInterface {
	var client ServerClientInterface
	var support = false
	if vender == nil {
		*vender = "Huawei"
	}
	if *vender == "HP" {
		client = hp.GetInstance(hostname, username, password)
		if client.Support() {
			support = true
		}
	} else if *vender == "Dell" {
		client = dell.GetInstance(hostname, username, password)
		if client.Support() {
			support = true
		}
	} else if *vender == "Huawei" {
		client = huawei.GetInstance(hostname, username, password)
		if client.Support() {
			support = true
		}
	} else if *vender == "Mock" {
		client = mock.GetInstance(hostname)
		if client.Support() {
			support = true
		}
	}
	if client != nil {
		return client
	}
	log.WithFields(log.Fields{"client": client, "vender": *vender, "support": support}).Warn("Client find best client failed.")
	return nil
}

// TODO.
func getServerManagementAccount(server *model.Server) (string, string) {
	if server.OriginUsername != nil && *server.OriginUsername != "" {
		return *server.OriginUsername, *server.OriginPassword
	}
	s := strings.Split(server.Credential, " ")
	if len(s) == 2 {
		return s[0], s[1]
	}
	log.Warn("getServerManagementAccount(), failed to get username and password.")
	return "", ""
}

// GetServerClient will return the server client based on server which is saved.
func GetServerClient(server *model.Server) ServerClientInterface {
	username, password := getServerManagementAccount(server)
	if server.Vender == "Huawei" {
		return huawei.GetInstance(server.Hostname, username, password)
	} else if server.Vender == "HP" {
		return hp.GetInstance(server.Hostname, username, password)
	} else if server.Vender == "Dell" {
		return dell.GetInstance(server.Hostname, username, password)
	} else if server.Vender == "Mock" {
		return mock.GetInstance(server.Hostname)
	}
	log.WithFields(log.Fields{"hostname": server.Hostname, "vender": server.Vender}).Warn("Client find best client failed.")
	return nil
}
