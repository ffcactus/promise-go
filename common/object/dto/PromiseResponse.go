package dto

import (
	"time"
)

// PromiseResponse is the base type of a response in Promise. 
type PromiseResponse struct {
	ID        string    `json:"ID"`
	URI       string    `json:"URI"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}
