package db

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/server/object/constvalue"
	"promise/server/object/entity"
	"promise/server/object/model"
)

// Server is the concrete DB for server.
type Server struct {
	base.DB
}

// ResourceName get the resource name.
func (impl *Server) ResourceName() string {
	return "servergroup"
}

// NewEntity return the a new entity.
func (impl *Server) NewEntity() base.EntityInterface {
	e := new(entity.Server)
	e.Entity.TemplateImpl = e
	return e
}

// NewEntityCollection return a collection of entity.
func (impl *Server) NewEntityCollection() interface{} {
	return new([]entity.Server)
}

// GetConnection return the DB connection.
func (impl *Server) GetConnection() *gorm.DB {
	return base.GetConnection()
}

// NeedCheckDuplication return if need check duplication for entity.
func (impl *Server) NeedCheckDuplication() bool {
	return true
}

// ConvertFindResultToCollection convert the Find() result to collection mode.
func (impl *Server) ConvertFindResultToCollection(start int64, total int64, result interface{}) (*base.CollectionModel, error) {
	collection, ok := result.(*[]entity.Server)
	if !ok {
		log.Error("Server.ConvertFindResult() failed, convert data failed.")
		return nil, base.ErrorDataConvert
	}
	ret := base.CollectionModel{}
	ret.Start = start
	ret.Count = int64(len(*collection))
	ret.Total = total
	for _, v := range *collection {
		ret.Members = append(ret.Members, v.ToCollectionMember())
	}
	return &ret, nil
}

// ConvertFindResultToModel convert the Find() result to model slice
func (impl *Server) ConvertFindResultToModel(result interface{}) ([]base.ModelInterface, error) {
	collection, ok := result.(*[]entity.Server)
	if !ok {
		log.Error("Server.ConvertFindResult() failed, convert data failed.")
		return nil, base.ErrorDataConvert
	}
	ret := make([]base.ModelInterface, 0)
	for _, v := range *collection {
		m := v.ToModel()
		ret = append(ret, m)
	}
	return ret, nil
}

// TODO There is a cap between we find the server with state "Added" and lock it.

// FindServerStateAdded Find the first server with state "Added" and return the server ID.
func (impl *Server) FindServerStateAdded() string {
	var (
		c      = impl.TemplateImpl.GetConnection()
		server = new(entity.Server)
	)
	if notFound := c.Where("\"State\" = ?", constvalue.ServerStateAdded).First(server).RecordNotFound(); notFound {
		return ""
	}
	return server.ID
}

// IsServerExist Check the existance of the server, if found, return it.
func (impl *Server) IsServerExist(s *model.Server) (bool, base.ModelInterface) {
	var (
		c      = impl.TemplateImpl.GetConnection()
		server = new(entity.Server)
	)

	if s == nil {
		return false, nil
	}

	c.Where("\"Physical_UUID\" = ?", s.PhysicalUUID).First(server)
	if s.PhysicalUUID == server.PhysicalUUID {
		return true, server.ToModel()
	}
	return false, nil
}

// CreateServer will save the server to the DB.
// Since a server should belongs to the default servergroup, we need create server and servergroup
// in DB at the same time, we also need return them both back for event dispatch.
func (impl *Server) CreateServer(s *model.Server) (base.ModelInterface, base.ModelInterface, error) {
	var (
		c      = impl.TemplateImpl.GetConnection()
		server = new(entity.Server)
	)
	if exist, _ := impl.IsServerExist(s); exist {
		// TODO get the server to return.
		log.WithFields(log.Fields{"hostname": s.Hostname}).Info("Post server in DB failed, server exist.")
		return nil, nil, base.ErrorResourceExist
	}
	server.Load(s)
	var serverServerGroup = new(entity.ServerServerGroup)

	// Generate the UUID.
	server.ID = uuid.New().String()
	serverServerGroup.ID = uuid.New().String()
	serverServerGroup.Category = base.CategoryServerServerGroup
	serverServerGroup.ServerID = server.ID
	serverServerGroup.ServerGroupID = DefaultServerGroupID
	// We need make sure save server and add it to the default server group both success or failure.
	tx := c.Begin()
	if err := tx.Error; err != nil {
		log.WithFields(log.Fields{"hostname": s.Hostname, "error": err}).Info("Post server in DB failed, start transaction failed.")
		return nil, nil, err
	}
	if err := tx.Create(server).Error; err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{"hostname": s.Hostname, "error": err}).Info("Post server in DB failed, create server failed, transaction rollback.")
		return nil, nil, err
	}
	if err := tx.Create(serverServerGroup).Error; err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{"hostname": s.Hostname, "error": err}).Info("Post server in DB failed, create server-servergroup failed, transaction rollback.")
		return nil, nil, err
	}
	if err := tx.Commit().Error; err != nil {
		log.WithFields(log.Fields{"hostname": s.Hostname, "error": err}).Info("Post server in DB failed, commit failed.")
		return nil, nil, err
	}
	return server.ToModel(), serverServerGroup.ToModel(), nil
}

