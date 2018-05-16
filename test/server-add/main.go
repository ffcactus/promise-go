package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type addServerRequest struct {
	Name     string `json:"Name"`
	Hostname  string `json:"Hostname"`
	Username string `json:"Username"`
	Password string `json:"Password"`
}

func main() {
	var (
		count  = 1
		Client = &http.Client{}
	)

	for i := 0; i < count; i++ {
		time.Sleep(time.Duration(0) * time.Millisecond)
		hostname := "Mock" + fmt.Sprintf("%05d", i)
		dto := addServerRequest{
			Name: hostname,
			Hostname: hostname,
			Username: hostname,
			Password: hostname,
		}
		b := new(bytes.Buffer)
		json.NewEncoder(b).Encode(dto)
		req, _ := http.NewRequest(http.MethodPost, "http://localhost/promise/v1/server/action/discover", b)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "application/json")
		if resp, err := Client.Do(req); err != nil {
			fmt.Printf("Post server %s failed. error = %s, response = %#v\n", hostname, err, resp)
		} else {
			resp.Body.Close()
			if resp.StatusCode != http.StatusAccepted {
				fmt.Printf("post server %s failed, status code = %d\n", hostname, resp.StatusCode)
			} else {
				fmt.Printf("post server %s done.\n", hostname)
			}
		}
	}
}
