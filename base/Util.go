package base

import (
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
