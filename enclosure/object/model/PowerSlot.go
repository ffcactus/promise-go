package model

// PowerSlot represents the power slot info.
type PowerSlot struct {
	PowerSlotCommon
}

// PowerSlotCommon holds the common properties for model, dto and entity.
type PowerSlotCommon struct {
	Index           int    `gorm:"column:Index"`
	Inserted        bool   `gorm:"column:Inserted"`
	Health          string `gorm:"column:Health" json:",omitempty" `
	PowerSupplyType string `gorm:"column:PowerSupplyType" json:",omitempty"`
	SerialNumber    string `gorm:"column:SerialNumber" json:",omitempty"`
	FirmwareVersion string `gorm:"column:FirmwareVersion" json:",omitempty"`
	SleepStatus     string `gorm:"column:SleepStatus" json:",omitempty"`
}
