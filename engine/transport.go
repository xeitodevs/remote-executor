package engine


type TransportAdapter interface {
	Connect() error
	Close()
	Run(command string) string
}
