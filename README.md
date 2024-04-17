# Go BLE Scan Response

[Fork of github.com/go-ble/ble](https://github.com/go-ble/ble)


## Added Features

Advertises device name, and specified service UUIDs.  
It tres to fit the UUIDs in the advertising packet as much as possible.  
Advertises the given manufacturer data in the scan response.  

```go
AdvertiseNameAndServicesWithScanResponse(ctx context.Context, name string, companyId uint16, b []byte, uuids ...UUID) error
```

## Go Ble

**ble** is a Golang [Bluetooth Low Energy](https://en.wikipedia.org/wiki/Bluetooth_Low_Energy) package for Linux and Mac OS.

**Note:** Is not being actively maintained, should use the original lib.
