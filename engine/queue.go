package engine

type Queue interface {
	Queue() chan Task
	Add(task Task)
	Close()
}

type taskQueue struct {
	taskQueue chan Task
}

func (q *taskQueue) Queue() chan Task {
	return q.taskQueue
}

func (q *taskQueue) Add(command Task) {
	q.Queue() <- command
}

func (q *taskQueue) Close() {
	close(q.taskQueue)
}

func NewCommandQueue(delta uint) *taskQueue {
	return &taskQueue{
		make(chan Task, delta),
	}
}
