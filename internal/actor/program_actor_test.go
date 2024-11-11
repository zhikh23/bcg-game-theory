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
	f := actor.NewProgramFactory("invalid")
	_, err := f.New()
	require.Error(t, err)
}

func TestEmptyActor_Start(t *testing.T) {
	f := actor.NewProgramFactory(emptyPythonScriptPath)
	a, err := f.New()
	require.NoError(t, err)
	// Скрипт пустой и он должен скоро окончить работу.
	require.Eventually(t, func() bool {
		return a.Running() == false
	}, 100*time.Millisecond, 10*time.Millisecond)
}

func TestEchoActor_Start(t *testing.T) {
	f := actor.NewProgramFactory(echoPythonScriptPath)
	a, err := f.New()
	require.NoError(t, err)
	require.True(t, a.Running())
	require.NoError(t, a.Terminate())
	require.Eventually(t, func() bool {
		return a.Running() == false
	}, 100*time.Millisecond, 10*time.Millisecond)
}

func TestEchoActor_TerminateTwice(t *testing.T) {
	f := actor.NewProgramFactory(echoPythonScriptPath)
	a, err := f.New()
	require.NoError(t, err)
	require.NoError(t, a.Terminate())
	require.ErrorIs(t, a.Terminate(), actor.ErrActorAlreadyTerminated)
}

func TestHelloActor_ReadLine(t *testing.T) {
	f := actor.NewProgramFactory(helloPythonScriptPath)
	a := f.MustNew()

	line, err := a.Receive()
	require.NoError(t, err)
	require.Equal(t, "Hello!", line)

	require.Eventually(t, func() bool {
		return a.Running() == false
	}, 100*time.Millisecond, 10*time.Millisecond)
}

func TestEchoActor_Send(t *testing.T) {
	f := actor.NewProgramFactory(echoPythonScriptPath)
	a := f.MustNew()

	sent := "Hello!"
	require.NoError(t, a.Send(sent))
	// Так как программа выполняет функцию эхо, ответ должен быть такой же.
	res, err := a.Receive()
	require.NoError(t, err)
	require.Equal(t, sent, res)

	require.NoError(t, a.Terminate())
}
