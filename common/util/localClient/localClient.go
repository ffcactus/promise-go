package localClient

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	commonDto "promise/common/object/dto"

	"github.com/astaxie/beego"
)

// LocalClient The client that connect to localhost.
type LocalClient struct {
	Client *http.Client
}

// Instance Geth LocalClient instance.
func Instance() *LocalClient {
	return &LocalClient{
		Client: &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}},
	}
}

// The REST operation.
func (c *LocalClient) rest(method string, uri string, request io.Reader) (*http.Response, error) {
	url := c.address(uri)
	beego.Trace("rest(), method = ", method, ", url = ", url)
	// Form the REST request.
	req, err := http.NewRequest(method, url, request)
	if err != nil {
		beego.Warning("NewRequest() failed,", " Method = ", method, " URI = ", uri, " error = ", err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	resp, err := c.Client.Do(req)
	if err != nil {
		beego.Warning("Do() failed,", " Method = ", method, " URI = ", uri, " error = ", err)
		return nil, err
	}
	return resp, err
}

func (c *LocalClient) isExpectedStatusCode(expectStatusCode []int, realStatusCode int) bool {
	for i := range expectStatusCode {
		if expectStatusCode[i] == realStatusCode {
			return true
		}
	}
	return false
}

// Rest Get and parse the object.
func (c *LocalClient) Rest(method string, uri string, requestDto interface{}, responseDtoP interface{}, expectStatusCode []int) ([]commonDto.Message, error) {
	var (
		resp *http.Response
		err  error
	)
	if requestDto != nil {
		b := new(bytes.Buffer)
		json.NewEncoder(b).Encode(requestDto)
		resp, err = c.rest(method, uri, b)
	} else {
		resp, err = c.rest(method, uri, nil)
	}
	defer resp.Body.Close()
	if err != nil {
		beego.Warning("rest() failed,", " Method = ", method, " URI = ", uri, " error = ", err)
		return nil, err
	}
	defer resp.Body.Close()
	// If is expected status code, turn the response to expectedDto, or turn to []Message.
	if c.isExpectedStatusCode(expectStatusCode, resp.StatusCode) {
		beego.Trace("isExpectedStatusCode() = true, status code = ", resp.StatusCode)
		if resp.Body == nil && responseDtoP != nil {
			beego.Warning("Resposne body is empty, ", method, " URI = ", uri)
			return nil, errors.New("response body is empty")
		}
		if err := json.NewDecoder(resp.Body).Decode(responseDtoP); err != nil {
			beego.Warning("Decode(responseDtoP) failed,", " Method = ", method, " URI = ", uri, " error = ", err)
			return nil, err
		}
		return nil, nil
	}
	beego.Trace("isExpectedStatusCode() = false, status code = ", resp.StatusCode)
	// TODO Not all the response body can be translate to messages.
	message := new([]commonDto.Message)
	if resp.Body == nil {
		beego.Warning("Resposne body is empty, ", method, " URI = ", uri)
		return nil, errors.New("response body is empty")
	}
	if err := json.NewDecoder(resp.Body).Decode(message); err != nil {
		beego.Warning("Decode(message) failed,", " Method = ", method, " URI = ", uri, " error = ", err)
		return nil, err
	}
	return *message, nil
}

// Get the complete address.
func (c *LocalClient) address(uri string) string {
	var buf bytes.Buffer
	buf.WriteString("http://localhost:8081")
	buf.WriteString(uri)

	return buf.String()
}
