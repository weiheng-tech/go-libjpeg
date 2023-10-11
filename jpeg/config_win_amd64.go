//go:build windows && amd64
// +build windows,amd64

package jpeg

// #cgo CFLAGS: -I./include
// #cgo LDFLAGS: -L./lib/win
import "C"
