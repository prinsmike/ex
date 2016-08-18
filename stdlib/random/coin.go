package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Result struct {
	Players []int
}

// Flips a coin.
func flip() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(2)
}

func playRound(r *Result) {
	var players []int
	for _, v := range r.Players {
		f := flip()
		if f == 1 {
			players = append(players, v)
		}
	}
	r.Players = players
}

func play(r *Result) {
	fmt.Println("Starting the game...")
	fmt.Printf("%d players.\n", len(r.Players))
	for ; len(r.Players) > 1; {
		playRound(r)
		fmt.Printf("%d players.\n", len(r.Players))
	}
	if len(r.Players) == 1 {
		fmt.Printf("Player %d won!\n", r.Players[0])
	} else {
		fmt.Printf("Game ended undetermined.\n")
	}
}

func main() {
	r := Result{}
	for i := 1; i <= 1000; i++ {
		r.Players = append(r.Players, i)
	}
	play(&r)
}