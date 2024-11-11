package actor

type Handler func(stdin <-chan string, stdout chan<- string)

type InternalFactory struct {
	handler   Handler
	stdinCap  int
	stdoutCap int
}

type Internal struct {
	handler Handler
	stdin   chan string
	stdout  chan string
	running bool
}

func NewInternalFactory(handler Handler, stdinCap int, stdoutCap int) *InternalFactory {
	return &InternalFactory{
		handler:   handler,
		stdinCap:  stdinCap,
		stdoutCap: stdoutCap,
	}
}

func (f *InternalFactory) New() (Actor, error) {
	stdin := make(chan string, f.stdinCap)
	stdout := make(chan string, f.stdoutCap)

	go f.handler(stdin, stdout)

	return &Internal{
		handler: f.handler,
		stdin:   stdin,
		stdout:  stdout,
		running: true,
	}, nil
}

func (f *InternalFactory) MustNew() Actor {
	a, err := f.New()
	if err != nil {
		panic(err)
	}
	return a
}

func (a *Internal) Running() bool {
	return a.running
}

func (a *Internal) Terminate() error {
	if !a.running {
		return ErrActorAlreadyTerminated
	}
	a.running = false
	return nil
}

func (a *Internal) Receive() (string, error) {
	if !a.running {
		return "", ErrActorIsNotRunning
	}
	return <-a.stdout, nil
}

func (a *Internal) Send(msg string) error {
	if !a.running {
		return ErrActorIsNotRunning
	}
	a.stdin <- msg
	return nil
}
