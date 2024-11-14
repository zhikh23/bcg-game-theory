package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/zhikh23/bcg-game-theory/internal/actor"
	"github.com/zhikh23/bcg-game-theory/internal/game"
)

type GameName string

const (
	Dilemma   GameName = "dilemma"
	Ultimatum GameName = "ultimatum"
	Trust     GameName = "trust"
)

var games map[GameName]game.BinaryGame

func init() {
	games = make(map[GameName]game.BinaryGame)
	games[Dilemma] = game.NewPrisonerDilemma(game.PrisonerDilemmaConfig{
		MutualDefects: 1,
		Defect:        8,
		Cooperate:     5,
	})
	games[Ultimatum] = game.NewUltimatumGame(game.UltimatumConfig{
		Sum: 10,
	})
	games[Trust] = game.NewTrust(game.TrustConfig{
		Sum: 10,
	})
}

func main() {
	gameName := os.Args[1]
	g, ok := games[GameName(gameName)]
	if !ok {
		log.Fatal("game not found: ", gameName)
	}

	prts := make([]*game.Participant, 0, len(os.Args)-2)
	for _, programPath := range os.Args[2:] {
		program := strings.TrimSuffix(path.Base(programPath), path.Ext(programPath))
		factory := actor.NewProgramFactory(programPath)
		prts = append(prts, game.NewParticipant(program, factory))
	}

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

	fmt.Printf("Program,Score\n")
	for _, prt := range prts {
		fmt.Printf("%s,%d\n", prt.Name(), prt.Score())
	}
}
