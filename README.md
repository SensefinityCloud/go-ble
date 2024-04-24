# Go BLE Scan Response

[Fork of github.com/go-ble/ble](https://github.com/go-ble/ble)

## Added Features

### General Advertisement function

```go
Advertise(ctx context.Context, adv AdvertisementData) error
```

Struct:

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

eg:

```go
advData := ble.AdvertisementData{
	ShortName: "Example",
	Services:  []ble.UUID{ble.MustParse("6E400001-B5A3-F393-E0A9-E50E24DCCA9E")},
	ScanResponse: &ble.ScanResponse{
		ManufacturerData: &ble.ManufacturerData{
			CompanyId: uint16(89),
			Data:      []byte("Example"),
		},
	},
}

// It blocks the current goroutine until the advertising process is finished.
err = ble.Advertise(ctx, advData)
if err != nil {
	panic(err)
}
```

### Added Notifier on HandleWrite

Feature: We can notify other characteristics when receiving new data

```go
rx.HandleWrite(ble.WriteHandlerFunc(func(req ble.Request, rsp ble.ResponseWriter, n func(h uint16, data []byte) (int, error))
```

eg:

```go
tx.HandleNotify(ble.NotifyHandlerFunc(func(req ble.Request, n ble.Notifier) {
		logger.Log(processServer, "TX Notify handler called", nil)

}))


rx.HandleWrite(ble.WriteHandlerFunc(func(req ble.Request, rsp ble.ResponseWriter, n func(h uint16, data []byte) (int, error)) {

		n(tx.ValueHandle, []byte("ok"))
}))
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
