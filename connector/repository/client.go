package repository

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
	"net/url"
)

var _client = &http.Client{
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	},
}

// Client implements EnclosureClient interface.
type Client struct {
	Protocol       string
	Addresses      []string
	Username       string
	Password       string
	CurrentAddress string
}

// String returns the debug info of the client.
func (c Client) String() string {
	return c.CurrentAddress
}

// ClientErrorImpl holds the error info.
// ClientErrorImpl implements Error interface.
type ClientErrorImpl struct {
	status          int
	body            []byte
	connectionError bool
	jsonError       error
	requestError    error
	responseError   error
	timeout         bool
	loginFailure    bool
}

// Status returns the HTTP status code
func (e ClientErrorImpl) Status() int {
	return e.status
}

// Body returns the response body.
func (e ClientErrorImpl) Body() []byte {
	return e.body
}

// ConnectionError returns if it is connection error.
func (e ClientErrorImpl) ConnectionError() bool {
	return e.connectionError
}

// Timeout returns if it is timeout.
func (e ClientErrorImpl) Timeout() bool {
	return e.timeout
}

// LoginFailure returns if it is login failure.
func (e ClientErrorImpl) LoginFailure() bool {
	return e.loginFailure
}

// Error implements the error interface.
func (e ClientErrorImpl) Error() string {
	if e.jsonError != nil {
		return fmt.Sprintf("json error = %v", e.jsonError)
	}
	if e.requestError != nil {
		return fmt.Sprintf("request error = %v", e.requestError)
	}
	if e.responseError != nil {
		return fmt.Sprintf("response error = %v", e.requestError)
	}
	if e.timeout {
		return "timeout"
	}
	if e.connectionError {
		return fmt.Sprintf("(connection error, status = %d)", e.status)
	}
	if e.loginFailure {
		return fmt.Sprintf("(login failure, status = %d)", e.status)
	}
	return fmt.Sprintf("status = %d", e.status)
}

// ToClientError translate the error from http.Client.Do to ClientError.
func ToClientError(err error) ClientError {
	var errorImpl ClientErrorImpl
	errorImpl.responseError = err
	if urlError, ok := err.(*url.Error); ok {
		if urlError.Timeout() {
			errorImpl.timeout = true
		}
	}
	return &errorImpl
}

// GetRequest creates http GET request.
func (c Client) GetRequest(url string) (*http.Request, ClientError) {
	var errorImpl ClientErrorImpl

	url = c.Protocol + "://" + c.CurrentAddress + url
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		errorImpl.requestError = err
		return nil, &errorImpl
	}
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(c.Username, c.Password)
	return req, nil
}

// DeleteRequest creates http GET request.
func (c Client) DeleteRequest(url string) (*http.Request, ClientError) {
	var errorImpl ClientErrorImpl

	url = c.Protocol + "://" + c.CurrentAddress + url
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		errorImpl.requestError = err
		return nil, &errorImpl
	}
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(c.Username, c.Password)
	return req, nil
}

// PostRequest creates http POST request.
func (c Client) PostRequest(url string, request interface{}) (*http.Request, ClientError) {
	var errorImpl ClientErrorImpl

	url = c.Protocol + "://" + c.CurrentAddress + url
	b := new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(request); err != nil {
		errorImpl.jsonError = err
		return nil, &errorImpl
	}

	req, err := http.NewRequest(http.MethodPost, url, b)
	if err != nil {
		errorImpl.requestError = err
		return nil, &errorImpl
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(c.Username, c.Password)
	return req, nil
}

// PatchRequest creates http POST request.
func (c Client) PatchRequest(url, etag string, request interface{}) (*http.Request, ClientError) {
	var errorImpl ClientErrorImpl

	url = c.Protocol + "://" + c.CurrentAddress + url
	b := new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(request); err != nil {
		errorImpl.jsonError = err
		return nil, &errorImpl
	}

	req, err := http.NewRequest(http.MethodPost, url, b)
	if err != nil {
		errorImpl.requestError = err
		return nil, &errorImpl
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	if etag != "" {
		req.Header.Set("Etag", etag)
	}
	req.SetBasicAuth(c.Username, c.Password)
	return req, nil
}

// Unmarshal parse the http response to DTO in case the status is 2xx.
// It returns client error if parse failed, or the status is not 2xx.
func (c Client) Unmarshal(resp *http.Response, response interface{}) ClientError {
	var errorImpl ClientErrorImpl
	var body []byte

	defer resp.Body.Close()
	body, _ = ioutil.ReadAll(resp.Body)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		errorImpl.status = resp.StatusCode
		errorImpl.body = body
		return &errorImpl
	}
	if err := json.Unmarshal(body, response); err != nil {
		errorImpl.status = resp.StatusCode
		errorImpl.body = body
		errorImpl.jsonError = err
		return &errorImpl
	}
	return nil
}

// Do is will call http.Client.Do() and unmarshal the response.
// It helps on unify the error process.
func (c Client) Do(request *http.Request, response interface{}) ClientError {
	log.WithFields(log.Fields{"method": request.Method, "URL": request.URL}).Info("Client perform operation.")
	httpResponse, err := _client.Do(request)
	if err != nil {
		log.WithFields(log.Fields{"method": request.Method, "URL": request.URL, "error": err}).Warn("Client operation failed.")
		return ToClientError(err)
	}
	return c.Unmarshal(httpResponse, response)
}

// Get do http GET to uri, and unmarshal the response dto.
func (c Client) Get(uri string, response interface{}) ClientError {
	httpRequest, err := c.GetRequest(uri)
	if err != nil {
		return err
	}
	return c.Do(httpRequest, response)
}

// Post do http POST to uri, and unmarshal the response to dto.
func (c Client) Post(uri string, request, response interface{}) ClientError {
	httpRequest, err := c.PostRequest(uri, request)
	if err != nil {
		return err
	}
	return c.Do(httpRequest, response)
}
