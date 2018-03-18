package message

import (
	commonM "promise/common/object/model"
)

const (
	// MessageIDGroupExist is message ID.
	MessageIDGroupExist = "MessageIDGroupExist"
	// MessageIDGroupNotExist is message ID.
	MessageIDGroupNotExist = "MessageIDGroupNotExist"
)

// NewGroupExist return a new message.
func NewGroupExist() commonM.Message {
	ret := commonM.NewMessage(commonM.CategoryServer)
	ret.ID = MessageIDGroupExist
	ret.Severity = commonM.SeverityNormal
	ret.Description = "Server group already exists."
	return ret
}

// NewGroupNotExist return a new message.
func NewGroupNotExist() commonM.Message {
	ret := commonM.NewMessage(commonM.CategoryServer)
	ret.ID = MessageIDGroupNotExist
	ret.Severity = commonM.SeverityNormal
	ret.Description = "Server group not exists."
	return ret
}
