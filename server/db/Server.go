package db

import (
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
	return "server"
}

// NewEntity return the a new entity.
func (impl *Server) NewEntity() base.EntityInterface {
	return new(entity.Server)
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
func (impl *Server) ConvertFindResultToCollection(start int64, total int64, result interface{}) (*base.CollectionModel, *base.Message) {
	collection, ok := result.(*[]entity.Server)
	if !ok {
		log.Error("Server.ConvertFindResult() failed, convert data failed.")
		return nil, base.NewMessageInternalError()
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
func (impl *Server) ConvertFindResultToModel(result interface{}) ([]base.ModelInterface, *base.Message) {
	collection, ok := result.(*[]entity.Server)
	if !ok {
		log.Error("Server.ConvertFindResult() failed, convert data failed.")
		return nil, base.NewMessageInternalError()
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
	if notFound := c.Where("\"State\" = ?", constvalue.ServerStateAdded).Order("\"Name\" asc").First(server).RecordNotFound(); notFound {
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

	c.Where("\"PhysicalUUID\" = ?", s.PhysicalUUID).First(server)
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
	tx.Exec("SET TRANSACTION ISOLATION LEVEL SERIALIZABLE;")
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
func (impl *Server) SetServerState(ID string, state string) (base.ModelInterface, error) {
	var (
		c      = impl.TemplateImpl.GetConnection()
		server = new(entity.Server)
	)
	tx := c.Begin()
	if err := tx.Error; err != nil {
		log.WithFields(log.Fields{
			"id":    ID,
			"op":    "SetServerState",
			"error": err,
		}).Warn("DB operation failed , start transaction failed.")
		return nil, base.ErrorTransaction
	}
	found, err := impl.GetInternal(tx, ID, server)
	if !found || err != nil {
		log.WithFields(log.Fields{
			"id": ID,
			"op": "SetServerState",
		}).Warn("DB operation failed , load server failed.")
		return nil, base.ErrorTransaction
	}
	if err := tx.Model(server).UpdateColumn("State", state).Error; err != nil {
		log.WithFields(log.Fields{
			"id":    ID,
			"op":    "SetServerState",
			"error": err,
		}).Warn("DB opertion failed, update server failed, transaction rollback.")
		return nil, base.ErrorTransaction
	}
	if err := tx.Commit().Error; err != nil {
		log.WithFields(log.Fields{
			"id":    ID,
			"op":    "SetServerState",
			"error": err,
		}).Warn("DB opertion failed, commit failed.")
		return nil, base.ErrorTransaction
	}
	return server.ToModel(), nil
}

// SetServerHealth Set server health.
func (impl *Server) SetServerHealth(ID string, health string) (base.ModelInterface, error) {
	var (
		c      = impl.TemplateImpl.GetConnection()
		server = new(entity.Server)
	)
	tx := c.Begin()
	if err := tx.Error; err != nil {
		log.WithFields(log.Fields{
			"id":    ID,
			"op":    "SetServerHealth",
			"error": err,
		}).Warn("DB operation failed , start transaction failed.")
		return nil, base.ErrorTransaction
	}
	found, err := impl.GetInternal(tx, ID, server)
	if !found || err != nil {
		log.WithFields(log.Fields{
			"id": ID,
			"op": "SetServerHealth",
		}).Warn("DB operation failed , load server failed.")
		return nil, base.ErrorTransaction
	}
	if err := tx.Model(server).UpdateColumn("Health", health).Error; err != nil {
		log.WithFields(log.Fields{
			"id":    ID,
			"op":    "SetServerHealth",
			"error": err,
		}).Warn("DB opertion failed, update server failed, transaction rollback.")
		return nil, base.ErrorTransaction
	}
	if err := tx.Commit().Error; err != nil {
		log.WithFields(log.Fields{
			"id":    ID,
			"op":    "SetServerHealth",
			"error": err,
		}).Warn("DB opertion failed, commit failed.")
		return nil, base.ErrorTransaction
	}
	return server.ToModel(), nil
}

func (impl *Server) deleteProcessors(c *gorm.DB, server *entity.Server) error {
	for i := range server.Processors {
		if err := c.Delete(server.Processors[i]).Error; err != nil {
			return base.ErrorTransaction
		}
	}
	return nil
}

// UpdateProcessors Update processors info.
func (impl *Server) UpdateProcessors(ID string, processors []model.Processor) (base.ModelInterface, error) {
	var (
		c      = impl.TemplateImpl.GetConnection()
		server = new(entity.Server)
	)
	tx := c.Begin()
	if err := tx.Error; err != nil {
		log.WithFields(log.Fields{
			"id":    ID,
			"op":    "UpdateProcessors",
			"error": err,
		}).Warn("DB operation failed , start transaction failed.")
		return nil, base.ErrorTransaction
	}
	found, err := impl.GetInternal(tx, ID, server)
	if !found || err != nil {
		log.WithFields(log.Fields{
			"id": ID,
			"op": "UpdateProcessors",
		}).Warn("DB operation failed , load server failed.")
		return nil, base.ErrorTransaction
	}
	if err := impl.deleteProcessors(tx, server); err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"id":    ID,
			"op":    "UpdateProcessors",
			"error": err,
		}).Warn("DB operation failed , delete association failed, transaction rollback.")
		return nil, base.ErrorTransaction
	}
	server.Processors = []entity.Processor{}
	for _, v := range processors {
		each := entity.Processor{}
		each.Load(&v)
		server.Processors = append(server.Processors, each)
	}
	if err := tx.Save(server).Error; err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"id":    ID,
			"op":    "UpdateProcessors",
			"error": err,
		}).Warn("DB opertion failed, save server failed, transaction rollback.")
		return nil, base.ErrorTransaction
	}
	if err := tx.Commit().Error; err != nil {
		log.WithFields(log.Fields{
			"id":    ID,
			"op":    "UpdateProcessors",
			"error": err,
		}).Warn("DB opertion failed, commit failed.")
		return nil, base.ErrorTransaction
	}
	return server.ToModel(), nil
}

