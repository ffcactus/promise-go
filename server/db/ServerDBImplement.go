package db

import (
	"errors"
	"fmt"
	"promise/common/app"
	commonDB "promise/common/db"
	commonUtil "promise/common/util"
	"promise/server/object/entity"
	"promise/server/object/model"
	"promise/server/util"
	"github.com/astaxie/beego"
	"github.com/google/uuid"
)

// ServerDBImplement The DB implementation.
type ServerDBImplement struct {
}

// GetDBInstance Get a DB implementation instance.
func GetDBInstance() ServerDBInterface {
	return new(ServerDBImplement)
}

// IsServerExist Check the existance of the server, if found, return it.
func (i *ServerDBImplement) IsServerExist(s *model.Server) (bool, *model.Server) {
	if s == nil {
		return false, nil
	}
	c := commonDB.GetConnection()
	server := new(entity.Server)
	c.Where("Physical_UUID = ?", s.PhysicalUUID).First(server)
	if s.PhysicalUUID == server.PhysicalUUID {
		return true, createServerModel(server)
	}
	return false, nil
}

// PostServer Post server from server basic info.
func (i *ServerDBImplement) PostServer(s *model.Server) (*model.Server, error) {
	c := commonDB.GetConnection()
	if exist, _ := i.IsServerExist(s); exist {
		// TODO get the server to return.
		beego.Trace("PostServer(), server exist.")
		return nil, errors.New("server exist")
	}
	var server = createServerEntityFromServer(s)
	// Generate the UUID.
	server.ID = uuid.New().String()
	c.Create(server)
	return createServerModel(server), nil
}

// GetServer Get server by ID.
func (i *ServerDBImplement) GetServer(ID string) *model.Server {
	c := commonDB.GetConnection()
	var s = new(entity.Server)
	c.Where("ID = ?", ID).First(s)
	if s.ID != ID {
		return nil
	}
	return createServerModel(s)
}

// GetServerCollection Get server collection by start and count.
func (i *ServerDBImplement) GetServerCollection(start int, count int) (*model.ServerCollection, error) {
	var (
		total  int
		server []entity.Server
		ret    = new(model.ServerCollection)
	)

	c := commonDB.GetConnection()
	c.Table("server").Count(total)
	c.Order("Name asc").Limit(count).Offset(start).Select([]string{"ID", "Name", "State", "Health"}).Find(&server)
	ret.Start = start
	ret.Count = len(server)
	ret.Total = total
	for i := range server {
		ret.Members = append(ret.Members, model.ServerMember{
			URI:    toServerURI(server[i].ID),
			Name:   server[i].Name,
			State:  server[i].State,
			Health: server[i].Health,
		})
	}
	return ret, nil
}

// GetServerFull Get server full info.
func (i *ServerDBImplement) GetServerFull(ID string) *model.Server {
	c := commonDB.GetConnection()
	var s = new(entity.Server)
	c.Where("ID = ?", ID).First(s)
	if s.ID != ID {
		return nil
	}
	c.Where("ID = ?", ID).
		Preload("Processors").
		Preload("Memory").
		Preload("EthernetInterfaces").
		Preload("EthernetInterfaces.IPv4Addresses").
		Preload("EthernetInterfaces.IPv6Addresses").
		Preload("EthernetInterfaces.VLANs").
		Preload("NetworkInterfaces").
		Preload("Storages").
		Preload("Storages.StorageControllers").
		Preload("Power").
		Preload("Power.PowerControl").
		Preload("Power.Voltages").
		Preload("Power.PowerSupplies").
		Preload("Power.Redundancy").
		Preload("Thermal").
		Preload("Thermal.Temperatures").
		Preload("Thermal.Fans").
		Preload("OemHuaweiBoards").
		Preload("NetworkAdapters").
		Preload("NetworkAdapters.Controllers").
		Preload("NetworkAdapters.Controllers.NetworkPorts").
		Preload("Drives").
		Preload("Drives.Location").
		Preload("Drives.Location.PostalAddress").
		Preload("Drives.Location.Placement").
		Preload("PCIeDevices").
		Preload("PCIeDevices.PCIeFunctions").
		First(s)
	return createServerModel(s)
}

// FindServerStateAdded Find the first server with state "Added" and return the server ID.
func (i *ServerDBImplement) FindServerStateAdded() string {
	c := commonDB.GetConnection()
	var s = new(entity.Server)
	if notFound := c.Where("State = ?", util.ServerStateAdded).First(s).RecordNotFound(); notFound {
		return ""
	}
	return s.ID
}

// GetAndLockServer Get and lock the server.
func (i *ServerDBImplement) GetAndLockServer(ID string) (bool, *model.Server) {
	c := commonDB.GetConnection()
	// Transaction start.
	tx := c.Begin()
	var s = new(entity.Server)
	tx.Where("ID = ?", ID).First(s)
	if s.ID != ID {
		// Can't find server, rollback.
		beego.Info("GetAndLockServer() failed, server not exist.")
		tx.Rollback()
		return false, nil
	}
	if !util.ServerLockable(s.State) {
		// Server not ready, rollback.
		tx.Rollback()
		return false, createServerModel(s)
	}
	// Change the state.
	tx.Model(s).UpdateColumn("State", util.ServerStateLocked)
	// Commit.
	tx.Commit()
	errs := c.GetErrors()
	if len(errs) != 0 {
		beego.Warning("GetAndLockServer() failed, server ID = ", ID, ", error = ", errs)
		return false, nil
	}
	return true, createServerModel(s)
}

// SetServerState Set server state.
func (i *ServerDBImplement) SetServerState(ID string, state string) bool {
	c := commonDB.GetConnection()
	var s = new(entity.Server)
	if c.Where("ID = ?", ID).First(s).RecordNotFound() {
		return false
	}
	if err := c.Model(s).UpdateColumn("State", state).Error; err != nil {
		beego.Warning("SetServerState() failed, server ID = ", ID, ", error = ", err)
		return false
	}
	return true
}

// SetServerHealth Set server health.
func (i *ServerDBImplement) SetServerHealth(ID string, health string) bool {
	c := commonDB.GetConnection()
	var s = new(entity.Server)
	if c.Where("ID = ?", ID).First(s).RecordNotFound() {
		return false
	}
	if err := c.Model(s).UpdateColumn("Health", health).Error; err != nil {
		beego.Warning("SetServerHealth() failed, server ID = ", ID, ", error = ", err)
		return false
	}
	return true
}

// SetServerTask Set server task.
func (i *ServerDBImplement) SetServerTask(ID string, taskURI string) bool {
	c := commonDB.GetConnection()
	var s = new(entity.Server)
	if c.Where("ID = ?", ID).First(s).RecordNotFound() {
		return false
	}
	if err := c.Model(s).UpdateColumn("CurrentTask", taskURI).Error; err != nil {
		beego.Warning("SetServerTask() failed, server ID = ", ID, ", error = ", err)
	}
	return true
}

// UpdateProcessors Update processors info.
func (i *ServerDBImplement) UpdateProcessors(ID string, processors []model.Processor) error {
	server := new(entity.Server)

	c := commonDB.GetConnection()
	notFound := c.Where("ID = ?", ID).
		Preload("Processors").
		First(server).
		RecordNotFound()
	if notFound {
		return fmt.Errorf("Can't find server %s", ID)
	}
	for i := range server.Processors {
		c.Delete(server.Processors[i])
	}
	server.Processors = []entity.Processor{}
	for i := range processors {
		server.Processors = append(server.Processors, *createProcessor(&processors[i]))
	}
	if err := c.Save(server).Error; err != nil {
		beego.Warning("UpdateProcessors() failed, server ID = ", ID, ", error = ", err)
		return err
	}
	return nil
}

// UpdateMemory Update server memory
func (i *ServerDBImplement) UpdateMemory(ID string, memory []model.Memory) error {
	server := new(entity.Server)

	c := commonDB.GetConnection()
	notFound := c.Where("ID = ?", ID).
		Preload("Memory").
		First(server).
		RecordNotFound()
	if notFound {
		return fmt.Errorf("Can't find server %s", ID)
	}
	for i := range server.Memory {
		c.Delete(server.Memory[i])
	}
	server.Memory = []entity.Memory{}
	for i := range memory {
		server.Memory = append(server.Memory, *createMemory(&memory[i]))
	}
	if err := c.Save(server).Error; err != nil {
		beego.Warning("UpdateMemory() failed, server ID = ", ID, ", error = ", err)
		return err
	}
	return nil
}

