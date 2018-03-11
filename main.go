package main

import (
	"log"
	"os"
	"github.com/xeitodevs/remote-executor/files"
	"github.com/xeitodevs/remote-executor/transport"
	"fmt"
)

func main() {

	args := os.Args[1:]
	configFile := args[0]
	command := args[1]

	hosts, err := files.FileLinesExtractor(configFile)
	if err != nil {
		log.Fatal("Cannot parse input hosts file.")
	}

	responses := make(chan string)
	calls := 0
	for _, host := range hosts {
		go transport.ExecuteCommand(command, host, responses)
		calls++
	}
	for i := 0; i < calls; i++ {
		fmt.Println(<-responses)
	}
}
