package entity

import (
	commonUtil "promise/common/util"
	"promise/server/object/model"
)

// StorageController This schema defines a storage controller and its respective properties.  A storage controller represents a storage device (physical or virtual) that produces Volumes.
type StorageController struct {
	StorageRef uint
	Member
	ProductInfo
	SpeedGbps                int    // The speed of the storage controller interface.
	FirmwareVersion          string // The firmware version of this storage Controller.
	SupportedDeviceProtocols string // This represents the protocols which the storage controller can use to communicate with attached devices.
}

// Storage This schema defines a storage subsystem and its respective properties.  A storage subsystem represents a set of storage controllers (physical or virtual) and the resources such as volumes that can be accessed from that subsystem.
type Storage struct {
	ServerRef string
	EmbeddedResource
	StorageControllers []StorageController `gorm:"ForeignKey:StorageRef"` // The set of storage controllers represented by this resource.
	DriveURIs          string              // The set of drives attached to the storage controllers represented by this resource.
}

// ToModel will create a new model from entity.
func (e *Storage) ToModel() *model.Storage {
	m := model.Storage{}
	createResourceModel(&e.EmbeddedResource, &m.Resource)
	a := []string{}
	commonUtil.StringToStruct(e.DriveURIs, &a)
	m.DriveURIs = a
	for i := range e.StorageControllers {
		eachM := model.StorageController{}
		eachE := e.StorageControllers[i]
		createMemberModel(&eachE.Member, &eachM.Member)
		createProductInfoModel(&eachE.ProductInfo, &eachM.ProductInfo)
		eachM.SpeedGbps = eachE.SpeedGbps
		eachM.FirmwareVersion = eachE.FirmwareVersion
		a := []string{}
		commonUtil.StringToStruct(eachE.SupportedDeviceProtocols, &a)
		eachM.SupportedDeviceProtocols = a
		m.StorageControllers = append(m.StorageControllers, eachM)
	}
	return &m
}

// Load will load data from model.
func (e *Storage) Load(m *model.Storage) {
	updateResourceEntity(&e.EmbeddedResource, &m.Resource)
	s := commonUtil.StructToString(m.DriveURIs)
	e.DriveURIs = s
	for i := range m.StorageControllers {
		each := StorageController{}
		updateMemberEntity(&each.Member, &m.StorageControllers[i].Member)
		updateProductInfoEntity(&each.ProductInfo, &m.StorageControllers[i].ProductInfo)
		each.SpeedGbps = m.StorageControllers[i].SpeedGbps
		each.FirmwareVersion = m.StorageControllers[i].FirmwareVersion
		s := commonUtil.StructToString(m.StorageControllers[i].SupportedDeviceProtocols)
		each.SupportedDeviceProtocols = s
		e.StorageControllers = append(e.StorageControllers, each)
	}
}