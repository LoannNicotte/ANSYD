package main

import (
	"fmt"
	"time"
)

type Ball struct {
	vect []int
}

func main() {
	table := make(chan *Ball)
	go player("Alice", table)
	go player("Bob", table)
	ball := new(Ball)
	ball.vect = []int{0, 0, 0, 0, 0}
	table <- ball // game on; toss the ball
	time.Sleep(1 * time.Second)
	<-table // game over; grab the ball
}

func inc(ball Ball) {
	for i := range ball.vect {
		ball.vect[i] += i + 1
	}
}

func player(name string, table chan *Ball) {
	for {
		ball := <-table
		inc(*ball)
		fmt.Println(name, ball.vect)
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}