func (impl *Server) deleteMemory(c *gorm.DB, server *entity.Server) error {
	for i := range server.Memory {
		if err := c.Delete(server.Memory[i]).Error; err != nil {
			return base.ErrorTransaction
		}
	}
	return nil
}

// UpdateMemory Update server memory
func (impl *Server) UpdateMemory(ID string, memory []model.Memory) (base.ModelInterface, error) {
	var (
		c      = impl.TemplateImpl.GetConnection()
		server = new(entity.Server)
	)
	tx := c.Begin()
	if err := tx.Error; err != nil {
		log.WithFields(log.Fields{
			"id":    ID,
			"op":    "UpdateMemory",
			"error": err,
		}).Warn("DB operation failed , start transaction failed.")
		return nil, base.ErrorTransaction
	}
	found, err := impl.GetInternal(tx, ID, server)
	if !found || err != nil {
		log.WithFields(log.Fields{
			"id": ID,
			"op": "UpdateMemory",
		}).Warn("DB operation failed , load server failed.")
		return nil, base.ErrorTransaction
	}
	if err := impl.deleteMemory(tx, server); err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"id":    ID,
			"op":    "UpdateMemory",
			"error": err,
		}).Warn("DB operation failed , delete association failed, transaction rollback.")
		return nil, base.ErrorTransaction
	}
	server.Memory = []entity.Memory{}
	for _, v := range memory {
		each := entity.Memory{}
		each.Load(&v)
		server.Memory = append(server.Memory, each)
	}
	if err := tx.Save(server).Error; err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"id":    ID,
			"op":    "UpdateMemory",
			"error": err,
		}).Warn("DB opertion failed, save server failed, transaction rollback.")
		return nil, base.ErrorTransaction
	}
	if err := tx.Commit().Error; err != nil {
		log.WithFields(log.Fields{
			"id":    ID,
			"op":    "UpdateMemory",
			"error": err,
		}).Warn("DB opertion failed, commit failed.")
		return nil, base.ErrorTransaction
	}
	return server.ToModel(), nil
}

func (impl *Server) deleteEthernetInterfaces(c *gorm.DB, server *entity.Server) error {
	for i := range server.EthernetInterfaces {
		for k := range server.EthernetInterfaces[i].IPv4Addresses {
			if err := c.Delete(&(server.EthernetInterfaces[i].IPv4Addresses)[k]).Error; err != nil {
				return base.ErrorTransaction
			}
		}
		for k := range server.EthernetInterfaces[i].IPv6Addresses {
			if err := c.Delete(&(server.EthernetInterfaces[i].IPv6Addresses)[k]).Error; err != nil {
				return base.ErrorTransaction
			}
		}
		for k := range server.EthernetInterfaces[i].VLANs {
			if err := c.Delete(&(server.EthernetInterfaces[i].VLANs)[k]).Error; err != nil {
				return base.ErrorTransaction
			}
		}
		if err := c.Delete(server.EthernetInterfaces[i]).Error; err != nil {
			return base.ErrorTransaction
		}
	}
	return nil
}