// UpdateEthernetInterfaces Update server ethernet interface.
// Each ethernet interface has many IPv4, IPv6 and vLAN, so it's very hard to check if 2 ethernet interface
// are the same. So we just remove all of them and recreate them.
func (i *ServerDBImplement) UpdateEthernetInterfaces(ID string, ethernet []model.EthernetInterface) error {
	server := new(entity.Server)
	c := commonDB.GetConnection()
	// Get Server and it's ethernet interface.
	notFound := c.Where("ID = ?", ID).
		Preload("EthernetInterfaces").
		Preload("EthernetInterfaces.IPv4Addresses").
		Preload("EthernetInterfaces.IPv6Addresses").
		Preload("EthernetInterfaces.VLANs").
		First(server).
		RecordNotFound()
	if notFound {
		return fmt.Errorf("Can't find server %s", ID)
	}
	// Delete all ethernet interface.
	for i := range server.EthernetInterfaces {
		for k := range server.EthernetInterfaces[i].IPv4Addresses {
			c.Delete(&(server.EthernetInterfaces[i].IPv4Addresses)[k])
		}
		for k := range server.EthernetInterfaces[i].IPv6Addresses {
			c.Delete(&(server.EthernetInterfaces[i].IPv6Addresses)[k])
		}
		for k := range server.EthernetInterfaces[i].VLANs {
			c.Delete(&(server.EthernetInterfaces[i].VLANs)[k])
		}
		c.Delete(server.EthernetInterfaces[i])
	}
	// Regenerate ethernet interface.
	server.EthernetInterfaces = []entity.EthernetInterface{}
	for i := range ethernet {
		server.EthernetInterfaces = append(server.EthernetInterfaces, *createEthernetInterface(&ethernet[i]))
	}
	if err := c.Save(server).Error; err != nil {
		beego.Warning("UpdateEthernetInterfaces() failed, server ID = ", ID, ", error = ", err)
		return err
	}
	return nil
}

// UpdateNetworkInterfaces Update network interfaces.
func (i *ServerDBImplement) UpdateNetworkInterfaces(ID string, networkInterface []model.NetworkInterface) error {
	server := new(entity.Server)
	c := commonDB.GetConnection()
	// Get resources.
	notFound := c.Where("ID = ?", ID).
		Preload("NetworkInterfaces").
		First(server).
		RecordNotFound()
	if notFound {
		return fmt.Errorf("Can't find server %s", ID)
	}
	// Delete them.
	for i := range server.NetworkInterfaces {
		c.Delete(server.NetworkInterfaces[i])
	}
	networkInterfacesE := []entity.NetworkInterface{}
	for i := range networkInterface {
		networkInterfacesE = append(networkInterfacesE, *createNetworkInterface(&networkInterface[i]))
	}
	server.NetworkInterfaces = networkInterfacesE
	if err := c.Save(server).Error; err != nil {
		beego.Warning("UpdateNetworkInterfaces() failed, server ID = ", ID, ", error = ", err)
		return err
	}
	return nil
}

// UpdateStorages Update storages.
func (i *ServerDBImplement) UpdateStorages(ID string, storages []model.Storage) error {
	server := new(entity.Server)
	c := commonDB.GetConnection()
	notFound := c.Where("ID = ?", ID).
		Preload("Storages").
		Preload("Storages.StorageControllers").
		First(server).
		RecordNotFound()
	if notFound {
		return fmt.Errorf("Can't find server %s", ID)
	}
	for i := range server.Storages {
		for j := range server.Storages[i].StorageControllers {
			c.Delete(&(server.Storages[i].StorageControllers)[j])
		}
		c.Delete(server.Storages[i])
	}
	storagesE := []entity.Storage{}
	for i := range storages {
		storagesE = append(storagesE, *createStorage(&storages[i]))
	}
	server.Storages = storagesE
	if err := c.Save(server).Error; err != nil {
		beego.Warning("UpdateStorages() failed, server ID = ", ID, ", error = ", err)
		return err
	}
	return nil
}

// UpdatePower Update power
func (i *ServerDBImplement) UpdatePower(ID string, power *model.Power) error {
	server := new(entity.Server)
	c := commonDB.GetConnection()
	notFound := c.Where("ID = ?", ID).
		Preload("Power").
		Preload("Power.PowerControl").
		Preload("Power.Voltages").
		Preload("Power.PowerSupplies").
		Preload("Power.Redundancy").
		First(server).
		RecordNotFound()
	if notFound {
		return fmt.Errorf("Can't find server %s", ID)
	}
	for i := range server.Power.PowerControl {
		c.Delete(server.Power.PowerControl[i])
	}
	for i := range server.Power.Voltages {
		c.Delete(server.Power.Voltages[i])
	}
	for i := range server.Power.PowerSupplies {
		c.Delete(server.Power.PowerSupplies[i])
	}
	for i := range server.Power.Redundancy {
		c.Delete(server.Power.Redundancy[i])
	}
	c.Delete(server.Power)
	server.Power = *createPower(power)
	if err := c.Save(server).Error; err != nil {
		beego.Warning("UpdatePower() failed, server ID = ", ID, ", error = ", err)
		return err
	}
	return nil
}

// UpdateThermal Update thermal
func (i *ServerDBImplement) UpdateThermal(ID string, thermal *model.Thermal) error {
	server := new(entity.Server)
	c := commonDB.GetConnection()
	notFound := c.Where("ID = ?", ID).
		Preload("Thermal").
		Preload("Thermal.Temperatures").
		Preload("Thermal.Fans").
		First(server).
		RecordNotFound()
	if notFound {
		return fmt.Errorf("Can't find server %s", ID)
	}
	for i := range server.Thermal.Temperatures {
		c.Delete(server.Thermal.Temperatures[i])
	}
	for i := range server.Thermal.Fans {
		c.Delete(server.Thermal.Fans[i])
	}
	c.Delete(server.Thermal)
	server.Thermal = *createThermal(thermal)
	if err := c.Save(server).Error; err != nil {
		beego.Warning("UpdateThermal() failed, server ID = ", ID, ", error = ", err)
		return err
	}
	return nil
}

// UpdateOemHuaweiBoards Update OEM Huawei boards.
func (i *ServerDBImplement) UpdateOemHuaweiBoards(ID string, boards []model.OemHuaweiBoard) error {
	c := commonDB.GetConnection()
	var server = new(entity.Server)
	notFound := c.Where("ID = ?", ID).
		Preload("OemHuaweiBoards").
		First(server).
		RecordNotFound()
	if notFound {
		return fmt.Errorf("Can't find server %s", ID)
	}
	for i := range server.OemHuaweiBoards {
		c.Delete(server.OemHuaweiBoards[i])
	}
	boardsE := []entity.OemHuaweiBoard{}
	for i := range boards {
		boardsE = append(boardsE, *createOemHuaweiBoard(&boards[i]))
	}
	server.OemHuaweiBoards = boardsE
	if err := c.Save(server).Error; err != nil {
		beego.Warning("UpdateOemHuaweiBoards() failed, server ID = ", ID, ", error = ", err)
		return err
	}
	return nil
}

// UpdateNetworkAdapters Update network adapters.
func (i *ServerDBImplement) UpdateNetworkAdapters(ID string, networkAdapters []model.NetworkAdapter) error {
	c := commonDB.GetConnection()
	var server = new(entity.Server)
	notFound := c.Where("ID = ?", ID).
		Preload("NetworkAdapters").
		Preload("NetworkAdapters.Controllers").
		Preload("NetworkAdapters.Controllers.NetworkPorts").
		First(server).
		RecordNotFound()
	if notFound {
		return fmt.Errorf("Can't find server %s", ID)
	}
	for i := range server.NetworkAdapters {
		adapter := server.NetworkAdapters[i]
		for j := range adapter.Controllers {
			controller := adapter.Controllers[j]
			for k := range controller.NetworkPorts {
				c.Delete(&controller.NetworkPorts[k])
			}
			c.Delete(&controller)
		}
		c.Delete(&adapter)
	}
	networkAdaptersE := []entity.NetworkAdapter{}
	for i := range networkAdapters {
		networkAdaptersE = append(networkAdaptersE, *createNetworkAdapter(&networkAdapters[i]))
	}
	server.NetworkAdapters = networkAdaptersE
	if err := c.Save(server).Error; err != nil {
		beego.Warning("UpdateNetworkAdapters() failed, server ID = ", ID, ", error = ", err)
		return err
	}
	return nil
}

