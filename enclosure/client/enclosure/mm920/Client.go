package mm920

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"promise/base"
	"promise/enclosure/object/model"
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
	Addresses      []string
	Username       string
	Password       string
	currentAddress string
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

// String returns the debug info for the client error.
func (e ClientErrorImpl) String() string {
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
func ToClientError(err error) base.ClientError {
	var errorImpl ClientErrorImpl
	errorImpl.responseError = err
	if err.Timeout() {
		errorImpl.timeout = true
	}
	return &errorImpl
}

// NewClient creates a client for enclosure.
func NewClient(enclosure *model.Enclosure) *Client {
	client := Client{}
	// TODO or get username and password from keystore.
	if enclosure.Credential.URL == "" {
		client.Username = enclosure.Credential.Username
		client.Password = enclosure.Credential.Password
	}
	if len(enclosure.Addresses) > 0 {
		client.currentAddress = enclosure.Addresses[0]
		client.Addresses = enclosure.Addresses
	}
	return &client
}

// GetRequest creates http GET request.
func (c Client) GetRequest(url string) (*http.Request, base.ClientError) {
	var errorImpl ClientErrorImpl

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
func (c Client) DeleteRequest(url string) (*http.Request, base.ClientError) {
	var errorImpl ClientErrorImpl

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
func (c Client) PostRequest(url string, dto interface{}) (*http.Request, base.ClientError) {
	var errorImpl ClientErrorImpl

	b := new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(dto); err != nil {
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
func (c Client) PatchRequest(url, etag string, dto interface{}) (*http.Request, base.ClientError) {
	var errorImpl ClientErrorImpl

	b := new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(dto); err != nil {
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
func (c Client) Unmarshal(resp *http.Response, dto interface{}) base.ClientError {
	var errorImpl ClientErrorImpl
	var body []byte

	defer resp.Body.Close()
	body, _ = ioutil.ReadAll(resp.Body)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		errorImpl.status = resp.StatusCode
		errorImpl.body = body
		return &errorImpl
	}
	if err := json.Unmarshal(body, dto); err != nil {
		errorImpl.status = resp.StatusCode
		errorImpl.body = body
		errorImpl.jsonError = err
		return &errorImpl
	}
	return nil
}

// Do is will call http.Client.Do() and unmarshal the response.
// It helps on unify the error process.
func (c Client) Do(request *http.Request, dto interface{}) base.ClientError {
	httpResponse, err := _client.Do(request)
	if err != nil {
		return ToClientError(err)
	}
	return c.Unmarshal(httpResponse, dto)
}

// Get do http GET to uri, and unmarshal the response to dto.
func (c Client) Get(uri string, dto interface{}) base.ClientError {
	if httpRequest, err := c.GetRequest(uri); err != nil {
		return err
	} else {
		return c.Do(httpRequest, dto)
	}
}

// String returns the debug info of the client.
func (c Client) String() string {
	return c.currentAddress
}

// Ready returns if the enclosure is ready.
func (c Client) Ready() bool {
	return false
}

// Claim should make make a flag on the enclosure that indicate it is exclusively managed by this system.
func (c Client) Claim() base.ClientError {
	return nil
}

// Unclaim should remove the flag that indicate the enclosure is managed by this system.
func (c Client) Unclaim() base.ClientError {
	return nil
}

// DeviceIdentity returns the device identity.
func (c Client) DeviceIdentity() (*base.DeviceIdentity, base.ClientError) {
	var (
		httpRequest      *http.Request
		httpResponse     *http.Response
		identity         base.DeviceIdentity
		redfishV1        GetRedfishV1Response
		chassisEnclosure GetChassisEnclosureResponse
	)

	// Get UUID.
	if err := c.Get("/redfish/v1", &redfishV1); err != nil {
		return nil, err
	}
	identity.UUID = redfishV1.UUID

	// Get SerialNumber and PartNumber
	if err := c.Get("/redfish/v1/Chassis/Enclosure", &chassisEnclosure); err != nil {
		return nil, err
	}
	identity.SerialNumber = chassisEnclosure.SerialNumber
	identity.PartNumber = chassisEnclosure.PartNumber

	return &identity, nil
}

// BladeSlot returns the blade slot info.
func (c Client) BladeSlot() ([]model.BladeSlot, base.ClientError) {
	log.WithFields(log.Fields{"client": c}).Info("Client get blade slot.")
	return nil, nil
}

// SwitchSlot returns the switch ade slot info.
func (c Client) SwitchSlot() ([]model.SwitchSlot, base.ClientError) {
	log.WithFields(log.Fields{"client": c}).Info("Client get switch slot.")
	return nil, nil
}

// FanSlot returns the fan slot info.
func (c Client) FanSlot() ([]model.FanSlot, base.ClientError) {
	log.WithFields(log.Fields{"client": c}).Info("Client get fan slot.")
	return nil, nil
}

// PowerSlot returns the power slot info.
func (c Client) PowerSlot() ([]model.PowerSlot, base.ClientError) {
	log.WithFields(log.Fields{"client": c}).Info("Client get power slot.")
	return nil, nil
}

// ManagerSlot returns the manager slot info.
func (c Client) ManagerSlot() ([]model.ManagerSlot, base.ClientError) {
	log.WithFields(log.Fields{"client": c}).Info("Client get manager slot.")
	return nil, nil
}

// ApplianceSlot returns the manager slot info.
func (c Client) ApplianceSlot() ([]model.ApplianceSlot, base.ClientError) {
	log.WithFields(log.Fields{"client": c}).Info("Client get appliance slot.")
	return nil, nil
}
