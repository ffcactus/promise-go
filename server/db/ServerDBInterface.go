package db

import (
	"promise/server/object/model"
)

// ServerDBInterface The DB interface
type ServerDBInterface interface {
	IsServerExist(s *model.Server) (bool, *model.Server)
	PostServer(s *model.Server) (*model.Server, error)
	GetServer(id string) *model.Server
	GetServerCollection(start int, count int) (*model.ServerCollection, error)
	GetServerFull(id string) *model.Server
	FindServerStateAdded() string
	GetAndLockServer(id string) (bool, *model.Server)
	SetServerState(id string, state string) bool
	SetServerHealth(id string, health string) bool
	SetServerTask(id string, taskURI string) bool
	UpdateProcessors(id string, processors []model.Processor) error
	UpdateMemory(id string, memory []model.Memory) error
	UpdateEthernetInterfaces(id string, ethernet []model.EthernetInterface) error
	UpdateNetworkInterfaces(id string, networkInterface []model.NetworkInterface) error
	UpdateStorages(id string, storages []model.Storage) error
	UpdatePower(id string, power *model.Power) error
	UpdateThermal(id string, thermal *model.Thermal) error
	UpdateOemHuaweiBoards(id string, boards []model.OemHuaweiBoard) error
	UpdateNetworkAdapters(id string, networkAdapter []model.NetworkAdapter) error
	UpdateDrives(id string, drives []model.Drive) error
	UpdatePCIeDevices(id string, pcieDevices []model.PCIeDevice) error
}
