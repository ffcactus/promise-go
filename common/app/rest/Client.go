package rest

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

var (
	client = http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
)

func rest(method string, url string, request io.Reader) (*http.Response, error) {
	beego.Trace("rest(), method =", method, "URL =", url)
	// Form the REST request.
	req, err := http.NewRequest(method, url, request)
	if err != nil {
		beego.Warning("NewRequest() failed", "Method =", method, "URL =", url, "error =", err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		beego.Warning("Do() failed", "Method =", method, "URL =", url, "error =", err)
		return nil, err
	}
	return resp, err
}

func isExpectedStatusCode(expectStatusCode []int, realStatusCode int) bool {
	for i := range expectStatusCode {
		if expectStatusCode[i] == realStatusCode {
			return true
		}
	}
	return false
}

// Do Get and parse the object.
func Do(method string, uri string, requestDto interface{}, responseDtoP interface{}, expectStatusCode []int) ([]commonDto.Message, error) {
	var (
		resp *http.Response
		err  error
	)
	if requestDto != nil {
		b := new(bytes.Buffer)
		json.NewEncoder(b).Encode(requestDto)
		resp, err = rest(method, uri, b)
	} else {
		resp, err = rest(method, uri, nil)
	}
	if err != nil {
		beego.Warning("rest() failed,", "Method =", method, "URI =", uri, "error =", err)
		return nil, err
	}
	// Only when err == nil should the resp can be dereferenced.
	defer resp.Body.Close()
	// If is expected status code, turn the response to expectedDto, or turn to []Message.
	if isExpectedStatusCode(expectStatusCode, resp.StatusCode) {
		beego.Trace("isExpectedStatusCode() = true, status code =", resp.StatusCode)
		if resp.Body == nil && responseDtoP != nil {
			beego.Warning("Resposne body is empty, ", method, " URI = ", uri)
			return nil, errors.New("response body is empty")
		}
		if err := json.NewDecoder(resp.Body).Decode(responseDtoP); err != nil {
			beego.Warning("Decode(responseDtoP) failed", "Method =", method, "URI =", uri, "error =", err)
			return nil, err
		}
		return nil, nil
	}
	beego.Trace("isExpectedStatusCode() = false, status code =", resp.StatusCode)
	// TODO Not all the response body can be translate to messages.
	message := new([]commonDto.Message)
	if resp.Body == nil {
		beego.Warning("Resposne body is empty", method, "URI =", uri)
		return nil, errors.New("response body is empty")
	}
	if err := json.NewDecoder(resp.Body).Decode(message); err != nil {
		beego.Warning("Decode(message) failed", "Method =", method, "URI =", uri, "error =", err)
		return nil, err
	}
	return *message, nil
}

