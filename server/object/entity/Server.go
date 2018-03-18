package entity

import (
	"time"
)

// Server Server entity.
type Server struct {
	Entity
	State              string
	Health             string
	CreatedAt          time.Time
	UpdatedAt          time.Time
	URI                string
	Name               string
	Description        string
	OriginURIsChassis  *string
	OriginURIsSystem   *string
	PhysicalUUID       string
	Hostname           string
	Type               string
	Protocol           string
	Credential         string
	CurrentTask        string
	Processors         []Processor         `gorm:"ForeignKey:ServerRef"`
	Memory             []Memory            `gorm:"ForeignKey:ServerRef"`
	EthernetInterfaces []EthernetInterface `gorm:"ForeignKey:ServerRef"`
	NetworkInterfaces  []NetworkInterface  `gorm:"ForeignKey:ServerRef"`
	Storages           []Storage           `gorm:"ForeignKey:ServerRef"`
	Power              Power               `gorm:"ForeignKey:ServerRef"`
	Thermal            Thermal             `gorm:"ForeignKey:ServerRef"`
	OemHuaweiBoards    []OemHuaweiBoard    `gorm:"ForeignKey:ServerRef"`
	Drives             []Drive             `gorm:"ForeignKey:ServerRef"`
	PCIeDevices        []PCIeDevice        `gorm:"ForeignKey:ServerRef"`
	NetworkAdapters    []NetworkAdapter    `gorm:"ForeignKey:ServerRef"`
}
