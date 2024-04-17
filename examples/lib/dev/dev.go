package dev

import (
	"github.com/sensefinitycloud/go-ble/ble"
)

// NewDevice ...
func NewDevice(impl string, opts ...ble.Option) (d ble.Device, err error) {
	return DefaultDevice(opts...)
}
