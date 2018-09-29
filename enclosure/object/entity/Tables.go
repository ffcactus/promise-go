package entity

import (
	"promise/base"
)

var (
	// Tables The tables used by this project.
	Tables = []base.TableInfo{
		{Name: "Enclosure", Info: new(Enclosure)},
		{Name: "BladeSlot", Info: new(BladeSlot)},
		{Name: "SwitchSlot", Info: new(SwitchSlot)},
		{Name: "ManagerSlot", Info: new(ManagerSlot)},
		{Name: "ApplianceSlot", Info: new(ApplianceSlot)},
		{Name: "FanSlot", Info: new(FanSlot)},
		{Name: "PowerSlot", Info: new(PowerSlot)},
	}
)
