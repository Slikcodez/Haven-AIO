package constants

import (
	"fmt"
	"syscall"
	"time"
	"unsafe"
)

var Carts = 0
var Declines = 0
var Checkouts = 0
var Version = "0.5.992- req 2/3 fix"

func RunCmdLoop() {
	go func() {
		CmdTitles()
	}()
	return
}

func CmdTitles() {

	for {

		kernel32 := syscall.NewLazyDLL("kernel32.dll")

		proc := kernel32.NewProc("SetConsoleTitleW")

		ret, _, err := proc.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(fmt.Sprintf("Haven %s | Carts %d | Checkouts %d | Declines %d", Version, Carts, Checkouts, Declines)))))
		if ret == 0 {
			panic(err)
		}
		time.Sleep(1 * time.Second)
	}

}
