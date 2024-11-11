package game

import (
	"fmt"
	"github.com/zhikh23/bcg-game-theory/internal/actor"
)

type PrisonerDilemmaConfig struct {
	MutualDefects Score
	Defect        Score
	Cooperate     Score
}

type PrisonerDilemma struct {
	cfg PrisonerDilemmaConfig
}

func NewPrisonerDilemma(cfg PrisonerDilemmaConfig) *PrisonerDilemma {
	return &PrisonerDilemma{
		cfg: cfg,
	}
}

func (d *PrisonerDilemma) Play(rounds int, a, b *Participant) error {
	actorA, err := a.Actor()
	if err != nil {
		return err
	}

	actorB, err := b.Actor()
	if err != nil {
		return err
	}

	for range rounds {
		err = d.round(actorA, actorB, a, b)
		if err != nil {
			return err
		}
	}

	return nil
}

func (d *PrisonerDilemma) round(a, b actor.Actor, pa, pb *Participant) error {
	stepA, err := a.Receive()
	if err != nil {
		return err
	}

	stepB, err := b.Receive()
	if err != nil {
		return err
	}

	switch {
	case stepA == "Y" && stepB == "Y":
		pa.Award(d.cfg.Cooperate)
		pb.Award(d.cfg.Cooperate)
	case stepA == "Y" && stepB == "N":
		pb.Award(d.cfg.Defect)
	case stepA == "N" && stepB == "Y":
		pa.Award(d.cfg.Defect)
	case stepA == "N" && stepB == "N":
		pa.Award(d.cfg.MutualDefects)
		pb.Award(d.cfg.MutualDefects)
	default:
		return fmt.Errorf("invalid input: %s = %s, %s = %s", pa.Name(), stepA, pb.Name(), stepB)
	}

	err = b.Send(stepA)
	if err != nil {
		return err
	}
	err = a.Send(stepB)
	if err != nil {
		return err
	}

	return nil
}