// UpdateEthernetInterfaces Update server ethernet interface.
// Each ethernet interface has many IPv4, IPv6 and vLAN, so it's very hard to check if 2 ethernet interface
// are the same. So we just remove all of them and recreate them.
func (impl *Server) UpdateEthernetInterfaces(ID string, ethernet []model.EthernetInterface) (base.ModelInterface, error) {
	var (
		c      = impl.TemplateImpl.GetConnection()
		server = new(entity.Server)
	)
	tx := c.Begin()
	if err := tx.Error; err != nil {
		log.WithFields(log.Fields{
			"id":    ID,
			"op":    "UpdateEthernetInterfaces",
			"error": err,
		}).Warn("DB operation failed , start transaction failed.")
		return nil, base.ErrorTransaction
	}
	found, err := impl.GetInternal(tx, ID, server)
	if !found || err != nil {
		log.WithFields(log.Fields{
			"id": ID,
			"op": "UpdateEthernetInterfaces",
		}).Warn("DB operation failed , load server failed.")
		return nil, base.ErrorTransaction
	}
	if err := impl.deleteEthernetInterfaces(tx, server); err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"id":    ID,
			"op":    "UpdateEthernetInterfaces",
			"error": err,
		}).Warn("DB operation failed , delete association failed, transaction rollback.")
		return nil, base.ErrorTransaction
	}
	server.EthernetInterfaces = []entity.EthernetInterface{}
	for _, v := range ethernet {
		each := entity.EthernetInterface{}
		each.Load(&v)
		server.EthernetInterfaces = append(server.EthernetInterfaces, each)
	}
	if err := tx.Save(server).Error; err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"id":    ID,
			"op":    "UpdateEthernetInterfaces",
			"error": err,
		}).Warn("DB opertion failed, save server failed, transaction rollback.")
		return nil, base.ErrorTransaction
	}
	if err := tx.Commit().Error; err != nil {
		log.WithFields(log.Fields{
			"id":    ID,
			"op":    "UpdateEthernetInterfaces",
			"error": err,
		}).Warn("DB opertion failed, commit failed.")
		return nil, base.ErrorTransaction
	}
	return server.ToModel(), nil
}

func (impl *Server) deleteNetworkInterfaces(c *gorm.DB, server *entity.Server) error {
	for i := range server.NetworkInterfaces {
		if err := c.Delete(server.NetworkInterfaces[i]).Error; err != nil {
			return base.ErrorTransaction
		}
	}
	return nil
}

// UpdateNetworkInterfaces Update network interfaces.
func (impl *Server) UpdateNetworkInterfaces(ID string, networkInterface []model.NetworkInterface) (base.ModelInterface, error) {
	var (
		c      = impl.TemplateImpl.GetConnection()
		server = new(entity.Server)
	)
	tx := c.Begin()
	if err := tx.Error; err != nil {
		log.WithFields(log.Fields{
			"id":    ID,
			"op":    "UpdateNetworkInterfaces",
			"error": err,
		}).Warn("DB operation failed , start transaction failed.")
		return nil, base.ErrorTransaction
	}
	found, err := impl.GetInternal(tx, ID, server)
	if !found || err != nil {
		log.WithFields(log.Fields{
			"id": ID,
			"op": "UpdateNetworkInterfaces",
		}).Warn("DB operation failed , load server failed.")
		return nil, base.ErrorTransaction
	}
	if err := impl.deleteNetworkInterfaces(tx, server); err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"id":    ID,
			"op":    "UpdateNetworkInterfaces",
			"error": err,
		}).Warn("DB operation failed , delete association failed, transaction rollback.")
		return nil, base.ErrorTransaction
	}
	networkInterfacesE := []entity.NetworkInterface{}
	for _, v := range networkInterface {
		each := entity.NetworkInterface{}
		each.Load(&v)
		networkInterfacesE = append(networkInterfacesE, each)
	}
	server.NetworkInterfaces = networkInterfacesE
	if err := tx.Save(server).Error; err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"id":    ID,
			"op":    "UpdateNetworkInterfaces",
			"error": err,
		}).Warn("DB opertion failed, save server failed, transaction rollback.")
		return nil, base.ErrorTransaction
	}
	if err := tx.Commit().Error; err != nil {
		log.WithFields(log.Fields{
			"id":    ID,
			"op":    "UpdateNetworkInterfaces",
			"error": err,
		}).Warn("DB opertion failed, commit failed.")
		return nil, base.ErrorTransaction
	}
	return server.ToModel(), nil
}

