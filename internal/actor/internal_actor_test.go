package actor_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/zhikh23/bcg-game-theory/internal/actor"
)

func noOp(_ string) string {
	return ""
}

func echo(s string) string {
	return s
}

func TestInternalActor_Start(t *testing.T) {
	a := actor.NewInternalActor(noOp)
	require.False(t, a.Started())
	require.NoError(t, a.Start())
	require.True(t, a.Started())
}

func TestInternalActor_Terminate(t *testing.T) {
	a := actor.NewInternalActor(noOp)
	require.NoError(t, a.Start())
	require.NoError(t, a.Terminate())
	require.False(t, a.Started())
}

func TestInternalActor_SendAndReceive(t *testing.T) {
	a := actor.NewInternalActor(echo)

	sent := "Hello!"
	require.ErrorIs(t, a.Send(sent), actor.ErrNotStarted)

	require.NoError(t, a.Start())
	require.NoError(t, a.Send(sent))

	res, err := a.Receive()
	require.NoError(t, err)
	require.Equal(t, sent, res)
}
