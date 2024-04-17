package dev

import (
	"github.com/sensefinitycloud/go-ble"
	"github.com/sensefinitycloud/go-ble/darwin"
)

// DefaultDevice ...
func DefaultDevice(opts ...ble.Option) (d ble.Device, err error) {
	return &darwin.Device{}, nil
}
