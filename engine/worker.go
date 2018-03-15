package engine


type Worker interface {
	Run(commandsQueue Queue)
}

type worker struct {
	commandsQueue <-chan *Command
}

func (_ *worker) Run(commandsQueue Queue) {
	for nextCommand := range commandsQueue.Queue() {
		nextCommand.Exec()
	}
}

func New() *worker {
	return &worker{}
}
