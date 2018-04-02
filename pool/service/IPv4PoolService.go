package service

import (
	"promise/common/category"
	commonMessage "promise/common/object/message"
	"promise/pool/db"
	"promise/pool/object/dto"
	"promise/pool/object/model"
	wsSDK "promise/sdk/ws"
)

// PostIPv4Pool post a IPv4 pool.
func PostIPv4Pool(request *dto.PostIPv4PoolRequest) (*model.IPv4Pool, []commonMessage.Message) {
	dbImpl := db.GetPoolDB()

	exist, posted, commited, err := dbImpl.PostIPv4Pool(request.ToModel())
	if exist {
		return nil, []commonMessage.Message{commonMessage.NewResourceDuplicate()}
	}
	if err != nil || !commited {
		return nil, []commonMessage.Message{commonMessage.NewTransactionError()}
	}
	var poolDTO dto.GetIPv4PoolResponse
	poolDTO.Load(posted)
	wsSDK.DispatchResourceCreateEvent(&poolDTO)
	return posted, nil
}

// GetIPv4Pool will get IPv4 pool by ID.
func GetIPv4Pool(id string) (*model.IPv4Pool, []commonMessage.Message) {
	dbImpl := db.GetPoolDB()

	ipv4Pool := dbImpl.GetIPv4Pool(id)
	if ipv4Pool == nil {
		return nil, []commonMessage.Message{commonMessage.NewResourceNotExist()}
	}
	return ipv4Pool, nil
}

// GetIPv4PoolCollection will get IPv4 pool collection.
func GetIPv4PoolCollection(start int, count int, filter string) (*model.IPv4PoolCollection, []commonMessage.Message) {
	dbImpl := db.GetPoolDB()
	ret, err := dbImpl.GetIPv4PoolCollection(start, count, filter)
	if err != nil {
		return nil, []commonMessage.Message{commonMessage.NewTransactionError()}
	}
	return ret, nil
}

// DeleteIPv4Pool will delete IPv4 pool by ID.
func DeleteIPv4Pool(id string) []commonMessage.Message {
	dbImpl := db.GetPoolDB()

	exist, previous, commited, err := dbImpl.DeleteIPv4Pool(id)

	if !exist {
		return []commonMessage.Message{commonMessage.NewResourceNotExist()}
	}
	if err != nil || !commited {
		return []commonMessage.Message{commonMessage.NewTransactionError()}
	}
	var poolDTO dto.GetIPv4PoolResponse
	poolDTO.Load(previous)
	wsSDK.DispatchResourceDeleteEvent(&poolDTO)
	return nil
}

// DeleteIPv4PoolCollection will delete all the IPv4 pool.
func DeleteIPv4PoolCollection() []commonMessage.Message {
	dbImpl := db.GetPoolDB()
	records, commited, err := dbImpl.DeleteIPv4PoolCollection()
	if err != nil || !commited {
		return []commonMessage.Message{commonMessage.NewTransactionError()}
	}
	for _, each := range records {
		eachDTO := dto.GetIPv4PoolResponse{}
		eachDTO.Load(&each)
		wsSDK.DispatchResourceDeleteEvent(&eachDTO)
	}
	wsSDK.DispatchResourceCollectionDeleteEvent(category.PoolIPv4)
	return nil
}

// AllocateIPv4Address will allocate an IP from pool.
func AllocateIPv4Address(id string, request dto.AllocateIPv4Request) (*model.IPv4Pool, []commonMessage.Message) {
	dbImpl := db.GetPoolDB()

	ipv4Pool := dbImpl.GetIPv4Pool(id)
	if ipv4Pool == nil {
		return nil, []commonMessage.Message{commonMessage.NewResourceNotExist()}
	}
	return ipv4Pool, nil
}

// FreeIPv4Address will free an IP from pool.
func FreeIPv4Address(id string, request dto.FreeIPv4Request) (*model.IPv4Pool, []commonMessage.Message) {
	dbImpl := db.GetPoolDB()

	ipv4Pool := dbImpl.GetIPv4Pool(id)
	if ipv4Pool == nil {
		return nil, []commonMessage.Message{commonMessage.NewResourceNotExist()}
	}
	return ipv4Pool, nil
}
