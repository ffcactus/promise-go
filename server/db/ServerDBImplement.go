package db

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"promise/common/category"
	commonDB "promise/common/db"
	commonConstError "promise/common/object/constError"
	commonUtil "promise/common/util"
	"promise/server/object/constValue"
	"promise/server/object/entity"
	"promise/server/object/model"
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
		return true, server.ToModel()
	}
	return false, nil
}

// PostServer will save the server to the DB.
// Since a server should belongs to the default servergroup, we need create server and servergroup
// in DB at the same time, we also need return them both back for event dispatch.
func (i *ServerDBImplement) PostServer(s *model.Server) (*model.Server, *model.ServerServerGroup, error) {
	c := commonDB.GetConnection()
	if exist, _ := i.IsServerExist(s); exist {
		// TODO get the server to return.
		log.WithFields(log.Fields{"hostname": s.Hostname}).Info("Post server in DB failed, server exist.")
		return nil, nil, errors.New("server exist")
	}
	server := new(entity.Server)
	server.Load(s)
	var serverServerGroup = new(entity.ServerServerGroup)

	// Generate the UUID.
	server.ID = uuid.New().String()
	serverServerGroup.ID = uuid.New().String()
	serverServerGroup.Category = category.ServerServerGroup
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

// GetServer Get server by ID.
func (i *ServerDBImplement) GetServer(ID string) *model.Server {
	c := commonDB.GetConnection()
	var s = new(entity.Server)
	c.Where("\"ID\" = ?", ID).First(s)
	if s.ID != ID {
		return nil
	}
	return s.ToModel()
}

// GetServerCollection Get server collection by start and count.
func (i *ServerDBImplement) GetServerCollection(start int, count int) (*model.ServerCollection, error) {
	var (
		total  int
		server []entity.Server
		ret    = new(model.ServerCollection)
	)

	c := commonDB.GetConnection()
	c.Table("server").Count(&total)
	c.Order("Name asc").Limit(count).Offset(start).Select([]string{"ID", "Name", "State", "Health"}).Find(&server)
	ret.Start = start
	ret.Count = len(server)
	ret.Total = total
	for i := range server {
		ret.Members = append(ret.Members, model.ServerMember{
			ID:     server[i].ID,
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
	c.Where("\"ID\" = ?", ID).First(s)
	if s.ID != ID {
		return nil
	}
	c.Where("\"ID\" = ?", ID).
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
	return s.ToModel()
}

// FindServerStateAdded Find the first server with state "Added" and return the server ID.
func (i *ServerDBImplement) FindServerStateAdded() string {
	c := commonDB.GetConnection()
	var s = new(entity.Server)
	if notFound := c.Where("State = ?", constValue.ServerStateAdded).First(s).RecordNotFound(); notFound {
		return ""
	}
	return s.ID
}

// DeleteServer will delete server by ID.
// Since we also need to delete the server-servergroup association, we need transaction here.
// This function will return the if the server exist, deleted server, the slice of deleted server-servergroup, whether operation commited and error if any.
func (i *ServerDBImplement) DeleteServer(id string) (bool, *model.Server, []model.ServerServerGroup, bool, error) {
	var (
		s = new(entity.Server)
		deletedServer = new(model.Server)
		deletedSSG = make([]model.ServerServerGroup, 0)
		deletedSSGEntity = make([]entity.ServerServerGroup, 0)
	)

	c := commonDB.GetConnection()
	tx := c.Begin()
	if err := tx.Error; err != nil {
		log.WithFields(log.Fields{
			"error": err}).
			Warn("Delete server in DB failed, start transaction failed.")
		// Assume server exist.
		return true, nil, nil, false, err
	}
	// Check if the server exist.
	if tx.Where("\"ID\" = ?", id).First(s).RecordNotFound() {
		tx.Rollback()
		log.WithFields(log.Fields{"id": id}).
			Debug("Delete server from DB failed, server does not exist, transaction rollback.")
		return false, nil, nil, false, commonConstError.ErrorResourceNotExist
	}
	// Load full server info so we can delete them all together.
	if err := tx.Where("\"ID\" = ?", id).
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
		Find(s).Error; err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{"id": id}).
			Debug("Delete server in DB failed, can not load full server info, transaction rollback.")
		return true, nil, nil, false, err
	}
	deletedServer = s.ToModel()
	// Delete all.
	i.deleteProcessors(tx, s)
	i.deleteMemory(tx, s)
	i.deleteEthernetInterfaces(tx, s)
	i.deleteNetworkInterfaces(tx, s)
	i.deleteStorages(tx, s)
	i.deletePower(tx, s)
	i.deleteThermal(tx, s)
	i.deleteOemHuaweiBoards(tx, s)
	i.deleteNetworkAdapters(tx, s)
	i.deleteDrives(tx, s)
	i.deletePCIeDevices(tx, s)
	if err := tx.Delete(s).Error; err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{"id": id}).
			Debug("Delete server in DB failed, delete resource failed, transaction rollback.")
		return true, nil, nil, false, err
	}
	

	// Delete the server-servergroup association.
	// But we need record them first.
	if err := tx.Where("\"ServerID\" = ?", id).Find(&deletedSSGEntity).Error; err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{"id": id}).
			Warn("Delete server in DB failed, record server-servergroup association failed, transaction rollback.")
	}
	for _, each := range deletedSSGEntity {
		deletedSSG = append(deletedSSG, *each.ToModel())
	}
	if err := tx.Where("\"ServerID\" = ?", id).Delete(entity.ServerServerGroup{}).Error; err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{"id": id}).
			Warn("Delete server in DB failed, delete server-servergroup association failed, transaction rollback.")
		return true, nil, nil, false, err
	}
	// Commit.
	if err := tx.Commit().Error; err != nil {
		log.WithFields(log.Fields{
			"error": err}).
			Warn("Delete server in DB failed, commit failed.")
		return true, nil, nil, false, err
	}
	return true, deletedServer, deletedSSG, true, nil
}

