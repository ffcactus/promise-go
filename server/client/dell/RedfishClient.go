package dell

import (
	"errors"
	// "promise/base"
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/server/client/hp/dto"
	"promise/server/object/constvalue"
	"promise/server/object/model"
	"strings"
)

// RedfishClient is the redfish client for HP servers.
type RedfishClient struct {
	base.Client
}

// GetInstance Get a new instance of Redfish client.
func GetInstance(address string, username string, password string) *RedfishClient {
	client := RedfishClient{}
	client.Protocol = "https"
	client.CurrentAddress = address
	client.Username = username
	client.Password = password
	return &client
}

// Support check if support.
func (c *RedfishClient) Support() bool {
	if err := c.Get("/redfish/v1", nil); err != nil {
		return false
	}
	return true
}

// String returns the client info.
func (c RedfishClient) String() string {
	return "Dell Redfish " + c.CurrentAddress
}

// GetProtocol Get the protocal used by this client.
func (c *RedfishClient) GetProtocol() string {
	return constvalue.RedfishV1
}

// GetBasicInfo Get server basic info.
// Just set parts of the properties.
func (c *RedfishClient) GetBasicInfo() (*model.ServerBasicInfo, error) {
	// First set the server type.
	var chassisCollection = dto.Collection{}
	var systemURI, chassisURI string

	if err := c.Get("/redfish/v1/Chassis", &chassisCollection); err != nil {
		return nil, err
	}

	var systemCollection = dto.Collection{}
	if err := c.Get("/redfish/v1/Systems", &systemCollection); err != nil {
		return nil, err
	}

	if systemCollection.Count != 1 {
		log.WithFields(log.Fields{"client": c}).Warn("GetBasicInfo failed, systems count not equal 1.")
		return nil, errors.New("can not find system URI")
	}
	systemURI = systemCollection.Members[0].Id

	for _, uri := range chassisCollection.Members {
		if strings.Contains(uri.Id, "System.Embedded") {
			chassisURI = uri.Id
		}
	}
	if chassisURI == "" {
		log.WithFields(log.Fields{"client": c}).Warn("GetBasicInfo failed, can not find chassis URI.")
		return nil, errors.New("can not find chassis URI")
	}

	// Get info from Computer system.
	var system = dto.GetSystemResponse{}
	if err := c.Get(systemURI, &system); err != nil {
		return nil, err
	}
	ret := model.ServerBasicInfo{}
	ret.Vender = "Dell"
	ret.OriginURIs.System = systemURI
	ret.OriginURIs.Chassis = chassisURI
	ret.PhysicalUUID = system.PhysicalUUID
	ret.Protocol = constvalue.RedfishV1
	// Get info from chassis.
	var chassis = dto.GetChassisResponse{}
	if err := c.Get(chassisURI, &chassis); err != nil {
		return nil, err
	}
	if *chassis.ChassisType == "" {
		log.WithFields(log.Fields{"client": c}).Warn("Get basic info failed, failed to get chassis type.")
		return nil, errors.New("failed to get server type")
	}
	ret.Type = *chassis.ChassisType
	return &ret, nil

}

// CreateManagementAccount Create Management account.
func (c *RedfishClient) CreateManagementAccount(username string, password string) error {
	requestBody := dto.PostAccountRequest{
		UserName: username,
		Password: password,
		RoleId:   "Administrator",
	}
	return c.Post("/redfish/v1/AccountService/Accounts", requestBody, nil)
}

// GetProcessors Get server's process info.
func (c *RedfishClient) GetProcessors(systemID string) ([]model.Processor, error) {
	collection := dto.Collection{}
	if err := c.Get(systemID+"/Processors", &collection); err != nil {
		return nil, err
	}
	var ret []model.Processor
	for i := range collection.Members {
		each := new(dto.GetProcessorResponse)
		if err := c.Get(collection.Members[i].Id, each); err != nil {
			return nil, err
		}
		ret = append(ret, *createProcessorModel(each))
	}
	return ret, nil
}

