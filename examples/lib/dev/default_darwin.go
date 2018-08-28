package dev

import (
	"github.com/cinello/ble"
	"github.com/cinello/ble/darwin"
)

// DefaultDevice ...
func DefaultDevice() (d ble.Device, err error) {
	return darwin.NewDevice()
}
