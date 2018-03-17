package main

import (
	"github.com/xeitodevs/remote-executor/input"
	"fmt"
	"time"
	"github.com/xeitodevs/remote-executor/transport/ssh"
	"github.com/xeitodevs/remote-executor/engine"
	"os"
)

func main() {

	args := input.ParseArgs()
	hosts := input.GetHosts(os.Stdin)

	delta := len(hosts)
	responses := make(chan string)
	commandQueue := engine.NewCommandQueue(delta)
	for worker := 1; worker <= args.MaxConcurrency; worker++ {
		go engine.New().Run(commandQueue)
	}
	for id, host := range hosts {
		commandQueue.Add(&engine.Command{
			Value:            args.Command,
			CreatedOn:        time.Now(),
			Id:               id,
			ChannelResponse:  responses,
			TransportAdapter: ssh.New(host, args.User, args.PrivateKeyPath), // In this case, is SSH , but could be http ...
		})
	}
	commandQueue.Close()
	for i := 0; i < len(hosts); i++ {
		fmt.Println(<-responses)
	}
}