// DeleteServerCollection will delete all the server from DB.
// Since there are many tables, we need transaction here.
func (i *ServerDBImplement) DeleteServerCollection() error {
	c := commonDB.GetConnection()
	tx := c.Begin()
	if err := tx.Error; err != nil {
		log.WithFields(log.Fields{
			"error": err}).
			Warn("Delete server collection in DB failed, start transaction failed.")
		return err
	}
	for i := range entity.ServerTables {
		if err := tx.Delete(entity.ServerTables[i].Info).Error; err != nil {
			tx.Rollback()
			log.WithFields(log.Fields{
				"table": entity.ServerTables[i].Name,
				"error": err}).
				Warn("Delete server collection in DB failed, delete resources failed, transaction rollback.")
			return err
		}
	}
	// When we delete all the servers we also need delete all the server-servergroup.
	if err := tx.Delete(entity.ServerServerGroup{}).Error; err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"error": err}).
			Warn("Delete server collection in DB failed, delete server-servergroup collection failed, transaction rollback.")
		return err
	}
	if err := tx.Commit().Error; err != nil {
		log.WithFields(log.Fields{
			"error": err}).
			Warn("Delete server collection in DB failed, commit failed.")
		return err
	}
	return nil
}

// GetAndLockServer Get and lock the server.
func (i *ServerDBImplement) GetAndLockServer(ID string) (bool, *model.Server) {
	c := commonDB.GetConnection()
	// Transaction start.
	tx := c.Begin()
	if err := tx.Error; err != nil {
		log.WithFields(log.Fields{
			"id":    ID,
			"error": err}).
			Warn("Get and lock server in DB failed, start transaction failed.")
		return false, nil
	}
	var s = new(entity.Server)
	if tx.Where("\"ID\" = ?", ID).First(s).RecordNotFound() {
		tx.Rollback()
		log.WithFields(log.Fields{
			"id": ID}).
			Debug("Get and lock server in DB failed, server does not exist.")
		return false, nil
	}
	if !constValue.ServerLockable(s.State) {
		// Server not ready, rollback.
		tx.Rollback()
		log.WithFields(log.Fields{
			"id":    ID,
			"state": s.State}).
			Debug("Get and lock server in DB failed, server not lockable.")
		return false, s.ToModel()
	}
	// Change the state.
	if err := tx.Model(s).UpdateColumn("State", constValue.ServerStateLocked).Error; err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"id":    ID,
			"state": s.State}).
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
	return true, s.ToModel()
}

