//go:build linux && arm
// +build linux,arm

package jpeg

// #cgo CFLAGS: -I./include
// #cgo LDFLAGS: -static-libgcc -static-libstdc++ -L./lib/armv7l
import "C"
