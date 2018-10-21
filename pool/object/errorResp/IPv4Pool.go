package errorResp

import (
	"promise/base"
)

const (
	// ErrorResponseIPv4PoolEmpty is errorResp ID
	ErrorResponseIPv4PoolEmpty = "IPv4.ErrorResponse.PoolEmpty"
	// ErrorResponseIPv4AddressNotExist is errorResp ID.
	ErrorResponseIPv4AddressNotExist = "IPv4.ErrorResponse.AddressNotExist"
	// ErrorResponseIPv4FormateError is errorResp ID.
	ErrorResponseIPv4FormateError = "IPv4.ErrorResponse.FormatError"
	// ErrorResponseIPv4RangeEndAddressError is errorResp ID.
	ErrorResponseIPv4RangeEndAddressError = "IPv4.ErrorResponse.RangeEndAddressError"
	// ErrorResponseIPv4RangeSizeError is errorResp ID.
	ErrorResponseIPv4RangeSizeError = "IPv4.ErrorResponse.RangeSizeError"
	// ErrorResponseIPv4RangeCountError is errorResp ID.
	ErrorResponseIPv4RangeCountError = "IPv4.ErrorResponse.RangeCountError"
	// ErrorResponseIPv4NotAllocatedError is errorResp ID.
	ErrorResponseIPv4NotAllocatedError = "IPv4.ErrorResponse.NotAllocatedError"
)

// NewErrorResponseIPv4PoolEmpty will return an error response.
func NewErrorResponseIPv4PoolEmpty() *base.ErrorResponse {
	ret := base.NewErrorResponse()
	ret.ID = ErrorResponseIPv4PoolEmpty
	ret.Severity = base.SeverityWarning
	ret.Description = "No more IPv4 address can be allocated."
	return ret
}

// NewErrorResponseIPv4AddressNotExistError will return an error response.
func NewErrorResponseIPv4AddressNotExistError() *base.ErrorResponse {
	ret := base.NewErrorResponse()
	ret.ID = ErrorResponseIPv4AddressNotExist
	ret.Severity = base.SeverityWarning
	ret.Description = "The address does not exist in this pool."
	return ret
}

// NewErrorResponseIPv4FormatError will return an error response.
func NewErrorResponseIPv4FormatError() *base.ErrorResponse {
	ret := base.NewErrorResponse()
	ret.ID = ErrorResponseIPv4FormateError
	ret.Severity = base.SeverityWarning
	ret.Description = "Unknown IPv4 format."
	return ret
}

// NewErrorResponseIPv4RangeEndAddressError will return an error response.
func NewErrorResponseIPv4RangeEndAddressError() *base.ErrorResponse {
	ret := base.NewErrorResponse()
	ret.ID = ErrorResponseIPv4RangeEndAddressError
	ret.Severity = base.SeverityWarning
	ret.Description = "The end address in a range should equal or big then start address"
	return ret
}

// NewErrorResponseIPv4RangeSizeError will return an error response.
func NewErrorResponseIPv4RangeSizeError() *base.ErrorResponse {
	ret := base.NewErrorResponse()
	ret.ID = ErrorResponseIPv4RangeSizeError
	ret.Severity = base.SeverityWarning
	ret.Description = "The number of addresses in a range should not more than 256."
	return ret
}

// NewErrorResponseIPv4RangeCountError will return an error response.
func NewErrorResponseIPv4RangeCountError() *base.ErrorResponse {
	ret := base.NewErrorResponse()
	ret.ID = ErrorResponseIPv4RangeCountError
	ret.Severity = base.SeverityWarning
	ret.Description = "IPv4 pool should contain one range at least."
	return ret
}

// NewErrorResponseIPv4NotAllocatedError will return an error response.
func NewErrorResponseIPv4NotAllocatedError() *base.ErrorResponse {
	ret := base.NewErrorResponse()
	ret.ID = ErrorResponseIPv4NotAllocatedError
	ret.Severity = base.SeverityWarning
	ret.Description = "IP is not allocated before."
	return ret
}
