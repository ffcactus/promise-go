package model

// ApplianceSlot represents the appliance slot info.
type ApplianceSlot struct {
	ApplianceSlotCommon
}

// ApplianceSlotCommon holds the common properties for model, dto and entity.
type ApplianceSlotCommon struct {
	Index           int    `gorm:"column:Index"`
	Inserted        bool   `gorm:"column:Inserted"`
	SerialNumber    string `gorm:"column:SerialNumber" json:",omitempty"`
	FirmwareVersion string `gorm:"column:FirmwareVersion" json:",omitempty"`
	BIOSVersion     string `gorm:"column:BIOSVersion" json:",omitempty"`
}
