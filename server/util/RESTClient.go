package util

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type RESTClient struct {
	Client       *http.Client
	Address      string
	Username     string
	Password     string
	UseBasicAuth bool
}

func GetInstance(address string, username string, password string, useBasicAuth bool) *RESTClient {
	return &RESTClient{
		Client:       &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}},
		Address:      address,
		Username:     username,
		Password:     password,
		UseBasicAuth: useBasicAuth,
	}
}

// The REST operation.
func (this *RESTClient) Rest(method string, uri string, body io.Reader) (resp *http.Response, err error) {
	url := this.address(uri)
	log.Info("Http REST ", method, " to ", url)
	// Form the REST request.
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Debug("REST %s %s create request failed, error = %#v\n", method, url, err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	// For basic auth, we pend credential first.
	if this.UseBasicAuth {
		// For basic auth.
		log.Debug("Use basic auth with %s\n", this.Username)
		req.SetBasicAuth(this.Username, this.Password)
	}
	if resp, err := this.Client.Do(req); err != nil {
		log.Debug("REST %s %s operation failed, error = %#v\n", method, url, err)
		return nil, err
	} else {
		return resp, err
	}
}

// Get and parse the object.
func (this *RESTClient) GetObject(uri string, req interface{}, o interface{}) (int, error) {
	var (
		resp *http.Response
		err  error
	)
	if req != nil {
		b := new(bytes.Buffer)
		json.NewEncoder(b).Encode(req)
		resp, err = this.Rest(http.MethodGet, uri, b)
	} else {
		resp, err = this.Rest(http.MethodGet, uri, nil)
	}
	defer resp.Body.Close()
	if err != nil {
		log.Debug("Get object from %s failed, error = %#v\n", uri, err)
		return resp.StatusCode, err
	} else {
		if resp.Body == nil {
			log.Debug("Get object from %s failed, resposne body is empty.\n", uri)
			return resp.StatusCode, errors.New("Response body is empty.")
		}
		if err := json.NewDecoder(resp.Body).Decode(o); err != nil {
			log.Debug("Get object from %s failed, failed to decode the object, error = %#v\n", uri, err)
			return resp.StatusCode, err
		}
		return resp.StatusCode, nil
	}
}

func (this *RESTClient) PostObject(uri string, req interface{}, o interface{}) (int, error) {
	var (
		resp *http.Response
		err  error
	)
	if req != nil {
		b := new(bytes.Buffer)
		json.NewEncoder(b).Encode(req)
		resp, err = this.Rest(http.MethodPost, uri, b)
	} else {
		resp, err = this.Rest(http.MethodPost, uri, nil)
	}
	defer resp.Body.Close()
	if err != nil {
		log.Debug("Post object to %s failed, error = %#v\n", uri, err)
		return resp.StatusCode, err
	} else {
		if resp.Body == nil {
			return resp.StatusCode, nil
		}
		// Decode only when the client asked for.
		if o != nil {
			if err := json.NewDecoder(resp.Body).Decode(o); err != nil {
				log.Debug("Post object to %s failed, failed to decode the response, error = %#v\n", uri, err)
				return resp.StatusCode, err
			}
		}
		return resp.StatusCode, nil
	}
}

func (this *RESTClient) PutObject(uri string, req interface{}, o interface{}) (int, error) {
	var (
		resp *http.Response
		err  error
	)
	if req != nil {
		b := new(bytes.Buffer)
		json.NewEncoder(b).Encode(req)
		resp, err = this.Rest(http.MethodPut, uri, b)
	} else {
		resp, err = this.Rest(http.MethodPut, uri, nil)
	}
	defer resp.Body.Close()
	if err != nil {
		log.Debug("Put object to %s failed, error = %#v\n", uri, err)
		return resp.StatusCode, err
	} else {
		if resp.Body == nil {
			return resp.StatusCode, nil
		}
		// Decode only when the client asked for.
		if o != nil {
			if err := json.NewDecoder(resp.Body).Decode(o); err != nil {
				log.Debug("Put object to %s failed, failed to decode the response, error = %#v\n", uri, err)
				return resp.StatusCode, err
			}
		}
		return resp.StatusCode, nil
	}
}

func (this *RESTClient) DeleteObject(uri string, req interface{}, o interface{}) (int, error) {
	var (
		resp *http.Response
		err  error
	)
	if req != nil {
		b := new(bytes.Buffer)
		json.NewEncoder(b).Encode(req)
		resp, err = this.Rest(http.MethodDelete, uri, b)
	} else {
		resp, err = this.Rest(http.MethodDelete, uri, nil)
	}
	defer resp.Body.Close()
	if err != nil {
		log.Debug("Delete object from %s failed, error = %#v\n", uri, err)
		return resp.StatusCode, err
	} else {
		if resp.Body == nil {
			return resp.StatusCode, nil
		}
		// Decode only when the client asked for.
		if o != nil {
			if err := json.NewDecoder(resp.Body).Decode(o); err != nil {
				log.Debug("Delete object to %s failed, failed to decode the response, error = %#v\n", uri, err)
				return resp.StatusCode, err
			}
		}
		return resp.StatusCode, nil
	}
}

// Get the complete address.
func (this *RESTClient) address(uri string) string {
	var buf bytes.Buffer
	buf.WriteString(this.Address)
	buf.WriteString(uri)

	return buf.String()
}