// UpdateDrives Update drives.
func (i *ServerDBImplement) UpdateDrives(ID string, drives []model.Drive) error {
	c := commonDB.GetConnection()
	var server = new(entity.Server)
	notFound := c.Where("ID = ?", ID).
		Preload("Drives").
		Preload("Drives.Location").
		Preload("Drives.Location.PostalAddress").
		Preload("Drives.Location.Placement").
		First(server).
		RecordNotFound()
	if notFound {
		return fmt.Errorf("Can't find server %s", ID)
	}
	for i := range server.Drives {
		drive := server.Drives[i]
		for j := range drive.Location {
			if drive.Location[j].PostalAddress != nil {
				c.Delete(&drive.Location[j].PostalAddress)
			}
			if drive.Location[j].Placement != nil {
				c.Delete(&drive.Location[j].Placement)
			}
			c.Delete(&drive.Location[j])
		}
		c.Delete(&drive)
	}
	drivesE := []entity.Drive{}
	for i := range drives {
		drivesE = append(drivesE, *createDrive(&drives[i]))
	}
	server.Drives = drivesE
	if err := c.Save(server).Error; err != nil {
		beego.Warning("UpdateDrives() failed, server ID = ", ID, ", error = ", err)
		return err
	}
	return nil
}

// UpdatePCIeDevices Updagte PCIe devices.
func (i *ServerDBImplement) UpdatePCIeDevices(ID string, pcieDevices []model.PCIeDevice) error {
	c := commonDB.GetConnection()
	var server = new(entity.Server)
	notFound := c.Where("ID = ?", ID).
		Preload("PCIeDevices").
		Preload("PCIeDevices.PCIeFunctions").
		First(server).
		RecordNotFound()
	if notFound {
		return fmt.Errorf("Can't find server %s", ID)
	}
	for i := range server.PCIeDevices {
		pcieDevice := server.PCIeDevices[i]
		for j := range pcieDevice.PCIeFunctions {
			c.Delete(&pcieDevice.PCIeFunctions[j])
		}
		c.Delete(&pcieDevice)
	}
	pcieDevicesE := new([]entity.PCIeDevice)
	for i := range pcieDevices {
		*pcieDevicesE = append(*pcieDevicesE, *createPCIeDevice(&pcieDevices[i]))
	}
	server.PCIeDevices = *pcieDevicesE
	if err := c.Save(server).Error; err != nil {
		beego.Warning("UpdatePCIeDevices() failed, server ID = ", ID, ", error = ", err)
		return err
	}
	return nil
}

func toServerURI(ID string) string {
	s := app.RootURL + "/server" + ID
	return s
}

func createServerModel(e *entity.Server) *model.Server {
	m := model.Server{}
	m.ID = e.ID
	m.OriginURIs.Chassis = e.OriginURIsChassis
	m.OriginURIs.System = e.OriginURIsSystem
	m.PhysicalUUID = e.PhysicalUUID
	m.Name = e.Name
	m.Description = e.Description
	m.URI = toServerURI(e.ID)
	m.Address = e.Address
	m.Type = e.Type
	m.Protocol = e.Protocol
	m.Credential = e.Credential
	m.State = e.State
	m.Health = e.Health
	m.CurrentTask = e.CurrentTask
	// ComputerSystem.Processors
	processors := []model.Processor{}
	for i := range e.Processors {
		processors = append(processors, *createProcessorModel(&e.Processors[i]))
	}
	m.ComputerSystem.Processors = processors
	// ComputerSystem.Memory
	memory := []model.Memory{}
	for i := range e.Memory {
		memory = append(memory, *createMemoryModel(&e.Memory[i]))
	}
	m.ComputerSystem.Memory = memory
	// ComputerSystem.EthernetInterfaces
	ethernetInterfaces := []model.EthernetInterface{}
	for i := range e.EthernetInterfaces {
		ethernetInterfaces = append(ethernetInterfaces, *createEthernetInterfaceModel(&e.EthernetInterfaces[i]))
	}
	m.ComputerSystem.EthernetInterfaces = ethernetInterfaces
	// ComputerSystem.NetworkInterfaces
	networkInterfaces := []model.NetworkInterface{}
	for i := range e.NetworkInterfaces {
		networkInterfaces = append(networkInterfaces, *createNetworkInterfaceModel(&e.NetworkInterfaces[i]))
	}
	m.ComputerSystem.NetworkInterfaces = networkInterfaces
	// ComputerSystem.Storages
	storages := []model.Storage{}
	for i := range e.Storages {
		storages = append(storages, *createStorageModel(&e.Storages[i]))
	}
	m.ComputerSystem.Storages = storages
	// Chassis.Power
	createResourceModel(&e.Power.EmbeddedResource, &m.Chassis.Power.Resource)
	powerControl := []model.PowerControl{}
	for i := range e.Power.PowerControl {
		powerControl = append(powerControl, *createPowerControlModel(&e.Power.PowerControl[i]))
	}
	m.Chassis.Power.PowerControl = &powerControl

	voltages := []model.Voltage{}
	for i := range e.Power.Voltages {
		voltages = append(voltages, *createVoltageModel(&e.Power.Voltages[i]))
	}
	m.Chassis.Power.Voltages = &voltages

	powerSupplies := []model.PowerSupply{}
	for i := range e.Power.PowerSupplies {
		powerSupplies = append(powerSupplies, *createPowerSupplyModel(&e.Power.PowerSupplies[i]))
	}
	m.Chassis.Power.PowerSupplies = &powerSupplies

	redundancy := []model.Redundancy{}
	for i := range e.Power.Redundancy {
		redundancy = append(redundancy, *createRedundancyModel(&e.Power.Redundancy[i]))
	}
	m.Chassis.Power.Redundancy = &redundancy
	// Chassis.Thermal
	createResourceModel(&e.Thermal.EmbeddedResource, &m.Chassis.Thermal.Resource)
	temperatures := []model.Temperature{}
	for i := range e.Thermal.Temperatures {
		temperatures = append(temperatures, *createTemperatureModel(&e.Thermal.Temperatures[i]))
	}
	m.Chassis.Thermal.Temperatures = temperatures
	fans := []model.Fan{}
	for i := range e.Thermal.Fans {
		fans = append(fans, *createFanModel(&e.Thermal.Fans[i]))
	}
	m.Chassis.Thermal.Fans = fans
	// Chassis.OemHuaweiBoards
	boards := []model.OemHuaweiBoard{}
	for i := range e.OemHuaweiBoards {
		boards = append(boards, *createOemHuaweiBoardModel(&e.OemHuaweiBoards[i]))
	}
	m.Chassis.OemHuaweiBoards = boards
	// Chassis.NetworkAdapters
	networkAdapters := []model.NetworkAdapter{}
	for i := range e.NetworkAdapters {
		networkAdapters = append(networkAdapters, *createNetworkAdapterModel(&e.NetworkAdapters[i]))
	}
	m.Chassis.NetworkAdapters = networkAdapters
	// Chassis.Drives
	drives := []model.Drive{}
	for i := range e.Drives {
		drives = append(drives, *createDriveModel(&e.Drives[i]))
	}
	m.Chassis.Drives = drives
	// Chassis.PCIeDevices
	pcieDevices := []model.PCIeDevice{}
	for i := range e.PCIeDevices {
		pcieDevices = append(pcieDevices, *createPCIeDeviceModel(&e.PCIeDevices[i]))
	}
	m.Chassis.PCIeDevices = pcieDevices
	return &m
}

