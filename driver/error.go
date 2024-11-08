package driver

import "fmt"

var (
	ErrDeviceNotFound  = fmt.Errorf("device not found")
	ErrDeviceOffline   = fmt.Errorf("device offline")
	ErrMultipleDevices = fmt.Errorf("multiple devices found")
	ErrFileNotFound    = fmt.Errorf("file not found")
)