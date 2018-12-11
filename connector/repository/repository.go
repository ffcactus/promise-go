package repository

import (
	"connector/dto"
)

type RedfishRepository interface {
	dto.Chassis[] Chassis()
	dto.System[] System()
	dto.Manager[] Manager()
	dto.FirmwareInventory FirmwareInventory()
}

type RedfishRepositoryImpl struct {
	client Client
}

func (impl *RedfishRepositoryImpl) Chassis() {
	
}

func (impl *RedfishRepositoryImpl) System() {

}

func (impl *RedfishRepositoryImpl) Manager() {

}

func (impl *RedfishRepositoryImpl) FirmwareInventory() {

}