func createResourceModel(e *entity.EmbeddedResource, m *model.Resource) {
	m.URI = e.URI
	m.OriginID = e.OriginID
	m.Name = e.Name
	m.Description = e.Description
	m.State = e.State
	m.Health = e.Health
	m.PhysicalState = e.PhysicalState
	m.PhysicalHealth = e.PhysicalHealth
}

func createMemberModel(e *entity.Member, m *model.Member) {
	m.URI = e.URI
	m.OriginMemberID = e.OriginMemberID
	m.MemberID = *e.OriginMemberID
	m.Name = e.Name
	m.Description = e.Description
	m.State = e.State
	m.Health = e.Health
	m.PhysicalState = e.PhysicalState
	m.PhysicalHealth = e.PhysicalHealth
}

func createProductInfoModel(e *entity.ProductInfo, m *model.ProductInfo) {
	m.Model = e.Model
	m.Manufacturer = e.Manufacturer
	m.SKU = e.SKU
	m.SerialNumber = e.SerialNumber
	m.SparePartNumber = e.SparePartNumber
	m.PartNumber = e.PartNumber
	m.AssetTag = e.AssetTag
}

func createThresholdModel(e *entity.Threshold, m *model.Threshold) {
	m.UpperThresholdNonCritical = e.UpperThresholdNonCritical
	m.UpperThresholdCritical = e.UpperThresholdCritical
	m.UpperThresholdFatal = e.UpperThresholdFatal
	m.LowerThresholdNonCritical = e.LowerThresholdNonCritical
	m.LowerThresholdCritical = e.LowerThresholdCritical
	m.LowerThresholdFatal = e.LowerThresholdFatal
}

func createLocationModel(e *entity.Location, m *model.Location) {
	m.Info = e.Info
	m.InfoFormat = e.InfoFormat
	if e.PostalAddress != nil {
		postalAddress := new(model.PostalAddress)
		postalAddress.Country = e.PostalAddress.Country
		postalAddress.Territory = e.PostalAddress.Territory
		postalAddress.District = e.PostalAddress.District
		postalAddress.City = e.PostalAddress.City
		postalAddress.Division = e.PostalAddress.Division
		postalAddress.Neighborhood = e.PostalAddress.Neighborhood
		postalAddress.LeadingStreetDirection = e.PostalAddress.LeadingStreetDirection
		postalAddress.Street = e.PostalAddress.Street
		postalAddress.TrailingStreetSuffix = e.PostalAddress.TrailingStreetSuffix
		postalAddress.StreetSuffix = e.PostalAddress.StreetSuffix
		postalAddress.HouseNumber = e.PostalAddress.HouseNumber
		postalAddress.HouseNumberSuffix = e.PostalAddress.HouseNumberSuffix
		postalAddress.Landmark = e.PostalAddress.Landmark
		postalAddress.Location = e.PostalAddress.Location
		postalAddress.Floor = e.PostalAddress.Floor
		postalAddress.Name = e.PostalAddress.Name
		postalAddress.PostalCode = e.PostalAddress.PostalCode
		postalAddress.Building = e.PostalAddress.Building
		postalAddress.Unit = e.PostalAddress.Unit
		postalAddress.Room = e.PostalAddress.Room
		postalAddress.Seat = e.PostalAddress.Seat
		postalAddress.PlaceType = e.PostalAddress.PlaceType
		postalAddress.Community = e.PostalAddress.Community
		postalAddress.POBox = e.PostalAddress.POBox
		postalAddress.AdditionalCode = e.PostalAddress.AdditionalCode
		postalAddress.Road = e.PostalAddress.Road
		postalAddress.RoadSection = e.PostalAddress.RoadSection
		postalAddress.RoadBranch = e.PostalAddress.RoadBranch
		postalAddress.RoadSubBranch = e.PostalAddress.RoadSubBranch
		postalAddress.RoadPreModifier = e.PostalAddress.RoadPreModifier
		postalAddress.RoadPostModifier = e.PostalAddress.RoadPostModifier
		postalAddress.GPSCoords = e.PostalAddress.GPSCoords

		m.PostalAddress = postalAddress
	}
	if e.Placement != nil {
		placement := new(model.Placement)
		placement.Row = e.Placement.Row
		placement.Rack = e.Placement.Rack
		placement.RackOffsetUnits = e.Placement.RackOffsetUnits
		placement.RackOffset = e.Placement.RackOffset

		m.Placement = placement
	}
}

func createProcessorModel(e *entity.Processor) *model.Processor {
	m := model.Processor{}
	createResourceModel(&e.EmbeddedResource, &m.Resource)
	createProductInfoModel(&e.ProductInfo, &m.ProductInfo)
	m.Socket = e.Socket
	m.ProcessorType = e.ProcessorType
	m.ProcessorArchitecture = e.ProcessorArchitecture
	m.InstructionSet = e.InstructionSet
	m.MaxSpeedMHz = e.MaxSpeedMHz
	m.TotalCores = e.TotalCores
	m.TotalThreads = e.TotalThreads
	m.ProcessorID = new(model.ProcessorID)
	if e.ProcessorIDVendorID != nil ||
		e.ProcessorIDIdentificationRegisters != nil ||
		e.ProcessorIDEffectiveFamily != nil ||
		e.ProcessorIDEffectiveModel != nil ||
		e.ProcessorIDStep != nil ||
		e.ProcessorIDMicrocodeInfo != nil {
		m.ProcessorID = new(model.ProcessorID)
		m.ProcessorID.VendorID = e.ProcessorIDVendorID
		m.ProcessorID.IdentificationRegisters = e.ProcessorIDIdentificationRegisters
		m.ProcessorID.EffectiveFamily = e.ProcessorIDEffectiveFamily
		m.ProcessorID.EffectiveModel = e.ProcessorIDEffectiveModel
		m.ProcessorID.Step = e.ProcessorIDStep
		m.ProcessorID.MicrocodeInfo = e.ProcessorIDMicrocodeInfo
	}
	return &m
}

func createMemoryModel(e *entity.Memory) *model.Memory {
	m := model.Memory{}
	createResourceModel(&e.EmbeddedResource, &m.Resource)
	createProductInfoModel(&e.ProductInfo, &m.ProductInfo)
	m.CapacityMiB = e.CapacityMiB
	m.OperatingSpeedMhz = e.OperatingSpeedMhz
	m.MemoryDeviceType = e.MemoryDeviceType
	m.DataWidthBits = e.DataWidthBits
	m.RankCount = e.RankCount
	m.DeviceLocator = e.DeviceLocator
	if e.MemoryLocationSocket != nil ||
		e.MemoryLocationController != nil ||
		e.MemoryLocationChannel != nil ||
		e.MemoryLocationSlot != nil {
		m.MemoryLocation = new(model.MemoryLocation)
		m.MemoryLocation.Socket = e.MemoryLocationSocket
		m.MemoryLocation.Controller = e.MemoryLocationController
		m.MemoryLocation.Channel = e.MemoryLocationChannel
		m.MemoryLocation.Slot = e.MemoryLocationSlot
	}
	return &m
}

func createEthernetInterfaceModel(e *entity.EthernetInterface) *model.EthernetInterface {
	m := model.EthernetInterface{}
	createResourceModel(&e.EmbeddedResource, &m.Resource)
	m.UefiDevicePath = e.UefiDevicePath
	m.InterfaceEnabled = e.InterfaceEnabled
	m.PermanentMACAddress = e.PermanentMACAddress
	m.MACAddress = e.MACAddress
	m.SpeedMbps = e.SpeedMbps
	m.AutoNeg = e.AutoNeg
	m.FullDuplex = e.FullDuplex
	m.MTUSize = e.MTUSize
	m.HostName = e.HostName
	m.FQDN = e.FQDN
	m.MaxIPv6StaticAddresses = e.MaxIPv6StaticAddresses
	m.LinkStatus = e.LinkStatus
	m.IPv4Addresses = []model.IPv4Address{}
	for i := range e.IPv4Addresses {
		each := model.IPv4Address{}
		each.Address = e.IPv4Addresses[i].Address
		each.SubnetMask = e.IPv4Addresses[i].SubnetMask
		each.AddressOrigin = e.IPv4Addresses[i].AddressOrigin
		each.Gateway = e.IPv4Addresses[i].Gateway
		m.IPv4Addresses = append(m.IPv4Addresses, each)

	}
	m.IPv6Addresses = []model.IPv6Address{}
	for i := range e.IPv6Addresses {
		each := model.IPv6Address{}
		each.Address = e.IPv6Addresses[i].Address
		each.PrefixLength = e.IPv6Addresses[i].PrefixLength
		each.AddressOrigin = e.IPv6Addresses[i].AddressOrigin
		each.AddressState = e.IPv6Addresses[i].AddressState
		m.IPv6Addresses = append(m.IPv6Addresses, each)

	}
	m.VLANs = []model.VLanNetworkInterface{}
	for i := range e.VLANs {
		each := model.VLanNetworkInterface{}
		each.VLANEnable = e.VLANs[i].VLANEnable
		each.VLANID = e.VLANs[i].VLANID
		m.VLANs = append(m.VLANs, each)
	}
	return &m
}