func (impl *Server) deleteStorages(c *gorm.DB, server *entity.Server) error {
	for i := range server.Storages {
		for j := range server.Storages[i].StorageControllers {
			if err := c.Delete(&(server.Storages[i].StorageControllers)[j]).Error; err != nil {
				return base.ErrorTransaction
			}
		}
		if err := c.Delete(server.Storages[i]).Error; err != nil {
			return base.ErrorTransaction
		}
	}
	return nil
}

// UpdateStorages Update storages.
func (impl *Server) UpdateStorages(ID string, storages []model.Storage) (base.ModelInterface, error) {
	var (
		c      = impl.TemplateImpl.GetConnection()
		server = new(entity.Server)
	)
	tx := c.Begin()
	if err := tx.Error; err != nil {
		log.WithFields(log.Fields{
			"id":    ID,
			"op":    "UpdateStorages",
			"error": err,
		}).Warn("DB operation failed , start transaction failed.")
		return nil, base.ErrorTransaction
	}
	found, err := impl.GetInternal(tx, ID, server)
	if !found || err != nil {
		log.WithFields(log.Fields{
			"id": ID,
			"op": "UpdateStorages",
		}).Warn("DB operation failed , load server failed.")
		return nil, base.ErrorTransaction
	}
	if err := impl.deleteStorages(tx, server); err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"id":    ID,
			"op":    "UpdateStorages",
			"error": err,
		}).Warn("DB operation failed , delete association failed, transaction rollback.")
		return nil, base.ErrorTransaction
	}
	storagesE := []entity.Storage{}
	for _, v := range storages {
		each := entity.Storage{}
		each.Load(&v)
		storagesE = append(storagesE, each)
	}
	server.Storages = storagesE
	if err := tx.Save(server).Error; err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"id":    ID,
			"op":    "UpdateStorages",
			"error": err,
		}).Warn("DB opertion failed, save server failed, transaction rollback.")
		return nil, base.ErrorTransaction
	}
	if err := tx.Commit().Error; err != nil {
		log.WithFields(log.Fields{
			"id":    ID,
			"op":    "UpdateStorages",
			"error": err,
		}).Warn("DB opertion failed, commit failed.")
		return nil, base.ErrorTransaction
	}
	return server.ToModel(), nil
}

func (impl *Server) deletePower(c *gorm.DB, server *entity.Server) error {
	for i := range server.Power.PowerControl {
		if err := c.Delete(server.Power.PowerControl[i]).Error; err != nil {
			return base.ErrorTransaction
		}
	}
	for i := range server.Power.Voltages {
		if err := c.Delete(server.Power.Voltages[i]).Error; err != nil {
			return base.ErrorTransaction
		}
	}
	for i := range server.Power.PowerSupplies {
		if err := c.Delete(server.Power.PowerSupplies[i]).Error; err != nil {
			return base.ErrorTransaction
		}
	}
	for i := range server.Power.Redundancy {
		if err := c.Delete(server.Power.Redundancy[i]).Error; err != nil {
			return base.ErrorTransaction
		}
	}
	if err := c.Delete(server.Power).Error; err != nil {
		return base.ErrorTransaction
	}
	return nil
}

// UpdatePower Update power
func (impl *Server) UpdatePower(ID string, power *model.Power) (base.ModelInterface, error) {
	var (
		c      = impl.TemplateImpl.GetConnection()
		server = new(entity.Server)
	)
	tx := c.Begin()
	if err := tx.Error; err != nil {
		log.WithFields(log.Fields{
			"id":    ID,
			"op":    "UpdatePower",
			"error": err,
		}).Warn("DB operation failed , start transaction failed.")
		return nil, base.ErrorTransaction
	}
	found, err := impl.GetInternal(tx, ID, server)
	if !found || err != nil {
		log.WithFields(log.Fields{
			"id": ID,
			"op": "UpdatePower",
		}).Warn("DB operation failed , load server failed.")
		return nil, base.ErrorTransaction
	}
	if err := impl.deletePower(tx, server); err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"id":    ID,
			"op":    "UpdatePower",
			"error": err,
		}).Warn("DB operation failed , delete association failed, transaction rollback.")
		return nil, base.ErrorTransaction
	}
	server.Power.Load(power)
	if err := tx.Save(server).Error; err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"id":    ID,
			"op":    "UpdatePower",
			"error": err,
		}).Warn("DB opertion failed, save server failed, transaction rollback.")
		return nil, base.ErrorTransaction
	}
	if err := tx.Commit().Error; err != nil {
		log.WithFields(log.Fields{
			"id":    ID,
			"op":    "UpdatePower",
			"error": err,
		}).Warn("DB opertion failed, commit failed.")
		return nil, base.ErrorTransaction
	}
	return server.ToModel(), nil
}

