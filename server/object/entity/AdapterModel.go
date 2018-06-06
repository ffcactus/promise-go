package entity

import (
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/server/object/model"
)

// RAIDCapability describe the capability of an FCoE adapter.
type RAIDCapability struct {
	base.ElementEntity
	AdapterCapabilityRef base.ElementEntityRefType `gorm:"column:AdapterCapabilityRef"`
	Version              int                       `gorm:"column:Version"`
}

// TableName will set the table name.
func (RAIDCapability) TableName() string {
	return "RAIDCapability"
}

// Load will load data from model.
func (e *RAIDCapability) Load(m model.RAIDCapability) {
	e.Version = m.Version
}

// ToModel will create a new model from entity.
func (e *RAIDCapability) ToModel() *model.RAIDCapability {
	m := model.RAIDCapability{}
	m.Version = e.Version
	return &m
}

// EthernetCapability describe the capability of an FCoE adapter.
type EthernetCapability struct {
	base.ElementEntity
	AdapterCapabilityRef base.ElementEntityRefType `gorm:"column:AdapterCapabilityRef"`
	Version              int                       `gorm:"column:Version"`
}

// TableName will set the table name.
func (EthernetCapability) TableName() string {
	return "EthernetCapability"
}

// Load will load data from model.
func (e *EthernetCapability) Load(m model.EthernetCapability) {
	e.Version = m.Version
}

// ToModel will create a new model from entity.
func (e *EthernetCapability) ToModel() *model.EthernetCapability {
	m := model.EthernetCapability{}
	m.Version = e.Version
	return &m
}

// FCoECapability describe the capability of an FCoE adapter.
type FCoECapability struct {
	base.ElementEntity
	AdapterCapabilityRef base.ElementEntityRefType `gorm:"column:AdapterCapabilityRef"`
	Version              int                       `gorm:"column:Version"`
}

// TableName will set the table name.
func (FCoECapability) TableName() string {
	return "FCoECapability"
}

// Load will load data from model.
func (e *FCoECapability) Load(m model.FCoECapability) {
	e.Version = m.Version
}

// ToModel will create a new model from entity.
func (e *FCoECapability) ToModel() *model.FCoECapability {
	m := model.FCoECapability{}
	m.Version = e.Version
	return &m
}

// AdapterCapability describe the capability of an adapter.
type AdapterCapability struct {
	base.ElementEntity
	AdapterModelRef base.EntityRefType  `gorm:"column:AdapterModelRef"`
	Version         int                 `gorm:"column:Version"`
	RAID            *RAIDCapability     `gorm:"column:RAID;ForeignKey:AdapterCapabilityRef"`
	Ethernet        *EthernetCapability `gorm:"column:Ethernet;ForeignKey:AdapterCapabilityRef"`
	FCoE            *FCoECapability     `gorm:"column:FCoE;ForeignKey:AdapterCapabilityRef"`
}

// TableName will set the table name.
func (AdapterCapability) TableName() string {
	return "AdapterCapability"
}

// Load will load data from model.
func (e *AdapterCapability) Load(m model.AdapterCapability) {
	e.Version = m.Version
	if m.RAID != nil {
		e.RAID = new(RAIDCapability)
		e.RAID.Load(*m.RAID)
	}
	if m.Ethernet != nil {
		e.Ethernet = new(EthernetCapability)
		e.Ethernet.Load(*m.Ethernet)
	}
	if m.FCoE != nil {
		e.FCoE = new(FCoECapability)
		e.FCoE.Load(*m.FCoE)
	}
}

// ToModel will create a new model from entity.
func (e *AdapterCapability) ToModel() model.AdapterCapability {
	m := model.AdapterCapability{}
	m.Version = e.Version
	if e.RAID != nil {
		m.RAID = e.RAID.ToModel()
	}
	if e.Ethernet != nil {
		m.Ethernet = e.Ethernet.ToModel()
	}
	if e.FCoE != nil {
		m.FCoE = e.FCoE.ToModel()
	}
	return m
}

// AdapterModel is the entity.
type AdapterModel struct {
	base.Entity
	Name       string            `gorm:"column:Name"`
	Type       string            `gorm:"column:Type"`
	Capability AdapterCapability `gorm:"column:Capability;ForeignKey:AdapterModelRef"`
}

// TableName will set the table name.
func (AdapterModel) TableName() string {
	return "AdapterModel"
}

// DebugInfo return the debug name of this entity.
func (e *AdapterModel) DebugInfo() string {
	return e.Name
}

// PropertyNameForDuplicationCheck return the property name used for duplication check.
func (e *AdapterModel) PropertyNameForDuplicationCheck() string {
	return "Name"
}

// Preload return the property names that need to be preload.
func (e *AdapterModel) Preload() []string {
	return []string{
		"Capability",
		"Capability.RAID",
		"Capability.Ethernet",
		"Capability.FCoE",
	}
}

// Association return all the assocations that need to delete when deleting a resource.
func (e *AdapterModel) Association() []interface{} {
	ret := []interface{}{}
	if e.Capability.RAID != nil {
		ret = append(ret, &e.Capability.RAID)
	}
	if e.Capability.Ethernet != nil {
		ret = append(ret, &e.Capability.Ethernet)
	}
	if e.Capability.FCoE != nil {
		ret = append(ret, &e.Capability.FCoE)
	}
	ret = append(ret, &e.Capability)
	return ret
}

// Tables returns the tables to delete when you want delete all the resources.
func (e *AdapterModel) Tables() []interface{} {
	return []interface{}{
		new(AdapterModel),
		new(AdapterCapability),
		new(RAIDCapability),
		new(EthernetCapability),
		new(FCoECapability),
	}
}

// FilterNameList return all the property name that can be used in filter.
func (e *AdapterModel) FilterNameList() []string {
	return []string{"Name", "Type"}
}

// Load will load data from model. this function is used on POST.
func (e *AdapterModel) Load(i base.ModelInterface) error {
	m, ok := i.(*model.AdapterModel)
	if !ok {
		log.Error("entity.AdapterModel.Load() failed, convert interface failed.")
		return base.ErrorDataConvert
	}
	base.EntityLoad(&e.Entity, &m.Model)
	e.Name = m.Name
	e.Type = m.Type
	e.Capability.Load(m.Capability)
	return nil
}

// ToModel will create a new model from entity.
func (e *AdapterModel) ToModel() base.ModelInterface {
	m := model.AdapterModel{}
	base.EntityToModel(&e.Entity, &m.Model)
	m.Name = e.Name
	m.Type = e.Type
	m.Capability = e.Capability.ToModel()
	return &m
}

// ToCollectionMember convert the entity to member.
func (e *AdapterModel) ToCollectionMember() base.CollectionMemberModelInterface {
	m := new(model.AdapterModelCollectionMember)
	base.EntityToMember(&e.Entity, &m.CollectionMemberModel)
	m.Name = e.Name
	m.Type = e.Type
	return m
}
