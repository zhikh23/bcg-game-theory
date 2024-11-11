package actor

import (
	"bufio"
	"os/exec"
	"strings"
)

type Program struct {
	cmd    *exec.Cmd
	stdin  *bufio.Writer
	stdout *bufio.Reader
	stderr *bufio.Reader
	run    bool
}

func NewProgramActor(path string) (*Program, error) {
	cmd := exec.Command(path)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil, err
	}

	return &Program{
		cmd:    cmd,
		stdin:  bufio.NewWriter(stdin),
		stdout: bufio.NewReader(stdout),
		stderr: bufio.NewReader(stderr),
		run:    false,
	}, cmd.Err
}

func MustNewProgramActor(path string) *Program {
	a, err := NewProgramActor(path)
	if err != nil {
		panic(err)
	}
	return a
}

func (a *Program) Start() error {
	if a.run {
		return ErrAlreadyStarted
	}

	a.run = true
	err := a.cmd.Start()
	if err == nil {
		go func() {
			_ = a.cmd.Wait()
			a.run = false
		}()
	}
	return err
}

func (a *Program) Running() bool {
	return a.run
}

func (a *Program) Terminate() error {
	return a.cmd.Process.Kill()
}

func (a *Program) Receive() (string, error) {
	if !a.Running() {
		return "", ErrNotRunning
	}

	s, err := a.stdout.ReadString('\n')
	if err != nil {
		return "", err
	}

	s = strings.TrimSuffix(s, "\n")
	return s, nil
}

func (a *Program) Send(msg string) error {
	if !a.Running() {
		return ErrNotRunning
	}

	_, err := a.stdin.WriteString(msg + "\n")
	if err != nil {
		return err
	}

	return a.stdin.Flush()
}
