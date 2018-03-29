package entity

import (
	"promise/server/object/model"
)

// OemHuaweiBoard OEM Huawei board.
type OemHuaweiBoard struct {
	ServerRef string
	EmbeddedResource
	ProductInfo
	CardNo          *int
	DeviceLocator   *string
	DeviceType      *string
	Location        *string
	CPLDVersion     *string
	PCBVersion      *string
	BoardName       *string
	BoardID         *string
	ManufactureDate *string
}

// ToModel will create a new model from entity.
func (e *OemHuaweiBoard) ToModel() *model.OemHuaweiBoard {
	m := new(model.OemHuaweiBoard)
	createResourceModel(&e.EmbeddedResource, &m.Resource)
	createProductInfoModel(&e.ProductInfo, &m.ProductInfo)
	m.CardNo = e.CardNo
	m.DeviceLocator = e.DeviceLocator
	m.DeviceType = e.DeviceType
	m.Location = e.Location
	m.CPLDVersion = e.CPLDVersion
	m.PCBVersion = e.PCBVersion
	m.BoardName = e.BoardName
	m.BoardID = e.BoardID
	m.ManufactureDate = e.ManufactureDate
	return m
}
