package game

import "github.com/zhikh23/bcg-game-theory/internal/actor"

type Name string
type Score int

type Participant struct {
	factory actor.Factory
	name    Name
	score   Score
}

func NewParticipant(name string, factory actor.Factory) *Participant {
	return &Participant{
		factory: factory,
		name:    Name(name),
		score:   0,
	}
}

func (p *Participant) Actor() (actor.Actor, error) {
	return p.factory.New()
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
