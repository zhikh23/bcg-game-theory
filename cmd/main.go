package main

import (
	"fmt"
	"log"

	"github.com/zhikh23/bcg-game-theory/internal/actor"
	"github.com/zhikh23/bcg-game-theory/internal/game"
)

func main() {
	kindActor, err := actor.NewProgramActor("./tests/kind.py")
	if err != nil {
		log.Fatal(err)
	}
	kind := game.NewParticipant(kindActor, "Kind")

	titForTatActor, err := actor.NewProgramActor("./tests/tit_for_tat.py")
	if err != nil {
		log.Fatal(err)
	}
	titForTat := game.NewParticipant(titForTatActor, "Tit for tat")

	evilActor, err := actor.NewProgramActor("./tests/evil.py")
	if err != nil {
		log.Fatal(err)
	}
	_ = game.NewParticipant(evilActor, "Evil")

	g := game.NewPrisonerDilemma(titForTat, kind)

	err = g.Start()
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 10; i++ {
		err = g.Play()
		if err != nil {
			log.Fatal(err)
		}
	}

	for name, score := range g.Results() {
		fmt.Printf("%s: %d\n", name, score)
	}
}
