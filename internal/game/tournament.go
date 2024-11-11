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
	names := t.participantsNames()
	for i, aName := range names[:len(names)-1] {
		for _, bName := range names[i+1:] {
			a := t.prts[aName]
			b := t.prts[bName]
			err := t.game.Play(rounds, a, b)
			if err != nil {
				errs = errors.Join(errs, err)
				break
			}
		}
	}
	return errs
}

func (t *Tournament) participantsNames() []Name {
	names := make([]Name, 0, len(t.prts))
	for name := range t.prts {
		names = append(names, name)
	}
	return names
}
