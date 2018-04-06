package message

import (
	"promise/common/category"
	"promise/common/object/constvalue"
	commonMessage "promise/common/object/message"
)

const (
	// MessageIPv4PoolEmpty is message ID
	MessageIPv4PoolEmpty = "IPPool.Message.IPv4PoolEmpty"
	// MessageIPv4AddressNotExist is message ID.
	MessageIPv4AddressNotExist = "IPPool.Message.AddressNotExist"
	// MessageIPv4FormateError is message ID.
	MessageIPv4FormateError = "IPPool.Message.IPv4FormatError"
	// MessageIPv4RangeEndAddressError is message ID.
	MessageIPv4RangeEndAddressError = "IPPool.Message.IPv4RangeEndAddressError"
	// MessageIPv4RangeSizeError is message ID.
	MessageIPv4RangeSizeError = "IPPool.Message.IPv4RangeSizeError"
	// MessageIPv4RangeCountError is message ID.
	MessageIPv4RangeCountError = "IPPool.Message.IPv4RangeCountError"
	// MessageIPv4NotAllocatedError is message ID.
	MessageIPv4NotAllocatedError = "IPPool.Message.IPv4NotAllocatedError"
)

// NewIPv4PoolEmpty will return a message.
func NewIPv4PoolEmpty() commonMessage.Message {
	ret := commonMessage.NewMessage(category.PoolIPv4)
	ret.ID = MessageIPv4PoolEmpty
	ret.Severity = constvalue.SeverityWarning
	ret.Description = "No more IPv4 address can be allocated."
	return ret
}

// NewIPv4AddressNotExist will return a message.
func NewIPv4AddressNotExist() commonMessage.Message {
	ret := commonMessage.NewMessage(category.PoolIPv4)
	ret.ID = MessageIPv4AddressNotExist
	ret.Severity = constvalue.SeverityWarning
	ret.Description = "The address does not exist in this pool."
	return ret
}

// NewIPv4FormatError will return a message.
func NewIPv4FormatError() commonMessage.Message {
	ret := commonMessage.NewMessage(category.PoolIPv4)
	ret.ID = MessageIPv4FormateError
	ret.Severity = constvalue.SeverityWarning
	ret.Description = "Unknown IPv4 format."
	return ret
}

// NewIPv4RangeEndAddressError will return a message.
func NewIPv4RangeEndAddressError() commonMessage.Message {
	ret := commonMessage.NewMessage(category.PoolIPv4)
	ret.ID = MessageIPv4RangeEndAddressError
	ret.Severity = constvalue.SeverityWarning
	ret.Description = "The end address in a range should equal or big then start address"
	return ret
}

// NewIPv4RangeSizeError will return a message.
func NewIPv4RangeSizeError() commonMessage.Message {
	ret := commonMessage.NewMessage(category.PoolIPv4)
	ret.ID = MessageIPv4RangeSizeError
	ret.Severity = constvalue.SeverityWarning
	ret.Description = "The number of addresses in a range should not more than 256."
	return ret
}

// NewIPv4RangeCountError will return a message.
func NewIPv4RangeCountError() commonMessage.Message {
	ret := commonMessage.NewMessage(category.PoolIPv4)
	ret.ID = MessageIPv4RangeCountError
	ret.Severity = constvalue.SeverityWarning
	ret.Description = "IPv4 pool should contain one range at least."
	return ret
}

// NewIPv4NotAllocatedError will return a message.
func NewIPv4NotAllocatedError() commonMessage.Message {
	ret := commonMessage.NewMessage(category.PoolIPv4)
	ret.ID = MessageIPv4NotAllocatedError
	ret.Severity = constvalue.SeverityWarning
	ret.Description = "IP is not allocated before."
	return ret
}