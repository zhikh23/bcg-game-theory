package actor

type Handler func(stdin <-chan string, stdout chan<- string)

type InternalActor struct {
	handler Handler
	stdin   chan string
	stdout  chan string
	run     bool
}

func NewInternalActor(handler Handler) *InternalActor {
	return &InternalActor{
		handler: handler,
		stdin:   make(chan string, 10),
		stdout:  make(chan string, 10),
	}
}

func (a *InternalActor) Start() error {
	if a.run {
		return ErrAlreadyStarted
	}
	a.run = true
	go func() {
		a.handler(a.stdin, a.stdout)
		a.run = false
	}()
	return nil
}

func (a *InternalActor) Running() bool {
	return a.run
}

func (a *InternalActor) Terminate() error {
	a.run = false
	return nil
}

func (a *InternalActor) Receive() (string, error) {
	if !a.run {
		return "", ErrNotRunning
	}
	return <-a.stdout, nil
}

func (a *InternalActor) Send(msg string) error {
	if !a.run {
		return ErrNotRunning
	}
	a.stdin <- msg
	return nil
}
