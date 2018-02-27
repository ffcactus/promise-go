package redfish

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	. "promise/server/client/redfish/dto"
	"promise/server/object/model"

	"github.com/astaxie/beego"
)

const (
	COMMON_HEAD string = "application/json; charset=utf-8"
)

type RedfishClient struct {
	Client       *http.Client
	Address      string
	Username     string
	Password     string
	token        string
	UseBasicAuth bool
}

// Get a new instance of Redfish client.
func GetInstance(address string, username string, password string, useBasicAuth bool) *RedfishClient {
	return &RedfishClient{
		Client:       &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}},
		Address:      address,
		Username:     username,
		Password:     password,
		UseBasicAuth: useBasicAuth,
	}
}

func (this *RedfishClient) Support() bool {
	// Form the REST request.
	req, err := http.NewRequest(http.MethodGet, this.address("/redfish/v1"), nil)
	if err != nil {
		beego.Warning("NewRequest() failed, error = ", err)
		return false
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	if resp, err := this.Client.Do(req); err != nil {
		return false
	} else {
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusOK {
			return true
		} else {
			return false
		}
	}
}

// Get the protocal used by this client.
func (this *RedfishClient) GetProtocol() *string {
	return &model.RedfishV1
}

// Get server basic info.
// Just set parts of the properties.
func (this *RedfishClient) GetBasicInfo() (*model.ServerBasicInfo, error) {
	// First set the server type.
	var chassisCollection = Collection{}
	if err := this.getObject("/redfish/v1/chassis", &chassisCollection); err != nil {
		return nil, err
	}

	var systemCollection = Collection{}
	if err := this.getObject("/redfish/v1/systems", &systemCollection); err != nil {
		return nil, err
	}

	// If chassis count = 1, and computer system count = 1, it must be a blade or rack.
	if chassisCollection.Count == 1 && systemCollection.Count == 1 {
		// Get info from Computer system.
		var system = GetSystemResponse{}
		if err := this.getObject(systemCollection.Members[0].Id, &system); err != nil {
			return nil, err
		}
		ret := model.ServerBasicInfo{}
		ret.OriginURIs.System = &systemCollection.Members[0].Id
		ret.OriginURIs.Chassis = &chassisCollection.Members[0].Id
		ret.PhysicalUUID = system.PhysicalUUID
		ret.Protocol = model.RedfishV1
		// Get info from chassis.
		var chassis = GetChassisResponse{}
		if err := this.getObject(chassisCollection.Members[0].Id, &chassis); err != nil {
			return nil, err
		}
		if *chassis.ChassisType == "" {
			beego.Warning("GetBasicInfo() failed, failed to get chassis type.")
			return nil, errors.New("Failed to get server type.")
		}
		ret.Type = *chassis.ChassisType
		return &ret, nil
	} else {

	}
	return nil, nil
}

// Create Management account.
func (this *RedfishClient) CreateManagementAccount(username string, password string) error {
	requestBody := PostAccountRequest{
		UserName: username,
		Password: password,
		RoleId:   "Administrator",
	}
	if err := this.postObject("/redfish/v1/AccountService/Accounts", requestBody, nil); err != nil {
		return errors.New("Create management account failed.")
	} else {
		return nil
	}
}

// Get server's process info.
func (this *RedfishClient) GetProcessors(systemID string) ([]model.Processor, error) {
	collection := Collection{}
	if err := this.getObject(systemID+"/processors", &collection); err != nil {
		return nil, err
	}
	var ret []model.Processor
	for i, _ := range collection.Members {
		each := new(GetProcessorResponse)
		if err := this.getObject(collection.Members[i].Id, each); err != nil {
			return nil, err
		}
		ret = append(ret, *createProcessorModel(each))
	}
	return ret, nil
}

// Get server's memory info.
func (this *RedfishClient) GetMemory(systemID string) ([]model.Memory, error) {
	collection := Collection{}
	if err := this.getObject(systemID+"/memory", &collection); err != nil {
		return nil, err
	}

	var ret []model.Memory
	for i, _ := range collection.Members {
		each := new(GetMemoryResponse)
		if err := this.getObject(collection.Members[i].Id, each); err != nil {
			return nil, err
		}
		ret = append(ret, *createMemoryModel(each))
	}
	return ret, nil
}

// Get server's ethernet interface info.
func (this *RedfishClient) GetEthernetInterfaces(systemID string) ([]model.EthernetInterface, error) {
	collection := Collection{}
	if err := this.getObject(systemID+"/EthernetInterfaces", &collection); err != nil {
		return nil, err
	}
	var ret []model.EthernetInterface
	for i, _ := range collection.Members {
		eachEthernet := new(GetEthernetInterfaceResponse)
		if err := this.getObject(collection.Members[i].Id, eachEthernet); err != nil {
			return nil, err
		}
		// Get the VLANs
		vlanCollection := Collection{}
		vlanCollectionPageURI := systemID + "/EthernetInterfaces/" + *eachEthernet.Id + "/vlans"
		if err := this.getObject(vlanCollectionPageURI, &vlanCollection); err != nil {
			return nil, err
		}
		var vlans []model.VLanNetworkInterface
		for j, _ := range vlanCollection.Members {
			eachVlan := new(GetVLANResponse)
			if err := this.getObject(collection.Members[j].Id, eachVlan); err != nil {
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

func (this *RedfishClient) GetNetworkInterfaces(systemID string) ([]model.NetworkInterface, error) {
	collection := Collection{}
	if err := this.getObject(systemID+"/NetworkInterfaces", &collection); err != nil {
		return nil, err
	}
	var ret []model.NetworkInterface
	for i, _ := range collection.Members {
		networkInterface := new(GetNetworkInterfaceResponse)
		if err := this.getObject(collection.Members[i].Id, networkInterface); err != nil {
			return nil, err
		}

		ret = append(ret, *createNetworkInterfaceModel(networkInterface))
	}
	return ret, nil
}

func (this *RedfishClient) GetStorages(systemID string) ([]model.Storage, error) {
	collection := Collection{}
	if err := this.getObject(systemID+"/storages", &collection); err != nil {
		return nil, err
	}
	ret := []model.Storage{}
	for i, _ := range collection.Members {
		storage := new(GetStorageResponse)
		if err := this.getObject(collection.Members[i].Id, storage); err != nil {
			return nil, err
		}

		ret = append(ret, *createStorageModel(storage))
	}
	return ret, nil
}

func (this *RedfishClient) GetPower(chassisID string) (*model.Power, error) {
	power := new(GetPowerResponse)
	if err := this.getObject(chassisID+"/power", power); err != nil {
		return nil, err
	}
	model := createPowerModel(power)
	return model, nil
}

func (this *RedfishClient) GetThermal(chassisID string) (*model.Thermal, error) {
	thermal := new(GetThermalResponse)
	if err := this.getObject(chassisID+"/thermal", thermal); err != nil {
	}
	model := createThermalModel(thermal)
	return model, nil
}

func (this *RedfishClient) GetOemHuaweiBoards(chassisID string) ([]model.OemHuaweiBoard, error) {
	collection := Collection{}
	if err := this.getObject(chassisID+"/boards", &collection); err != nil {
		return nil, err
	}
	ret := []model.OemHuaweiBoard{}
	for i, _ := range collection.Members {
		board := new(GetOemHuaweiBoardResponse)
		if err := this.getObject(collection.Members[i].Id, board); err != nil {
			return nil, err
		}
		ret = append(ret, *createOemHuaweiBoardModel(board))
	}
	return ret, nil
}

func (this *RedfishClient) GetNetworkAdapters(chassisID string) ([]model.NetworkAdapter, error) {
	collection := Collection{}
	if err := this.getObject(chassisID+"/NetworkAdapters", &collection); err != nil {
		return nil, err
	}
	var ret []model.NetworkAdapter
	for i, _ := range collection.Members {
		resp := new(GetNetworkAdapterResponse)
		if err := this.getObject(collection.Members[i].Id, resp); err != nil {
			return nil, err
		}
		networkAdpter := createNetworkAdapterModel(resp)
		for j, _ := range resp.Controllers {
			eachController := createControllerModel(&resp.Controllers[j])
			portsResp := resp.Controllers[j].Links.NetworkPorts
			for k, _ := range portsResp {
				portPageURI := portsResp[k].OdataId
				portResp := new(NetworkPort)
				if err := this.getObject(portPageURI, portResp); err != nil {
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

func (this *RedfishClient) GetDrives(chassisID string) ([]model.Drive, error) {
	// Get the Drive links from chassis.
	chassis := new(GetChassisResponse)
	if err := this.getObject(chassisID, chassis); err != nil {
		return nil, err
	}
	ret := []model.Drive{}
	for i, _ := range chassis.Links.Drives {
		uri := chassis.Links.Drives[i].OdataId
		drive := new(GetDriveResponse)
		if err := this.getObject(uri, drive); err != nil {
			return nil, err
		}
		ret = append(ret, *createDriveModel(drive))
	}
	return ret, nil
}

func (this *RedfishClient) GetPCIeDevices(chassisID string) ([]model.PCIeDevice, error) {
	// Get the Drive links from chassis.
	chassis := new(GetChassisResponse)
	if err := this.getObject(chassisID, chassis); err != nil {
		return nil, err
	}
	ret := []model.PCIeDevice{}
	for i, _ := range chassis.Links.PCIeDevices {
		uri := chassis.Links.PCIeDevices[i].OdataId
		pcieDevice := new(GetPCIeDeviceResponse)
		if err := this.getObject(uri, pcieDevice); err != nil {
			return nil, err
		}
		pcieFunctions := new([]GetPCIeFunctionResponse)
		for j, _ := range pcieDevice.Links.PCIeFunctions {
			uri := pcieDevice.Links.PCIeFunctions[j].OdataId
			pcieFunction := new(GetPCIeFunctionResponse)
			if err := this.getObject(uri, pcieFunction); err != nil {
				return nil, err
			}
			*pcieFunctions = append(*pcieFunctions, *pcieFunction)
		}

		ret = append(ret, *createPCIeDeviceModel(pcieDevice, pcieFunctions))
	}
	return ret, nil
}

func (this *RedfishClient) GetNetworkPorts(uri string) ([]model.NetworkPort, error) {
	collection := Collection{}
	if err := this.getObject(uri, &collection); err != nil {
		return nil, err
	}
	var ret []model.NetworkPort
	for i, _ := range collection.Members {
		resp := new(NetworkPort)
		if err := this.getObject(collection.Members[i].Id, resp); err != nil {
			return nil, err
		}
		ret = append(ret, *createNetworkPortModel(resp))
	}
	return ret, nil
}

// The REST operation.
func (this *RedfishClient) rest(method string, uri string, body io.Reader) (resp *http.Response, err error) {
	url := this.address(uri)
	// Form the REST request.
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		beego.Warning("NewRequest() failed, method = ", method, ", URL = ", url, ", error = ", err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	// For basic auth, we pend credential first.
	if this.UseBasicAuth {
		// For basic auth.
		req.SetBasicAuth(this.Username, this.Password)
	}
	if resp, err := this.Client.Do(req); err != nil {
		beego.Warning("Do() failed, method = ", method, ", URL = ", url, ", error = ", err)
		return nil, err
	} else {
		return resp, err
	}
}

//
// Private methods.
//

// Get the complete address.
func (this *RedfishClient) address(uri string) string {
	var buf bytes.Buffer

	buf.WriteString("https://")
	buf.WriteString(this.Address)
	buf.WriteString(uri)

	return buf.String()
}

// Get and parse the object.
func (this *RedfishClient) getObject(uri string, o interface{}) error {
	if resp, err := this.rest(http.MethodGet, uri, nil); err != nil {
		return err
	} else {
		defer resp.Body.Close()
		if resp.Body == nil {
			beego.Warning("getObject() failed, resposne body is empty, URI = ", uri)
			return errors.New("Response body is empty.")
		}
		if resp.StatusCode != http.StatusOK {
			beego.Warning("getObject() failed, URI = ", uri, ", response code = ", resp.StatusCode)
			return fmt.Errorf("Response code = %d.", resp.StatusCode)
		}
		if err := json.NewDecoder(resp.Body).Decode(o); err != nil {
			beego.Warning("NewDecoder() failed, URI = ", uri, ", error = ", err)
			return err
		}
		return nil
	}
}

func (this *RedfishClient) postObject(uri string, req interface{}, resp interface{}) error {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(req)
	if resp, err := this.rest(http.MethodPost, uri, b); err != nil {
		return err
	} else {
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusCreated {
			// respBody, _ := ioutil.ReadAll(resp.Body)
			beego.Warning("postObject() failed, URI = ", uri, ", response code = ", resp.StatusCode)
			return fmt.Errorf("Response code = %d.", resp.StatusCode)
		}
		if resp.Body == nil {
			return nil
		}
		// Decode only when the client asked for.
		if resp != nil {
			if err := json.NewDecoder(resp.Body).Decode(resp); err != nil {
				beego.Warning("NewDecoder() failed, URI = ", uri, ", error = ", err)
				return err
			}
		}
		return nil
	}
}

func createProcessorModel(d *GetProcessorResponse) *model.Processor {
	ret := model.Processor{}
	createResourceModel(&d.Resource, &ret.Resource)
	createProductInfoModel(&d.ProductInfo, &ret.ProductInfo)
	ret.Socket = d.Socket
	ret.ProcessorType = d.ProcessorType
	ret.ProcessorArchitecture = d.ProcessorArchitecture
	ret.InstructionSet = d.InstructionSet
	ret.MaxSpeedMHz = d.MaxSpeedMHz
	ret.TotalCores = d.TotalCores
	if d.ProcessorID != nil {
		ret.ProcessorID = new(model.ProcessorID)
		ret.ProcessorID.VendorID = d.ProcessorID.VendorID
		ret.ProcessorID.MicrocodeInfo = d.ProcessorID.MicrocodeInfo
		ret.ProcessorID.Step = d.ProcessorID.Step
		ret.ProcessorID.IdentificationRegisters = d.ProcessorID.IdentificationRegisters
		ret.ProcessorID.EffectiveFamily = d.ProcessorID.EffectiveFamily
		ret.ProcessorID.EffectiveModel = d.ProcessorID.EffectiveModel
	}
	return &ret
}

func createResourceModel(d *Resource, m *model.Resource) {
	m.URI = d.OdataID
	m.Name = d.Name
	m.Description = d.Description
	m.OriginID = d.Id
	if d.Status != nil {
		m.PhysicalState = d.Status.State
		m.PhysicalHealth = d.Status.Health
	}
}

func createMemberModel(d *Member, m *model.Member) {
	m.URI = d.OdataID
	m.Name = d.Name
	m.Description = d.Description
	m.OriginMemberID = d.MemberId
	if d.Status != nil {
		m.PhysicalState = d.Status.State
		m.PhysicalHealth = d.Status.Health
	}
}

func createThresholdModel(d *Threshold, m *model.Threshold) {
	m.UpperThresholdNonCritical = d.UpperThresholdNonCritical
	m.UpperThresholdCritical = d.UpperThresholdCritical
	m.UpperThresholdFatal = d.UpperThresholdFatal
	m.LowerThresholdNonCritical = d.LowerThresholdNonCritical
	m.LowerThresholdCritical = d.LowerThresholdCritical
	m.LowerThresholdFatal = d.LowerThresholdFatal
}

func createProductInfoModel(d *ProductInfo, m *model.ProductInfo) {
	m.Model = d.Model
	m.Manufacturer = d.Manufacturer
	m.SKU = d.SKU
	m.SerialNumber = d.SerialNumber
	m.PartNumber = d.PartNumber
	m.SparePartNumber = d.SparePartNumber
	m.AssetTag = d.AssetTag
}

func createLocationModel(d *Location, m *model.Location) {
	m.Info = d.Info
	m.InfoFormat = d.InfoFormat
}

func createMemoryModel(d *GetMemoryResponse) *model.Memory {
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

func createEthernetInterfaceModel(d *GetEthernetInterfaceResponse) *model.EthernetInterface {
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
		for i, _ := range *d.IPv4Addresses {
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
		for i, _ := range *d.IPv6Addresses {
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

func createVLanModel(d *GetVLANResponse) *model.VLanNetworkInterface {
	ret := model.VLanNetworkInterface{}
	createResourceModel(&d.Resource, &ret.Resource)
	ret.VLANEnable = d.VLANEnable
	ret.VLANID = d.VLANID
	return &ret
}

func createNetworkInterfaceModel(d *GetNetworkInterfaceResponse) *model.NetworkInterface {
	ret := model.NetworkInterface{}
	createResourceModel(&d.Resource, &ret.Resource)
	ret.NetworkAdapterURI = d.Links.NetworkAdapter.OdataId
	return &ret
}

func createStorageControllerModel(d *StorageController) *model.StorageController {
	ret := model.StorageController{}
	createMemberModel(&d.Member, &ret.Member)
	createProductInfoModel(&d.ProductInfo, &ret.ProductInfo)
	ret.SpeedGbps = d.SpeedGbps
	ret.FirmwareVersion = d.FirmwareVersion
	ret.SupportedDeviceProtocols = d.SupportedDeviceProtocols
	return &ret
}

func createStorageModel(d *GetStorageResponse) *model.Storage {
	ret := model.Storage{}
	createResourceModel(&d.Resource, &ret.Resource)
	for i, _ := range d.Drives {
		ret.DriveURIs = append(ret.DriveURIs, d.Drives[i].OdataId)
	}
	for i, _ := range d.StorageControllers {
		ret.StorageControllers = append(ret.StorageControllers, *createStorageControllerModel(&d.StorageControllers[i]))
	}
	return &ret
}

func createPowerModel(d *GetPowerResponse) *model.Power {
	dto := *d
	ret := model.Power{}
	createResourceModel(&d.Resource, &ret.Resource)
	// PowerControl
	powerControl := []model.PowerControl{}
	for i, _ := range *dto.PowerControl {
		eachModel := model.PowerControl{}
		eachDto := (*dto.PowerControl)[i]
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
	ret.PowerControl = &powerControl

	// Voltages
	voltages := []model.Voltage{}
	for i, _ := range *dto.Voltages {
		eachModel := model.Voltage{}
		eachDto := (*dto.Voltages)[i]
		createResourceModel(&eachDto.Resource, &eachModel.Resource)
		createThresholdModel(&eachDto.Threshold, &eachModel.Threshold)
		eachModel.SensorNumber = eachDto.SensorNumber
		eachModel.ReadingVolts = eachDto.ReadingVolts
		eachModel.MinReadingRange = eachDto.MinReadingRange
		eachModel.MaxReadingRange = eachDto.MaxReadingRange
		eachModel.PhysicalContext = eachDto.PhysicalContext
		voltages = append(voltages, eachModel)
	}
	ret.Voltages = &voltages

	// PowerSupplies
	powerSupplies := []model.PowerSupply{}
	for i, _ := range *dto.PowerSupplies {
		eachModel := model.PowerSupply{}
		eachDto := (*dto.PowerSupplies)[i]
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
	ret.PowerSupplies = &powerSupplies

	// Redundancy
	redundancy := []model.Redundancy{}
	for i, _ := range *dto.Redundancy {
		eachModel := model.Redundancy{}
		eachDto := (*dto.Redundancy)[i]
		createResourceModel(&eachDto.Resource, &eachModel.Resource)
		eachModel.Mode = eachDto.Mode
		eachModel.MaxNumSupported = eachDto.MaxNumSupported
		eachModel.MinNumNeeded = eachDto.MinNumNeeded
		eachModel.RedundancyEnabled = eachDto.RedundancyEnabled
		// only name is needed in the name of redundancy set.
		redundancySet := []string{}
		for j, _ := range *eachDto.RedundancySet {
			for k, _ := range *dto.PowerSupplies {
				redundancyOdataId := (*eachDto.RedundancySet)[j].OdataId
				powerSupply := (*dto.PowerSupplies)[k]
				if *powerSupply.OdataID == redundancyOdataId {
					redundancySet = append(redundancySet, *powerSupply.Name)
				}
			}
		}
		eachModel.RedundancySet = &redundancySet
		redundancy = append(redundancy, eachModel)
	}
	ret.Redundancy = &redundancy
	return &ret
}

func createThermalModel(d *GetThermalResponse) *model.Thermal {
	ret := new(model.Thermal)
	createResourceModel(&d.Resource, &ret.Resource)
	temperatures := []model.Temperature{}
	for i, _ := range d.Temperatures {
		each := model.Temperature{}
		createMemberModel(&d.Temperatures[i].Member, &each.Member)
		createThresholdModel(&d.Temperatures[i].Threshold, &each.Threshold)
		each.SensorNumber = d.Temperatures[i].SensorNumber
		each.ReadingCelsius = d.Temperatures[i].ReadingCelsius
		each.MinReadingRangeTemp = d.Temperatures[i].MinReadingRangeTemp
		each.MaxReadingRangeTemp = d.Temperatures[i].MaxReadingRangeTemp
		temperatures = append(temperatures, each)
	}
	ret.Temperatures = temperatures

	fans := []model.Fan{}
	for i, _ := range d.Fans {
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

func createOemHuaweiBoardModel(d *GetOemHuaweiBoardResponse) *model.OemHuaweiBoard {
	ret := new(model.OemHuaweiBoard)
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

func createNetworkAdapterModel(d *GetNetworkAdapterResponse) *model.NetworkAdapter {
	ret := new(model.NetworkAdapter)
	createResourceModel(&d.Resource, &ret.Resource)
	createProductInfoModel(&d.ProductInfo, &ret.ProductInfo)
	return ret
}

func createControllerModel(d *Controller) *model.Controller {
	ret := new(model.Controller)
	ret.FirmwarePackageVersion = d.FirmwarePackageVersion
	ret.ControllerCapabilities.NetworkPortCount = d.ControllerCapabilities.NetworkPortCount
	return ret
}

func createNetworkPortModel(d *NetworkPort) *model.NetworkPort {
	ret := new(model.NetworkPort)
	createResourceModel(&d.Resource, &ret.Resource)
	ret.PhysicalPortNumber = d.PhysicalPortNumber
	ret.LinkStatus = d.LinkStatus
	ret.AssociatedNetworkAddresses = d.AssociatedNetworkAddresses
	return ret
}

func createDriveModel(d *GetDriveResponse) *model.Drive {
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
	for i, _ := range d.Location {
		m := new(model.Location)
		createLocationModel(&d.Location[i], m)
		ret.Location = append(ret.Location, *m)
	}
	return ret
}

func createPCIeDeviceModel(device *GetPCIeDeviceResponse, functions *[]GetPCIeFunctionResponse) *model.PCIeDevice {
	ret := new(model.PCIeDevice)
	createResourceModel(&device.Resource, &ret.Resource)
	createProductInfoModel(&device.ProductInfo, &ret.ProductInfo)
	ret.DeviceType = device.DeviceType
	ret.FirmwareVersion = device.FirmwareVersion
	for i, _ := range *functions {
		ret.PCIeFunctions = append(ret.PCIeFunctions, *createPCIeFunctionModel(&(*functions)[i]))
	}
	return ret
}

func createPCIeFunctionModel(d *GetPCIeFunctionResponse) *model.PCIeFunction {
	ret := new(model.PCIeFunction)
	createResourceModel(&d.Resource, &ret.Resource)
	ret.DeviceClass = d.DeviceClass
	ret.DeviceID = d.DeviceID
	ret.VendorID = d.VendorID
	ret.SubsystemID = d.SubsystemID
	ret.SubsystemVendorID = d.SubsystemVendorID
	for i, _ := range d.Links.EthernetInterfaces {
		ret.EthernetInterfaces = append(ret.EthernetInterfaces, d.Links.EthernetInterfaces[i].OdataId)
	}
	return ret
}
