package actor

import (
	"bufio"
	"errors"
	"os/exec"
)

var ErrNotStarted = errors.New("actor is not started")
var ErrAlreadyStarted = errors.New("actor already started")

type Actor struct {
	cmd    *exec.Cmd
	stdin  *bufio.Writer
	stdout *bufio.Reader
	stderr *bufio.Reader
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
		stdin:  bufio.NewWriter(stdin),
		stdout: bufio.NewReader(stdout),
		stderr: bufio.NewReader(stderr),
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
	return a.stdout.ReadString('\n')
}

func (a *Actor) Send(msg string) error {
	if !a.Started() {
		return ErrNotStarted
	}
	_, err := a.stdin.WriteString(msg)
	if err != nil {
		return err
	}
	return a.stdin.Flush()
}
