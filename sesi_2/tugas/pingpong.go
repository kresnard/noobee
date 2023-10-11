package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	Name string
	Hit  int
}

const BreakPoint = 11

func main() {
	player := make(chan *Player)
	done := make(chan *Player)

	a := &Player{Name: "Player 1", Hit: 0}
	b := &Player{Name: "Player 2", Hit: 0}

	players := []*Player{a, b}
	finished := make(chan bool)

	for _, p := range players {
		go play(p, player, done, finished)
	}

	player <- a

	go func() {
		for i := 0; i < len(players); i++ {
			<-finished
		}
		close(finished)
	}()

	finish(done)
}

func play(player *Player, playerChan, done chan *Player, finished chan bool) {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 100

	defer func() {
		finished <- true
	}()

	for {
		select {
		case p := <-playerChan:
			v := rand.Intn(max-min) + 1
			time.Sleep(500 * time.Microsecond)
			player.Hit++
			fmt.Println(player.Name, "hit", player.Hit, "dengan value", v)

			if v%BreakPoint == 0 {
				done <- player
				return
			}

			playerChan <- p
		}

	}
}

func finish(done chan *Player) {
	for d := range done {
		fmt.Println(d.Name, "kalah pada hit ke", d.Hit)
	}

}
