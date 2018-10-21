package model

// ServerSlot represents the blade slot info.
type ServerSlot struct {
	ServerSlotCommon
}

// ServerSlotCommon holds the common properties for model, dto and entity.
type ServerSlotCommon struct {
	Index        int    `gorm:"column:Index"`
	Inserted     bool   `gorm:"column:Inserted"`
	ProductName  string `gorm:"column:ProductName" json:",omitempty"`
	SerialNumber string `gorm:"column:SerialNumber" json:",omitempty"`
	Height       int    `gorm:"column:Height" json:",omitempty"`
	Width        int    `gorm:"column:Width" json:",omitempty"`
	ServerURL    string `gorm:"column:ServerURL" json:",omitempty"`
}
