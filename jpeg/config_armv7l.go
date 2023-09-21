//go:build linux && arm
// +build linux,arm

package jpeg

// #cgo CFLAGS: -I./include/armv7l
// #cgo LDFLAGS: -L./lib/armv7l
import "C"
