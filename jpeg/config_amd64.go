//go:build linux && amd64
// +build linux,amd64

package jpeg

// #cgo CFLAGS: -I./include
// #cgo LDFLAGS: -static-libgcc -static-libstdc++ -L./lib/amd64
import "C"
