package main

import (
	"fmt"
	"log"

	"github.com/zhikh23/bcg-game-theory/internal/actor"
	"github.com/zhikh23/bcg-game-theory/internal/game"
)

func main() {
	var prts []*game.Participant

	factoryTrustful := actor.NewProgramFactory("./tests/trust/trustful.py")
	prts = append(prts, game.NewParticipant("Trustful 1", factoryTrustful))
	prts = append(prts, game.NewParticipant("Trustful 2", factoryTrustful))
	prts = append(prts, game.NewParticipant("Trustful 3", factoryTrustful))
	factoryEvil := actor.NewProgramFactory("./tests/trust/evil.py")
	prts = append(prts, game.NewParticipant("Evil 1", factoryEvil))

	g := game.NewTrust(game.TrustConfig{
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
