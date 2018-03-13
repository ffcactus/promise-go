package rest

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	commonDto "promise/common/object/dto"

	log "github.com/sirupsen/logrus"
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
	log.Debug("rest(), method =", method, "URL =", url)
	// Form the REST request.
	req, err := http.NewRequest(method, url, request)
	if err != nil {
		log.Warn("NewRequest() failed", "Method =", method, "URL =", url, "error =", err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Warn("Do() failed", "Method =", method, "URL =", url, "error =", err)
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
		log.Warn("rest() failed,", "Method =", method, "URI =", uri, "error =", err)
		return nil, err
	}
	// Only when err == nil should the resp can be dereferenced.
	defer resp.Body.Close()
	// If is expected status code, turn the response to expectedDto, or turn to []Message.
	if isExpectedStatusCode(expectStatusCode, resp.StatusCode) {
		log.Debug("isExpectedStatusCode() = true, status code =", resp.StatusCode)
		if resp.Body == nil && responseDtoP != nil {
			log.Warn("Resposne body is empty, ", method, " URI = ", uri)
			return nil, errors.New("response body is empty")
		}
		if err := json.NewDecoder(resp.Body).Decode(responseDtoP); err != nil {
			log.Warn("Decode(responseDtoP) failed", "Method =", method, "URI =", uri, "error =", err)
			return nil, err
		}
		return nil, nil
	}
	log.Debug("isExpectedStatusCode() = false, status code =", resp.StatusCode)
	// TODO Not all the response body can be translate to messages.
	message := new([]commonDto.Message)
	if resp.Body == nil {
		log.Warn("Resposne body is empty", method, "URI =", uri)
		return nil, errors.New("response body is empty")
	}
	if err := json.NewDecoder(resp.Body).Decode(message); err != nil {
		log.Warn("Decode(message) failed", "Method =", method, "URI =", uri, "error =", err)
		return nil, err
	}
	return *message, nil
}
