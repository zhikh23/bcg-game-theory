package game

import "github.com/zhikh23/bcg-game-theory/internal/actor"

type Participant struct {
	actor.Actor
	score int
}

func NewParticipant(a actor.Actor) *Participant {
	return &Participant{
		Actor: a,
		score: 0,
	}
}

func (p *Participant) Score() int {
	return p.score
}

func (p *Participant) Award(score int) {
	p.score += score
}
