package entity

import (
	commonModel "promise/common/object/model"
	commonEntity "promise/common/object/entity"
	commonUtil "promise/common/util"
	"promise/pool/object/model"
)

// IPv4Address is the entity to represents an IPv4 address.
type IPv4Address struct {
	commonEntity.Element
	IPv4RangeRef commonEntity.ElementRefType `gorm:"column:IPv4RangeRef"`
	Key          string                      `gorm:"column:Key"`
	Address      string                      `gorm:"column:Address"`
	Allocated    bool                        `gorm:"column:Allocated"`
}

// TableName will set the table name.
func (IPv4Address) TableName() string {
	return "IPv4Address"
}

// IPv4Range is the IPv4 range.
type IPv4Range struct {
	commonEntity.Element
	IPv4PoolRef string        `gorm:"column:IPv4PoolRef"`
	Start       string        `gorm:"column:Start"`
	End         string        `gorm:"column:End"`
	Total       uint32        `gorm:"column:Total"`
	Free        uint32        `gorm:"column:Free"`
	Allocatable uint32        `gorm:"column:Allocatable"`
	Addresses   []IPv4Address `gorm:"column:Addresses;ForeignKey:IPv4RangeRef"`
}

// TableName will set the table name.
func (IPv4Range) TableName() string {
	return "IPv4Range"
}

// IPv4Pool is the entity.
type IPv4Pool struct {
	commonEntity.PromiseEntity
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

// ToModel will create a new model from entity.
func (e *IPv4Pool) ToModel() *model.IPv4Pool {
	m := new(model.IPv4Pool)
	m.PromiseModel = e.PromiseEntity.ToModel()
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
		commonUtil.StringToStruct(*e.DNSServers, &dns)
		if dns == nil {
			dns = make([]string, 0)
		}
		m.DNSServers = &dns
	}

	return m
}

// ToMember change the entity to a collection member.
func (e *IPv4Pool) ToMember() commonModel.PromiseMemberInterface {
	m := model.IPv4PoolMember{}
	m.ID = e.ID
	m.Category = e.Category
	m.Name = e.Name
	return &m
}

// Load will load data from model. this function is used on POST.
func (e *IPv4Pool) Load(m *model.IPv4Pool) {
	e.PromiseEntity.Load(m.PromiseModel)
	e.Name = m.Name
	e.Description = m.Description
	e.SubnetMask = m.SubnetMask
	e.Gateway = m.Gateway
	e.Domain = m.Domain
	if m.DNSServers == nil {
		e.DNSServers = nil
	} else {
		s := commonUtil.StructToString(*m.DNSServers)
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
}
