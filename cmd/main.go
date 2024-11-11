package main

import (
	"fmt"
	"log"

	"github.com/zhikh23/bcg-game-theory/internal/actor"
	"github.com/zhikh23/bcg-game-theory/internal/game"
)

func main() {
	var prts []*game.Participant

	factoryKind := actor.NewProgramFactory("./tests/ultimatum/random_kind.py")
	prts = append(prts, game.NewParticipant("Kind 1", factoryKind))
	prts = append(prts, game.NewParticipant("Kind 2", factoryKind))
	prts = append(prts, game.NewParticipant("Kind 3", factoryKind))
	prts = append(prts, game.NewParticipant("Kind 4", factoryKind))
	factoryEvil := actor.NewProgramFactory("./tests/ultimatum/random_evil.py")
	prts = append(prts, game.NewParticipant("Evil 1", factoryEvil))
	prts = append(prts, game.NewParticipant("Evil 2", factoryEvil))
	prts = append(prts, game.NewParticipant("Evil 3", factoryEvil))
	prts = append(prts, game.NewParticipant("Evil 4", factoryEvil))

	g := game.NewUltimatumGame(game.UltimatumConfig{
		Sum: 10,
	})
	t := game.NewTournament(g)
	for _, prt := range prts {
		err := t.AddParticipant(prt)
		if err != nil {
			log.Fatal(err)
		}
	}

	err := t.Tour(5)
	if err != nil {
		log.Fatal(err)
	}

	for _, prt := range prts {
		fmt.Printf("%s: %d\n", prt.Name(), prt.Score())
	}
}
