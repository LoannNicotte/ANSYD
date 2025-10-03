package main

type Message struct {
	max       int
	terminate bool
}

func node(id int, channelinput chan *Message,
	channeloutput chan *Message, Decision chan *Message) {
	State := new(Message)
	State.terminate = false
	State.max = id
	for {
		select {}
	}
}
func main() {
	c1 := make(chan *Message, 1)
	c2 := make(chan *Message, 1)
	c3 := make(chan *Message, 1)

	m1 := make(chan *Message)
	m2 := make(chan *Message)
	m3 := make(chan *Message)

	go node(1, c3, c2, m1)
	go node(2, c1, c3, m2)
	go node(3, c2, c1, m3)

	Input1 := <-m1
	Input2 := <-m2
	Input3 := <-m3
	println(Input1.max, Input3.max, Input2.max)
}