func (impl *Server) deleteThermal(c *gorm.DB, server *entity.Server) error {
	for i := range server.Thermal.Temperatures {
		if err := c.Delete(server.Thermal.Temperatures[i]).Error; err != nil {
			return base.ErrorTransaction
		}
	}
	for i := range server.Thermal.Fans {
		if err := c.Delete(server.Thermal.Fans[i]).Error; err != nil {
			return base.ErrorTransaction
		}
	}
	if err := c.Delete(server.Thermal).Error; err != nil {
		return base.ErrorTransaction
	}
	return nil
}

// UpdateThermal Update thermal
func (impl *Server) UpdateThermal(ID string, thermal *model.Thermal) (base.ModelInterface, error) {
	var (
		c      = impl.TemplateImpl.GetConnection()
		server = new(entity.Server)
	)
	tx := c.Begin()
	if err := tx.Error; err != nil {
		log.WithFields(log.Fields{
			"id":    ID,
			"op":    "UpdateThermal",
			"error": err,
		}).Warn("DB operation failed , start transaction failed.")
		return nil, base.ErrorTransaction
	}
	found, err := impl.GetInternal(tx, ID, server)
	if !found || err != nil {
		log.WithFields(log.Fields{
			"id": ID,
			"op": "UpdateThermal",
		}).Warn("DB operation failed , load server failed.")
		return nil, base.ErrorTransaction
	}
	if err := impl.deleteThermal(tx, server); err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"id":    ID,
			"op":    "UpdateThermal",
			"error": err,
		}).Warn("DB operation failed , delete association failed, transaction rollback.")
		return nil, base.ErrorTransaction
	}
	server.Thermal.Load(thermal)
	if err := tx.Save(server).Error; err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"id":    ID,
			"op":    "UpdateThermal",
			"error": err,
		}).Warn("DB opertion failed, save server failed, transaction rollback.")
		return nil, base.ErrorTransaction
	}
	if err := tx.Commit().Error; err != nil {
		log.WithFields(log.Fields{
			"id":    ID,
			"op":    "UpdateThermal",
			"error": err,
		}).Warn("DB opertion failed, commit failed.")
		return nil, base.ErrorTransaction
	}
	return server.ToModel(), nil
}

func (impl *Server) deleteOemHuaweiBoards(c *gorm.DB, server *entity.Server) error {
	for i := range server.OemHuaweiBoards {
		if err := c.Delete(server.OemHuaweiBoards[i]).Error; err != nil {
			return base.ErrorTransaction
		}
	}
	return nil
}

// UpdateOemHuaweiBoards Update OEM Huawei boards.
func (impl *Server) UpdateOemHuaweiBoards(ID string, boards []model.OemHuaweiBoard) (base.ModelInterface, error) {
	var (
		c      = impl.TemplateImpl.GetConnection()
		server = new(entity.Server)
	)
	tx := c.Begin()
	if err := tx.Error; err != nil {
		log.WithFields(log.Fields{
			"id":    ID,
			"op":    "UpdateOemHuaweiBoards",
			"error": err,
		}).Warn("DB operation failed , start transaction failed.")
		return nil, base.ErrorTransaction
	}
	found, err := impl.GetInternal(tx, ID, server)
	if !found || err != nil {
		log.WithFields(log.Fields{
			"id": ID,
			"op": "UpdateOemHuaweiBoards",
		}).Warn("DB operation failed , load server failed.")
		return nil, base.ErrorTransaction
	}
	if err := impl.deleteOemHuaweiBoards(tx, server); err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"id":    ID,
			"op":    "UpdateOemHuaweiBoards",
			"error": err,
		}).Warn("DB operation failed , delete association failed, transaction rollback.")
		return nil, base.ErrorTransaction
	}
	boardsE := []entity.OemHuaweiBoard{}
	for _, v := range boards {
		each := entity.OemHuaweiBoard{}
		each.Load(&v)
		boardsE = append(boardsE, each)
	}
	server.OemHuaweiBoards = boardsE
	if err := tx.Save(server).Error; err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"id":    ID,
			"op":    "UpdateOemHuaweiBoards",
			"error": err,
		}).Warn("DB opertion failed, save server failed, transaction rollback.")
		return nil, base.ErrorTransaction
	}
	if err := tx.Commit().Error; err != nil {
		log.WithFields(log.Fields{
			"id":    ID,
			"op":    "UpdateOemHuaweiBoards",
			"error": err,
		}).Warn("DB opertion failed, commit failed.")
		return nil, base.ErrorTransaction
	}
	return server.ToModel(), nil
}