// GetMemory Get server's memory info.
func (c *RedfishClient) GetMemory(systemID string) ([]model.Memory, error) {
	collection := dto.Collection{}
	if err := c.Get(systemID+"/Memory", &collection); err != nil {
		return nil, err
	}

	var ret []model.Memory
	for i := range collection.Members {
		each := new(dto.GetMemoryResponse)
		if err := c.Get(collection.Members[i].Id, each); err != nil {
			return nil, err
		}
		ret = append(ret, *createMemoryModel(each))
	}
	return ret, nil
}

// GetEthernetInterfaces Get server's ethernet interface info.
func (c *RedfishClient) GetEthernetInterfaces(systemID string) ([]model.EthernetInterface, error) {
	collection := dto.Collection{}
	if err := c.Get(systemID+"/EthernetInterfaces", &collection); err != nil {
		return nil, err
	}
	var ret []model.EthernetInterface
	for i := range collection.Members {
		eachEthernet := new(dto.GetEthernetInterfaceResponse)
		if err := c.Get(collection.Members[i].Id, eachEthernet); err != nil {
			return nil, err
		}
		// Get the VLANs
		vlanCollection := dto.Collection{}
		vlanCollectionPageURI := systemID + "/EthernetInterfaces/" + *eachEthernet.Id + "/VLANs"
		if err := c.Get(vlanCollectionPageURI, &vlanCollection); err != nil {
			return nil, err
		}
		var vlans []model.VLanNetworkInterface
		for j := range vlanCollection.Members {
			eachVlan := new(dto.GetVLANResponse)
			if err := c.Get(collection.Members[j].Id, eachVlan); err != nil {
				return nil, err
			}
			vlans = append(vlans, *createVLanModel(eachVlan))
		}
		ethernetMode := *createEthernetInterfaceModel(eachEthernet)
		ethernetMode.VLANs = vlans
		ret = append(ret, ethernetMode)
	}
	return ret, nil
}

// GetNetworkInterfaces get network interfaces.
func (c *RedfishClient) GetNetworkInterfaces(systemID string) ([]model.NetworkInterface, error) {
	collection := dto.Collection{}
	if err := c.Get(systemID+"/NetworkInterfaces", &collection); err != nil {
		return nil, err
	}
	var ret []model.NetworkInterface
	for i := range collection.Members {
		networkInterface := new(dto.GetNetworkInterfaceResponse)
		if err := c.Get(collection.Members[i].Id, networkInterface); err != nil {
			return nil, err
		}

		ret = append(ret, *createNetworkInterfaceModel(networkInterface))
	}
	return ret, nil
}

// GetStorages get storages.
func (c *RedfishClient) GetStorages(systemID string) ([]model.Storage, error) {
	collection := dto.Collection{}
	if err := c.Get(systemID+"/Storages", &collection); err != nil {
		return nil, err
	}
	ret := []model.Storage{}
	for i := range collection.Members {
		storage := new(dto.GetStorageResponse)
		if err := c.Get(collection.Members[i].Id, storage); err != nil {
			return nil, err
		}

		ret = append(ret, *createStorageModel(storage))
	}
	return ret, nil
}

// GetPower get power.
func (c *RedfishClient) GetPower(chassisID string) (*model.Power, error) {
	power := new(dto.GetPowerResponse)
	if err := c.Get(chassisID+"/Power", power); err != nil {
		return nil, err
	}
	model := createPowerModel(power)
	return model, nil
}

// GetThermal get thermal.
func (c *RedfishClient) GetThermal(chassisID string) (*model.Thermal, error) {
	thermal := new(dto.GetThermalResponse)
	if err := c.Get(chassisID+"/Thermal", thermal); err != nil {
	}
	model := createThermalModel(thermal)
	return model, nil
}

