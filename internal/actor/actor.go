package actor

import (
	"bufio"
	"errors"
	"io"
	"os/exec"
)

var ErrNotStarted = errors.New("actor is not started")
var ErrAlreadyStarted = errors.New("actor already started")

type Actor struct {
	cmd    *exec.Cmd
	stdin  io.Writer
	stdout io.Reader
	stderr io.Reader
	run    bool
}

func FromProgram(path string) (*Actor, error) {
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

	return &Actor{
		cmd:    cmd,
		stdin:  stdin,
		stdout: stdout,
		stderr: stderr,
		run:    false,
	}, cmd.Err
}

func MustFromProgram(path string) *Actor {
	a, err := FromProgram(path)
	if err != nil {
		panic(err)
	}
	return a
}

func (a *Actor) Start() error {
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

func (a *Actor) Started() bool {
	return a.run
}

func (a *Actor) Terminate() error {
	return a.cmd.Process.Kill()
}

func (a *Actor) ReadLine() (string, error) {
	if !a.Started() {
		return "", ErrNotStarted
	}

	reader := bufio.NewReader(a.stdout)
	return reader.ReadString('\n')
}