// SetServerState Set server state.
func (i *ServerDBImplement) SetServerState(ID string, state string) bool {
	c := commonDB.GetConnection()
	var s = new(entity.Server)
	if c.Where("\"ID\" = ?", ID).First(s).RecordNotFound() {
		return false
	}
	if err := c.Model(s).UpdateColumn("State", state).Error; err != nil {
		log.WithFields(log.Fields{"id": ID, "op": "SetServerState", "error": err}).Warn("DB opertion failed.")
		return false
	}
	return true
}

// SetServerHealth Set server health.
func (i *ServerDBImplement) SetServerHealth(ID string, health string) bool {
	c := commonDB.GetConnection()
	var s = new(entity.Server)
	if c.Where("\"ID\" = ?", ID).First(s).RecordNotFound() {
		return false
	}
	if err := c.Model(s).UpdateColumn("Health", health).Error; err != nil {
		log.WithFields(log.Fields{"id": ID, "op": "SetServerHealth", "error": err}).Warn("DB opertion failed.")
		return false
	}
	return true
}

func (i *ServerDBImplement) deleteProcessors(c *gorm.DB, server *entity.Server) error {
	for i := range server.Processors {
		c.Delete(server.Processors[i])
	}
	return nil
}

// UpdateProcessors Update processors info.
func (i *ServerDBImplement) UpdateProcessors(ID string, processors []model.Processor) error {
	server := new(entity.Server)

	c := commonDB.GetConnection()
	notFound := c.Where("\"ID\" = ?", ID).
		Preload("Processors").
		First(server).
		RecordNotFound()
	if notFound {
		return fmt.Errorf("can not find server %s", ID)
	}
	i.deleteProcessors(c, server)
	server.Processors = []entity.Processor{}
	for i := range processors {
		server.Processors = append(server.Processors, *createProcessor(&processors[i]))
	}
	if err := c.Save(server).Error; err != nil {
		log.WithFields(log.Fields{"id": ID, "op": "UpdateProcessors", "error": err}).Warn("DB opertion failed.")
		return err
	}
	return nil
}

func (i *ServerDBImplement) deleteMemory(c *gorm.DB, server *entity.Server) error {
	for i := range server.Memory {
		c.Delete(server.Memory[i])
	}
	return nil
}

// UpdateMemory Update server memory
func (i *ServerDBImplement) UpdateMemory(ID string, memory []model.Memory) error {
	server := new(entity.Server)

	c := commonDB.GetConnection()
	notFound := c.Where("\"ID\" = ?", ID).
		Preload("Memory").
		First(server).
		RecordNotFound()
	if notFound {
		return fmt.Errorf("Can't find server %s", ID)
	}
	i.deleteMemory(c, server)
	server.Memory = []entity.Memory{}
	for i := range memory {
		server.Memory = append(server.Memory, *createMemory(&memory[i]))
	}
	if err := c.Save(server).Error; err != nil {
		log.WithFields(log.Fields{"id": ID, "op": "UpdateMemory", "error": err}).Warn("DB opertion failed.")
		return err
	}
	return nil
}

