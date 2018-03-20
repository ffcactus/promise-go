package entity

import (
	"promise/common/object/model"
	"time"
)

// PromiseEntity is the common entity used in Promise.
type PromiseEntity struct {
	ID        string    `gorm:"column:ID"`
	CreatedAt time.Time `gorm:"column:CreatedAt"`
	UpdatedAt time.Time `gorm:"column:UpdatedAt"`
}

// ToModel change the entity to model.
func (e *PromiseEntity) ToModel() model.PromiseModel {
	var ret model.PromiseModel

	ret.ID = e.ID
	ret.CreatedAt = e.CreatedAt
	ret.UpdatedAt = e.UpdatedAt
	return ret
}
