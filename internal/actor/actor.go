package actor

import (
	"errors"
)

var ErrNotStarted = errors.New("actor is not started")
var ErrAlreadyStarted = errors.New("actor already started")

type Actor interface {
	Start() error
	Started() bool
	Terminate() error
	Receive() (string, error)
	Send(msg string) error
}