// GetBoards get oem huawei boards.
func (c *RedfishClient) GetBoards(chassisID string) ([]model.Board, error) {
	collection := dto.Collection{}
	if err := c.Get(chassisID+"/Boards", &collection); err != nil {
		return nil, err
	}
	ret := []model.Board{}
	for i := range collection.Members {
		board := new(dto.GetBoardsResponse)
		if err := c.Get(collection.Members[i].Id, board); err != nil {
			return nil, err
		}
		ret = append(ret, *createBoardsModel(board))
	}
	return ret, nil
}

// GetNetworkAdapters get networkadapters.
func (c *RedfishClient) GetNetworkAdapters(chassisID string) ([]model.NetworkAdapter, error) {
	collection := dto.Collection{}
	if err := c.Get(chassisID+"/NetworkAdapters", &collection); err != nil {
		return nil, err
	}
	var ret []model.NetworkAdapter
	for i := range collection.Members {
		resp := new(dto.GetNetworkAdapterResponse)
		if err := c.Get(collection.Members[i].Id, resp); err != nil {
			return nil, err
		}
		networkAdpter := createNetworkAdapterModel(resp)
		for j := range resp.Controllers {
			eachController := createControllerModel(&resp.Controllers[j])
			portsResp := resp.Controllers[j].Links.NetworkPorts
			for k := range portsResp {
				portPageURI := portsResp[k].OdataId
				portResp := new(dto.NetworkPort)
				if err := c.Get(portPageURI, portResp); err != nil {
					return nil, err
				}
				eachController.NetworkPorts = append(eachController.NetworkPorts, *createNetworkPortModel(portResp))
			}
			networkAdpter.Controllers = append(networkAdpter.Controllers, *eachController)
		}
		ret = append(ret, *networkAdpter)
	}
	// util.PrintJson(ret)
	return ret, nil
}

// GetDrives get drives.
func (c *RedfishClient) GetDrives(chassisID string) ([]model.Drive, error) {
	// Get the Drive links from chassis.
	chassis := new(dto.GetChassisResponse)
	if err := c.Get(chassisID, chassis); err != nil {
		return nil, err
	}
	ret := []model.Drive{}
	for i := range chassis.Links.Drives {
		uri := chassis.Links.Drives[i].OdataId
		drive := new(dto.GetDriveResponse)
		if err := c.Get(uri, drive); err != nil {
			return nil, err
		}
		ret = append(ret, *createDriveModel(drive))
	}
	return ret, nil
}

// GetPCIeDevices get PCIeDevices.
func (c *RedfishClient) GetPCIeDevices(chassisID string) ([]model.PCIeDevice, error) {
	// Get the Drive links from chassis.
	chassis := new(dto.GetChassisResponse)
	if err := c.Get(chassisID, chassis); err != nil {
		return nil, err
	}
	ret := []model.PCIeDevice{}
	for i := range chassis.Links.PCIeDevices {
		uri := chassis.Links.PCIeDevices[i].OdataId
		pcieDevice := new(dto.GetPCIeDeviceResponse)
		if err := c.Get(uri, pcieDevice); err != nil {
			return nil, err
		}
		pcieFunctions := new([]dto.GetPCIeFunctionResponse)
		for j := range pcieDevice.Links.PCIeFunctions {
			uri := pcieDevice.Links.PCIeFunctions[j].OdataId
			pcieFunction := new(dto.GetPCIeFunctionResponse)
			if err := c.Get(uri, pcieFunction); err != nil {
				return nil, err
			}
			*pcieFunctions = append(*pcieFunctions, *pcieFunction)
		}

		ret = append(ret, *createPCIeDeviceModel(pcieDevice, pcieFunctions))
	}
	return ret, nil
}

// GetNetworkPorts get network ports.
func (c *RedfishClient) GetNetworkPorts(uri string) ([]model.NetworkPort, error) {
	collection := dto.Collection{}
	if err := c.Get(uri, &collection); err != nil {
		return nil, err
	}
	var ret []model.NetworkPort
	for i := range collection.Members {
		resp := new(dto.NetworkPort)
		if err := c.Get(collection.Members[i].Id, resp); err != nil {
			return nil, err
		}
		ret = append(ret, *createNetworkPortModel(resp))
	}
	return ret, nil
}

