package model

import (
	"time"
)

// PromiseModel is the common model in Promise.
type PromiseModel struct {
	ID        string
	Category  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
