package dev

import (
	"github.com/sensefinitycloud/go-ble"
	"github.com/sensefinitycloud/go-ble/linux"
)

// DefaultDevice ...
func DefaultDevice(opts ...ble.Option) (d ble.Device, err error) {
	return linux.NewDevice(opts...)
}
