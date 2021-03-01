package cmd

// Importieren des Testing-Paketes
import (
	"testing"
)

// Showcase for long tests and skipping in short mode

func TestWaitAndPrint(t *testing.T) {
	// tell the testengine to skip the test in short mode
	if testing.Short() {
		t.Skip("This test is to long, skip in short mode")
	}
	// start test
	if "Hello World" != waitAndPrint() {
		t.Error("Waiting and printing Hello World has failed")
	}

}
