//go:build linux && arm
// +build linux,arm

package jpeg

// #cgo CFLAGS: -I./include
// #cgo LDFLAGS: -L./lib/armv7l
import "C"
