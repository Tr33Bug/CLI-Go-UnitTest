package cmd

// Importieren des Testing-Paketes
import (
	"testing"
)

// Showcase for running specific tests

func TestPrintHello(t *testing.T) {
	if "Hello" != printHello() {
		t.Error("Printing Hello has failed!")
	}

}

func TestPrintWorld(t *testing.T) {
	t.Run("subtestPrintWorld", func(t *testing.T) {
		if "World" != printWorld() {
			t.Error("Printing World has failed!")
		}
	})
}
