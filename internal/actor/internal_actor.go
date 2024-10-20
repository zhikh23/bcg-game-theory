package actor

type InternalActor struct {
	handler func(string) string
	out     chan string
	run     bool
}

func NewInternalActor(handler func(string) string) *InternalActor {
	return &InternalActor{
		handler: handler,
		out:     make(chan string, 10),
	}
}

func (a *InternalActor) Start() error {
	if a.run {
		return ErrAlreadyStarted
	}
	a.run = true
	return nil
}

func (a *InternalActor) Started() bool {
	return a.run
}

func (a *InternalActor) Terminate() error {
	a.run = false
	return nil
}

func (a *InternalActor) Receive() (string, error) {
	if !a.run {
		return "", ErrNotStarted
	}
	return <-a.out, nil
}

func (a *InternalActor) Send(msg string) error {
	if !a.run {
		return ErrNotStarted
	}
	a.out <- a.handler(msg)
	return nil
}
