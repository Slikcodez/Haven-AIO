package constants

import (
	"fmt"
	"syscall"
	"unsafe"
)

var carts = 0
var declines = 0
var checkouts = 0
var version = 0

func CmdTitles() {
	for {

		kernel32, err := syscall.NewLazyDLL("kernel32.dll")
		if err != nil {
			panic(err)
		}

		proc, err := kernel32.NewProc("SetConsoleTitleW")
		if err != nil {
			panic(err)
		}

		ret, _, err := proc.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(fmt.Sprintf("Haven %d | Carts %d | Checkouts %d | Declines %d", version, carts, checkouts, declines)))))
		if ret == 0 {
			panic(err)
		}
	}
}
