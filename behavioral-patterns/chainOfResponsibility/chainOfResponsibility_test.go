package chainOfResponsibility

import (
	"testing"
	"strings"
)

func TestCreateDefaultChain(t *testing.T) {
	writer := TestWriter{}
	writerLogger := WriterLogger{ Writer: &writer}
	second := SecondLogger{NextChain: &writerLogger}
	chain := FirstLogger{NextChain: &second}

	t.Run("3 loggers, 2 of the writes to console, second only if it founds the word 'hello', third writes to some variable if second found 'hello'",
	func(t *testing.T) {
		chain.Next("message that breaks the chain \n")
		if writer.receivedMessage != "" {
			t.Fatal("Last link should not receive any message")
		}

		chain.Next("Hello\n")
		if !strings.Contains(writer.receivedMessage, "Hello") {
			t.Fatal("Last link didn't received expected message")
		}
	})
}