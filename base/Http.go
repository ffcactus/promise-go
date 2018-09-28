package base

import (
	"net/http"
	"bytes"
	"encoding/json"
)

// NewGetRequest creates http GET request.
func NewGetRequest(url string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// NewDeleteRequest creates http GET request.
func NewDeleteRequest(url string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}


// NewPostRequest creates http POST request.
func NewPostRequest(url string, dto interface{}) (*http.Request, error) {
	b := new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(dto); err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, b)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// NewPatchRequest creates http POST request.
func NewPatchRequest(url, etag string, dto interface{}) (*http.Request, error) {
	b := new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(dto); err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, b)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	if etag != "" {
		req.Header.Set("Etag", etag)
	}
	return req, nil
}