func createProcessorModel(d *dto.GetProcessorResponse) *model.Processor {
	ret := model.Processor{}
	createResourceModel(&d.Resource, &ret.Resource)
	createProductInfoModel(&d.ProductInfo, &ret.ProductInfo)
	ret.Socket = d.Socket
	ret.ProcessorType = d.ProcessorType
	ret.ProcessorArchitecture = d.ProcessorArchitecture
	ret.InstructionSet = d.InstructionSet
	ret.MaxSpeedMHz = d.MaxSpeedMHz
	ret.TotalCores = d.TotalCores
	return &ret
}

func createResourceModel(d *dto.Resource, m *model.Resource) {
	m.URI = d.OdataID
	m.Name = d.Name
	m.Description = d.Description
	m.OriginID = d.Id
	if d.Status != nil {
		m.PhysicalState = d.Status.State
		m.PhysicalHealth = d.Status.Health
	}
}

func createMemberModel(d *dto.Member, m *model.Member) {
	m.URI = d.OdataID
	m.Name = d.Name
	m.Description = d.Description
	m.OriginMemberID = d.MemberId
	if d.Status != nil {
		m.PhysicalState = d.Status.State
		m.PhysicalHealth = d.Status.Health
	}
}

func createThresholdModel(d *dto.Threshold, m *model.Threshold) {
	m.UpperThresholdNonCritical = d.UpperThresholdNonCritical
	m.UpperThresholdCritical = d.UpperThresholdCritical
	m.UpperThresholdFatal = d.UpperThresholdFatal
	m.LowerThresholdNonCritical = d.LowerThresholdNonCritical
	m.LowerThresholdCritical = d.LowerThresholdCritical
	m.LowerThresholdFatal = d.LowerThresholdFatal
}

func createProductInfoModel(d *dto.ProductInfo, m *model.ProductInfo) {
	m.Model = d.Model
	m.Manufacturer = d.Manufacturer
	m.SKU = d.SKU
	m.SerialNumber = d.SerialNumber
	m.PartNumber = d.PartNumber
	m.SparePartNumber = d.SparePartNumber
	m.AssetTag = d.AssetTag
}

func createLocationModel(d *dto.Location, m *model.Location) {
	m.Info = d.Info
	m.InfoFormat = d.InfoFormat
}

func createMemoryModel(d *dto.GetMemoryResponse) *model.Memory {
	ret := model.Memory{}
	createResourceModel(&d.Resource, &ret.Resource)
	createProductInfoModel(&d.ProductInfo, &ret.ProductInfo)
	ret.CapacityMiB = d.CapacityMiB
	ret.OperatingSpeedMhz = d.OperatingSpeedMhz
	ret.MemoryDeviceType = d.MemoryDeviceType
	ret.DataWidthBits = d.DataWidthBits
	ret.RankCount = d.RankCount
	ret.DeviceLocator = d.DeviceLocator

	if d.MemoryLocation != nil {
		ret.MemoryLocation = new(model.MemoryLocation)
		ret.MemoryLocation.Socket = d.MemoryLocation.Socket
		ret.MemoryLocation.Controller = d.MemoryLocation.Controller
		ret.MemoryLocation.Channel = d.MemoryLocation.Channel
		ret.MemoryLocation.Slot = d.MemoryLocation.Slot
	}
	return &ret
}