func createNetworkInterfaceModel(e *entity.NetworkInterface) *model.NetworkInterface {
	m := model.NetworkInterface{}
	createResourceModel(&e.EmbeddedResource, &m.Resource)
	m.NetworkAdapterURI = e.NetworkAdapterURI
	return &m
}

func createStorageModel(e *entity.Storage) *model.Storage {
	m := model.Storage{}
	createResourceModel(&e.EmbeddedResource, &m.Resource)
	a := []string{}
	commonUtil.StringToStruct(e.DriveURIs, &a)
	m.DriveURIs = a
	for i := range e.StorageControllers {
		eachM := model.StorageController{}
		eachE := e.StorageControllers[i]
		createMemberModel(&eachE.Member, &eachM.Member)
		createProductInfoModel(&eachE.ProductInfo, &eachM.ProductInfo)
		eachM.SpeedGbps = eachE.SpeedGbps
		eachM.FirmwareVersion = eachE.FirmwareVersion
		a := []string{}
		commonUtil.StringToStruct(eachE.SupportedDeviceProtocols, &a)
		eachM.SupportedDeviceProtocols = a
		m.StorageControllers = append(m.StorageControllers, eachM)
	}
	return &m
}

func createPowerControlModel(e *entity.PowerControl) *model.PowerControl {
	m := model.PowerControl{}
	createResourceModel(&e.EmbeddedResource, &m.Resource)
	createProductInfoModel(&e.ProductInfo, &m.ProductInfo)
	m.PowerConsumedWatts = e.PowerConsumedWatts
	m.PowerRequestedWatts = e.PowerRequestedWatts
	m.PowerAvailableWatts = e.PowerAvailableWatts
	m.PowerCapacityWatts = e.PowerCapacityWatts
	m.PowerAllocatedWatts = e.PowerAllocatedWatts

	m.PowerMetrics = new(model.PowerMetrics)
	m.PowerMetrics.MinConsumedWatts = e.PowerMetricsMinConsumedWatts
	m.PowerMetrics.MaxConsumedWatts = e.PowerMetricsMaxConsumedWatts
	m.PowerMetrics.AverageConsumedWatts = e.PowerMetricsAverageConsumedWatts

	m.PowerLimit = new(model.PowerLimit)
	m.PowerLimit.LimitInWatts = e.PowerLimitLimitInWatts
	m.PowerLimit.LimitException = e.PowerLimitLimitException
	m.PowerLimit.CorrectionInMs = e.PowerLimitCorrectionInMs
	return &m
}

func createVoltageModel(e *entity.Voltage) *model.Voltage {
	m := model.Voltage{}
	createResourceModel(&e.EmbeddedResource, &m.Resource)
	createThresholdModel(&e.Threshold, &m.Threshold)
	m.SensorNumber = e.SensorNumber
	m.ReadingVolts = e.ReadingVolts
	m.MinReadingRange = e.MinReadingRange
	m.MaxReadingRange = e.MaxReadingRange
	m.PhysicalContext = e.PhysicalContext
	return &m
}

func createPowerSupplyModel(e *entity.PowerSupply) *model.PowerSupply {
	m := model.PowerSupply{}
	createResourceModel(&e.EmbeddedResource, &m.Resource)
	createProductInfoModel(&e.ProductInfo, &m.ProductInfo)
	m.PowerSupplyType = e.PowerSupplyType
	m.LineInputVoltageType = e.LineInputVoltageType
	m.LineInputVoltage = e.LineInputVoltage
	m.PowerCapacityWatts = e.PowerCapacityWatts
	m.LastPowerOutputWatts = e.LastPowerOutputWatts
	m.FirmwareVersion = e.FirmwareVersion
	m.IndicatorLed = e.IndicatorLed
	return &m
}

func createRedundancyModel(e *entity.Redundancy) *model.Redundancy {
	m := model.Redundancy{}
	createResourceModel(&e.EmbeddedResource, &m.Resource)
	m.Mode = e.Mode
	m.MaxNumSupported = e.MaxNumSupported
	m.MinNumNeeded = e.MinNumNeeded
	m.RedundancyEnabled = e.RedundancyEnabled
	a := []string{}
	commonUtil.StringToStruct(*e.RedundancySet, &a)
	m.RedundancySet = &a
	return &m
}

func createTemperatureModel(e *entity.Temperature) *model.Temperature {
	m := new(model.Temperature)
	createMemberModel(&e.Member, &m.Member)
	createThresholdModel(&e.Threshold, &m.Threshold)
	m.SensorNumber = e.SensorNumber
	m.ReadingCelsius = e.ReadingCelsius
	m.MinReadingRangeTemp = e.MinReadingRangeTemp
	m.MaxReadingRangeTemp = e.MaxReadingRangeTemp
	return m
}

func createFanModel(e *entity.Fan) *model.Fan {
	m := new(model.Fan)
	createMemberModel(&e.Member, &m.Member)
	createProductInfoModel(&e.ProductInfo, &m.ProductInfo)
	createThresholdModel(&e.Threshold, &m.Threshold)
	m.Reading = e.Reading
	m.MinReadingRange = e.MinReadingRange
	m.MaxReadingRange = e.MaxReadingRange
	m.ReadingUnits = e.ReadingUnits
	return m
}