func (impl *Server) deleteNetworkAdapters(c *gorm.DB, server *entity.Server) error {
	for i := range server.NetworkAdapters {
		adapter := server.NetworkAdapters[i]
		for j := range adapter.Controllers {
			controller := adapter.Controllers[j]
			for k := range controller.NetworkPorts {
				if err := c.Delete(&controller.NetworkPorts[k]).Error; err != nil {
					return base.ErrorTransaction
				}
			}
			if err := c.Delete(&controller).Error; err != nil {
				return base.ErrorTransaction
			}
		}
		if err := c.Delete(&adapter).Error; err != nil {
			return base.ErrorTransaction
		}
	}
	return nil
}

// UpdateNetworkAdapters Update network adapters.
func (impl *Server) UpdateNetworkAdapters(ID string, networkAdapters []model.NetworkAdapter) (base.ModelInterface, error) {
	var (
		c      = impl.TemplateImpl.GetConnection()
		server = new(entity.Server)
	)
	tx := c.Begin()
	if err := tx.Error; err != nil {
		log.WithFields(log.Fields{
			"id":    ID,
			"op":    "UpdateNetworkAdapters",
			"error": err,
		}).Warn("DB operation failed , start transaction failed.")
		return nil, base.ErrorTransaction
	}
	found, err := impl.GetInternal(tx, ID, server)
	if !found || err != nil {
		log.WithFields(log.Fields{
			"id": ID,
			"op": "UpdateNetworkAdapters",
		}).Warn("DB operation failed , load server failed.")
		return nil, base.ErrorTransaction
	}
	if err := impl.deleteNetworkAdapters(tx, server); err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"id":    ID,
			"op":    "UpdateNetworkAdapters",
			"error": err,
		}).Warn("DB operation failed , delete association failed, transaction rollback.")
		return nil, base.ErrorTransaction
	}
	networkAdaptersE := []entity.NetworkAdapter{}
	for _, v := range networkAdapters {
		each := entity.NetworkAdapter{}
		each.Load(&v)
		networkAdaptersE = append(networkAdaptersE, each)
	}
	server.NetworkAdapters = networkAdaptersE
	if err := tx.Save(server).Error; err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"id":    ID,
			"op":    "UpdateNetworkAdapters",
			"error": err,
		}).Warn("DB opertion failed, save server failed, transaction rollback.")
		return nil, base.ErrorTransaction
	}
	if err := tx.Commit().Error; err != nil {
		log.WithFields(log.Fields{
			"id":    ID,
			"op":    "UpdateNetworkAdapters",
			"error": err,
		}).Warn("DB opertion failed, commit failed.")
		return nil, base.ErrorTransaction
	}
	return server.ToModel(), nil
}

func (impl *Server) deleteDrives(c *gorm.DB, server *entity.Server) error {
	for i := range server.Drives {
		drive := server.Drives[i]
		for j := range drive.Location {
			if drive.Location[j].PostalAddress != nil {
				if err := c.Delete(&drive.Location[j].PostalAddress).Error; err != nil {
					return base.ErrorTransaction
				}
			}
			if drive.Location[j].Placement != nil {
				if err := c.Delete(&drive.Location[j].Placement).Error; err != nil {
					return base.ErrorTransaction
				}
			}
			if err := c.Delete(&drive.Location[j]).Error; err != nil {
				return base.ErrorTransaction
			}
		}
		if err := c.Delete(&drive).Error; err != nil {
			return base.ErrorTransaction
		}
	}
	return nil
}

// UpdateDrives Update drives.
func (impl *Server) UpdateDrives(ID string, drives []model.Drive) (base.ModelInterface, error) {
	var (
		c      = impl.TemplateImpl.GetConnection()
		server = new(entity.Server)
	)
	tx := c.Begin()
	if err := tx.Error; err != nil {
		log.WithFields(log.Fields{
			"id":    ID,
			"op":    "UpdateDrives",
			"error": err,
		}).Warn("DB operation failed , start transaction failed.")
		return nil, base.ErrorTransaction
	}
	found, err := impl.GetInternal(tx, ID, server)
	if !found || err != nil {
		log.WithFields(log.Fields{
			"id": ID,
			"op": "UpdateDrives",
		}).Warn("DB operation failed , load server failed.")
		return nil, base.ErrorTransaction
	}
	if err := impl.deleteDrives(tx, server); err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"id":    ID,
			"op":    "UpdateDrives",
			"error": err,
		}).Warn("DB operation failed , delete association failed, transaction rollback.")
		return nil, base.ErrorTransaction
	}
	drivesE := []entity.Drive{}
	for _, v := range drives {
		each := entity.Drive{}
		each.Load(&v)
		drivesE = append(drivesE, each)
	}
	server.Drives = drivesE
	if err := tx.Save(server).Error; err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"id":    ID,
			"op":    "UpdateDrives",
			"error": err,
		}).Warn("DB opertion failed, save server failed, transaction rollback.")
		return nil, base.ErrorTransaction
	}
	if err := tx.Commit().Error; err != nil {
		log.WithFields(log.Fields{
			"id":    ID,
			"op":    "UpdateDrives",
			"error": err,
		}).Warn("DB opertion failed, commit failed.")
		return nil, base.ErrorTransaction
	}
	return server.ToModel(), nil
}

