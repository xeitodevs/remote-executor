package transport


type TransportAdapter interface {
	Connect() error
	Close()
	Run(command string) string
}
