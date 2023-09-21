//go:build linux && amd64
// +build linux,amd64

package jpeg

// #cgo CFLAGS: -I./include/amd64
// #cgo LDFLAGS: -L./lib/amd64
import "C"
