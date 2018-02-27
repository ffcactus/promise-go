package context

import (
	commonDto "promise/common/object/dto"
	commonM "promise/common/object/model"
)

// ErrorHandlerInterface The interface of error handler.
type ErrorHandlerInterface interface {
	AppendErrorMessage(message commonM.Message)
	ErrorMessages() *[]commonM.Message
}

// ErrorHandler The implementation of error handler.
type ErrorHandler struct {
	messages []commonM.Message
}

// AppendMessage Append message.
func (c *ServerContext) AppendMessage(message commonM.Message) {
	c.messages = append(c.messages, message)
}

// AppendMessages Append messages.
func (c *ServerContext) AppendMessages(messages []commonDto.Message) {
	for i := range messages {
		c.messages = append(c.messages, *messages[i].Model())
	}
}

// Messages Get messages.
func (c *ServerContext) Messages() []commonM.Message {
	return c.messages
}
