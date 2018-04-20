package context

import (
	commonDto "promise/common/object/dto"
	commonMessage "promise/common/object/message"
)

// ErrorHandlerInterface The interface of error handler.
type ErrorHandlerInterface interface {
	AppendErrorMessage(message commonMessage.Message)
	ErrorMessages() *[]commonMessage.Message
}

// ErrorHandler The implementation of error handler.
type ErrorHandler struct {
	messages []commonMessage.Message
}

// AppendMessage Append message.
func (c *ServerContext) AppendMessage(message commonMessage.Message) {
	c.messages = append(c.messages, message)
}

// AppendMessages Append messages.
func (c *ServerContext) AppendMessages(messages []commonDto.Message) {
	for i := range messages {
		c.messages = append(c.messages, *messages[i].Model())
	}
}

// Messages Get messages.
func (c *ServerContext) Messages() []commonMessage.Message {
	return c.messages
}
