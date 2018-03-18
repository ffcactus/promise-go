package dto

import (
	"promise/server/object/model"
)

// OemHuaweiBoard is DTO.
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

// Load will load data from model.
func (dto *OemHuaweiBoard) Load(m *model.OemHuaweiBoard) {
	dto.LoadResourceResponse(&m.Resource)
	dto.LoadProductInfoResponse(&m.ProductInfo)
	dto.CardNo = m.CardNo
	dto.DeviceLocator = m.DeviceLocator
	dto.DeviceType = m.DeviceType
	dto.Location = m.Location
	dto.CPLDVersion = m.CPLDVersion
	dto.PCBVersion = m.PCBVersion
	dto.BoardName = m.BoardName
	dto.BoardID = m.BoardID
	dto.ManufactureDate = m.ManufactureDate
}
