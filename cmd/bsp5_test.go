package cmd

// Importieren des Testing-Paketes
import (
	"testing"
)

// Showcase for raceconditions in tests
var superCount int

func TestStartRace(t *testing.T) {
	// Start subtest
	t.Run("raceDown", func(t *testing.T) {
		t.Parallel()
		countDown(t)
	})
	// Start second subtest
	t.Run("raceUp", func(t *testing.T) {
		t.Parallel()
		countUp(t)
	})

}

func countUp(t *testing.T) {
	superCount++
	t.Log(superCount)
}

func countDown(t *testing.T) {
	superCount--
	t.Log(superCount)
}
