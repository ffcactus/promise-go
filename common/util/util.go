package util

import (
	"bytes"
	"encoding/json"
	"github.com/astaxie/beego"
	"log"
	"os"
)

func init() {
	Logger.SetOutput(os.Stdout)
}

var (
	logBuf bytes.Buffer
	Logger = log.New(&logBuf, "", log.LstdFlags|log.Lshortfile)
)

func PrintJson(data interface{}) {
	b, _ := json.MarshalIndent(data, "", "    ")
	beego.Info(string(b))
}

func StructToString(data interface{}) string {
	b, _ := json.Marshal(data)
	return string(b)
}

func StringToStruct(s string, p interface{}) error {
	return json.Unmarshal([]byte(s), p)
}
