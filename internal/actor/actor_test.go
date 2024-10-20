package actor_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/zhikh23/bcg-game-theory/internal/actor"
)

const (
	emptyPythonScriptPath = "../../tests/empty.py"
	echoPythonScriptPath  = "../../tests/echo.py"
	helloPythonScriptPath = "../../tests/hello.py"
)

func TestInvalidActor_FromProgram(t *testing.T) {
	_, err := actor.FromProgram("invalid")
	require.Error(t, err)
}

func TestEmptyActor_Start(t *testing.T) {
	a, err := actor.FromProgram(emptyPythonScriptPath)
	require.NoError(t, err)
	require.NoError(t, a.Start())
	// Скрипт пустой и он должен скоро окончить работу.
	require.Eventually(t, func() bool {
		return a.Started() == false
	}, 100*time.Millisecond, 10*time.Millisecond)
}

func TestEchoActor_Start(t *testing.T) {
	a, err := actor.FromProgram(echoPythonScriptPath)
	require.NoError(t, err)
	require.NoError(t, a.Start())
	require.True(t, a.Started())
	require.NoError(t, a.Terminate())
	require.Eventually(t, func() bool {
		return a.Started() == false
	}, 100*time.Millisecond, 10*time.Millisecond)
}

func TestEchoActor_StartTwice(t *testing.T) {
	a, err := actor.FromProgram(echoPythonScriptPath)
	require.NoError(t, err)
	require.NoError(t, a.Start())
	require.ErrorIs(t, a.Start(), actor.ErrAlreadyStarted)
}

func TestEchoActor_TerminateTwice(t *testing.T) {
	a, err := actor.FromProgram(echoPythonScriptPath)
	require.NoError(t, err)
	require.NoError(t, a.Start())
	require.NoError(t, a.Terminate())
	require.NoError(t, a.Terminate())
}

func TestHelloActor_ReadLine(t *testing.T) {
	a := actor.MustFromProgram(helloPythonScriptPath)
	require.NoError(t, a.Start())
	require.True(t, a.Started())
	line, err := a.ReadLine()
	require.NoError(t, err)
	require.Equal(t, "Hello!\n", line)
	require.Eventually(t, func() bool {
		return a.Started() == false
	}, 100*time.Millisecond, 10*time.Millisecond)
}
