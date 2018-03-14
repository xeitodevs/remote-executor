package command

type Queue interface {
	Queue() chan *Command
	Add(command *Command)
	Close()
}

type commandQueue struct {
	commandQueue chan *Command
}

func (q *commandQueue) Queue() chan *Command {
	return q.commandQueue
}

func (q *commandQueue) Add(command *Command) {
	q.Queue() <- command
}

func (q *commandQueue) Close() {
	close(q.commandQueue)
}

func NewCommandQueue(delta int) *commandQueue {
	return &commandQueue{
		make(chan *Command, delta),
	}
}
