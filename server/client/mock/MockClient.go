package mock

import (
	"fmt"
	"math/rand"
	"promise/server/object/model"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/google/uuid"
	"promise/server/object/constValue"
)

var (
	delay = 5000
)

// MockClient is a mock client.
type MockClient struct {
	Hostname string
}

// GetInstance will return a mock client.
func GetInstance(address string) *MockClient {
	delay, _ = beego.AppConfig.Int("MockClientDelay")
	return &MockClient{
		Hostname: address,
	}
}

func mockDelay() {
	time.Sleep(time.Duration(delay) * time.Millisecond)
}

// Support return if the server support this client.
func (c *MockClient) Support() bool {
	if strings.HasPrefix(c.Hostname, constValue.MockType) {
		return true
	}
	return false

}

// GetProtocol will the the protocol
func (c *MockClient) GetProtocol() string {
	return constValue.MockProtocol
}

// GetBasicInfo return the basic info.
func (c *MockClient) GetBasicInfo() (*model.ServerBasicInfo, error) {
	ret := model.ServerBasicInfo{}
	ret.OriginURIs.Chassis = randString()
	ret.OriginURIs.System = randString()
	ret.PhysicalUUID = uuid.New().String()
	ret.Name = c.Hostname
	ret.Description = *randString()
	ret.Type = constValue.MockType
	ret.Protocol = constValue.MockProtocol
	return &ret, nil
}

// CreateManagementAccount is a mock method.
func (c *MockClient) CreateManagementAccount(username string, password string) error {
	return nil
}

// GetProcessors is a mock method.
func (c *MockClient) GetProcessors(systemID string) ([]model.Processor, error) {
	var ret []model.Processor
	ret = append(ret, *randProcessor("1"))
	ret = append(ret, *randProcessor("2"))
	mockDelay()
	return ret, nil
}

// GetMemory is a mock method.
func (c *MockClient) GetMemory(systemID string) ([]model.Memory, error) {
	var ret []model.Memory
	ret = append(ret, *randMemory("1"))
	ret = append(ret, *randMemory("2"))
	mockDelay()
	return ret, nil
}

// GetEthernetInterfaces is a mock method.
func (c *MockClient) GetEthernetInterfaces(systemID string) ([]model.EthernetInterface, error) {
	var ret []model.EthernetInterface
	ret = append(ret, *randEthernetInterface("eth0"))
	ret = append(ret, *randEthernetInterface("eth1"))
	mockDelay()
	return ret, nil
}

// GetNetworkInterfaces is a mock method.
func (c *MockClient) GetNetworkInterfaces(systemID string) ([]model.NetworkInterface, error) {
	var ret []model.NetworkInterface
	ret = append(ret, *randNetworkInterface("NetworkInterface0"))
	ret = append(ret, *randNetworkInterface("NetworkInterface1"))
	mockDelay()
	return ret, nil
}

// GetStorages is a mock method.
func (c *MockClient) GetStorages(systemID string) ([]model.Storage, error) {
	var ret []model.Storage
	ret = append(ret, *randStorage("Storage0"))
	ret = append(ret, *randStorage("Storage1"))
	mockDelay()
	return ret, nil
}

// GetPower is a mock method.
func (c *MockClient) GetPower(chassisID string) (*model.Power, error) {
	return &model.Power{}, nil
}

// GetThermal is a mock method.
func (c *MockClient) GetThermal(chassisID string) (*model.Thermal, error) {
	return &model.Thermal{}, nil
}

// GetOemHuaweiBoards is a mock method.
func (c *MockClient) GetOemHuaweiBoards(chassisID string) ([]model.OemHuaweiBoard, error) {
	return []model.OemHuaweiBoard{}, nil
}

// GetNetworkAdapters is a mock method.
func (c *MockClient) GetNetworkAdapters(chassisID string) ([]model.NetworkAdapter, error) {
	var ret []model.NetworkAdapter
	return ret, nil
}

// GetDrives is a mock method.
func (c *MockClient) GetDrives(chassisID string) ([]model.Drive, error) {
	return []model.Drive{}, nil
}

// GetPCIeDevices is a mock method.
func (c *MockClient) GetPCIeDevices(chassisID string) ([]model.PCIeDevice, error) {
	return []model.PCIeDevice{}, nil
}

func randResource(id string, m *model.Resource) {
	m.OriginID = &id
	m.State = randState()
	m.PhysicalState = randState()
	m.Health = randHealth()
	m.PhysicalHealth = randHealth()
	m.Name = randString()
	m.Description = randString()
}

func randProductInfo(m *model.ProductInfo) {
	m.Model = randString()
	m.Manufacturer = randString()
	m.SKU = randString()
	m.SerialNumber = randString()
	m.SparePartNumber = randString()
	m.PartNumber = randString()
	m.AssetTag = randString()
}

func randMember(ID string, m *model.Member) {
	m.MemberID = ID
	m.OriginMemberID = &ID
	m.State = randState()
	m.PhysicalState = randState()
	m.Health = randHealth()
	m.PhysicalHealth = randHealth()
	m.Name = randString()
	m.Description = randString()
}

