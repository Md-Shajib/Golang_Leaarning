package logger

type ConsoleLogger struct{}

func NewConsoleLogger() *ConsoleLogger {
	return &ConsoleLogger{}
}

func (c *ConsoleLogger) Info(message string) {
	println("[INFO]: " + message)
}

func (c *ConsoleLogger) Error(message string) {
	println("[ERROR]: " + message)
}