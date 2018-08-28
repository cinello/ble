package dev

import (
	"github.com/cinello/ble"
	"github.com/cinello/ble/linux/hci"
)

// NewDevice ...
func NewDevice(impl string, opts ...hci.Option) (d ble.Device, err error) {
	return DefaultDevice(opts...)
}
