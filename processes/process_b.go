package processes

import (
	"../message"
	"fmt"
)

// ProcessB waits for a message and then prints it
type ProcessB struct {
	m message.Message
}

// Start a ProcessB, wait for a message on the channel and print it
func (p *ProcessB) Start(c <-chan message.Message) {
	p.m = <-c
	fmt.Println(p.m.String())
}