func randMemory(ID string) *model.Memory {
	ret := model.Memory{}
	randResource(ID, &ret.Resource)
	randProductInfo(&ret.ProductInfo)
	ret.CapacityMiB = randInt()
	ret.OperatingSpeedMhz = randInt()
	ret.MemoryDeviceType = randString()
	ret.DataWidthBits = randInt()
	ret.RankCount = randInt()
	ret.DeviceLocator = randString()

	ret.MemoryLocation = new(model.MemoryLocation)
	ret.MemoryLocation.Socket = randInt()
	ret.MemoryLocation.Controller = randInt()
	ret.MemoryLocation.Channel = randInt()
	ret.MemoryLocation.Slot = randInt()

	return &ret
}

func randProcessor(ID string) *model.Processor {
	ret := model.Processor{}
	randResource(ID, &ret.Resource)
	randProductInfo(&ret.ProductInfo)
	ret.Socket = randInt()
	ret.ProcessorType = randString()
	ret.ProcessorArchitecture = randString()
	ret.InstructionSet = randString()
	ret.MaxSpeedMHz = randInt()
	ret.TotalCores = randInt()
	ret.ProcessorID = new(model.ProcessorID)
	ret.ProcessorID.VendorID = randString()
	ret.ProcessorID.MicrocodeInfo = randString()
	ret.ProcessorID.Step = randString()
	ret.ProcessorID.IdentificationRegisters = randString()
	ret.ProcessorID.EffectiveFamily = randString()
	ret.ProcessorID.EffectiveModel = randString()
	return &ret
}

func randEthernetInterface(name string) *model.EthernetInterface {
	ret := model.EthernetInterface{}
	ret.Name = &name
	ret.PermanentMACAddress = randString()
	ret.IPv4Addresses = append(ret.IPv4Addresses, *randIPv4Address())
	ret.IPv4Addresses = append(ret.IPv4Addresses, *randIPv4Address())
	ret.IPv6Addresses = append(ret.IPv6Addresses, *randIPv6Address())
	ret.IPv6Addresses = append(ret.IPv6Addresses, *randIPv6Address())
	ret.VLANs = append(ret.VLANs, *randVLANs())
	ret.VLANs = append(ret.VLANs, *randVLANs())
	ret.VLANs = append(ret.VLANs, *randVLANs())
	ret.VLANs = append(ret.VLANs, *randVLANs())
	return &ret
}

func randNetworkInterface(name string) *model.NetworkInterface {
	ret := model.NetworkInterface{}
	ret.Name = &name
	ret.NetworkAdapterURI = "NetworkAdapters/0"
	return &ret
}

func randStorage(name string) *model.Storage {
	ret := model.Storage{}
	randResource(name, &ret.Resource)
	ret.Name = &name
	ret.StorageControllers = append(ret.StorageControllers, *randStorageController("Controller0"))
	ret.StorageControllers = append(ret.StorageControllers, *randStorageController("Controller1"))
	return &ret
}

func randStorageController(name string) *model.StorageController {
	ret := model.StorageController{}
	randMember(name, &ret.Member)
	randProductInfo(&ret.ProductInfo)
	ret.Name = &name
	ret.SpeedGbps = *randInt()
	ret.FirmwareVersion = *randString()
	ret.SupportedDeviceProtocols = append(ret.SupportedDeviceProtocols, *randString())
	ret.SupportedDeviceProtocols = append(ret.SupportedDeviceProtocols, *randString())
	return &ret
}

func randIPv4Address() *model.IPv4Address {
	ret := model.IPv4Address{}
	ret.Address = randIPv4()
	ret.Gateway = randIPv4()
	ret.SubnetMask = randIPv4()
	return &ret
}

func randIPv6Address() *model.IPv6Address {
	ret := model.IPv6Address{}
	ret.Address = randIPv6()
	ret.AddressState = randState()
	ret.PrefixLength = randInt()
	return &ret
}

func randVLANs() *model.VLanNetworkInterface {
	ret := model.VLanNetworkInterface{}
	ret.VLANEnable = randBool()
	ret.VLANID = randInt()
	return &ret
}

func randIPv4() *string {
	s := fmt.Sprintf("%d.%d.%d.%d", rand.Intn(255), rand.Intn(255), rand.Intn(255), rand.Intn(255))
	return &s
}

func randIPv6() *string {
	s := fmt.Sprintf("%d:%d:%d:%d:%d:%d", rand.Intn(99), rand.Intn(99), rand.Intn(99), rand.Intn(99), rand.Intn(99), rand.Intn(99))
	return &s
}

func randBool() *bool {
	i := rand.Intn(100)
	var r bool
	if i < 50 {
		r = false
	} else {
		r = true
	}
	return &r
}

func randString() *string {
	ret := uuid.New().String()
	return &ret
}

func randInt() *int {
	i := rand.Intn(100)
	return &i
}

var state = []string{
	"Enabled",
	"Disabled",
	"StandbyOffline",
	"StandbySpare",
	"InTest",
	"Starting",
	"Absent",
	"UnavailableOffline",
	"Deferring",
	"Quiesced",
	"Updating",
}

var health = []string{
	"OK",
	"Warning",
	"Critical",
}

var powerState = []string{
	"On",
	"Off",
	"PoweringOn",
	"PoweringOff",
}

func randState() *string {
	return &state[rand.Intn(len(state))]
}

func randHealth() *string {
	return &state[rand.Intn(len(health))]
}

func randPowerState() *string {
	return &state[rand.Intn(len(powerState))]
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// RandStringBytesMaskImpr return a rand string
func RandStringBytesMaskImpr(n int) string {
	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}
