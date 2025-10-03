package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Token struct {
	token int
}

func init_chan() [][]chan *Token {
	c1 := make(chan *Token)
	c2 := make(chan *Token)
	c3 := make(chan *Token)
	c4 := make(chan *Token)

	return [][]chan *Token{{c1, c2}, {c1, c3}, {c2, c3, c4}, {c4}}

}

func main() {
	chanels := init_chan()
	go player("Player 1", chanels[0])
	go player("Player 2", chanels[1])
	go player("Player 3", chanels[2])
	go player("Player 4", chanels[3])

	chanels[3][0] <- new(Token)
	time.Sleep(3 * time.Second)
	<-chanels[3][0]
}

func player(name string, chans []chan *Token) {
	for {
		for _, ch := range chans {
			select {
			case tok := <-ch:
				fmt.Println("Reçu :", name)
				time.Sleep(200 * time.Millisecond)
				chanel_choice := chans[rand.Intn(len(chans))]
				chanel_choice <- tok

			default:
				// chan suivant
			}
		}
	}
}