func createEthernetInterfaceModel(d *dto.GetEthernetInterfaceResponse) *model.EthernetInterface {
	ret := model.EthernetInterface{}
	createResourceModel(&d.Resource, &ret.Resource)
	ret.UefiDevicePath = d.UefiDevicePath
	ret.InterfaceEnabled = d.InterfaceEnabled
	ret.PermanentMACAddress = d.PermanentMACAddress
	ret.MACAddress = d.MACAddress
	ret.SpeedMbps = d.SpeedMbps
	ret.AutoNeg = d.AutoNeg
	ret.FullDuplex = d.FullDuplex
	ret.MTUSize = d.MTUSize
	ret.HostName = d.HostName
	ret.FQDN = d.FQDN
	ret.MaxIPv6StaticAddresses = d.MaxIPv6StaticAddresses
	ret.LinkStatus = d.LinkStatus
	if d.IPv4Addresses != nil {
		ipv4 := []model.IPv4Address{}
		for i := range *d.IPv4Addresses {
			each := model.IPv4Address{}
			each.Address = (*d.IPv4Addresses)[i].Address
			each.SubnetMask = (*d.IPv4Addresses)[i].SubnetMask
			each.AddressOrigin = (*d.IPv4Addresses)[i].AddressOrigin
			each.Gateway = (*d.IPv4Addresses)[i].Gateway
			ipv4 = append(ipv4, each)
		}
		ret.IPv4Addresses = ipv4
	}
	if d.IPv6Addresses != nil {
		ipv6 := []model.IPv6Address{}
		for i := range *d.IPv6Addresses {
			each := model.IPv6Address{}
			each.Address = (*d.IPv6Addresses)[i].Address
			each.PrefixLength = (*d.IPv6Addresses)[i].PrefixLength
			each.AddressOrigin = (*d.IPv6Addresses)[i].AddressOrigin
			each.AddressState = (*d.IPv6Addresses)[i].AddressState
			ipv6 = append(ipv6, each)
		}
		ret.IPv6Addresses = ipv6
	}
	return &ret
}

func createVLanModel(d *dto.GetVLANResponse) *model.VLanNetworkInterface {
	ret := model.VLanNetworkInterface{}
	createResourceModel(&d.Resource, &ret.Resource)
	ret.VLANEnable = d.VLANEnable
	ret.VLANID = d.VLANID
	return &ret
}

func createNetworkInterfaceModel(d *dto.GetNetworkInterfaceResponse) *model.NetworkInterface {
	ret := model.NetworkInterface{}
	createResourceModel(&d.Resource, &ret.Resource)
	ret.NetworkAdapterURI = d.Links.NetworkAdapter.OdataId
	return &ret
}

func createStorageControllerModel(d *dto.StorageController) *model.StorageController {
	ret := model.StorageController{}
	createMemberModel(&d.Member, &ret.Member)
	createProductInfoModel(&d.ProductInfo, &ret.ProductInfo)
	ret.SpeedGbps = d.SpeedGbps
	ret.FirmwareVersion = d.FirmwareVersion
	ret.SupportedDeviceProtocols = d.SupportedDeviceProtocols
	return &ret
}

func createStorageModel(d *dto.GetStorageResponse) *model.Storage {
	ret := model.Storage{}
	createResourceModel(&d.Resource, &ret.Resource)
	for i := range d.Drives {
		ret.DriveURIs = append(ret.DriveURIs, d.Drives[i].OdataId)
	}
	for i := range d.StorageControllers {
		ret.StorageControllers = append(ret.StorageControllers, *createStorageControllerModel(&d.StorageControllers[i]))
	}
	return &ret
}

