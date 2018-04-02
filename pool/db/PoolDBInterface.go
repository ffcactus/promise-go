package db

import (
	"promise/pool/object/model"
)

// PoolDBInterface is the DB interface.
type PoolDBInterface interface {
	PostIPv4Pool(*model.IPv4Pool) (bool, *model.IPv4Pool, bool, error)
	GetIPv4Pool(id string) *model.IPv4Pool
	GetIPv4PoolCollection(start int, count int, filter string) (*model.IPv4PoolCollection, error)
	DeleteIPv4Pool(id string) (bool, *model.IPv4Pool, bool, error)
	DeleteIPv4PoolCollection() ([]model.IPv4Pool, bool, error)
}