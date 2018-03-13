package main

import (
	"bytes"
	"encoding/json"
	"github.com/astaxie/beego"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type AddServerRequest struct {
	Address  string `json:"address"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	var (
		count  int          = 100
		Client *http.Client = &http.Client{}
	)

	for i := 0; i < count; i++ {
		time.Sleep(time.Duration(0) * time.Millisecond)
		address := "Mock" + Logger.Sprintf("%05d", i)
		dto := AddServerRequest{
			Address: address,
		}
		b := new(bytes.Buffer)
		json.NewEncoder(b).Encode(dto)
		req, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/director/rich/v1/server", b)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "application/json")
		if resp, err := Client.Do(req); err != nil {
			log.Info("Post server %s failed. error = %s, response = %#v\n", address, err, resp)
		} else {
			resp.Body.Close()
			if resp.StatusCode != http.StatusOK {
				log.Info("post server %s failed, status code = %d\n", address, resp.StatusCode)
			} else {
				log.Info("post server %s done.\n", address)
			}
		}
	}
}
