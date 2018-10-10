package entity

import (
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/enclosure/object/model"
)

// Enclosure is the entity.
type Enclosure struct {
	base.Entity
	base.DeviceIdentity
	Name           string          `gorm:"column:Name"`
	Description    string          `gorm:"column:Description"`
	Type           string          `gorm:"column:Type"`
	State          string          `gorm:"column:State"`
	StateReason    string          `gorm:"column:StateReason"`
	Health         string          `gorm:"column:Health"`
	ServerSlots    []ServerSlot    `gorm:"column:ServerSlots;ForeignKey:EnclosureRef"`
	SwitchSlots    []SwitchSlot    `gorm:"column:SwitchSlots;ForeignKey:EnclosureRef"`
	ManagerSlots   []ManagerSlot   `gorm:"column:ManagerSlots;ForeignKey:EnclosureRef"`
	ApplianceSlots []ApplianceSlot `gorm:"column:ApplianceSlots;ForeignKey:EnclosureRef"`
	FanSlots       []FanSlot       `gorm:"column:FanSlots;ForeignKey:EnclosureRef"`
	PowerSlots     []PowerSlot     `gorm:"column:PowerSlots;ForeignKey:EnclosureRef"`
	CredentialURL  string          `gorm:"column:CredentialURL"`
}

// TableName will set the table name.
func (Enclosure) TableName() string {
	return "Enclosure"
}

// String return the debug name of this entity.
func (e Enclosure) String() string {
	return e.Name
}

// PropertyNameForDuplicationCheck return the property name used for duplication check.
func (e *Enclosure) PropertyNameForDuplicationCheck() string {
	return "Name"
}

// Preload return the property names that need to be preload.
func (e *Enclosure) Preload() []string {
	return []string{
		"ServerSlots",
		"SwitchSlots",
		"ManagerSlots",
		"ApplianceSlots",
		"FanSlots",
		"PowerSlots",
	}
}

// Association return all the assocations that need to delete when deleting a resource.
func (e *Enclosure) Association() []interface{} {
	ret := []interface{}{}
	for _, v := range e.PowerSlots {
		ret = append(ret, v)
	}
	for _, v := range e.FanSlots {
		ret = append(ret, v)
	}
	for _, v := range e.ApplianceSlots {
		ret = append(ret, v)
	}
	for _, v := range e.ManagerSlots {
		ret = append(ret, v)
	}
	for _, v := range e.SwitchSlots {
		ret = append(ret, v)
	}
	for _, v := range e.ServerSlots {
		ret = append(ret, v)
	}
	return ret
}

// Tables returns the tables to delete when you want delete all the resources.
func (e *Enclosure) Tables() []interface{} {
	return []interface{}{
		new(PowerSlot),
		new(FanSlot),
		new(ApplianceSlot),
		new(ManagerSlot),
		new(SwitchSlot),
		new(ServerSlot),
		new(Enclosure),
	}
}

// FilterNameList return all the property name that can be used in filter.
func (e *Enclosure) FilterNameList() []string {
	return []string{"Name", "State", "State", "Health", "UUID", "SerialNumber", "PartNumber"}
}

// Load will load data from model. this function is used on POST.
func (e *Enclosure) Load(i base.ModelInterface) error {
	m, ok := i.(*model.Enclosure)
	if !ok {
		log.Error("entity.Enclosure.Load() failed, convert interface failed.")
		return base.ErrorDataConvert
	}
	base.EntityLoad(&e.Entity, &m.Model)
	e.DeviceIdentity = m.DeviceIdentity
	e.Name = m.Name
	e.Description = m.Description
	e.Type = m.Type
	e.State = m.State
	e.Health = m.Health
	// blade
	e.ServerSlots = make([]ServerSlot, 0)
	for _, v := range m.ServerSlots {
		k := ServerSlot{}
		k.Load(&v)
		e.ServerSlots = append(e.ServerSlots, k)
	}
	// switch
	e.SwitchSlots = make([]SwitchSlot, 0)
	for _, v := range m.SwitchSlots {
		k := SwitchSlot{}
		k.Load(&v)
		e.SwitchSlots = append(e.SwitchSlots, k)
	}
	// manager
	e.ManagerSlots = make([]ManagerSlot, 0)
	for _, v := range m.ManagerSlots {
		k := ManagerSlot{}
		k.Load(&v)
		e.ManagerSlots = append(e.ManagerSlots, k)
	}
	// appliance
	e.ApplianceSlots = make([]ApplianceSlot, 0)
	for _, v := range m.ApplianceSlots {
		k := ApplianceSlot{}
		k.Load(&v)
		e.ApplianceSlots = append(e.ApplianceSlots, k)
	}
	// fan
	e.FanSlots = make([]FanSlot, 0)
	for _, v := range m.FanSlots {
		k := FanSlot{}
		k.Load(&v)
		e.FanSlots = append(e.FanSlots, k)
	}
	// power
	e.PowerSlots = make([]PowerSlot, 0)
	for _, v := range m.PowerSlots {
		k := PowerSlot{}
		k.Load(&v)
		e.PowerSlots = append(e.PowerSlots, k)
	}
	e.CredentialURL = m.Credential.URL
	return nil
}

// ToModel will create a new model from entity.
func (e *Enclosure) ToModel() base.ModelInterface {
	m := model.Enclosure{}
	base.EntityToModel(&e.Entity, &m.Model)
	m.DeviceIdentity = e.DeviceIdentity
	m.Name = e.Name
	m.Description = e.Description
	m.Type = e.Type
	m.State = e.State
	m.Health = e.Health
	// blade
	m.ServerSlots = make([]model.ServerSlot, 0)
	for _, v := range e.ServerSlots {
		m.ServerSlots = append(m.ServerSlots, *v.ToModel())
	}
	// switch
	m.SwitchSlots = make([]model.SwitchSlot, 0)
	for _, v := range e.SwitchSlots {
		m.SwitchSlots = append(m.SwitchSlots, *v.ToModel())
	}
	// manager
	m.ManagerSlots = make([]model.ManagerSlot, 0)
	for _, v := range e.ManagerSlots {
		m.ManagerSlots = append(m.ManagerSlots, *v.ToModel())
	}
	// appliance
	m.ApplianceSlots = make([]model.ApplianceSlot, 0)
	for _, v := range e.ApplianceSlots {
		m.ApplianceSlots = append(m.ApplianceSlots, *v.ToModel())
	}
	// fan
	m.FanSlots = make([]model.FanSlot, 0)
	for _, v := range e.FanSlots {
		m.FanSlots = append(m.FanSlots, *v.ToModel())
	}
	// power
	m.PowerSlots = make([]model.PowerSlot, 0)
	for _, v := range e.PowerSlots {
		m.PowerSlots = append(m.PowerSlots, *v.ToModel())
	}
	e.CredentialURL = m.Credential.URL
	return &m
}

// ToCollectionMember convert the entity to member.
func (e *Enclosure) ToCollectionMember() base.CollectionMemberModelInterface {
	m := new(model.EnclosureCollectionMember)
	base.EntityToMember(&e.Entity, &m.CollectionMemberModel)
	m.Name = e.Name
	m.Description = e.Description
	m.State = e.State
	m.Health = e.Health
	return m
}
