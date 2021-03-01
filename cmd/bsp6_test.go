package cmd

// Importieren des Testing-Paketes
import (
	"testing"
)

// Showcase for stress testing

func TestRandomReturnHello(t *testing.T) {
	answer := randomReturnHello()
	if answer != "..." {
		t.Errorf("The test failed, probably try out one more time...")
	}

}

// go test -v -run=^TestRandomReturn -count=100 .\cmd\
