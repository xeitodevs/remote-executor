package transport

import "github.com/xeitodevs/remote-executor/transport/command"

type Transport interface {
	Worker(commandsQueue command.Queue)
}

type transport struct {
	commandsQueue <-chan *command.Command
}

func (_ *transport) Worker(commandsQueue command.Queue) {
	for nextCommand := range commandsQueue.Queue() {
		nextCommand.Exec()
	}
}

func New() *transport {
	return &transport{}
}
