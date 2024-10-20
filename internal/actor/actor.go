package actor

import (
	"errors"
	"os/exec"
)

var ErrAlreadyStarted = errors.New("actor already started")

type Actor struct {
	cmd *exec.Cmd
	run bool
}

func FromProgram(path string) (*Actor, error) {
	cmd := exec.Command(path)
	return &Actor{
		cmd: cmd,
		run: false,
	}, cmd.Err
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
