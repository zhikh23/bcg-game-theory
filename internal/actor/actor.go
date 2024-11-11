package actor

import "errors"

var ErrActorAlreadyTerminated = errors.New("actor already terminated")
var ErrActorIsNotRunning = errors.New("actor is not running")

type Actor interface {
	Running() bool
	Terminate() error
	Receive() (string, error)
	Send(msg string) error
}
