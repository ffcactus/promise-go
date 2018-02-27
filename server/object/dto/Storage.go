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

func (this *StorageController) Load(m *model.StorageController) {
	this.LoadMemberResponse(&m.Member)
	this.LoadProductInfoResponse(&m.ProductInfo)
	this.SpeedGbps = m.SpeedGbps
	this.FirmwareVersion = m.FirmwareVersion
	this.SupportedDeviceProtocols = m.SupportedDeviceProtocols
}

// This schema defines a storage subsystem and its respective properties.  A storage subsystem represents a set of storage controllers (physical or virtual) and the resources such as volumes that can be accessed from that subsystem.
type Storage struct {
	ResourceResponse
	StorageControllers []StorageController `json:"StorageControllers"` // The set of storage controllers represented by this resource.
	Drives             []ResourceRef       `json:"Drives"`             // The set of drives attached to the storage controllers represented by this resource.
}

func (this *Storage) Load(m *model.Storage, drives []model.Drive) {
	this.LoadResourceResponse(&m.Resource)
	for i, _ := range m.StorageControllers {
		each := StorageController{}
		each.Load(&m.StorageControllers[i])
		this.StorageControllers = append(this.StorageControllers, each)
	}
	for i, _ := range m.DriveURIs {
		src := m.DriveURIs[i]
		for j, _ := range drives {
			target := drives[j]
			if (target.URI != nil) && (src == *target.URI) {
				ref := ResourceRef{}
				ref.Ref = fmt.Sprintf("#/Chassis/Drives/%d", j)
				this.Drives = append(this.Drives, ref)
			}
		}
	}
}
