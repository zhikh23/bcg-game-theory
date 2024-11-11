package game

import "github.com/zhikh23/bcg-game-theory/internal/actor"

type Name string
type Score int

type Participant struct {
	actor.Actor
	name  Name
	score Score
}

func NewParticipant(a actor.Actor, name string) *Participant {
	return &Participant{
		Actor: a,
		name:  Name(name),
		score: 0,
	}
}

func (p *Participant) Name() Name {
	return p.name
}

func (p *Participant) Score() Score {
	return p.score
}

func (p *Participant) Award(score Score) {
	p.score += score
}
