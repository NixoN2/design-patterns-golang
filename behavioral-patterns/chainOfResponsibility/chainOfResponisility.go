package chainOfResponsibility

import (
	"io"
	"fmt"
	"strings"
)

type ChainLogger interface {
	Next(string)
}

type FirstLogger struct {
	NextChain ChainLogger
}

func (f *FirstLogger) Next(s string) {
	fmt.Printf("First logger: %s\n", s)
	if f.NextChain != nil {
		f.NextChain.Next(s)
	}
}

type SecondLogger struct {
	NextChain ChainLogger
}

func (se *SecondLogger) Next(s string) {
	if strings.Contains(strings.ToLower(s), "hello") {
		fmt.Printf("Second logger: %s\n",s)
		if se.NextChain != nil {
			se.NextChain.Next(s)
		}
		return
	}

	fmt.Printf("Finishing in second logging\n\n")
}

type WriterLogger struct {
	NextChain ChainLogger
	Writer io.Writer
}

func (w *WriterLogger) Next(s string) {
	if w.Writer != nil {
		w.Writer.Write([]byte("WriterLogger: " + s))
	}

	if w.NextChain != nil {
		w.NextChain.Next(s)
	}
}

type TestWriter struct {
	receivedMessage string
}

func (m *TestWriter) Write(p []byte) (int, error) {
	m.receivedMessage += string(p)
	return len(p), nil
}

func (m *TestWriter) Next(s string) {
	m.Write([]byte(s))
}