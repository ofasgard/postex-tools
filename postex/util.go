package postex
//Contains general purpose helper functions.

import "os"
import "runtime"

func Exists(filepath string) error {
	_,err := os.Stat(filepath)
	return err
}

func CheckShell() string {
	if runtime.GOOS == "linux" {
		if Exists("/bin/bash") == nil {
			return "/bin/bash"
		}
	}
	if runtime.GOOS == "windows" {
		if Exists("C:\\Windows\\System32\\WindowsPowerShell\\v1.0\\powershell.exe") == nil {
			return "C:\\Windows\\System32\\WindowsPowerShell\\v1.0\\powershell.exe"
		}
		if Exists("C:\\Windows\\System32\\cmd.exe") == nil {
			return "C:\\Windows\\System32\\cmd.exe"
		}
	}
	return "/bin/sh" //default, better than nothing
}
