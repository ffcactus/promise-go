package message

import (
	"promise/base"
)

const (
	// MessageIPv4PoolEmpty is message ID
	MessageIPv4PoolEmpty = "IPv4.Message.PoolEmpty"
	// MessageIPv4AddressNotExist is message ID.
	MessageIPv4AddressNotExist = "IPv4.Message.AddressNotExist"
	// MessageIPv4FormateError is message ID.
	MessageIPv4FormateError = "IPv4.Message.FormatError"
	// MessageIPv4RangeEndAddressError is message ID.
	MessageIPv4RangeEndAddressError = "IPv4.Message.RangeEndAddressError"
	// MessageIPv4RangeSizeError is message ID.
	MessageIPv4RangeSizeError = "IPv4.Message.RangeSizeError"
	// MessageIPv4RangeCountError is message ID.
	MessageIPv4RangeCountError = "IPv4.Message.RangeCountError"
	// MessageIPv4NotAllocatedError is message ID.
	MessageIPv4NotAllocatedError = "IPv4.Message.NotAllocatedError"
)

// NewMessageIPv4PoolEmpty will return a message.
func NewMessageIPv4PoolEmpty() *base.Message {
	ret := base.NewMessage()
	ret.ID = MessageIPv4PoolEmpty
	ret.Severity = base.SeverityWarning
	ret.Description = "No more IPv4 address can be allocated."
	return ret
}

// NewMessageIPv4AddressNotExistError will return a message.
func NewMessageIPv4AddressNotExistError() *base.Message {
	ret := base.NewMessage()
	ret.ID = MessageIPv4AddressNotExist
	ret.Severity = base.SeverityWarning
	ret.Description = "The address does not exist in this pool."
	return ret
}

// NewMessageIPv4FormatError will return a message.
func NewMessageIPv4FormatError() *base.Message {
	ret := base.NewMessage()
	ret.ID = MessageIPv4FormateError
	ret.Severity = base.SeverityWarning
	ret.Description = "Unknown IPv4 format."
	return ret
}

// NewMessageIPv4RangeEndAddressError will return a message.
func NewMessageIPv4RangeEndAddressError() *base.Message {
	ret := base.NewMessage()
	ret.ID = MessageIPv4RangeEndAddressError
	ret.Severity = base.SeverityWarning
	ret.Description = "The end address in a range should equal or big then start address"
	return ret
}

// NewMessageIPv4RangeSizeError will return a message.
func NewMessageIPv4RangeSizeError() *base.Message {
	ret := base.NewMessage()
	ret.ID = MessageIPv4RangeSizeError
	ret.Severity = base.SeverityWarning
	ret.Description = "The number of addresses in a range should not more than 256."
	return ret
}

// NewMessageIPv4RangeCountError will return a message.
func NewMessageIPv4RangeCountError() *base.Message {
	ret := base.NewMessage()
	ret.ID = MessageIPv4RangeCountError
	ret.Severity = base.SeverityWarning
	ret.Description = "IPv4 pool should contain one range at least."
	return ret
}

// NewMessageIPv4NotAllocatedError will return a message.
func NewMessageIPv4NotAllocatedError() *base.Message {
	ret := base.NewMessage()
	ret.ID = MessageIPv4NotAllocatedError
	ret.Severity = base.SeverityWarning
	ret.Description = "IP is not allocated before."
	return ret
}