func createPowerModel(d *dto.GetPowerResponse) *model.Power {
	r := *d
	ret := model.Power{}
	createResourceModel(&d.Resource, &ret.Resource)
	// PowerControl
	powerControl := []model.PowerControl{}
	for i := range *r.PowerControl {
		eachModel := model.PowerControl{}
		eachDto := (*r.PowerControl)[i]
		createResourceModel(&eachDto.Resource, &eachModel.Resource)
		createProductInfoModel(&eachDto.ProductInfo, &eachModel.ProductInfo)
		eachModel.PowerConsumedWatts = eachDto.PowerConsumedWatts
		eachModel.PowerRequestedWatts = eachDto.PowerRequestedWatts
		eachModel.PowerAvailableWatts = eachDto.PowerAvailableWatts
		eachModel.PowerCapacityWatts = eachDto.PowerCapacityWatts
		eachModel.PowerAllocatedWatts = eachDto.PowerAllocatedWatts
		if eachDto.PowerMetrics != nil {
			powerMetrics := model.PowerMetrics{}
			powerMetrics.MinConsumedWatts = eachDto.PowerMetrics.MinConsumedWatts
			powerMetrics.MaxConsumedWatts = eachDto.PowerMetrics.MaxConsumedWatts
			powerMetrics.AverageConsumedWatts = eachDto.PowerMetrics.AverageConsumedWatts
			eachModel.PowerMetrics = &powerMetrics
		}
		if eachDto.PowerLimit != nil {
			powerLimit := model.PowerLimit{}
			powerLimit.LimitInWatts = eachDto.PowerLimit.LimitInWatts
			powerLimit.LimitException = eachDto.PowerLimit.LimitException
			powerLimit.CorrectionInMs = eachDto.PowerLimit.CorrectionInMs
			eachModel.PowerLimit = &powerLimit
		}
		powerControl = append(powerControl, eachModel)
	}
	ret.PowerControl = powerControl

	// PowerSupplies
	powerSupplies := []model.PowerSupply{}
	for i := range *r.PowerSupplies {
		eachModel := model.PowerSupply{}
		eachDto := (*r.PowerSupplies)[i]
		createResourceModel(&eachDto.Resource, &eachModel.Resource)
		createProductInfoModel(&eachDto.ProductInfo, &eachModel.ProductInfo)
		eachModel.PowerSupplyType = eachDto.PowerSupplyType
		eachModel.LineInputVoltageType = eachDto.LineInputVoltageType
		eachModel.LineInputVoltage = eachDto.LineInputVoltage
		eachModel.PowerCapacityWatts = eachDto.PowerCapacityWatts
		eachModel.LastPowerOutputWatts = eachDto.LastPowerOutputWatts
		eachModel.FirmwareVersion = eachDto.FirmwareVersion
		eachModel.IndicatorLed = eachDto.IndicatorLed
		powerSupplies = append(powerSupplies, eachModel)
	}
	ret.PowerSupplies = powerSupplies

	// Redundancy
	redundancy := []model.Redundancy{}
	for i := range *r.Redundancy {
		eachModel := model.Redundancy{}
		eachDto := (*r.Redundancy)[i]
		createResourceModel(&eachDto.Resource, &eachModel.Resource)
		eachModel.Mode = eachDto.Mode
		eachModel.MaxNumSupported = eachDto.MaxNumSupported
		eachModel.MinNumNeeded = eachDto.MinNumNeeded
		eachModel.RedundancyEnabled = eachDto.RedundancyEnabled
		// only name is needed in the name of redundancy set.
		redundancySet := []string{}
		for j := range *eachDto.RedundancySet {
			for k := range *r.PowerSupplies {
				redundancyOdataID := (*eachDto.RedundancySet)[j].OdataId
				powerSupply := (*r.PowerSupplies)[k]
				if *powerSupply.OdataID == redundancyOdataID {
					redundancySet = append(redundancySet, *powerSupply.Name)
				}
			}
		}
		eachModel.RedundancySet = &redundancySet
		redundancy = append(redundancy, eachModel)
	}
	ret.Redundancy = redundancy
	return &ret
}

func createThermalModel(d *dto.GetThermalResponse) *model.Thermal {
	ret := new(model.Thermal)
	createResourceModel(&d.Resource, &ret.Resource)

	fans := []model.Fan{}
	for i := range d.Fans {
		each := model.Fan{}
		createMemberModel(&d.Fans[i].Member, &each.Member)
		createProductInfoModel(&d.Fans[i].ProductInfo, &each.ProductInfo)
		createThresholdModel(&d.Fans[i].Threshold, &each.Threshold)
		each.Reading = d.Fans[i].Reading
		each.MinReadingRange = d.Fans[i].MinReadingRange
		each.MaxReadingRange = d.Fans[i].MaxReadingRange
		each.ReadingUnits = d.Fans[i].ReadingUnits
		// Redundancy is needed for Enclosure.
		fans = append(fans, each)
	}
	ret.Fans = fans

	return ret
}

