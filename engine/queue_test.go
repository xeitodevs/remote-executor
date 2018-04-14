package engine

import (
	"testing"
	"time"
	"sync/atomic"
	"github.com/stretchr/testify/assert"
)

type FakeCommand struct {
	Id              uint
	ChannelResponse chan<- uint
	taskCounter     *uint32
}

func (c *FakeCommand) Exec() {

	duration := time.Duration(c.Id) * time.Second
	time.Sleep(duration)
	c.ChannelResponse <- c.Id
	atomic.AddUint32(c.taskCounter, uint32(1))
}

func TestTaskQueue(t *testing.T) {

	var taskCounter uint32 = 0
	const systemPrefetch = 6
	taskQueue := NewCommandQueue(systemPrefetch)
	responseCollector := make(chan uint, systemPrefetch)

	startWorkers(systemPrefetch, taskQueue)
	commandLaunch(systemPrefetch, taskQueue, responseCollector, &taskCounter)
	taskQueue.Close()

	assertResponseCollectorData(systemPrefetch, t, responseCollector)
	assert.Equal(t, uint32(systemPrefetch), taskCounter, "Number of tasks differs from commands launched.")
}

func commandLaunch(systemPrefetch int, taskQueue *taskQueue, responseCollector chan uint, taskCounter *uint32) {
	for command := 1; command <= systemPrefetch; command++ {
		taskQueue.Add(&FakeCommand{uint(command), responseCollector, taskCounter})
	}
}

func startWorkers(systemPrefetch int, taskQueue *taskQueue) {
	for worker := 1; worker <= systemPrefetch; worker++ {
		go New().Run(taskQueue)
	}
}

func assertResponseCollectorData(systemPrefetch int, t *testing.T, responseCollector chan uint) {
	for accomplished := 1; accomplished <= systemPrefetch; accomplished++ {
		assert.Equal(t, uint(accomplished), <-responseCollector, "There is and error in order of returning results")
	}
	assert.Equal(t, 0, len(responseCollector), "The responses channel must be empty at the end.")
}
