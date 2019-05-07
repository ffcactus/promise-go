package dto

type PrefixLength int

type SubnetMask string

// This type describes an IPv4 Address.
type IPv4Address struct {
	Address       *string // This is the IPv4 Address.
	SubnetMask    *string // This is the IPv4 Subnet mask.
	AddressOrigin *string // This indicates how the address was determined.
	Gateway       *string // This is the IPv4 gateway for this address.
}

// This type describes an IPv6 Address.
type IPv6Address struct {
	Address       *string // This is the IPv6 Address.
	PrefixLength  *int    // This is the IPv6 Address Prefix Length.
	AddressOrigin *string // This indicates how the address was determined.
	AddressState  *string // The current state of this address as defined in RFC 4862.
}

// This object represents a single IPv6 static address to be assigned on a network interface.
type IPv6StaticAddress struct {
	Address      *string // A valid IPv6 address.
	PrefixLength *int    // The Prefix Length of this IPv6 address.
}
