package entity

import (
	commonUtil "promise/common/util"
	"promise/common/object/entity"
	"promise/pool/object/model"
)

// IPv4Pool is the entity.
type IPv4Pool struct {
	entity.PromiseEntity
	Name string `gorm:"column:Name"`
	Description string `gorm:"column:Description"`
	Ranges string `gorm:"column:Ranges"`
	SubnetMask string `gorm:"column:SubnetMask"`
	Gateway string `gorm:"column:Gateway"`
	Domain string `gorm:"column:Domain"`
	DNSServers string `gorm:"column:DNSServers"`
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
	ranges := make([]model.IPv4Range, 0)
	commonUtil.StringToStruct(e.Ranges, &ranges)
	m.Ranges = ranges
	m.SubnetMask = e.SubnetMask
	m.Gateway = e.Gateway
	m.Domain = e.Domain
	dns := make([]string, 0)
	commonUtil.StringToStruct(e.DNSServers, &dns)
	if dns == nil {
		dns := make([]string, 0)
	}
	m.DNSServers = dns

	return m
}

// Load will load data from model.
func (e *IPv4Pool) Load(m *model.IPv4Pool) {
	e.PromiseEntity.Load(m.PromiseModel)
	e.Name = m.Name
	e.Description = m.Description
	e.SubnetMask = m.SubnetMask
	e.Gateway = m.Gateway
	e.Domain = m.Domain
	e.Ranges = commonUtil.StructToString(m.Ranges)
	e.DNSServers = commonUtil.StructToString(m.DNSServers)
}