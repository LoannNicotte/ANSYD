package main

import "fmt"

type Message struct {
	max    int
	leader bool
}

type NodeConfig struct {
	in  []chan *Message
	out []chan *Message
}

// Initialise les canaux du graphe et renvoie les configs des nœuds
func init_nodes() []NodeConfig {
	// Créer 2 canaux par arête (aller + retour)
	c13 := make(chan *Message, 1)
	c31 := make(chan *Message, 1)

	c23 := make(chan *Message, 1)
	c32 := make(chan *Message, 1)

	c35 := make(chan *Message, 1)
	c53 := make(chan *Message, 1)

	c14 := make(chan *Message, 1)
	c41 := make(chan *Message, 1)

	c45 := make(chan *Message, 1)
	c54 := make(chan *Message, 1)

	// Config des nœuds (in, out)
	n1 := NodeConfig{in: []chan *Message{c31, c41}, out: []chan *Message{c13, c14}}
	n2 := NodeConfig{in: []chan *Message{c32}, out: []chan *Message{c23}}
	n3 := NodeConfig{in: []chan *Message{c13, c23, c53}, out: []chan *Message{c31, c32, c35}}
	n4 := NodeConfig{in: []chan *Message{c14, c54}, out: []chan *Message{c41, c45}}
	n5 := NodeConfig{in: []chan *Message{c35, c45}, out: []chan *Message{c53, c54}}

	return []NodeConfig{n1, n2, n3, n4, n5}
}

func node(id int, rounds int, cfg NodeConfig, decision chan<- *Message) {
	state := &Message{max: id, leader: false}

	// envoi initial de son propre id
	for _, ch := range cfg.out {
		ch <- &Message{max: state.max}
	}

	for r := 0; r < rounds; r++ {
		// Recevoir
		for _, ch := range cfg.in {
			msg := <-ch
			if msg.max > state.max {
				state.max = msg.max
			}
		}
		// Envoyer
		for _, ch := range cfg.out {
			ch <- &Message{max: state.max}
		}
	}

	if state.max == id {
		state.leader = true
	}
	decision <- state
}

func main() {
	rounds := 2
	nodes := init_nodes()

	// Canaux de décision
	d1 := make(chan *Message)
	d2 := make(chan *Message)
	d3 := make(chan *Message)
	d4 := make(chan *Message)
	d5 := make(chan *Message)

	// Lancer les goroutines
	go node(1, rounds, nodes[0], d1)
	go node(2, rounds, nodes[1], d2)
	go node(3, rounds, nodes[2], d3)
	go node(4, rounds, nodes[3], d4)
	go node(5, rounds, nodes[4], d5)

	// Collecter résultats
	in1 := <-d1
	in2 := <-d2
	in3 := <-d3
	in4 := <-d4
	in5 := <-d5

	fmt.Println("Résultats finaux :")
	fmt.Println("Node 1:", in1.max, "leader =", in1.leader)
	fmt.Println("Node 2:", in2.max, "leader =", in2.leader)
	fmt.Println("Node 3:", in3.max, "leader =", in3.leader)
	fmt.Println("Node 4:", in4.max, "leader =", in4.leader)
	fmt.Println("Node 5:", in5.max, "leader =", in5.leader)
}
