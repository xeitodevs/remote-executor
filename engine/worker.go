package engine

type Worker interface {
	Run(taskQueue Queue)
}

type worker struct {
	taskQueue <-chan Task
}

func (_ *worker) Run(taskQueue Queue) {
	for nextTask := range taskQueue.Queue() {
		nextTask.Exec()
	}
}

func New() *worker {
	return &worker{}
}
