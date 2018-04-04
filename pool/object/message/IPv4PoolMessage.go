package message

import (
	"promise/common/category"
	"promise/common/object/constvalue"
	commonMessage "promise/common/object/message"
)

const (
	// MessageIPv4PoolEmpty is message ID
	MessageIPv4PoolEmpty = "Server.Message.IPv4PoolEmpty"
)

// NewIPv4PoolEmpty will return a message.
func NewIPv4PoolEmpty() commonMessage.Message {
	ret := commonMessage.NewMessage(category.PoolIPv4)
	ret.ID = MessageIPv4PoolEmpty
	ret.Severity = constvalue.SeverityWarning
	ret.Description = "No more IPv4 address can be allocated."
	return ret
}