// GetAndLockServer Get and lock the server.
func (impl *Server) GetAndLockServer(ID string) (bool, base.ModelInterface) {
	var (
		c      = impl.TemplateImpl.GetConnection()
		server = new(entity.Server)
	)
	// Transaction start.
	tx := c.Begin()
	if err := tx.Error; err != nil {
		log.WithFields(log.Fields{
			"id":    ID,
			"error": err}).
			Warn("Get and lock server in DB failed, start transaction failed.")
		return false, nil
	}

	if tx.Where("\"ID\" = ?", ID).First(server).RecordNotFound() {
		tx.Rollback()
		log.WithFields(log.Fields{
			"id": ID}).
			Debug("Get and lock server in DB failed, server does not exist.")
		return false, nil
	}
	if !constvalue.ServerLockable(server.State) {
		// Server not ready, rollback.
		tx.Rollback()
		log.WithFields(log.Fields{
			"id":    ID,
			"state": server.State}).
			Debug("Get and lock server in DB failed, server not lockable.")
		return false, server.ToModel()
	}
	// Change the state.
	if err := tx.Model(server).UpdateColumn("State", constvalue.ServerStateLocked).Error; err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"id":    ID,
			"state": server.State}).
			Debug("Get and lock server in DB failed, update state failed.")
		return false, nil
	}
	// Commit.
	if err := tx.Commit().Error; err != nil {
		log.WithFields(log.Fields{
			"id":    ID,
			"error": err}).
			Warn("Get and lock server in DB failed, commit failed.")
		return false, nil
	}
	return true, server.ToModel()
}

// SetServerState Set server state.
func (impl *Server) SetServerState(ID string, state string) bool {
	var (
		c      = impl.TemplateImpl.GetConnection()
		server = new(entity.Server)
	)
	if c.Where("\"ID\" = ?", ID).First(server).RecordNotFound() {
		return false
	}
	if err := c.Model(server).UpdateColumn("State", state).Error; err != nil {
		log.WithFields(log.Fields{"id": ID, "op": "SetServerState", "error": err}).Warn("DB opertion failed.")
		return false
	}
	return true
}

// SetServerHealth Set server health.
func (impl *Server) SetServerHealth(ID string, health string) bool {
	var (
		c      = impl.TemplateImpl.GetConnection()
		server = new(entity.Server)
	)
	if c.Where("\"ID\" = ?", ID).First(server).RecordNotFound() {
		return false
	}
	if err := c.Model(server).UpdateColumn("Health", health).Error; err != nil {
		log.WithFields(log.Fields{"id": ID, "op": "SetServerHealth", "error": err}).Warn("DB opertion failed.")
		return false
	}
	return true
}

func (impl *Server) deleteProcessors(c *gorm.DB, server *entity.Server) error {
	for i := range server.Processors {
		c.Delete(server.Processors[i])
	}
	return nil
}

