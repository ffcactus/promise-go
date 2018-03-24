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

func rest(method string, uri string, request io.Reader) (*http.Response, error) {
	log.WithFields(log.Fields{"method": method, "uri": uri}).Debug("rest() call.")
	// Form the REST request.
	req, err := http.NewRequest(method, uri, request)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
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
	log.WithFields(log.Fields{"method": method, "uri": uri}).Debug("Start REST call.")
	if requestDto != nil {
		b := new(bytes.Buffer)
		json.NewEncoder(b).Encode(requestDto)
		resp, err = rest(method, uri, b)
	} else {
		resp, err = rest(method, uri, nil)
	}
	if err != nil {
		log.WithFields(log.Fields{"method": method, "uri": uri, "err": err}).Warn("rest() call failed.")
		return nil, err
	}
	// Only when err == nil should the resp can be dereferenced.
	defer resp.Body.Close()
	// If is expected status code, turn the response to expectedDto, or turn to []Message.
	if isExpectedStatusCode(expectStatusCode, resp.StatusCode) {
		if resp.Body == nil && responseDtoP != nil {
			log.WithFields(log.Fields{"method": method, "uri": uri}).Warn("REST call failed, response body is empty")
			return nil, errors.New("response body is empty")
		}
		if resp.Body == nil || responseDtoP == nil {
			return nil, nil
		}
		if err := json.NewDecoder(resp.Body).Decode(responseDtoP); err != nil {
			log.WithFields(log.Fields{"method": method, "uri": uri}).Warn("REST call failed, can not decode response.")
			return nil, err
		}
		return nil, nil
	}
	log.WithFields(log.Fields{"method": method, "uri": uri, "status": resp.StatusCode, "expect": expectStatusCode}).Info("Not the expected http status code.")
	// TODO Not all the response body can be translate to messages.
	message := new([]commonDto.Message)
	if resp.Body == nil {
		log.WithFields(log.Fields{"method": method, "uri": uri}).Warn("REST call failed, message is empty.")
		return nil, errors.New("response body is empty")
	}
	if err := json.NewDecoder(resp.Body).Decode(message); err != nil {
		log.WithFields(log.Fields{"method": method, "uri": uri}).Warn("REST call failed, can not decode message.")
		return nil, err
	}
	return *message, nil
}
