package transport

type Transport interface {
	ExecuteCommand(command string, server string, channel chan<- string)
}
