package postex
//Contains functions for executing commands.

import "os/exec"
import "io"
import "bytes"

func spawnShell(shellpath string) (*shell,error) {
	session := shell{}
	session.command_shell = exec.Command(shellpath)
	session.stdout_pipe,_ = session.command_shell.StdoutPipe()
	session.stderr_pipe,_ = session.command_shell.StderrPipe()
	session.stdin_pipe,_ = session.command_shell.StdinPipe()
	return &session,nil
}

// SHELL STRUCT DEFINITION

type shell struct {
	command_shell *exec.Cmd
	stdout_pipe io.ReadCloser
	stderr_pipe io.ReadCloser
	stdin_pipe io.WriteCloser
}

func (s shell) launch() {
	s.command_shell.Start()
}

func (s shell) stop() {
	s.command_shell.Process.Kill()
}

func (s shell) send(cmd string) {
	//this blocks
	s.stdin_pipe.Write([]byte(cmd))
}

func (s shell) recv() string {
	//this blocks too
	out := bytes.Buffer{}
	buf := make([]byte, 500)
	written,_ := s.stdout_pipe.Read(buf)
	out.WriteString(string(buf))
	for written == 500 {
		buf = make([]byte, 500)
		written,_ = s.stdout_pipe.Read(buf)
		out.WriteString(string(buf))
	}
	return out.String()
}