// UpdateProcessors Update processors info.
func (impl *Server) UpdateProcessors(ID string, processors []model.Processor) error {
	var (
		c      = impl.TemplateImpl.GetConnection()
		server = new(entity.Server)
	)

	notFound := c.Where("\"ID\" = ?", ID).
		Preload("Processors").
		First(server).
		RecordNotFound()
	if notFound {
		return fmt.Errorf("can not find server %s", ID)
	}
	impl.deleteProcessors(c, server)
	server.Processors = []entity.Processor{}
	for _, v := range processors {
		each := entity.Processor{}
		each.Load(&v)
		server.Processors = append(server.Processors, each)
	}
	if err := c.Save(server).Error; err != nil {
		log.WithFields(log.Fields{"id": ID, "op": "UpdateProcessors", "error": err}).Warn("DB opertion failed.")
		return err
	}
	return nil
}

func (impl *Server) deleteMemory(c *gorm.DB, server *entity.Server) error {
	for i := range server.Memory {
		c.Delete(server.Memory[i])
	}
	return nil
}

// UpdateMemory Update server memory
func (impl *Server) UpdateMemory(ID string, memory []model.Memory) error {
	var (
		c      = impl.TemplateImpl.GetConnection()
		server = new(entity.Server)
	)

	notFound := c.Where("\"ID\" = ?", ID).
		Preload("Memory").
		First(server).
		RecordNotFound()
	if notFound {
		return fmt.Errorf("Can't find server %s", ID)
	}
	impl.deleteMemory(c, server)
	server.Memory = []entity.Memory{}
	for _, v := range memory {
		each := entity.Memory{}
		each.Load(&v)
		server.Memory = append(server.Memory, each)
	}
	if err := c.Save(server).Error; err != nil {
		log.WithFields(log.Fields{"id": ID, "op": "UpdateMemory", "error": err}).Warn("DB opertion failed.")
		return err
	}
	return nil
}

func (impl *Server) deleteEthernetInterfaces(c *gorm.DB, server *entity.Server) error {
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
	return nil
}

// UpdateEthernetInterfaces Update server ethernet interface.
// Each ethernet interface has many IPv4, IPv6 and vLAN, so it's very hard to check if 2 ethernet interface
// are the same. So we just remove all of them and recreate them.
func (impl *Server) UpdateEthernetInterfaces(ID string, ethernet []model.EthernetInterface) error {
	var (
		c      = impl.TemplateImpl.GetConnection()
		server = new(entity.Server)
	)

	// Get Server and it's ethernet interface.
	notFound := c.Where("\"ID\" = ?", ID).
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
	impl.deleteEthernetInterfaces(c, server)
	// Regenerate ethernet interface.
	server.EthernetInterfaces = []entity.EthernetInterface{}
	for _, v := range ethernet {
		each := entity.EthernetInterface{}
		each.Load(&v)
		server.EthernetInterfaces = append(server.EthernetInterfaces, each)
	}
	if err := c.Save(server).Error; err != nil {
		log.WithFields(log.Fields{"id": ID, "op": "UpdateEthernetInterfaces", "error": err}).Warn("DB opertion failed.")
		return err
	}
	return nil
}

func (impl *Server) deleteNetworkInterfaces(c *gorm.DB, server *entity.Server) error {
	for i := range server.NetworkInterfaces {
		c.Delete(server.NetworkInterfaces[i])
	}
	return nil
}

// UpdateNetworkInterfaces Update network interfaces.
func (impl *Server) UpdateNetworkInterfaces(ID string, networkInterface []model.NetworkInterface) error {
	var (
		c      = impl.TemplateImpl.GetConnection()
		server = new(entity.Server)
	)

	// Get resources.
	notFound := c.Where("\"ID\" = ?", ID).
		Preload("NetworkInterfaces").
		First(server).
		RecordNotFound()
	if notFound {
		return fmt.Errorf("Can't find server %s", ID)
	}
	// Delete them.
	impl.deleteEthernetInterfaces(c, server)
	networkInterfacesE := []entity.NetworkInterface{}
	for _, v := range networkInterface {
		each := entity.NetworkInterface{}
		each.Load(&v)
		networkInterfacesE = append(networkInterfacesE, each)
	}
	server.NetworkInterfaces = networkInterfacesE
	if err := c.Save(server).Error; err != nil {
		log.WithFields(log.Fields{"id": ID, "op": "UpdateNetworkInterfaces", "error": err}).Warn("DB opertion failed.")
		return err
	}
	return nil
}

