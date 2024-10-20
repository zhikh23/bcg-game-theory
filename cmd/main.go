package main

import (
	"github.com/zhikh23/bcg-game-theory/internal/actor"
	"github.com/zhikh23/bcg-game-theory/internal/game"
	"time"
)

func main() {
	kind := actor.NewInternalActor(func(stdin <-chan string, stdout chan<- string) {
		stdout <- "Y"
		for {
			select {
			case anotherStep := <-stdin:
				stdout <- anotherStep
			default:
				time.Sleep(time.Millisecond)
			}
		}
	})

	evil := actor.NewInternalActor(func(stdin <-chan string, stdout chan<- string) {
		stdout <- "N"
		for {
			select {
			case <-stdin:
				stdout <- "N"
			default:
				time.Sleep(time.Millisecond)
			}
		}
	})

	g := game.NewDilemma(kind, evil)
	g.Start()
	g.Round()
	g.Round()
	g.Round()
	g.Round()
}
