package main

import (
	"log"
	"os"
	"github.com/xeitodevs/remote-executor/files"
	"github.com/xeitodevs/remote-executor/transport"
	"github.com/xeitodevs/remote-executor/transport/command"
	"fmt"
	"time"
	"github.com/xeitodevs/remote-executor/transport/ssh"
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
	commandQueue := command.NewCommandQueue(delta)
	for worker := 1; worker <= numOfWorkers; worker++ {
		go transport.New().Worker(commandQueue)
	}
	for id, host := range hosts {
		commandQueue.Add(&command.Command{
			Value:           userCommand,
			CreatedOn:       time.Now(),
			Id:              id,
			ChannelResponse: responses,
			SSHAdapter:      ssh.New(host),
		})
	}
	commandQueue.Close()
	for i := 0; i < len(hosts); i++ {
		fmt.Println(<-responses)
	}
}
