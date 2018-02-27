package entity

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
