package actor

import (
	"errors"
)

var ErrNotRunning = errors.New("actor is not started")
var ErrAlreadyStarted = errors.New("actor already started")

type Actor interface {
	Start() error
	Running() bool
	Terminate() error
	Receive() (string, error)
	Send(msg string) error
}
