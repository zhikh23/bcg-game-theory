package main

import (
	"fmt"
	"log"

	"github.com/zhikh23/bcg-game-theory/internal/actor"
	"github.com/zhikh23/bcg-game-theory/internal/game"
)

func main() {
	var prts []*game.Participant

	kindFactory := actor.NewProgramFactory("./tests/kind.py")
	prts = append(prts, game.NewParticipant("Kind 1", kindFactory))

	titForTatFactory := actor.NewProgramFactory("./tests/tit_for_tat.py")
	prts = append(prts, game.NewParticipant("Tit for tat 1", titForTatFactory))

	evilFactory := actor.NewProgramFactory("./tests/evil.py")
	prts = append(prts, game.NewParticipant("Evil 1", evilFactory))
	prts = append(prts, game.NewParticipant("Evil 2", evilFactory))
	prts = append(prts, game.NewParticipant("Evil 3", evilFactory))

	g := game.NewPrisonerDilemma(game.PrisonerDilemmaConfig{
		Cooperate:     5,
		Defect:        10,
		MutualDefects: 1,
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
