# Go BLE Scan Response

[Fork of github.com/go-ble/ble](https://github.com/go-ble/ble)


## Added Features

### General Advertisement function

```go
Advertise(ctx context.Context, adv AdvertisementData) error
```

Using:

```go
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
```

### Advertise Name And Services With ScanResponse

Advertises device name, and specified service UUIDs.  
It tres to fit the UUIDs in the advertising packet as much as possible.  
Advertises the given manufacturer data in the scan response.  

```go
AdvertiseNameAndServicesWithScanResponse(ctx context.Context, name string, companyId uint16, b []byte, uuids ...UUID) error
```

## Go Ble

**ble** is a Golang [Bluetooth Low Energy](https://en.wikipedia.org/wiki/Bluetooth_Low_Energy) package for Linux and Mac OS.

**Note:** Is not being actively maintained, should use the original lib.