func createBoardsModel(d *dto.GetBoardsResponse) *model.Board {
	ret := new(model.Board)
	createResourceModel(&d.Resource, &ret.Resource)
	createProductInfoModel(&d.ProductInfo, &ret.ProductInfo)
	ret.CardNo = d.CardNo
	ret.DeviceLocator = d.DeviceLocator
	ret.DeviceType = d.DeviceType
	ret.Location = d.Location
	ret.CPLDVersion = d.CPLDVersion
	ret.PCBVersion = d.PCBVersion
	ret.BoardName = d.BoardName
	ret.BoardID = d.BoardID
	ret.ManufactureDate = d.ManufactureDate
	return ret
}

func createNetworkAdapterModel(d *dto.GetNetworkAdapterResponse) *model.NetworkAdapter {
	ret := new(model.NetworkAdapter)
	createResourceModel(&d.Resource, &ret.Resource)
	createProductInfoModel(&d.ProductInfo, &ret.ProductInfo)
	return ret
}

func createControllerModel(d *dto.Controller) *model.Controller {
	ret := new(model.Controller)
	ret.FirmwarePackageVersion = d.FirmwarePackageVersion
	ret.ControllerCapabilities.NetworkPortCount = d.ControllerCapabilities.NetworkPortCount
	return ret
}

func createNetworkPortModel(d *dto.NetworkPort) *model.NetworkPort {
	ret := new(model.NetworkPort)
	createResourceModel(&d.Resource, &ret.Resource)
	ret.PhysicalPortNumber = d.PhysicalPortNumber
	ret.LinkStatus = d.LinkStatus
	ret.AssociatedNetworkAddresses = d.AssociatedNetworkAddresses
	return ret
}

func createDriveModel(d *dto.GetDriveResponse) *model.Drive {
	ret := new(model.Drive)
	createResourceModel(&d.Resource, &ret.Resource)
	createProductInfoModel(&d.ProductInfo, &ret.ProductInfo)
	ret.StatusIndicator = d.StatusIndicator
	ret.IndicatorLED = d.IndicatorLED
	ret.Revision = d.Revision
	ret.CapacityBytes = d.CapacityBytes
	ret.FailurePredicted = d.FailurePredicted
	ret.Protocol = d.Protocol
	ret.MediaType = d.MediaType
	ret.HotspareType = d.HotspareType
	ret.CapableSpeedGbs = d.CapableSpeedGbs
	ret.NegotiatedSpeedGbs = d.NegotiatedSpeedGbs
	ret.PredictedMediaLifeLeftPercent = d.PredictedMediaLifeLeftPercent
	for i := range d.Location {
		m := new(model.Location)
		createLocationModel(&d.Location[i], m)
		ret.Location = append(ret.Location, *m)
	}
	return ret
}

func createPCIeDeviceModel(device *dto.GetPCIeDeviceResponse, functions *[]dto.GetPCIeFunctionResponse) *model.PCIeDevice {
	ret := new(model.PCIeDevice)
	createResourceModel(&device.Resource, &ret.Resource)
	createProductInfoModel(&device.ProductInfo, &ret.ProductInfo)
	ret.DeviceType = device.DeviceType
	ret.FirmwareVersion = device.FirmwareVersion
	for i := range *functions {
		ret.PCIeFunctions = append(ret.PCIeFunctions, *createPCIeFunctionModel(&(*functions)[i]))
	}
	return ret
}

func createPCIeFunctionModel(d *dto.GetPCIeFunctionResponse) *model.PCIeFunction {
	ret := new(model.PCIeFunction)
	createResourceModel(&d.Resource, &ret.Resource)
	ret.DeviceClass = d.DeviceClass
	ret.DeviceID = d.DeviceID
	ret.VendorID = d.VendorID
	ret.SubsystemID = d.SubsystemID
	ret.SubsystemVendorID = d.SubsystemVendorID
	for i := range d.Links.EthernetInterfaces {
		ret.EthernetInterfaces = append(ret.EthernetInterfaces, d.Links.EthernetInterfaces[i].OdataId)
	}
	return ret
}
