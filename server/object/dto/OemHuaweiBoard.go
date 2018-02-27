package dto

import (
	"promise/server/object/model"
)

type OemHuaweiBoard struct {
	ResourceResponse
	ProductInfoResponse
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

func (this *OemHuaweiBoard) Load(m *model.OemHuaweiBoard) {
	this.LoadResourceResponse(&m.Resource)
	this.LoadProductInfoResponse(&m.ProductInfo)
	this.CardNo = m.CardNo
	this.DeviceLocator = m.DeviceLocator
	this.DeviceType = m.DeviceType
	this.Location = m.Location
	this.CPLDVersion = m.CPLDVersion
	this.PCBVersion = m.PCBVersion
	this.BoardName = m.BoardName
	this.BoardID = m.BoardID
	this.ManufactureDate = m.ManufactureDate
}
