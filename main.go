package main

import (
	"./message"
	"./processes"
	"fmt"
	"github.com/akamensky/argparse"
	"os"
)

func main() {
	parser := argparse.NewParser("json_over_channel",
		"Open provided JSON file and sends it through a channel")
	// filename flag
	f := parser.String("f", "file", &argparse.Options{Required: true, Help: "JSON file to open"})

	err := parser.Parse(os.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Print(parser.Usage(err))
	}

	// initialize processes and channel
	var pA = processes.ProcessAFromJsonFile(*f)
	var pB processes.ProcessB
	var c = make(chan message.Message)

	go pA.Start(c)
	pB.Start(c)
}
