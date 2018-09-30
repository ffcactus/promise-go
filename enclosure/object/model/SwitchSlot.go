package model

// SwitchSlot represents the switch slot info.
type SwitchSlot struct {
	SwitchSlotCommon
}

// SwitchSlotCommon holds the common properties for model, dto and entity.
type SwitchSlotCommon struct {
	Index       int    `gorm:"column:Index"`
	Inserted    bool   `gorm:"column:Inserted"`
	ProductName string `gorm:"column:ProductName" json:",omitempty"`
}
