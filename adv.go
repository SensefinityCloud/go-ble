package ble

// AdvHandler handles advertisement.
type AdvHandler func(a Advertisement)

// AdvFilter returns true if the advertisement matches specified condition.
type AdvFilter func(a Advertisement) bool

// Advertisement ...
type Advertisement interface {
	LocalName() string
	ManufacturerData() []byte
	ServiceData() []ServiceData
	Services() []UUID
	OverflowService() []UUID
	TxPowerLevel() int
	Connectable() bool
	SolicitedService() []UUID

	RSSI() int
	Addr() Addr
}

// ServiceData ...
type ServiceData struct {
	UUID UUID
	Data []byte
}

// AdvertisementData represents the advertisement data for a Bluetooth Low Energy device.
type AdvertisementData struct {
	ShortName        string
	CompleteName     string
	ManufacturerData *ManufacturerData
	ScanResponse     *ScanResponse
	Services []UUID
}

type ScanResponse struct {
	ManufacturerData  *ManufacturerData
	IncludeDeviceName bool
}

// ManufacturerData represents the data structure for manufacturer-specific data in BLE advertising packets.
type ManufacturerData struct {
	CompanyId uint16
	Data      []byte
}
