package entity

import (
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/pool/object/model"
)

// IPv4Address is the entity to represents an IPv4 address.
type IPv4Address struct {
	base.ElementEntity
	IPv4RangeRef base.ElementEntityRefType `gorm:"column:IPv4RangeRef"`
	Key          string                    `gorm:"column:Key"`
	Address      string                    `gorm:"column:Address"`
	Allocated    bool                      `gorm:"column:Allocated"`
}

// TableName will set the table name.
func (IPv4Address) TableName() string {
	return "IPv4Address"
}

// IPv4Range is the IPv4 range.
type IPv4Range struct {
	base.ElementEntity
	IPv4PoolRef base.EntityRefType `gorm:"column:IPv4PoolRef"`
	Start       string             `gorm:"column:Start"`
	End         string             `gorm:"column:End"`
	Total       uint32             `gorm:"column:Total"`
	Free        uint32             `gorm:"column:Free"`
	Allocatable uint32             `gorm:"column:Allocatable"`
	Addresses   []IPv4Address      `gorm:"column:Addresses;ForeignKey:IPv4RangeRef"`
}

// TableName will set the table name.
func (IPv4Range) TableName() string {
	return "IPv4Range"
}

// IPv4Pool is the entity.
type IPv4Pool struct {
	base.Entity
	Name        string      `gorm:"column:Name"`
	Description *string     `gorm:"column:Description"`
	SubnetMask  *string     `gorm:"column:SubnetMask"`
	Gateway     *string     `gorm:"column:Gateway"`
	Domain      *string     `gorm:"column:Domain"`
	DNSServers  *string     `gorm:"column:DNSServers"`
	Ranges      []IPv4Range `gorm:"column:Ranges;ForeignKey:IPv4PoolRef"`
}

// TableName will set the table name.
func (IPv4Pool) TableName() string {
	return "IPv4Pool"
}

// GetDebugName return the debug name of this entity.
func (e *IPv4Pool) GetDebugName() string {
	return e.Name
}

// GetPropertyNameForDuplicationCheck return the property name used for duplication check.
func (e *IPv4Pool) GetPropertyNameForDuplicationCheck() string {
	return "Name"
}

// GetPreload return the property names that need to be preload.
func (e *IPv4Pool) GetPreload() []string {
	return []string{"Ranges", "Ranges.Addresses"}
}

// GetAssociation return all the assocations that need to delete when deleting a resource.
func (e *IPv4Pool) GetAssociation() []interface{} {
	ret := []interface{}{}
	for _, v := range e.Ranges {
		for _, j := range v.Addresses {
			ret = append(ret, j)
		}
		ret = append(ret, v)
	}
	return ret
}

// GetTables returns the tables to delete when you want delete all the resources.
func (e *IPv4Pool) GetTables() []interface{} {
	return []interface{}{new(IPv4Address), new(IPv4Range), new(IPv4Pool)}
}

// GetFilterNameList return all the property name that can be used in filter.
func (e *IPv4Pool) GetFilterNameList() []string {
	return []string{"Name"}
}

// Load will load data from model. this function is used on POST.
func (e *IPv4Pool) Load(i base.ModelInterface) error {
	m, ok := i.(*model.IPv4Pool)
	if !ok {
		log.Error("entity.IPv4Pool.Load() failed, convert interface failed.")
		return base.ErrorDataConvert
	}
	base.EntityLoad(&e.Entity, &m.Model)
	e.Name = m.Name
	e.Description = m.Description
	e.SubnetMask = m.SubnetMask
	e.Gateway = m.Gateway
	e.Domain = m.Domain
	if m.DNSServers == nil {
		e.DNSServers = nil
	} else {
		s := base.StructToString(*m.DNSServers)
		e.DNSServers = &s
	}

	e.Ranges = make([]IPv4Range, 0)
	for _, v := range m.Ranges {
		vv := IPv4Range{}
		vv.Start = v.Start
		vv.End = v.End
		vv.Total = v.Total
		vv.Free = v.Free
		vv.Allocatable = v.Allocatable
		addresses := make([]IPv4Address, 0)
		for _, k := range v.Addresses {
			kk := IPv4Address{}
			kk.Key = k.Key
			kk.Address = k.Address
			kk.Allocated = k.Allocated
			addresses = append(addresses, kk)
		}
		vv.Addresses = addresses
		e.Ranges = append(e.Ranges, vv)
	}
	return nil
}

// ToModel will create a new model from entity.
func (e *IPv4Pool) ToModel() base.ModelInterface {
	m := model.IPv4Pool{}
	base.EntityToModel(&e.Entity, &m.Model)
	m.Name = e.Name
	m.Description = e.Description
	m.SubnetMask = e.SubnetMask
	m.Gateway = e.Gateway
	m.Domain = e.Domain

	ranges := make([]model.IPv4Range, 0)
	for _, v := range e.Ranges {
		vv := model.IPv4Range{}
		vv.Start = v.Start
		vv.End = v.End
		vv.Total = v.Total
		vv.Free = v.Free
		vv.Allocatable = v.Allocatable
		addresses := make([]model.IPv4Address, 0)
		for _, k := range v.Addresses {
			kk := model.IPv4Address{}
			kk.Key = k.Key
			kk.Address = k.Address
			kk.Allocated = k.Allocated
			addresses = append(addresses, kk)
		}
		vv.Addresses = addresses
		ranges = append(ranges, vv)
	}
	m.Ranges = ranges

	if e.DNSServers == nil {
		m.DNSServers = nil
	} else {
		dns := make([]string, 0)
		base.StringToStruct(*e.DNSServers, &dns)
		if dns == nil {
			dns = make([]string, 0)
		}
		m.DNSServers = &dns
	}

	return &m
}

// ToCollectionMember convert the entity to member.
func (e *IPv4Pool) ToCollectionMember() base.CollectionMemberModelInterface {
	m := new(model.IPv4PoolCollectionMember)
	base.EntityToMember(&e.Entity, &m.CollectionMemberModel)
	m.Name = e.Name
	return m
}