func (i *ServerDBImplement) deleteEthernetInterfaces(c *gorm.DB, server *entity.Server) error {
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
func (i *ServerDBImplement) UpdateEthernetInterfaces(ID string, ethernet []model.EthernetInterface) error {
	server := new(entity.Server)
	c := commonDB.GetConnection()
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
	i.deleteEthernetInterfaces(c, server)
	// Regenerate ethernet interface.
	server.EthernetInterfaces = []entity.EthernetInterface{}
	for i := range ethernet {
		server.EthernetInterfaces = append(server.EthernetInterfaces, *createEthernetInterface(&ethernet[i]))
	}
	if err := c.Save(server).Error; err != nil {
		log.WithFields(log.Fields{"id": ID, "op": "UpdateEthernetInterfaces", "error": err}).Warn("DB opertion failed.")
		return err
	}
	return nil
}

func (i *ServerDBImplement) deleteNetworkInterfaces(c *gorm.DB, server *entity.Server) error {
	for i := range server.NetworkInterfaces {
		c.Delete(server.NetworkInterfaces[i])
	}
	return nil
}

// UpdateNetworkInterfaces Update network interfaces.
func (i *ServerDBImplement) UpdateNetworkInterfaces(ID string, networkInterface []model.NetworkInterface) error {
	server := new(entity.Server)
	c := commonDB.GetConnection()
	// Get resources.
	notFound := c.Where("\"ID\" = ?", ID).
		Preload("NetworkInterfaces").
		First(server).
		RecordNotFound()
	if notFound {
		return fmt.Errorf("Can't find server %s", ID)
	}
	// Delete them.
	i.deleteEthernetInterfaces(c, server)
	networkInterfacesE := []entity.NetworkInterface{}
	for i := range networkInterface {
		networkInterfacesE = append(networkInterfacesE, *createNetworkInterface(&networkInterface[i]))
	}
	server.NetworkInterfaces = networkInterfacesE
	if err := c.Save(server).Error; err != nil {
		log.WithFields(log.Fields{"id": ID, "op": "UpdateNetworkInterfaces", "error": err}).Warn("DB opertion failed.")
		return err
	}
	return nil
}

func (i *ServerDBImplement) deleteStorages(c *gorm.DB, server *entity.Server) error {
	for i := range server.Storages {
		for j := range server.Storages[i].StorageControllers {
			c.Delete(&(server.Storages[i].StorageControllers)[j])
		}
		c.Delete(server.Storages[i])
	}
	return nil
}

// UpdateStorages Update storages.
func (i *ServerDBImplement) UpdateStorages(ID string, storages []model.Storage) error {
	server := new(entity.Server)
	c := commonDB.GetConnection()
	notFound := c.Where("\"ID\" = ?", ID).
		Preload("Storages").
		Preload("Storages.StorageControllers").
		First(server).
		RecordNotFound()
	if notFound {
		return fmt.Errorf("Can't find server %s", ID)
	}
	i.deleteEthernetInterfaces(c, server)
	storagesE := []entity.Storage{}
	for i := range storages {
		storagesE = append(storagesE, *createStorage(&storages[i]))
	}
	server.Storages = storagesE
	if err := c.Save(server).Error; err != nil {
		log.WithFields(log.Fields{"id": ID, "op": "UpdateStorages", "error": err}).Warn("DB opertion failed.")
		return err
	}
	return nil
}

func (i *ServerDBImplement) deletePower(c *gorm.DB, server *entity.Server) error {
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
func (i *ServerDBImplement) UpdatePower(ID string, power *model.Power) error {
	server := new(entity.Server)
	c := commonDB.GetConnection()
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
	i.deletePower(c, server)
	server.Power = *createPower(power)
	if err := c.Save(server).Error; err != nil {
		log.WithFields(log.Fields{"id": ID, "op": "UpdatePower", "error": err}).Warn("DB opertion failed.")
		return err
	}
	return nil
}

func (i *ServerDBImplement) deleteThermal(c *gorm.DB, server *entity.Server) error {
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
func (i *ServerDBImplement) UpdateThermal(ID string, thermal *model.Thermal) error {
	server := new(entity.Server)
	c := commonDB.GetConnection()
	notFound := c.Where("\"ID\" = ?", ID).
		Preload("Thermal").
		Preload("Thermal.Temperatures").
		Preload("Thermal.Fans").
		First(server).
		RecordNotFound()
	if notFound {
		return fmt.Errorf("Can't find server %s", ID)
	}
	i.deleteThermal(c, server)
	server.Thermal = *createThermal(thermal)
	if err := c.Save(server).Error; err != nil {
		log.WithFields(log.Fields{"id": ID, "op": "UpdateThermal", "error": err}).Warn("DB opertion failed.")
		return err
	}
	return nil
}

func (i *ServerDBImplement) deleteOemHuaweiBoards(c *gorm.DB, server *entity.Server) error {
	for i := range server.OemHuaweiBoards {
		c.Delete(server.OemHuaweiBoards[i])
	}
	return nil
}

// UpdateOemHuaweiBoards Update OEM Huawei boards.
func (i *ServerDBImplement) UpdateOemHuaweiBoards(ID string, boards []model.OemHuaweiBoard) error {
	c := commonDB.GetConnection()
	var server = new(entity.Server)
	notFound := c.Where("\"ID\" = ?", ID).
		Preload("OemHuaweiBoards").
		First(server).
		RecordNotFound()
	if notFound {
		return fmt.Errorf("Can't find server %s", ID)
	}
	i.deleteOemHuaweiBoards(c, server)
	boardsE := []entity.OemHuaweiBoard{}
	for i := range boards {
		boardsE = append(boardsE, *createOemHuaweiBoard(&boards[i]))
	}
	server.OemHuaweiBoards = boardsE
	if err := c.Save(server).Error; err != nil {
		log.WithFields(log.Fields{"id": ID, "op": "UpdateOemHuaweiBoards", "error": err}).Warn("DB opertion failed.")
		return err
	}
	return nil
}

func (i *ServerDBImplement) deleteNetworkAdapters(c *gorm.DB, server *entity.Server) error {
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
func (i *ServerDBImplement) UpdateNetworkAdapters(ID string, networkAdapters []model.NetworkAdapter) error {
	c := commonDB.GetConnection()
	var server = new(entity.Server)
	notFound := c.Where("\"ID\" = ?", ID).
		Preload("NetworkAdapters").
		Preload("NetworkAdapters.Controllers").
		Preload("NetworkAdapters.Controllers.NetworkPorts").
		First(server).
		RecordNotFound()
	if notFound {
		return fmt.Errorf("Can't find server %s", ID)
	}
	i.deleteNetworkAdapters(c, server)
	networkAdaptersE := []entity.NetworkAdapter{}
	for i := range networkAdapters {
		networkAdaptersE = append(networkAdaptersE, *createNetworkAdapter(&networkAdapters[i]))
	}
	server.NetworkAdapters = networkAdaptersE
	if err := c.Save(server).Error; err != nil {
		log.WithFields(log.Fields{"id": ID, "op": "UpdateNetworkAdapters", "error": err}).Warn("DB opertion failed.")
		return err
	}
	return nil
}

func (i *ServerDBImplement) deleteDrives(c *gorm.DB, server *entity.Server) error {
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
func (i *ServerDBImplement) UpdateDrives(ID string, drives []model.Drive) error {
	c := commonDB.GetConnection()
	var server = new(entity.Server)
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
	i.deleteDrives(c, server)
	drivesE := []entity.Drive{}
	for i := range drives {
		drivesE = append(drivesE, *createDrive(&drives[i]))
	}
	server.Drives = drivesE
	if err := c.Save(server).Error; err != nil {
		log.WithFields(log.Fields{"id": ID, "op": "UpdateDrives", "error": err}).Warn("DB opertion failed.")
		return err
	}
	return nil
}

func (i *ServerDBImplement) deletePCIeDevices(c *gorm.DB, server *entity.Server) error {
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
func (i *ServerDBImplement) UpdatePCIeDevices(ID string, pcieDevices []model.PCIeDevice) error {
	c := commonDB.GetConnection()
	var server = new(entity.Server)
	notFound := c.Where("\"ID\" = ?", ID).
		Preload("PCIeDevices").
		Preload("PCIeDevices.PCIeFunctions").
		First(server).
		RecordNotFound()
	if notFound {
		return fmt.Errorf("Can't find server %s", ID)
	}
	i.deletePCIeDevices(c, server)
	pcieDevicesE := new([]entity.PCIeDevice)
	for i := range pcieDevices {
		*pcieDevicesE = append(*pcieDevicesE, *createPCIeDevice(&pcieDevices[i]))
	}
	server.PCIeDevices = *pcieDevicesE
	if err := c.Save(server).Error; err != nil {
		log.WithFields(log.Fields{"id": ID, "op": "UpdatePCIeDevices", "error": err}).Warn("DB opertion failed.")
		return err
	}
	return nil
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