func (impl *Server) deleteStorages(c *gorm.DB, server *entity.Server) error {
	for i := range server.Storages {
		for j := range server.Storages[i].StorageControllers {
			c.Delete(&(server.Storages[i].StorageControllers)[j])
		}
		c.Delete(server.Storages[i])
	}
	return nil
}

// UpdateStorages Update storages.
func (impl *Server) UpdateStorages(ID string, storages []model.Storage) error {
	var (
		c      = impl.TemplateImpl.GetConnection()
		server = new(entity.Server)
	)

	notFound := c.Where("\"ID\" = ?", ID).
		Preload("Storages").
		Preload("Storages.StorageControllers").
		First(server).
		RecordNotFound()
	if notFound {
		return fmt.Errorf("Can't find server %s", ID)
	}
	impl.deleteEthernetInterfaces(c, server)
	storagesE := []entity.Storage{}
	for _, v := range storages {
		each := entity.Storage{}
		each.Load(&v)
		storagesE = append(storagesE, each)
	}
	server.Storages = storagesE
	if err := c.Save(server).Error; err != nil {
		log.WithFields(log.Fields{"id": ID, "op": "UpdateStorages", "error": err}).Warn("DB opertion failed.")
		return err
	}
	return nil
}

func (impl *Server) deletePower(c *gorm.DB, server *entity.Server) error {
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
	return nil
}

// UpdatePower Update power
func (impl *Server) UpdatePower(ID string, power *model.Power) error {
	var (
		c      = impl.TemplateImpl.GetConnection()
		server = new(entity.Server)
	)

	notFound := c.Where("\"ID\" = ?", ID).
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
	impl.deletePower(c, server)
	server.Power.Load(power)
	if err := c.Save(server).Error; err != nil {
		log.WithFields(log.Fields{"id": ID, "op": "UpdatePower", "error": err}).Warn("DB opertion failed.")
		return err
	}
	return nil
}

func (impl *Server) deleteThermal(c *gorm.DB, server *entity.Server) error {
	for i := range server.Thermal.Temperatures {
		c.Delete(server.Thermal.Temperatures[i])
	}
	for i := range server.Thermal.Fans {
		c.Delete(server.Thermal.Fans[i])
	}
	c.Delete(server.Thermal)
	return nil
}

// UpdateThermal Update thermal
func (impl *Server) UpdateThermal(ID string, thermal *model.Thermal) error {
	var (
		c      = impl.TemplateImpl.GetConnection()
		server = new(entity.Server)
	)
	notFound := c.Where("\"ID\" = ?", ID).
		Preload("Thermal").
		Preload("Thermal.Temperatures").
		Preload("Thermal.Fans").
		First(server).
		RecordNotFound()
	if notFound {
		return fmt.Errorf("Can't find server %s", ID)
	}
	impl.deleteThermal(c, server)
	server.Thermal.Load(thermal)
	if err := c.Save(server).Error; err != nil {
		log.WithFields(log.Fields{"id": ID, "op": "UpdateThermal", "error": err}).Warn("DB opertion failed.")
		return err
	}
	return nil
}

func (impl *Server) deleteOemHuaweiBoards(c *gorm.DB, server *entity.Server) error {
	for i := range server.OemHuaweiBoards {
		c.Delete(server.OemHuaweiBoards[i])
	}
	return nil
}

// UpdateOemHuaweiBoards Update OEM Huawei boards.
func (impl *Server) UpdateOemHuaweiBoards(ID string, boards []model.OemHuaweiBoard) error {
	var (
		c      = impl.TemplateImpl.GetConnection()
		server = new(entity.Server)
	)

	notFound := c.Where("\"ID\" = ?", ID).
		Preload("OemHuaweiBoards").
		First(server).
		RecordNotFound()
	if notFound {
		return fmt.Errorf("Can't find server %s", ID)
	}
	impl.deleteOemHuaweiBoards(c, server)
	boardsE := []entity.OemHuaweiBoard{}
	for _, v := range boards {
		each := entity.OemHuaweiBoard{}
		each.Load(&v)
		boardsE = append(boardsE, each)
	}
	server.OemHuaweiBoards = boardsE
	if err := c.Save(server).Error; err != nil {
		log.WithFields(log.Fields{"id": ID, "op": "UpdateOemHuaweiBoards", "error": err}).Warn("DB opertion failed.")
		return err
	}
	return nil
}

