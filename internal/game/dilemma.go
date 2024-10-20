package game

import (
	"errors"
	"fmt"

	"github.com/zhikh23/bcg-game-theory/internal/actor"
)

type Dilemma struct {
	a actor.Actor
	b actor.Actor
}

func NewDilemma(a, b actor.Actor) *Dilemma {
	return &Dilemma{a: a, b: b}
}

func (d *Dilemma) Start() error {
	var err error
	err = errors.Join(err, d.a.Start())
	err = errors.Join(err, d.b.Start())
	return err
}

func (d *Dilemma) Round() error {
	stepA, err := d.a.Receive()
	if err != nil {
		return err
	}

	stepB, err := d.b.Receive()
	if err != nil {
		return err
	}

	fmt.Printf("A: %s\nB: %s\n", stepA, stepB)

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
