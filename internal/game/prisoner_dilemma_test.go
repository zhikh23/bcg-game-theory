package game_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/zhikh23/bcg-game-theory/internal/actor"
	"github.com/zhikh23/bcg-game-theory/internal/game"
)

const defaultStdCap = 10

func evilHandler(stdin <-chan string, stdout chan<- string) {
	stdout <- "N"
	for {
		select {
		case <-stdin:
			stdout <- "N"
		default:
			time.Sleep(time.Microsecond)
		}
	}
}

func kindHandler(stdin <-chan string, stdout chan<- string) {
	stdout <- "Y"
	for {
		select {
		case <-stdin:
			stdout <- "Y"
		default:
			time.Sleep(time.Microsecond)
		}
	}
}

func TestPrisonerDilemma_Round(t *testing.T) {
	g := game.NewPrisonerDilemma(game.PrisonerDilemmaConfig{
		MutualDefects: 1,
		Defect:        10,
		Cooperate:     5,
	})
	evilFactory := actor.NewInternalFactory(evilHandler, defaultStdCap, defaultStdCap)
	kindFactory := actor.NewInternalFactory(kindHandler, defaultStdCap, defaultStdCap)

	t.Run("prisoner's dilemma evil vs evil", func(t *testing.T) {
		evil1 := game.NewParticipant("Evil 1", evilFactory)
		evil2 := game.NewParticipant("Evil 2", evilFactory)

		err := g.Play(1, evil1, evil2)
		require.NoError(t, err)

		require.Equal(t, evil1.Score(), game.Score(1))
		require.Equal(t, evil2.Score(), game.Score(1))
	})

	t.Run("prisoner's dilemma kind vs kind", func(t *testing.T) {
		kind1 := game.NewParticipant("Kind 1", kindFactory)
		kind2 := game.NewParticipant("Kind 2", kindFactory)

		err := g.Play(1, kind1, kind2)
		require.NoError(t, err)

		require.Equal(t, kind1.Score(), game.Score(5))
		require.Equal(t, kind2.Score(), game.Score(5))
	})

	t.Run("prisoner's dilemma evil vs kind", func(t *testing.T) {
		evil1 := game.NewParticipant("Evil 1", evilFactory)
		kind2 := game.NewParticipant("Kind 2", kindFactory)

		err := g.Play(1, evil1, kind2)
		require.NoError(t, err)

		require.Equal(t, evil1.Score(), game.Score(10))
		require.Equal(t, kind2.Score(), game.Score(0))
	})

	t.Run("prisoner's dilemma kind vs evil", func(t *testing.T) {
		kind1 := game.NewParticipant("Kind 1", kindFactory)
		evil2 := game.NewParticipant("Evil 2", evilFactory)

		err := g.Play(1, kind1, evil2)
		require.NoError(t, err)

		require.Equal(t, kind1.Score(), game.Score(0))
		require.Equal(t, evil2.Score(), game.Score(10))
	})

	t.Run("prisoner's dilemma kind vs evil in several rounds", func(t *testing.T) {
		kind1 := game.NewParticipant("Kind 1", kindFactory)
		evil2 := game.NewParticipant("Evil 2", evilFactory)

		require.NoError(t, g.Play(3, kind1, evil2))

		require.Equal(t, kind1.Score(), game.Score(0))
		require.Equal(t, evil2.Score(), game.Score(30))
	})
}