func (impl *Server) deleteNetworkAdapters(c *gorm.DB, server *entity.Server) error {
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
	return nil
}

// UpdateNetworkAdapters Update network adapters.
func (impl *Server) UpdateNetworkAdapters(ID string, networkAdapters []model.NetworkAdapter) error {
	var (
		c      = impl.TemplateImpl.GetConnection()
		server = new(entity.Server)
	)

	notFound := c.Where("\"ID\" = ?", ID).
		Preload("NetworkAdapters").
		Preload("NetworkAdapters.Controllers").
		Preload("NetworkAdapters.Controllers.NetworkPorts").
		First(server).
		RecordNotFound()
	if notFound {
		return fmt.Errorf("Can't find server %s", ID)
	}
	impl.deleteNetworkAdapters(c, server)
	networkAdaptersE := []entity.NetworkAdapter{}
	for _, v := range networkAdapters {
		each := entity.NetworkAdapter{}
		each.Load(&v)
		networkAdaptersE = append(networkAdaptersE, each)
	}
	server.NetworkAdapters = networkAdaptersE
	if err := c.Save(server).Error; err != nil {
		log.WithFields(log.Fields{"id": ID, "op": "UpdateNetworkAdapters", "error": err}).Warn("DB opertion failed.")
		return err
	}
	return nil
}

func (impl *Server) deleteDrives(c *gorm.DB, server *entity.Server) error {
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
	return nil
}

// UpdateDrives Update drives.
func (impl *Server) UpdateDrives(ID string, drives []model.Drive) error {
	var (
		c      = impl.TemplateImpl.GetConnection()
		server = new(entity.Server)
	)
	notFound := c.Where("\"ID\" = ?", ID).
		Preload("Drives").
		Preload("Drives.Location").
		Preload("Drives.Location.PostalAddress").
		Preload("Drives.Location.Placement").
		First(server).
		RecordNotFound()
	if notFound {
		return fmt.Errorf("Can't find server %s", ID)
	}
	impl.deleteDrives(c, server)
	drivesE := []entity.Drive{}
	for _, v := range drives {
		each := entity.Drive{}
		each.Load(&v)
		drivesE = append(drivesE, each)
	}
	server.Drives = drivesE
	if err := c.Save(server).Error; err != nil {
		log.WithFields(log.Fields{"id": ID, "op": "UpdateDrives", "error": err}).Warn("DB opertion failed.")
		return err
	}
	return nil
}

func (impl *Server) deletePCIeDevices(c *gorm.DB, server *entity.Server) error {
	for i := range server.PCIeDevices {
		pcieDevice := server.PCIeDevices[i]
		for j := range pcieDevice.PCIeFunctions {
			c.Delete(&pcieDevice.PCIeFunctions[j])
		}
		c.Delete(&pcieDevice)
	}
	return nil
}

// UpdatePCIeDevices Updagte PCIe devices.
func (impl *Server) UpdatePCIeDevices(ID string, pcieDevices []model.PCIeDevice) error {
	var (
		c      = impl.TemplateImpl.GetConnection()
		server = new(entity.Server)
	)
	notFound := c.Where("\"ID\" = ?", ID).
		Preload("PCIeDevices").
		Preload("PCIeDevices.PCIeFunctions").
		First(server).
		RecordNotFound()
	if notFound {
		return fmt.Errorf("Can't find server %s", ID)
	}
	impl.deletePCIeDevices(c, server)
	pcieDevicesE := new([]entity.PCIeDevice)
	for _, v := range pcieDevices {
		each := entity.PCIeDevice{}
		each.Load(&v)
		*pcieDevicesE = append(*pcieDevicesE, each)
	}
	server.PCIeDevices = *pcieDevicesE
	if err := c.Save(server).Error; err != nil {
		log.WithFields(log.Fields{"id": ID, "op": "UpdatePCIeDevices", "error": err}).Warn("DB opertion failed.")
		return err
	}
	return nil
}
