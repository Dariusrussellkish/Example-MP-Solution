package processes

import (
	"../messages"
)

// ProcessA contains a messages and passes it onto a receiver process
type ProcessA struct {
	m messages.Message
}

// Initialize a ProcessA from a JSON file
func ProcessAFromJsonFile(filename string) ProcessA {
	m := messages.MessageFromJsonFile(filename)
	return ProcessA{
		m,
	}
}

// Start a ProcessA and send its messages over the channel
func (p *ProcessA) Start(c chan<- messages.Message) {
	c <- p.m
}
