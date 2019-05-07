package dto

type ChassisLinks struct {
	ComputerSystems []ResourceRef
	ManagedBy       []ResourceRef
	Drives          []ResourceRef
	PCIeDevices     []ResourceRef
}

type GetChassisResponse struct {
	Resource
	IndicatorLed *string
	ChassisType  *string `json:"ChassisType"`
	Links        ChassisLinks
}
