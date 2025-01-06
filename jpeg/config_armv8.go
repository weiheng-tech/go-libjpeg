//go:build linux && arm64
// +build linux,arm64

package jpeg

// #cgo CFLAGS: -I./include
// #cgo LDFLAGS: -static-libgcc -static-libstdc++ -L./lib/armv8
import "C"