func (impl *Server) deletePCIeDevices(c *gorm.DB, server *entity.Server) error {
	for i := range server.PCIeDevices {
		pcieDevice := server.PCIeDevices[i]
		for j := range pcieDevice.PCIeFunctions {
			if err := c.Delete(&pcieDevice.PCIeFunctions[j]).Error; err != nil {
				return base.ErrorTransaction
			}
		}
		if err := c.Delete(&pcieDevice).Error; err != nil {
			return base.ErrorTransaction
		}
	}
	return nil
}

// UpdatePCIeDevices Updagte PCIe devices.
func (impl *Server) UpdatePCIeDevices(ID string, pcieDevices []model.PCIeDevice) (base.ModelInterface, error) {
	var (
		c      = impl.TemplateImpl.GetConnection()
		server = new(entity.Server)
	)
	tx := c.Begin()
	if err := tx.Error; err != nil {
		log.WithFields(log.Fields{
			"id":    ID,
			"op":    "UpdatePCIeDevices",
			"error": err,
		}).Warn("DB operation failed , start transaction failed.")
		return nil, base.ErrorTransaction
	}
	found, err := impl.GetInternal(tx, ID, server)
	if !found || err != nil {
		log.WithFields(log.Fields{
			"id": ID,
			"op": "UpdatePCIeDevices",
		}).Warn("DB operation failed , load server failed.")
		return nil, base.ErrorTransaction
	}
	if err := impl.deletePCIeDevices(tx, server); err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"id":    ID,
			"op":    "UpdatePCIeDevices",
			"error": err,
		}).Warn("DB operation failed , delete association failed, transaction rollback.")
		return nil, base.ErrorTransaction
	}
	pcieDevicesE := new([]entity.PCIeDevice)
	for _, v := range pcieDevices {
		each := entity.PCIeDevice{}
		each.Load(&v)
		*pcieDevicesE = append(*pcieDevicesE, each)
	}
	server.PCIeDevices = *pcieDevicesE
	if err := tx.Save(server).Error; err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"id":    ID,
			"op":    "UpdatePCIeDevices",
			"error": err,
		}).Warn("DB opertion failed, save server failed, transaction rollback.")
		return nil, base.ErrorTransaction
	}
	if err := tx.Commit().Error; err != nil {
		log.WithFields(log.Fields{
			"id":    ID,
			"op":    "UpdatePCIeDevices",
			"error": err,
		}).Warn("DB opertion failed, commit failed.")
		return nil, base.ErrorTransaction
	}
	return server.ToModel(), nil
}

