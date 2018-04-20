package model

// OemHuaweiBoard OEM Huawei board object.
type OemHuaweiBoard struct {
	Resource
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
