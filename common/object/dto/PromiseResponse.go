package dto

import (
	"time"
)

type PromiseResponse struct {
	ID        string    `json:"ID"`
	URI       string    `json:"URI"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}
