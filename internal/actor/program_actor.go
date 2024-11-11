package actor

import (
	"bufio"
	"os/exec"
	"strings"
)

type ProgramFactory struct {
	path string
}

type Program struct {
	cmd    *exec.Cmd
	stdin  *bufio.Writer
	stdout *bufio.Reader
	stderr *bufio.Reader
	run    bool
}

func NewProgramFactory(path string) *ProgramFactory {
	return &ProgramFactory{
		path: path,
	}
}

func (f *ProgramFactory) New() (Actor, error) {
	cmd := exec.Command(f.path)
	if cmd.Err != nil {
		return nil, cmd.Err
	}

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

	p := &Program{
		cmd:    cmd,
		stdin:  bufio.NewWriter(stdin),
		stdout: bufio.NewReader(stdout),
		stderr: bufio.NewReader(stderr),
		run:    true,
	}

	err = cmd.Start()
	if err == nil {
		go func() {
			_ = cmd.Wait()
			p.run = false
		}()
	}
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (f *ProgramFactory) MustNew() Actor {
	a, err := f.New()
	if err != nil {
		panic(err)
	}
	return a
}

func (a *Program) Running() bool {
	return a.run
}

func (a *Program) Terminate() error {
	if !a.Running() {
		return ErrActorAlreadyTerminated
	}

	err := a.cmd.Process.Kill()
	if err != nil {
		return err
	}
	_ = a.cmd.Wait()

	return nil
}

func (a *Program) Receive() (string, error) {
	if !a.Running() {
		return "", ErrActorIsNotRunning
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
		return ErrActorIsNotRunning
	}

	_, err := a.stdin.WriteString(msg + "\n")
	if err != nil {
		return err
	}

	return a.stdin.Flush()
}
