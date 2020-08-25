package processes

import (
	"../message"
)

// ProcessA contains a message and passes it onto a receiver process
type ProcessA struct {
	m message.Message
}

// Initialize a ProcessA from a JSON file
func ProcessAFromJsonFile(filename string) ProcessA {
	m := message.MessageFromJsonFile(filename)
	return ProcessA{
		m,
	}
}

// Start a ProcessA and send its message over the channel
func (p *ProcessA) Start(c chan<- message.Message) {
	c <- p.m
}
