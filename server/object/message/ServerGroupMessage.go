package message

import (
	commonM "promise/common/object/model"
)

const (
	// MessageIDServerGroupExist is message ID.
	MessageIDServerGroupExist = "MessageIDServerGroupExist"
	// MessageIDServerGroupNotExist is message ID.
	MessageIDServerGroupNotExist = "MessageIDServerGroupNotExist"
)

// NewServerGroupExist return a new message.
func NewServerGroupExist() commonM.Message {
	ret := commonM.NewMessage(commonM.CategoryServer)
	ret.ID = MessageIDServerGroupExist
	ret.Severity = commonM.SeverityNormal
	ret.Description = "Server group already exists."
	return ret
}

// NewServerGroupNotExist return a new message.
func NewServerGroupNotExist() commonM.Message {
	ret := commonM.NewMessage(commonM.CategoryServer)
	ret.ID = MessageIDServerGroupNotExist
	ret.Severity = commonM.SeverityNormal
	ret.Description = "Server group not exists."
	return ret
}
