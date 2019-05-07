package dto

// This schema defines a storage controller and its respective properties.  A storage controller represents a storage device (physical or virtual) that produces Volumes.
type StorageController struct {
	Member
	ProductInfo
	SpeedGbps                int      // The speed of the storage controller interface.
	FirmwareVersion          string   // The firmware version of this storage Controller.
	SupportedDeviceProtocols []string // This represents the protocols which the storage controller can use to communicate with attached devices.
}

// This schema defines a storage subsystem and its respective properties.  A storage subsystem represents a set of storage controllers (physical or virtual) and the resources such as volumes that can be accessed from that subsystem.
type GetStorageResponse struct {
	Resource
	StorageControllers []StorageController // The set of storage controllers represented by this resource.
	Drives             []ResourceRef       // The set of drives attached to the storage controllers represented by this resource.
}
