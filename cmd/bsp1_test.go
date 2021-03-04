package cmd

// Importieren des Testing-Paketes
import (
	"testing"
)

// Showcase for running tests

func TestPrintHelloWorld(t *testing.T) {
	t.Log("This Log should only appear in verbose mode or if something went wrong")
	if "Hello World" != printHelloWorld() {
		t.Error("'Hello World' test failed!")
	}

}
