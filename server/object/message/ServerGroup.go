package message

import (
	"net/http"
	commonM "promise/common/object/model"
)

const (
	// MessageIDServerGroupExist is message ID.
	MessageIDServerGroupExist = "MessageIDServerGroupExist"
)

// NewServerGroupExist return a new message.
func NewServerGroupExist() commonM.Message {
	ret := commonM.NewMessage(commonM.CategoryServer)
	ret.ID = MessageIDServerGroupExist
	ret.StatusCode = http.StatusOK
	ret.Severity = commonM.SeverityNormal
	ret.Description = "Server group already exists."
	return ret
}