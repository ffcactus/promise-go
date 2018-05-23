package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type AddServerRequest struct {
	Name     string `json:"Name"`
	Hostname  string `json:"Hostname"`
	Username string `json:"Username"`
	Password string `json:"Password"`
}

type AddServerResponse struct {
	ID string `json:"ID"`
}

type AddServerGroupRequest struct {
	Name string `json:"Name"`
}

type AddServerGroupResponse struct {
	ID string `json"ID"`
}

type AddSSGRequestRequest struct {
	ServerID string `json:"ServerID"`
	ServerGroupID string `json:"ServerGroupID"`
}

func main() {
	var (
		count  = 1000
		serverID = make([]string, count)
		groupID = make([]string, 10)
		Client = &http.Client{}
	)
	groupID = make([]string, count)
	// Create 10 server group.
	for i := 0; i < 10; i++ {
		dto := AddServerGroupRequest{
			Name: "Group" + fmt.Sprintf(" 0-%d", (9 - i) * 100 - 1),
		}
		b := new(bytes.Buffer)
		json.NewEncoder(b).Encode(dto)
		req, _ := http.NewRequest(http.MethodPost, "http://localhost/promise/v1/servergroup", b)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "application/json")
		if resp, err := Client.Do(req); err != nil {
			fmt.Printf("Post group %s failed. error = %s, response = %#v\n", dto.Name, err, resp)
			return
		} else {
			if resp.StatusCode != http.StatusCreated {
				fmt.Printf("post group %s failed, status code = %d\n", dto.Name, resp.StatusCode)
			} else {
				respDTO := new(AddServerResponse)
				if err := json.NewDecoder(resp.Body).Decode(respDTO); err != nil {
					fmt.Printf("failed to get server response. err = %v\n", err)
					return
				}
				groupID[i] = respDTO.ID
				fmt.Printf("post group %s done.\n", dto.Name)				
			}
			resp.Body.Close()
		}		
	}

	for i := 0; i < count; i++ {
		hostname := "Mock" + fmt.Sprintf("%05d", i)
		dto := AddServerRequest{
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
			if resp.StatusCode != http.StatusAccepted {
				fmt.Printf("post server %s failed, status code = %d\n", hostname, resp.StatusCode)
			} else {
				respDTO := new(AddServerGroupResponse)
				if err := json.NewDecoder(resp.Body).Decode(respDTO); err != nil {
					fmt.Printf("failed to get group response.\n")
					return
				}
				serverID[i] = respDTO.ID
				fmt.Printf("post server %s done.\n", hostname)				
			}
			resp.Body.Close()
		}
	}

	for i := 0; i < 900; i++ {
		dto := AddSSGRequestRequest{
			ServerID: serverID[i],
			ServerGroupID: groupID[0],
		}
		b := new(bytes.Buffer)
		json.NewEncoder(b).Encode(dto)
		req, _ := http.NewRequest(http.MethodPost, "http://localhost/promise/v1/server-servergroup", b)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "application/json")
		if resp, err := Client.Do(req); err != nil {
			fmt.Printf("Post SSG failed. error = %s, response = %#v\n", err, resp)
		} else {
			if resp.StatusCode != http.StatusCreated {
				fmt.Printf("post SSG failed, status code = %d\n", resp.StatusCode)
			}
			resp.Body.Close()
		}		
	}

	for i := 0; i < 800; i++ {
		dto := AddSSGRequestRequest{
			ServerID: serverID[i],
			ServerGroupID: groupID[1],
		}
		b := new(bytes.Buffer)
		json.NewEncoder(b).Encode(dto)
		req, _ := http.NewRequest(http.MethodPost, "http://localhost/promise/v1/server-servergroup", b)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "application/json")
		if resp, err := Client.Do(req); err != nil {
			fmt.Printf("Post SSG failed. error = %s, response = %#v\n", err, resp)
		} else {
			if resp.StatusCode != http.StatusCreated {
				fmt.Printf("post SSG failed, status code = %d\n", resp.StatusCode)
			}
			resp.Body.Close()
		}		
	}
	
	for i := 0; i < 700; i++ {
		dto := AddSSGRequestRequest{
			ServerID: serverID[i],
			ServerGroupID: groupID[2],
		}
		b := new(bytes.Buffer)
		json.NewEncoder(b).Encode(dto)
		req, _ := http.NewRequest(http.MethodPost, "http://localhost/promise/v1/server-servergroup", b)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "application/json")
		if resp, err := Client.Do(req); err != nil {
			fmt.Printf("Post SSG failed. error = %s, response = %#v\n", err, resp)
		} else {
			if resp.StatusCode != http.StatusCreated {
				fmt.Printf("post SSG failed, status code = %d\n", resp.StatusCode)
			}
			resp.Body.Close()
		}		
	}
	
	for i := 0; i < 600; i++ {
		dto := AddSSGRequestRequest{
			ServerID: serverID[i],
			ServerGroupID: groupID[3],
		}
		b := new(bytes.Buffer)
		json.NewEncoder(b).Encode(dto)
		req, _ := http.NewRequest(http.MethodPost, "http://localhost/promise/v1/server-servergroup", b)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "application/json")
		if resp, err := Client.Do(req); err != nil {
			fmt.Printf("Post SSG failed. error = %s, response = %#v\n", err, resp)
		} else {
			if resp.StatusCode != http.StatusCreated {
				fmt.Printf("post SSG failed, status code = %d\n", resp.StatusCode)
			}
			resp.Body.Close()
		}		
	}
	
	for i := 0; i < 500; i++ {
		dto := AddSSGRequestRequest{
			ServerID: serverID[i],
			ServerGroupID: groupID[4],
		}
		b := new(bytes.Buffer)
		json.NewEncoder(b).Encode(dto)
		req, _ := http.NewRequest(http.MethodPost, "http://localhost/promise/v1/server-servergroup", b)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "application/json")
		if resp, err := Client.Do(req); err != nil {
			fmt.Printf("Post SSG failed. error = %s, response = %#v\n", err, resp)
		} else {
			if resp.StatusCode != http.StatusCreated {
				fmt.Printf("post SSG failed, status code = %d\n", resp.StatusCode)
			}
			resp.Body.Close()
		}		
	}
	
	for i := 0; i < 400; i++ {
		dto := AddSSGRequestRequest{
			ServerID: serverID[i],
			ServerGroupID: groupID[5],
		}
		b := new(bytes.Buffer)
		json.NewEncoder(b).Encode(dto)
		req, _ := http.NewRequest(http.MethodPost, "http://localhost/promise/v1/server-servergroup", b)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "application/json")
		if resp, err := Client.Do(req); err != nil {
			fmt.Printf("Post SSG failed. error = %s, response = %#v\n", err, resp)
		} else {
			if resp.StatusCode != http.StatusCreated {
				fmt.Printf("post SSG failed, status code = %d\n", resp.StatusCode)
			}
			resp.Body.Close()
		}		
	}
	
	for i := 0; i < 300; i++ {
		dto := AddSSGRequestRequest{
			ServerID: serverID[i],
			ServerGroupID: groupID[6],
		}
		b := new(bytes.Buffer)
		json.NewEncoder(b).Encode(dto)
		req, _ := http.NewRequest(http.MethodPost, "http://localhost/promise/v1/server-servergroup", b)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "application/json")
		if resp, err := Client.Do(req); err != nil {
			fmt.Printf("Post SSG failed. error = %s, response = %#v\n", err, resp)
		} else {
			if resp.StatusCode != http.StatusCreated {
				fmt.Printf("post SSG failed, status code = %d\n", resp.StatusCode)
			}
			resp.Body.Close()
		}		
	}
	
	for i := 0; i < 200; i++ {
		dto := AddSSGRequestRequest{
			ServerID: serverID[i],
			ServerGroupID: groupID[7],
		}
		b := new(bytes.Buffer)
		json.NewEncoder(b).Encode(dto)
		req, _ := http.NewRequest(http.MethodPost, "http://localhost/promise/v1/server-servergroup", b)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "application/json")
		if resp, err := Client.Do(req); err != nil {
			fmt.Printf("Post SSG failed. error = %s, response = %#v\n", err, resp)
		} else {
			if resp.StatusCode != http.StatusCreated {
				fmt.Printf("post SSG failed, status code = %d\n", resp.StatusCode)
			}
			resp.Body.Close()
		}		
	}
	
	for i := 0; i < 100; i++ {
		dto := AddSSGRequestRequest{
			ServerID: serverID[i],
			ServerGroupID: groupID[8],
		}
		b := new(bytes.Buffer)
		json.NewEncoder(b).Encode(dto)
		req, _ := http.NewRequest(http.MethodPost, "http://localhost/promise/v1/server-servergroup", b)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "application/json")
		if resp, err := Client.Do(req); err != nil {
			fmt.Printf("Post SSG failed. error = %s, response = %#v\n", err, resp)
		} else {
			if resp.StatusCode != http.StatusCreated {
				fmt.Printf("post SSG failed, status code = %d\n", resp.StatusCode)
			}
			resp.Body.Close()
		}		
	}	
}
