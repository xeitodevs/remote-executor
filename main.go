package main

import (
	"log"
	"os"
	"github.com/xeitodevs/remote-executor/files"
	"fmt"
	"time"
	"github.com/xeitodevs/remote-executor/transport/ssh"
	"github.com/xeitodevs/remote-executor/engine"
)

const numOfWorkers = 10

func main() {
	args := os.Args[1:]
	configFile := args[0]
	userCommand := args[1]

	hosts, err := files.FileLinesExtractor(configFile)
	if err != nil {
		log.Fatal("Cannot parse input hosts file.")
	}

	delta := len(hosts)
	responses := make(chan string)
	commandQueue := engine.NewCommandQueue(delta)
	for worker := 1; worker <= numOfWorkers; worker++ {
		go engine.New().Run(commandQueue)
	}
	for id, host := range hosts {
		commandQueue.Add(&engine.Command{
			Value:            userCommand,
			CreatedOn:        time.Now(),
			Id:               id,
			ChannelResponse:  responses,
			TransportAdapter: ssh.New(host), // In this case, is SSH , but could be http ...
		})
	}
	commandQueue.Close()
	for i := 0; i < len(hosts); i++ {
		fmt.Println(<-responses)
	}
}
