package postex
//Contains functions for executing commands.

import "os/exec"
import "io"
import "bytes"

func spawnShell(shellpath string) (*shell,error) {
	session := shell{}
	session.command_shell = exec.Command(shellpath)
	session.stdin_writer,_ = session.command_shell.StdinPipe()
	session.stdout_reader,_ = session.command_shell.StdoutPipe()
	return &session,nil
}

// SHELL STRUCT DEFINITION

type shell struct {
	command_shell *exec.Cmd
	stdin_writer io.Writer
	stdout_reader io.Reader
}

func (s shell) launch() {
	s.command_shell.Start()
}

func (s shell) stop() {
	s.command_shell.Process.Kill()
}

func (s shell) send(cmd string) {
	//this blocks
	io.WriteString(s.stdin_writer, cmd)
}

func (s shell) recv() string {
	//this blocks too
	out := bytes.Buffer{}
	buf := make([]byte, 500)
	written,_ := s.stdout_reader.Read(buf)
	out.WriteString(string(buf))
	for written == 500 {
		buf = make([]byte, 500)
		written,_ = s.stdout_reader.Read(buf)
		out.WriteString(string(buf))
	}
	return out.String()
}
