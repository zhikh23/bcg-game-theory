package game

import (
	"errors"
	"fmt"
	"log"

	"github.com/zhikh23/bcg-game-theory/internal/actor"
)

type PrisonerDilemma struct {
	a *Participant
	b *Participant
}

func NewPrisonerDilemma(a, b actor.Actor) *PrisonerDilemma {
	return &PrisonerDilemma{
		a: NewParticipant(a),
		b: NewParticipant(b),
	}
}

func (d *PrisonerDilemma) Start() error {
	var err error
	err = errors.Join(err, d.a.Start())
	err = errors.Join(err, d.b.Start())
	return err
}

func (d *PrisonerDilemma) Round() error {
	stepA, err := d.a.Receive()
	if err != nil {
		return err
	}

	stepB, err := d.b.Receive()
	if err != nil {
		return err
	}

	switch {
	case stepA == "Y" && stepB == "Y":
		d.a.Award(5)
		d.b.Award(5)
	case stepA == "Y" && stepB == "N":
		d.b.Award(10)
	case stepA == "N" && stepB == "Y":
		d.a.Award(10)
	case stepA == "N" && stepB == "N":
		d.a.Award(1)
		d.b.Award(1)
	default:
		log.Fatalf("Invalid input: A = %s, B = %s", stepA, stepB)
	}
	fmt.Printf(
		"A: %d\n"+
			"B: %d\n",
		d.a.Score(), d.b.Score(),
	)

	err = d.b.Send(stepA)
	if err != nil {
		return err
	}
	err = d.a.Send(stepB)
	if err != nil {
		return err
	}

	return nil
}
