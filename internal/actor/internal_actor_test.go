package actor_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/zhikh23/bcg-game-theory/internal/actor"
)

func noOp(_ <-chan string, _ chan<- string) {
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

func TestInternalActor_Start(t *testing.T) {
	a := actor.NewInternalActor(noOp)
	require.False(t, a.Running())
	require.NoError(t, a.Start())
	require.True(t, a.Running())
}

func TestInternalActor_Terminate(t *testing.T) {
	a := actor.NewInternalActor(noOp)
	require.NoError(t, a.Start())
	require.NoError(t, a.Terminate())
	require.False(t, a.Running())
}

func TestInternalActor_SendAndReceive(t *testing.T) {
	a := actor.NewInternalActor(echo)

	sent := "Hello!"
	require.ErrorIs(t, a.Send(sent), actor.ErrNotRunning)

	require.NoError(t, a.Start())
	require.NoError(t, a.Send(sent))

	res, err := a.Receive()
	require.NoError(t, err)
	require.Equal(t, sent, res)
}

func TestInternalActor_SendAndReceiveTwice(t *testing.T) {
	a := actor.NewInternalActor(echo)
	require.NoError(t, a.Start())

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
