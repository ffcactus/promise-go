package message

const (
	ServerGroupExist = "ServerGroupExist"
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