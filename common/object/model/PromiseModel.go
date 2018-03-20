package model

import (
	"time"
)

// PromiseModel is the common model in Promise.
type PromiseModel struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
}
