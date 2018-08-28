package dev

import (
	"github.com/cinello/ble"
	"github.com/cinello/ble/linux"
)

// DefaultDevice ...
func DefaultDevice() (d ble.Device, err error) {
	return linux.NewDevice()
}
