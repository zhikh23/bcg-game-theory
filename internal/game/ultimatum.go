package game

import (
	"fmt"
	"strconv"

	"github.com/zhikh23/bcg-game-theory/internal/actor"
)

type UltimatumConfig struct {
	Sum Score
}

type Ultimatum struct {
	cfg UltimatumConfig
}

func NewUltimatumGame(cfg UltimatumConfig) *Ultimatum {
	return &Ultimatum{cfg: cfg}
}

func (u *Ultimatum) Play(rounds int, a, b *Participant) error {
	actorA, err := a.Actor()
	if err != nil {
		return err
	}

	actorB, err := b.Actor()
	if err != nil {
		return err
	}

	for range rounds {
		err = u.round(actorA, actorB, a, b)
		if err != nil {
			return err
		}
	}

	return nil
}

func (u *Ultimatum) round(a, b actor.Actor, pa, pb *Participant) error {
	if err := a.Send("A"); err != nil {
		return fmt.Errorf("failed to send actor A: %w", err)
	}
	if err := b.Send("B"); err != nil {
		return fmt.Errorf("failed to send actor B: %w", err)
	}

	msg := fmt.Sprintf("%d", u.cfg.Sum)
	if err := a.Send(msg); err != nil {
		return fmt.Errorf("failed to send actor A: %w", err)
	}
	if err := b.Send(msg); err != nil {
		return fmt.Errorf("failed to send actor B: %w", err)
	}

	offerStr, err := a.Receive()
	if err != nil {
		return fmt.Errorf("failed to receive offer from A: %w", err)
	}
	offer, err := strconv.Atoi(offerStr)
	if err != nil {
		return fmt.Errorf("incorrect answer of A: %w", err)
	}
	if offer < 0 || offer > int(u.cfg.Sum) {
		return fmt.Errorf("invalid answer of A: %d", offer)
	}

	if err = b.Send(offerStr); err != nil {
		return fmt.Errorf("failed to send offer to B: %w", err)
	}

	answer, err := b.Receive()
	if err != nil {
		return fmt.Errorf("failed to receive answer from B: %w", err)
	}

	if answer == "Y" {
		pa.Award(u.cfg.Sum)
		pb.Award(u.cfg.Sum)
	} else if answer == "N" {
		/* nothing */
	} else {
		return fmt.Errorf("invalid answer of B: %s", answer)
	}

	return nil
}
