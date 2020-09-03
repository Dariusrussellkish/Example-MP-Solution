package processes

import (
	"../messages"
	"fmt"
)

// ProcessB waits for a messages and then prints it
type ProcessB struct {
	m messages.Message
}

// Start a ProcessB, wait for a messages on the channel and print it
func (p *ProcessB) Start(c <-chan messages.Message) {
	p.m = <-c
	fmt.Println(p.m.String())
}
