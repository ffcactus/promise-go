package model

import (
	commonModel "promise/common/object/model"
)

// IPv4PoolMember is the member in collection.
type IPv4PoolMember struct {
	commonModel.PromiseMember
	Name string
}

// IPv4PoolCollection is the model of collection.
type IPv4PoolCollection struct {
	commonModel.PromiseCollection
	Members     []commonModel.PromiseMemberInterface
}

// // GetMembers get the members in the collection.
// func (m *IPv4PoolCollection) GetMembers() []commonModel.PromiseMemberInterface {
// 	return make([]commonModel.PromiseMemberInterface, len(m.Members))
// }
