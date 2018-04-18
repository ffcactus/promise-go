package base

import (
	"bytes"
	"net"
	"strconv"
	"strings"
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

// PrintJSON will log the object.
func PrintJSON(data interface{}) {
	b, _ := json.MarshalIndent(data, "", "    ")
	log.Info(string(b))
}

// StructToString will turn struct to string.
func StructToString(data interface{}) string {
	b, _ := json.Marshal(data)
	return string(b)
}

// StringToStruct will turn string to struct.
func StringToStruct(s string, p interface{}) error {
	return json.Unmarshal([]byte(s), p)
}

// IPtoInt convert ip to int64.
func IPtoInt(ip net.IP) int64 {
	bits := strings.Split(ip.String(), ".")

	b0, _ := strconv.Atoi(bits[0])
	b1, _ := strconv.Atoi(bits[1])
	b2, _ := strconv.Atoi(bits[2])
	b3, _ := strconv.Atoi(bits[3])

	var sum int64

	sum += int64(b0) << 24
	sum += int64(b1) << 16
	sum += int64(b2) << 8
	sum += int64(b3)

	return sum
}

// IntToIP convert int64 ip to net.IP
func IntToIP(ip int64) net.IP {
	var bytes [4]byte
	bytes[0] = byte(ip & 0xFF)
	bytes[1] = byte((ip >> 8) & 0xFF)
	bytes[2] = byte((ip >> 16) & 0xFF)
	bytes[3] = byte((ip >> 24) & 0xFF)

	return net.IPv4(bytes[3], bytes[2], bytes[1], bytes[0])
}

// IPBetween check if the test ip included or between 'from' and 'to' net.IP.
func IPBetween(from net.IP, to net.IP, test net.IP) bool {
	if from == nil || to == nil || test == nil {
		return false
	}

	from16 := from.To16()
	to16 := to.To16()
	test16 := test.To16()
	if from16 == nil || to16 == nil || test16 == nil {
		return false
	}

	if bytes.Compare(test16, from16) >= 0 && bytes.Compare(test16, to16) <= 0 {
		return true
	}
	return false
}

// IPStringBetween check if the test ip included or between 'from' and 'to'.
func IPStringBetween(from string, to string, test string) bool {
	_from := net.ParseIP(from)
	_to := net.ParseIP(to)
	_test := net.ParseIP(test)
	return IPBetween(_from, _to, _test)
}
