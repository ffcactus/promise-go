package model

// FanSlot represents the fan slot info.
type FanSlot struct {
	FanSlotCommon
}

// FanSlotCommon holds the common properties for model, dto and entity.
type FanSlotCommon struct {
	Index           int    `gorm:"column:Index"`
	Inserted        bool   `gorm:"column:Inserted"`
	Health          string `gorm:"column:Health" json:",omitempty"`
	PCBVersion      string `gorm:"column:PCBVersion" json:",omitempty"`
	SoftwareVersion string `gorm:"column:SoftwareVersion" json:",omitempty"`
}
