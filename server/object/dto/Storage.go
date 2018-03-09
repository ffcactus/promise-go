package dto

import (
	"fmt"
	"promise/server/object/model"
)

// This schema defines a storage controller and its respective properties.  A storage controller represents a storage device (physical or virtual) that produces Volumes.
type StorageController struct {
	MemberResponse
	ProductInfoResponse
	SpeedGbps                int      `json:"SpeedGbps"`                // The speed of the storage controller interface.
	FirmwareVersion          string   `json:"FirmwareVersion"`          // The firmware version of this storage Controller.
	SupportedDeviceProtocols []string `json:"SupportedDeviceProtocols"` // This represents the protocols which the storage controller can use to communicate with attached devices.
}

func (dto *StorageController) Load(m *model.StorageController) {
	dto.LoadMemberResponse(&m.Member)
	dto.LoadProductInfoResponse(&m.ProductInfo)
	dto.SpeedGbps = m.SpeedGbps
	dto.FirmwareVersion = m.FirmwareVersion
	dto.SupportedDeviceProtocols = m.SupportedDeviceProtocols
}

// This schema defines a storage subsystem and its respective properties.  A storage subsystem represents a set of storage controllers (physical or virtual) and the resources such as volumes that can be accessed from that subsystem.
type Storage struct {
	ResourceResponse
	StorageControllers []StorageController `json:"StorageControllers"` // The set of storage controllers represented by this resource.
	Drives             []ResourceRef       `json:"Drives"`             // The set of drives attached to the storage controllers represented by this resource.
}

func (dto *Storage) Load(m *model.Storage, drives []model.Drive) {
	dto.LoadResourceResponse(&m.Resource)
	for i, _ := range m.StorageControllers {
		each := StorageController{}
		each.Load(&m.StorageControllers[i])
		dto.StorageControllers = append(dto.StorageControllers, each)
	}
	for i, _ := range m.DriveURIs {
		src := m.DriveURIs[i]
		for j, _ := range drives {
			target := drives[j]
			if (target.URI != nil) && (src == *target.URI) {
				ref := ResourceRef{}
				ref.Ref = fmt.Sprintf("#/Chassis/Drives/%d", j)
				dto.Drives = append(dto.Drives, ref)
			}
		}
	}
}
