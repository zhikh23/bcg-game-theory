package actor_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/zhikh23/bcg-game-theory/internal/actor"
)

const defaultStdCap = 10

func noOp(stdin <-chan string, _ chan<- string) {
	for {
		select {
		case <-stdin:
		default:
			time.Sleep(time.Microsecond)
		}
	}
}

func echo(stdin <-chan string, stdout chan<- string) {
	for {
		select {
		case s := <-stdin:
			stdout <- s
		default:
			time.Sleep(time.Microsecond)
		}
	}
}

func TestInternalActor_Create(t *testing.T) {
	f := actor.NewInternalFactory(noOp, defaultStdCap, defaultStdCap)
	a := f.MustNew()
	require.True(t, a.Running())
}

func TestInternalActor_Terminate(t *testing.T) {
	f := actor.NewInternalFactory(noOp, defaultStdCap, defaultStdCap)
	a := f.MustNew()
	require.NoError(t, a.Terminate())
	require.False(t, a.Running())
}

func TestInternalActor_SendAndReceive(t *testing.T) {
	f := actor.NewInternalFactory(echo, defaultStdCap, defaultStdCap)
	a := f.MustNew()

	sent := "Hello!"
	require.NoError(t, a.Send(sent))

	res, err := a.Receive()
	require.NoError(t, err)
	require.Equal(t, sent, res)
}

func TestInternalActor_SendAndReceiveTwice(t *testing.T) {
	f := actor.NewInternalFactory(echo, defaultStdCap, defaultStdCap)
	a := f.MustNew()

	sent := "Hello!"
	require.NoError(t, a.Send(sent))
	res, err := a.Receive()
	require.NoError(t, err)
	require.Equal(t, sent, res)

	require.NoError(t, a.Send(sent))
	res, err = a.Receive()
	require.NoError(t, err)
	require.Equal(t, sent, res)
}

func TestInternalActor_SendAfterTerminate(t *testing.T) {
	f := actor.NewInternalFactory(noOp, defaultStdCap, defaultStdCap)
	a := f.MustNew()

	require.NoError(t, a.Terminate())
	require.ErrorIs(t, a.Send("sth"), actor.ErrActorIsNotRunning)
}

func TestInternalActor_ReceiveAfterTerminate(t *testing.T) {
	f := actor.NewInternalFactory(noOp, defaultStdCap, defaultStdCap)
	a := f.MustNew()

	require.NoError(t, a.Terminate())
	_, err := a.Receive()
	require.ErrorIs(t, err, actor.ErrActorIsNotRunning)
}

func TestIntervalActor_TerminateTwice(t *testing.T) {
	f := actor.NewInternalFactory(noOp, defaultStdCap, defaultStdCap)
	a := f.MustNew()
	require.NoError(t, a.Terminate())
	require.ErrorIs(t, a.Terminate(), actor.ErrActorAlreadyTerminated)
}
