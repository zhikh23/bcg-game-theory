package game

import (
	"fmt"
	"strconv"

	"github.com/zhikh23/bcg-game-theory/internal/actor"
)

type TrustConfig struct {
	Sum Score
}

type Trust struct {
	cfg TrustConfig
}

func NewTrust(cfg TrustConfig) *Trust {
	return &Trust{cfg: cfg}
}

func (g *Trust) Play(rounds int, a, b *Participant) error {
	actorA, err := a.Actor()
	if err != nil {
		return err
	}

	actorB, err := b.Actor()
	if err != nil {
		return err
	}

	for range rounds {
		err = g.round(actorA, actorB, a, b)
		if err != nil {
			return err
		}
	}

	return nil
}

func (g *Trust) round(a, b actor.Actor, pa, pb *Participant) error {
	if err := a.Send("I"); err != nil {
		return fmt.Errorf("failed to send role to the actor A: %w", err)
	}

	if err := b.Send("T"); err != nil {
		return fmt.Errorf("failed to send role to the actor B: %w", err)
	}

	msg := fmt.Sprintf("%d", g.cfg.Sum)
	if err := a.Send(msg); err != nil {
		return fmt.Errorf("failed to send sum to the actor A: %w", err)
	}

	investmentStr, err := a.Receive()
	if err != nil {
		return fmt.Errorf("failed to receive investment from the actor A: %w", err)
	}
	investment, err := strconv.Atoi(investmentStr)
	if err != nil {
		return fmt.Errorf("failed to parse investment from the actor A: %w", err)
	}
	if investment < 1 || investment > int(g.cfg.Sum) {
		return fmt.Errorf("invalid investment: %d", investment)
	}

	msg = fmt.Sprintf("%d", investment)
	if err := b.Send(msg); err != nil {
		return fmt.Errorf("failed to send investment to the actor B: %w", err)
	}

	returnedStr, err := b.Receive()
	if err != nil {
		return fmt.Errorf("failed to receive returned from the actor B: %w", err)
	}
	returned, err := strconv.Atoi(returnedStr)
	if err != nil {
		return fmt.Errorf("failed to parse returned from the actor B: %w", err)
	}
	if returned < 0 || returned > 3*investment {
		return fmt.Errorf("returned out of range: %d", returned)
	}

	if err := a.Send(returnedStr); err != nil {
		return fmt.Errorf("failed to send returned from the actor A: %w", err)
	}

	pa.Award(g.cfg.Sum - Score(investment) + Score(returned))
	pb.Award(Score(3*investment - returned))

	return nil
}
