package game

import (
	"errors"
	"fmt"
)

var ErrParticipantAlreadyExists = errors.New("participant already exists")

type Tournament struct {
	prts map[Name]*Participant
	game BinaryGame
}

func NewTournament(game BinaryGame) *Tournament {
	return &Tournament{
		prts: make(map[Name]*Participant),
		game: game,
	}
}

func (t *Tournament) AddParticipant(prt *Participant) error {
	if _, ok := t.prts[prt.Name()]; ok {
		return fmt.Errorf("%w: %s", ErrParticipantAlreadyExists, prt.Name())
	}

	t.prts[prt.Name()] = prt

	return nil
}

func (t *Tournament) Tour(rounds int) error {
	var errs error
	for nameA, prtA := range t.prts {
		for nameB, prtB := range t.prts {
			if nameA == nameB {
				continue
			}
			err := t.game.Play(rounds, prtA, prtB)
			if err != nil {
				errs = errors.Join(errs, err)
				break
			}
		}
	}
	return errs
}
