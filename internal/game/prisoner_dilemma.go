package game

import (
	"errors"
	"fmt"
)

type PrisonerDilemma struct {
	a *Participant
	b *Participant
}

func NewPrisonerDilemma(a, b *Participant) *PrisonerDilemma {
	return &PrisonerDilemma{a: a, b: b}
}

func (d *PrisonerDilemma) Start() error {
	var err error
	err = errors.Join(err, d.a.Start())
	err = errors.Join(err, d.b.Start())
	return err
}

func (d *PrisonerDilemma) Play() error {
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
		return fmt.Errorf("invalid input: %s = %s, %s = %s", d.a.Name(), stepA, d.b.Name(), stepB)
	}

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

func (d *PrisonerDilemma) Results() map[Name]Score {
	return map[Name]Score{
		d.a.Name(): d.a.Score(),
		d.b.Name(): d.b.Score(),
	}
}