// DeleteServer will override the default process.
// When delete a server we should delete all the associations too and return the association back.
func (impl *Server) DeleteServer(id string) (base.ModelInterface, []base.ModelInterface, *base.Message) {
	var (
		name             = impl.ResourceName()
		record           = new(entity.Server)
		previous         = new(entity.Server)
		deletedSSG       = make([]base.ModelInterface, 0)
		deletedSSGEntity = make([]entity.ServerServerGroup, 0)
		c                = impl.GetConnection()
	)

	tx := c.Begin()
	if err := tx.Error; err != nil {
		log.WithFields(log.Fields{
			"resource": name,
			"id":       id,
			"error":    err,
		}).Warn("DB delete resource failed, start transaction failed.")
		return nil, nil, base.NewMessageTransactionError()
	}
	exist, err := impl.GetInternal(tx, id, previous)
	// Rollback in GetInternal.
	if !exist {
		return nil, nil, base.NewMessageNotExist()
	}
	if err != nil {
		return nil, nil, base.NewMessageTransactionError()
	}

	record.SetID(id)
	for _, v := range record.Association() {
		if err := tx.Delete(v).Error; err != nil {
			tx.Rollback()
			log.WithFields(log.Fields{
				"resource": name,
				"id":       id,
				"error":    err,
			}).Warn("DB delete resource failed, delete association failed, transaction rollback.")
			return nil, nil, base.NewMessageTransactionError()
		}
	}
	if err := tx.Delete(record).Error; err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"resource": name,
			"id":       id,
			"error":    err,
		}).Warn("DB delete resource failed, delete resource failed, transaction rollback.")
		return nil, nil, base.NewMessageTransactionError()
	}

	// Delete the server-servergroup association.
	// But we need record them first.
	if err := tx.Where("\"ServerID\" = ?", id).Find(&deletedSSGEntity).Error; err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"resource": name,
			"id":       id,
			"error":    err,
		}).Warn("DB delete resource failed, record server-servergroup association failed, transaction rollback.")
		return nil, nil, base.NewMessageTransactionError()
	}
	for _, each := range deletedSSGEntity {
		deletedSSG = append(deletedSSG, each.ToModel())
	}
	if err := tx.Where("\"ServerID\" = ?", id).Delete(entity.ServerServerGroup{}).Error; err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"resource": name,
			"id":       id,
			"error":    err,
		}).Warn("DB delete resource failed, delete server-servergroup association failed, transaction rollback.")
		return nil, nil, base.NewMessageTransactionError()
	}

	if err := tx.Commit().Error; err != nil {
		log.WithFields(log.Fields{
			"resource": name,
			"id":       id,
			"error":    err,
		}).Warn("DB delete resource failed, commit failed.")
		return nil, nil, base.NewMessageTransactionError()
	}
	return previous.ToModel(), deletedSSG, nil
}

// DeleteServerCollection override the default process.
// We need delete the servergroup relationship too.
// It returns all the deleted resources,
// It returnes if committed.
// It returens error if any.
func (impl *Server) DeleteServerCollection() ([]base.ModelInterface, []base.ModelInterface, *base.Message) {
	var (
		name             = impl.TemplateImpl.ResourceName()
		recordCollection = impl.TemplateImpl.NewEntityCollection()
		c                = impl.TemplateImpl.GetConnection()
		tables           = impl.TemplateImpl.NewEntity().Tables()
		deletedSSG       = make([]base.ModelInterface, 0)
		deletedSSGEntity = make([]entity.ServerServerGroup, 0)
	)

	// We need transaction to ensure the total and the query count is consistent.
	tx := c.Begin()
	if err := tx.Error; err != nil {
		log.WithFields(log.Fields{
			"resource": name,
			"error":    err,
		}).Warn("DB delete collection failed, start transaction failed.")
		return nil, nil, base.NewMessageTransactionError()
	}

	if err := tx.Find(recordCollection).Error; err != nil {
		log.WithFields(log.Fields{
			"resource": name,
			"error":    err,
		}).Warn("DB delete collection failed, find resource failed.")
		return nil, nil, base.NewMessageTransactionError()
	}
	for _, v := range tables {
		if err := tx.Delete(v).Error; err != nil {
			tx.Rollback()
			log.WithFields(log.Fields{
				"resource": name,
				"error":    err,
			}).Warn("DB delete collection failed, delete resources failed, transaction rollback.")
			return nil, nil, base.NewMessageTransactionError()
		}
	}
	// When we delete all the servers we also need delete all the server-servergroup.
	if err := tx.Find(&deletedSSGEntity).Error; err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"resource": name,
			"error":    err,
		}).Warn("DB delete resource failed, record server-servergroup association failed, transaction rollback.")
		return nil, nil, base.NewMessageTransactionError()
	}
	for _, each := range deletedSSGEntity {
		deletedSSG = append(deletedSSG, each.ToModel())
	}
	if err := tx.Delete(entity.ServerServerGroup{}).Error; err != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"error": err}).
			Warn("Delete server collection in DB failed, delete server-servergroup collection failed, transaction rollback.")
		return nil, nil, base.NewMessageTransactionError()
	}
	ret, message := impl.TemplateImpl.ConvertFindResultToModel(recordCollection)
	if message != nil {
		tx.Rollback()
		log.WithFields(log.Fields{
			"resource": name,
			"message":  message.ID,
		}).Warn("DB delete collection failed, convert find result failed, transaction rollback.")
		return nil, nil, base.NewMessageTransactionError()
	}
	if err := tx.Commit().Error; err != nil {
		log.WithFields(log.Fields{
			"resource": name,
			"error":    err,
		}).Warn("DB delete collection failed, commit failed.")
		return nil, nil, base.NewMessageTransactionError()
	}
	return ret, deletedSSG, nil
}
