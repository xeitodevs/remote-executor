package command

import (
	"time"
	"github.com/xeitodevs/remote-executor/transport/ssh"
)

type Command struct {
	Id              int
	Value           string
	CreatedOn       time.Time
	ChannelResponse chan<- string
	SSHAdapter      ssh.SSHAdapter
}

func (c *Command) Exec() {
	err := c.SSHAdapter.Connect()
	if err != nil {
		c.ChannelResponse <- "an error occurred" + err.Error()
		return
	}
	response := c.SSHAdapter.Run(c.Value)
	c.ChannelResponse <- response
}
