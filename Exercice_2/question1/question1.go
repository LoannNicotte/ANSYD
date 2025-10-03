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
	Counter := 1
	for {
		select {}
	}
}
func main() {
	// Add all the channels
	go node(1) //Add all the channels
	go node(2) //Add all the channels
	go node(3) //Add all the channels
	counter_main := 1
	Input := <- //to complete
	println(Input.max)
}
