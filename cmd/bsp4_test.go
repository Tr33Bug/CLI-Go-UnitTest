package cmd

// Importieren des Testing-Paketes
import (
	"testing"
	"time"
)

// Showcase for failing tests

func TestSumUp(t *testing.T) {
	a, b := 5, 5
	result := sumUp(a, b)
	t.Log("The following test should fail. If you dont want it to fail, change the test in ./cmd/bsp4_test.go")
	if result != 11 {
		t.Errorf("This test should fail. Dont worry!")
	}

	// Replace with te following to fix the test
	/*
		if result != 10 {
			t.Errorf("Failed to calculate the sum of 5 and 5")
		}
	*/

}

func TestRetTrue(t *testing.T) {
	time.Sleep(4 * time.Second)
	if !retTrue() {
		t.Errorf("This should never Fail...")
	}
}
