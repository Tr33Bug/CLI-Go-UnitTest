package cmd

// Importieren des Testing-Paketes
import (
	"testing"
)

// Showcase for running tests

func TestPrintHelloWorld(t *testing.T) {
	if "Hello World" != printHelloWorld() {
		t.Error("Hello World test failed!")
	}

}
