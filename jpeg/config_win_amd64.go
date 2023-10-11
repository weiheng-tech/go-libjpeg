//go:build windows && amd64
// +build windows,amd64

package jpeg

// #cgo CFLAGS: -I./include
// #cgo LDFLAGS: -static-libgcc -static-libstdc++ -L./lib/win
import "C"
