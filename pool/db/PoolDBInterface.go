package db

import (
	"promise/pool/object/model"
)

// PoolDBInterface is the DB interface.
type PoolDBInterface interface {
	PostIPv4Pool(*model.IPv4Pool) (bool, *model.IPv4Pool, bool, error)
	GetIPv4Pool(id string) *model.IPv4Pool
	GetIPv4PoolCollection(start int64, count int64, filter string) (*model.IPv4PoolCollection, error)
	DeleteIPv4Pool(id string) (bool, *model.IPv4Pool, bool, error)
	DeleteIPv4PoolCollection() ([]model.IPv4Pool, bool, error)
	AllocateIPv4Address(id string, key string) (exist bool, address string, pool *model.IPv4Pool, commited bool, err error)
	FreeIPv4Address(id string, address string) (exist bool, pool *model.IPv4Pool, commited bool, err error)
}
