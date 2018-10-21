package model

// ManagerSlot represents the manager slot info.
type ManagerSlot struct {
	ManagerSlotCommon
}

// ManagerSlotCommon holds the common properties for model, dto and entity.
type ManagerSlotCommon struct {
	Index           int    `gorm:"column:Index"`
	Inserted        bool   `gorm:"column:Inserted"`
	ProductName     string `gorm:"column:ProductName" json:",omitempty"`
	SerialNumber    string `gorm:"column:SerialNumber" json:",omitempty"`
	FirmwareVersion string `gorm:"column:FirmwareVersion" json:",omitempty"`
	CPLDVersion     string `gorm:"column:CPLDVersion" json:",omitempty"`
	ServerURL       string `gorm:"column:ServerURL" json:",omitempty"`
}
