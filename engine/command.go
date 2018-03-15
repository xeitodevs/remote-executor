package engine

import (
	"time"
	"github.com/xeitodevs/remote-executor/transport/ssh"
)

type Command struct {
	Id               int
	Value            string
	CreatedOn        time.Time
	ChannelResponse  chan<- string
	TransportAdapter ssh.TransportAdapter
}

func (c *Command) Exec() {
	err := c.TransportAdapter.Connect()
	if err != nil {
		c.ChannelResponse <- "an error occurred" + err.Error()
		return
	}
	response := c.TransportAdapter.Run(c.Value)
	c.ChannelResponse <- response
}
