package main

import (
	"log"

	"github.com/zhikh23/bcg-game-theory/internal/actor"
	"github.com/zhikh23/bcg-game-theory/internal/game"
)

func main() {
	_, err := actor.NewProgramActor("./tests/kind.py")
	if err != nil {
		log.Fatal(err)
	}
	echo, err := actor.NewProgramActor("./tests/tit_for_tat.py")
	if err != nil {
		log.Fatal(err)
	}
	evil, err := actor.NewProgramActor("./tests/evil.py")
	if err != nil {
		log.Fatal(err)
	}

	g := game.NewPrisonerDilemma(echo, evil)

	err = g.Start()
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 10; i++ {
		err = g.Round()
		if err != nil {
			log.Fatal(err)
		}
	}

}
