package dev

import (
	"github.com/cinello/ble"
	"github.com/cinello/ble/linux"
	"github.com/cinello/ble/linux/hci"
)

// DefaultDevice ...
func DefaultDevice(opts ...hci.Option) (d ble.Device, err error) {
	return linux.NewDevice(opts...)
}
