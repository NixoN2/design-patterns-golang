package command

import "testing"

func TestCommand(t *testing.T) {
	queue := CommandQueue{}
	queue.AddCommand(CreateCommand("First message"))
	queue.AddCommand(CreateCommand("Second message"))
	queue.AddCommand(CreateCommand("Third message"))
	queue.AddCommand(CreateCommand("Fourth message"))
	queue.AddCommand(CreateCommand("Fifth message"))
}