func createOemHuaweiBoardModel(e *entity.OemHuaweiBoard) *model.OemHuaweiBoard {
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

func createNetworkAdapterModel(e *entity.NetworkAdapter) *model.NetworkAdapter {
	m := new(model.NetworkAdapter)
	createResourceModel(&e.EmbeddedResource, &m.Resource)
	createProductInfoModel(&e.ProductInfo, &m.ProductInfo)
	for i := range e.Controllers {
		controllerE := e.Controllers[i]
		controllerM := model.Controller{}
		controllerM.FirmwarePackageVersion = controllerE.FirmwarePackageVersion
		controllerM.ControllerCapabilities.NetworkPortCount = controllerE.ControllerCapabilitiesNetworkPortCount
		for j := range controllerE.NetworkPorts {
			portE := controllerE.NetworkPorts[j]
			portM := model.NetworkPort{}
			portM.PhysicalPortNumber = portE.PhysicalPortNumber
			portM.LinkStatus = portE.LinkStatus
			a := []string{}
			commonUtil.StringToStruct(portE.AssociatedNetworkAddresses, &a)
			portM.AssociatedNetworkAddresses = a
			controllerM.NetworkPorts = append(controllerM.NetworkPorts, portM)
		}
		m.Controllers = append(m.Controllers, controllerM)
	}
	return m
}

func createDriveModel(e *entity.Drive) *model.Drive {
	m := new(model.Drive)
	createResourceModel(&e.EmbeddedResource, &m.Resource)
	createProductInfoModel(&e.ProductInfo, &m.ProductInfo)
	m.StatusIndicator = e.StatusIndicator
	m.IndicatorLED = e.IndicatorLED
	m.Revision = e.Revision
	m.CapacityBytes = e.CapacityBytes
	m.FailurePredicted = e.FailurePredicted
	m.Protocol = e.Protocol
	m.MediaType = e.MediaType
	m.HotspareType = e.HotspareType
	m.CapableSpeedGbs = e.CapableSpeedGbs
	m.NegotiatedSpeedGbs = e.NegotiatedSpeedGbs
	m.PredictedMediaLifeLeftPercent = e.PredictedMediaLifeLeftPercent
	for i := range e.Location {
		locationM := new(model.Location)
		createLocationModel(&e.Location[i], locationM)
		m.Location = append(m.Location, *locationM)
	}
	return m
}

func createPCIeDeviceModel(e *entity.PCIeDevice) *model.PCIeDevice {
	m := new(model.PCIeDevice)
	createResourceModel(&e.EmbeddedResource, &m.Resource)
	createProductInfoModel(&e.ProductInfo, &m.ProductInfo)
	m.DeviceType = e.DeviceType
	m.FirmwareVersion = e.FirmwareVersion
	for i := range e.PCIeFunctions {
		m.PCIeFunctions = append(m.PCIeFunctions, *createPCIeFunctionModel(&e.PCIeFunctions[i]))
	}
	return m
}

func createPCIeFunctionModel(e *entity.PCIeFunction) *model.PCIeFunction {
	m := new(model.PCIeFunction)
	createResourceModel(&e.EmbeddedResource, &m.Resource)
	m.DeviceClass = e.DeviceClass
	m.DeviceID = e.DeviceID
	m.VendorID = e.VendorID
	m.SubsystemID = e.SubsystemID
	m.SubsystemVendorID = e.SubsystemVendorID
	a := []string{}
	commonUtil.StringToStruct(e.EthernetInterfaces, &a)
	m.EthernetInterfaces = a
	return m
}

func createServerEntityFromServer(s *model.Server) *entity.Server {
	var server entity.Server
	server.State = s.State
	server.Health = s.Health
	server.Name = s.Name
	server.Description = s.Description
	server.OriginURIsChassis = s.OriginURIs.Chassis
	server.OriginURIsSystem = s.OriginURIs.System
	server.PhysicalUUID = s.PhysicalUUID
	server.Address = s.Address
	server.Credential = s.Credential
	server.Type = s.Type
	server.Protocol = s.Protocol
	server.CurrentTask = s.CurrentTask
	return &server
}

func updateResourceEntity(e *entity.EmbeddedResource, m *model.Resource) {
	// ID won't be updated.
	e.URI = m.URI
	e.OriginID = m.OriginID
	e.Name = m.Name
	e.Description = m.Description
	e.State = m.State
	e.Health = m.Health
	e.PhysicalState = m.PhysicalState
	e.PhysicalHealth = m.PhysicalHealth
}

func updateMemberEntity(e *entity.Member, m *model.Member) {
	e.URI = m.URI
	e.OriginMemberID = m.OriginMemberID
	e.Name = m.Name
	e.Description = m.Description
	e.State = m.State
	e.Health = m.Health
	e.PhysicalState = m.PhysicalState
	e.PhysicalHealth = m.PhysicalHealth
}

func updateProductInfoEntity(e *entity.ProductInfo, m *model.ProductInfo) {
	e.Model = m.Model
	e.Manufacturer = m.Manufacturer
	e.SerialNumber = m.SerialNumber
	e.PartNumber = m.PartNumber
	e.SparePartNumber = m.SparePartNumber
	e.SKU = m.SKU
	e.AssetTag = m.AssetTag
}

func updateThresholdEntity(e *entity.Threshold, m *model.Threshold) {
	e.UpperThresholdNonCritical = m.UpperThresholdNonCritical
	e.UpperThresholdCritical = m.UpperThresholdCritical
	e.UpperThresholdFatal = m.UpperThresholdFatal
	e.LowerThresholdNonCritical = m.LowerThresholdNonCritical
	e.LowerThresholdCritical = m.LowerThresholdCritical
	e.LowerThresholdFatal = m.LowerThresholdFatal
}

func updateProcessor(e *entity.Processor, m *model.Processor) {
	updateResourceEntity(&(*e).EmbeddedResource, &(*m).Resource)
	updateProductInfoEntity(&(*e).ProductInfo, &(*m).ProductInfo)
	e.Socket = m.Socket
	e.ProcessorType = m.ProcessorType
	e.ProcessorArchitecture = m.ProcessorArchitecture
	e.InstructionSet = m.InstructionSet
	e.MaxSpeedMHz = m.MaxSpeedMHz
	e.TotalCores = m.TotalCores
	e.TotalThreads = m.TotalThreads
	e.ProcessorIDVendorID = m.ProcessorID.VendorID
	e.ProcessorIDIdentificationRegisters = m.ProcessorID.IdentificationRegisters
	e.ProcessorIDEffectiveFamily = m.ProcessorID.EffectiveFamily
	e.ProcessorIDEffectiveModel = m.ProcessorID.EffectiveModel
	e.ProcessorIDStep = m.ProcessorID.Step
	e.ProcessorIDMicrocodeInfo = m.ProcessorID.MicrocodeInfo
}

func updateMemory(e *entity.Memory, m *model.Memory) {
	updateResourceEntity(&(*e).EmbeddedResource, &(*m).Resource)
	updateProductInfoEntity(&(*e).ProductInfo, &(*m).ProductInfo)
	e.MemoryType = m.MemoryType
	e.MemoryDeviceType = m.MemoryDeviceType
	e.BaseModuleType = m.BaseModuleType
	e.MemoryMedia = m.MemoryMedia
	e.CapacityMiB = m.CapacityMiB
	e.DataWidthBits = m.DataWidthBits
	e.BusWidthBits = m.BusWidthBits
	e.AllowedSpeedsMHz = m.AllowedSpeedsMHz
	e.FirmwareRevision = m.FirmwareRevision
	e.FirmwareAPIVersion = m.FirmwareAPIVersion
	e.VendorID = m.VendorID
	e.DeviceID = m.DeviceID
	e.SubsystemVendorID = m.SubsystemVendorID
	e.SubsystemDeviceID = m.SubsystemDeviceID
	e.SpareDeviceCount = m.SpareDeviceCount
	e.RankCount = m.RankCount
	e.DeviceLocator = m.DeviceLocator
	if m.MemoryLocation != nil {
		e.MemoryLocationSocket = m.MemoryLocation.Socket
		e.MemoryLocationController = m.MemoryLocation.Controller
		e.MemoryLocationChannel = m.MemoryLocation.Channel
		e.MemoryLocationSlot = m.MemoryLocation.Slot
	}
	e.ErrorCorrection = m.ErrorCorrection
	e.OperatingSpeedMhz = m.OperatingSpeedMhz
	e.VolatileRegionSizeLimitMiB = m.VolatileRegionSizeLimitMiB
	e.PersistentRegionSizeLimitMiB = m.PersistentRegionSizeLimitMiB
	e.OperatingMemoryModes = m.OperatingMemoryModes
	e.IsSpareDeviceEnabled = m.IsSpareDeviceEnabled
	e.IsRankSpareEnabled = m.IsRankSpareEnabled
	e.VolatileRegionNumberLimit = m.VolatileRegionNumberLimit
	e.PersistentRegionNumberLimit = m.PersistentRegionNumberLimit
	e.VolatileRegionSizeMaxMiB = m.VolatileRegionSizeMaxMiB
	e.PersistentRegionSizeMaxMiB = m.PersistentRegionSizeMaxMiB
	e.AllocationIncrementMiB = m.AllocationIncrementMiB
	e.AllocationAlignmentMiB = m.AllocationAlignmentMiB
}

func updateEthernetInterface(e *entity.EthernetInterface, m *model.EthernetInterface) {
	updateResourceEntity(&(*e).EmbeddedResource, &(*m).Resource)
	e.UefiDevicePath = m.UefiDevicePath
	e.InterfaceEnabled = m.InterfaceEnabled
	e.PermanentMACAddress = m.PermanentMACAddress
	e.MACAddress = m.MACAddress
	e.SpeedMbps = m.SpeedMbps
	e.AutoNeg = m.AutoNeg
	e.FullDuplex = m.FullDuplex
	e.MTUSize = m.MTUSize
	e.HostName = m.HostName
	e.FQDN = m.FQDN
	e.MaxIPv6StaticAddresses = m.MaxIPv6StaticAddresses
	e.LinkStatus = m.LinkStatus

	for i := range m.IPv4Addresses {
		each := entity.IPv4Address{}
		each.Address = m.IPv4Addresses[i].Address
		each.SubnetMask = m.IPv4Addresses[i].SubnetMask
		each.AddressOrigin = m.IPv4Addresses[i].AddressOrigin
		each.Gateway = m.IPv4Addresses[i].Gateway
		e.IPv4Addresses = append(e.IPv4Addresses, each)
	}

	for i := range m.IPv6Addresses {
		each := entity.IPv6Address{}
		each.Address = m.IPv6Addresses[i].Address
		each.PrefixLength = m.IPv6Addresses[i].PrefixLength
		each.AddressOrigin = m.IPv6Addresses[i].AddressOrigin
		each.AddressState = m.IPv6Addresses[i].AddressState
		e.IPv6Addresses = append(e.IPv6Addresses, each)
	}
	for i := range m.VLANs {
		each := entity.VLanNetworkInterface{}
		updateResourceEntity(&each.EmbeddedResource, &m.VLANs[i].Resource)
		each.VLANEnable = m.VLANs[i].VLANEnable
		each.VLANID = m.VLANs[i].VLANID
		e.VLANs = append(e.VLANs, each)
	}
}

func updateNetworkInterface(e *entity.NetworkInterface, m *model.NetworkInterface) {
	updateResourceEntity(&e.EmbeddedResource, &m.Resource)
	e.NetworkAdapterURI = m.NetworkAdapterURI
}

func updateStorage(e *entity.Storage, m *model.Storage) {
	updateResourceEntity(&e.EmbeddedResource, &m.Resource)
	s := commonUtil.StructToString(m.DriveURIs)
	e.DriveURIs = s
	for i := range m.StorageControllers {
		each := entity.StorageController{}
		updateMemberEntity(&each.Member, &m.StorageControllers[i].Member)
		updateProductInfoEntity(&each.ProductInfo, &m.StorageControllers[i].ProductInfo)
		each.SpeedGbps = m.StorageControllers[i].SpeedGbps
		each.FirmwareVersion = m.StorageControllers[i].FirmwareVersion
		s := commonUtil.StructToString(m.StorageControllers[i].SupportedDeviceProtocols)
		each.SupportedDeviceProtocols = s
		e.StorageControllers = append(e.StorageControllers, each)
	}
}

func updatePower(e *entity.Power, m *model.Power) {
	updateResourceEntity(&e.EmbeddedResource, &m.Resource)
	// PowerControl
	if m.PowerControl != nil {
		for i := range *m.PowerControl {
			each := entity.PowerControl{}
			powerControl := (*m.PowerControl)[i]
			updateResourceEntity(&each.EmbeddedResource, &powerControl.Resource)
			updateProductInfoEntity(&each.ProductInfo, &powerControl.ProductInfo)
			each.PowerConsumedWatts = powerControl.PowerConsumedWatts
			each.PowerRequestedWatts = powerControl.PowerRequestedWatts
			each.PowerAvailableWatts = powerControl.PowerAvailableWatts
			each.PowerCapacityWatts = powerControl.PowerCapacityWatts
			each.PowerAllocatedWatts = powerControl.PowerAllocatedWatts
			if powerControl.PowerMetrics != nil {
				each.PowerMetricsMinConsumedWatts = powerControl.PowerMetrics.MinConsumedWatts
				each.PowerMetricsMaxConsumedWatts = powerControl.PowerMetrics.MaxConsumedWatts
				each.PowerMetricsAverageConsumedWatts = powerControl.PowerMetrics.AverageConsumedWatts
			}
			if powerControl.PowerLimit != nil {
				each.PowerLimitLimitInWatts = powerControl.PowerLimit.LimitInWatts
				each.PowerLimitLimitException = powerControl.PowerLimit.LimitException
				each.PowerLimitCorrectionInMs = powerControl.PowerLimit.CorrectionInMs
			}
			e.PowerControl = append(e.PowerControl, each)
		}
	}
	// Voltages
	if m.Voltages != nil {
		for i := range *m.Voltages {
			each := entity.Voltage{}
			voltages := (*m.Voltages)[i]
			updateResourceEntity(&each.EmbeddedResource, &voltages.Resource)
			updateThresholdEntity(&each.Threshold, &voltages.Threshold)
			each.SensorNumber = voltages.SensorNumber
			each.ReadingVolts = voltages.ReadingVolts
			each.MinReadingRange = voltages.MinReadingRange
			each.MaxReadingRange = voltages.MaxReadingRange
			each.PhysicalContext = voltages.PhysicalContext
			e.Voltages = append(e.Voltages, each)
		}
	}
	// PowerSupplies
	if m.PowerSupplies != nil {
		for i := range *m.PowerSupplies {
			each := entity.PowerSupply{}
			powerSupplies := (*m.PowerSupplies)[i]
			updateResourceEntity(&each.EmbeddedResource, &powerSupplies.Resource)
			updateProductInfoEntity(&each.ProductInfo, &powerSupplies.ProductInfo)
			each.PowerSupplyType = powerSupplies.PowerSupplyType
			each.LineInputVoltageType = powerSupplies.LineInputVoltageType
			each.LineInputVoltage = powerSupplies.LineInputVoltage
			each.PowerCapacityWatts = powerSupplies.PowerCapacityWatts
			each.LastPowerOutputWatts = powerSupplies.LastPowerOutputWatts
			each.FirmwareVersion = powerSupplies.FirmwareVersion
			each.IndicatorLed = powerSupplies.IndicatorLed
			e.PowerSupplies = append(e.PowerSupplies, each)
		}
	}
	// Redundancy
	if m.Redundancy != nil {
		for i := range *m.Redundancy {
			each := entity.Redundancy{}
			redundancy := (*m.Redundancy)[i]
			updateResourceEntity(&each.EmbeddedResource, &redundancy.Resource)
			each.Mode = redundancy.Mode
			each.MaxNumSupported = redundancy.MaxNumSupported
			each.MinNumNeeded = redundancy.MinNumNeeded
			s := commonUtil.StructToString(redundancy.RedundancySet)
			each.RedundancySet = &s
			each.RedundancyEnabled = redundancy.RedundancyEnabled
			e.Redundancy = append(e.Redundancy, each)
		}
	}
}

func updateThermal(e *entity.Thermal, m *model.Thermal) {
	updateResourceEntity(&e.EmbeddedResource, &m.Resource)
	temperatures := []entity.Temperature{}
	for i := range m.Temperatures {
		e := entity.Temperature{}
		m := m.Temperatures[i]
		updateMemberEntity(&e.Member, &m.Member)
		updateThresholdEntity(&e.Threshold, &m.Threshold)
		e.SensorNumber = m.SensorNumber
		e.ReadingCelsius = m.ReadingCelsius
		e.MinReadingRangeTemp = m.MinReadingRangeTemp
		e.MaxReadingRangeTemp = m.MaxReadingRangeTemp
		temperatures = append(temperatures, e)
	}
	e.Temperatures = temperatures

	fans := []entity.Fan{}
	for i := range m.Fans {
		e := entity.Fan{}
		m := m.Fans[i]
		updateMemberEntity(&e.Member, &m.Member)
		updateProductInfoEntity(&e.ProductInfo, &m.ProductInfo)
		updateThresholdEntity(&e.Threshold, &m.Threshold)
		e.Reading = m.Reading
		e.MinReadingRange = m.MinReadingRange
		e.MaxReadingRange = m.MaxReadingRange
		e.ReadingUnits = m.ReadingUnits
		fans = append(fans, e)
	}
	e.Fans = fans
}

func updateOemHuaweiBoard(e *entity.OemHuaweiBoard, m *model.OemHuaweiBoard) {
	updateResourceEntity(&e.EmbeddedResource, &m.Resource)
	updateProductInfoEntity(&e.ProductInfo, &m.ProductInfo)
	e.CardNo = m.CardNo
	e.DeviceLocator = m.DeviceLocator
	e.DeviceType = m.DeviceType
	e.Location = m.Location
	e.CPLDVersion = m.CPLDVersion
	e.PCBVersion = m.PCBVersion
	e.BoardName = m.BoardName
	e.BoardID = m.BoardID
	e.ManufactureDate = m.ManufactureDate
}

func updateNetworkAdapter(e *entity.NetworkAdapter, m *model.NetworkAdapter) {
	updateResourceEntity(&e.EmbeddedResource, &m.Resource)
	updateProductInfoEntity(&e.ProductInfo, &m.ProductInfo)
	controllers := []entity.Controller{}
	for i := range m.Controllers {
		controllerE := entity.Controller{}
		controllerM := m.Controllers[i]
		controllerE.FirmwarePackageVersion = controllerM.FirmwarePackageVersion
		controllerE.ControllerCapabilitiesNetworkPortCount = controllerM.ControllerCapabilities.NetworkPortCount
		ports := []entity.NetworkPort{}
		for j := range controllerM.NetworkPorts {
			portE := entity.NetworkPort{}
			portM := controllerM.NetworkPorts[j]
			updateResourceEntity(&portE.EmbeddedResource, &portM.Resource)
			portE.PhysicalPortNumber = portM.PhysicalPortNumber
			portE.LinkStatus = portM.LinkStatus
			s := commonUtil.StructToString(portM.AssociatedNetworkAddresses)
			portE.AssociatedNetworkAddresses = s
			ports = append(ports, portE)
		}
		controllerE.NetworkPorts = ports
		controllers = append(controllers, controllerE)
	}
	e.Controllers = controllers
}

func updateLocation(e *entity.Location, m *model.Location) {
	e.Info = m.Info
	e.InfoFormat = m.InfoFormat
	if m.PostalAddress != nil {
		postalAddress := new(entity.PostalAddress)
		postalAddress.Country = m.PostalAddress.Country
		postalAddress.Territory = m.PostalAddress.Territory
		postalAddress.District = m.PostalAddress.District
		postalAddress.City = m.PostalAddress.City
		postalAddress.Division = m.PostalAddress.Division
		postalAddress.Neighborhood = m.PostalAddress.Neighborhood
		postalAddress.LeadingStreetDirection = m.PostalAddress.LeadingStreetDirection
		postalAddress.Street = m.PostalAddress.Street
		postalAddress.TrailingStreetSuffix = m.PostalAddress.TrailingStreetSuffix
		postalAddress.StreetSuffix = m.PostalAddress.StreetSuffix
		postalAddress.HouseNumber = m.PostalAddress.HouseNumber
		postalAddress.HouseNumberSuffix = m.PostalAddress.HouseNumberSuffix
		postalAddress.Landmark = m.PostalAddress.Landmark
		postalAddress.Location = m.PostalAddress.Location
		postalAddress.Floor = m.PostalAddress.Floor
		postalAddress.Name = m.PostalAddress.Name
		postalAddress.PostalCode = m.PostalAddress.PostalCode
		postalAddress.Building = m.PostalAddress.Building
		postalAddress.Unit = m.PostalAddress.Unit
		postalAddress.Room = m.PostalAddress.Room
		postalAddress.Seat = m.PostalAddress.Seat
		postalAddress.PlaceType = m.PostalAddress.PlaceType
		postalAddress.Community = m.PostalAddress.Community
		postalAddress.POBox = m.PostalAddress.POBox
		postalAddress.AdditionalCode = m.PostalAddress.AdditionalCode
		postalAddress.Road = m.PostalAddress.Road
		postalAddress.RoadSection = m.PostalAddress.RoadSection
		postalAddress.RoadBranch = m.PostalAddress.RoadBranch
		postalAddress.RoadSubBranch = m.PostalAddress.RoadSubBranch
		postalAddress.RoadPreModifier = m.PostalAddress.RoadPreModifier
		postalAddress.RoadPostModifier = m.PostalAddress.RoadPostModifier
		postalAddress.GPSCoords = m.PostalAddress.GPSCoords

		e.PostalAddress = postalAddress
	}
	if m.Placement != nil {
		placement := new(entity.Placement)
		placement.Row = m.Placement.Row
		placement.Rack = m.Placement.Rack
		placement.RackOffsetUnits = m.Placement.RackOffsetUnits
		placement.RackOffset = m.Placement.RackOffset

		e.Placement = placement
	}
}

func updateDrive(e *entity.Drive, m *model.Drive) {
	updateResourceEntity(&e.EmbeddedResource, &m.Resource)
	updateProductInfoEntity(&e.ProductInfo, &m.ProductInfo)
	e.StatusIndicator = m.StatusIndicator
	e.IndicatorLED = m.IndicatorLED
	e.Revision = m.Revision
	e.CapacityBytes = m.CapacityBytes
	e.FailurePredicted = m.FailurePredicted
	e.Protocol = m.Protocol
	e.MediaType = m.MediaType
	e.HotspareType = m.HotspareType
	e.CapableSpeedGbs = m.CapableSpeedGbs
	e.NegotiatedSpeedGbs = m.NegotiatedSpeedGbs
	e.PredictedMediaLifeLeftPercent = m.PredictedMediaLifeLeftPercent
	for i := range m.Location {
		locationE := new(entity.Location)
		updateLocation(locationE, &m.Location[i])
		e.Location = append(e.Location, *locationE)
	}
}

func updatePCIeDevice(e *entity.PCIeDevice, m *model.PCIeDevice) {
	updateResourceEntity(&e.EmbeddedResource, &m.Resource)
	updateProductInfoEntity(&e.ProductInfo, &m.ProductInfo)
	e.DeviceType = m.DeviceType
	e.FirmwareVersion = m.FirmwareVersion
	for i := range m.PCIeFunctions {
		pcieFunctionE := new(entity.PCIeFunction)
		updatePCIeFunction(pcieFunctionE, &m.PCIeFunctions[i])
		e.PCIeFunctions = append(e.PCIeFunctions, *pcieFunctionE)
	}
}

func updatePCIeFunction(e *entity.PCIeFunction, m *model.PCIeFunction) {
	updateResourceEntity(&e.EmbeddedResource, &m.Resource)
	e.DeviceClass = m.DeviceClass
	e.DeviceID = m.DeviceID
	e.VendorID = m.VendorID
	e.SubsystemID = m.SubsystemID
	e.SubsystemVendorID = m.SubsystemVendorID
	s := commonUtil.StructToString(m.EthernetInterfaces)
	e.EthernetInterfaces = s
}

func createProcessor(m *model.Processor) *entity.Processor {
	e := entity.Processor{}
	updateProcessor(&e, m)
	return &e
}

func createMemory(m *model.Memory) *entity.Memory {
	e := new(entity.Memory)
	updateMemory(e, m)
	return e
}

func createEthernetInterface(m *model.EthernetInterface) *entity.EthernetInterface {
	e := new(entity.EthernetInterface)
	updateEthernetInterface(e, m)
	return e
}

func createNetworkInterface(m *model.NetworkInterface) *entity.NetworkInterface {
	e := new(entity.NetworkInterface)
	updateNetworkInterface(e, m)
	return e
}

func createStorage(m *model.Storage) *entity.Storage {
	e := new(entity.Storage)
	updateStorage(e, m)
	return e
}

func createPower(m *model.Power) *entity.Power {
	e := new(entity.Power)
	updatePower(e, m)
	return e
}

func createThermal(m *model.Thermal) *entity.Thermal {
	e := new(entity.Thermal)
	updateThermal(e, m)
	return e
}

func createOemHuaweiBoard(m *model.OemHuaweiBoard) *entity.OemHuaweiBoard {
	e := new(entity.OemHuaweiBoard)
	updateOemHuaweiBoard(e, m)
	return e
}

func createNetworkAdapter(m *model.NetworkAdapter) *entity.NetworkAdapter {
	e := new(entity.NetworkAdapter)
	updateNetworkAdapter(e, m)
	return e
}

func createDrive(m *model.Drive) *entity.Drive {
	e := new(entity.Drive)
	updateDrive(e, m)
	return e
}

func createPCIeDevice(m *model.PCIeDevice) *entity.PCIeDevice {
	e := new(entity.PCIeDevice)
	updatePCIeDevice(e, m)
	return e
}
