package cmd

import (
	"testing"
)

func TestMult(t *testing.T) {
	total := mult(5) // alternative try mult2(5)
	if total != 25 {
		t.Errorf("Simple test failed, Tested 5 got %d", total)
	}
}

//testing a function with multiple arguments. --> defining what to expect
func TestMultList(t *testing.T) {
	tableMult := []struct {
		x      int
		result int
	}{
		{5, 25},
		{6, 36},
		{10, 100},
		{1, 1},
		{0, 0},
		{-1, 1},
	}

	for _, table := range tableMult {
		result := mult(table.x)
		if result != table.result {
			t.Errorf("The multiplication of %d with %d should give %d not %d", table.x, table.x, table.result, result)
		}
	}
}
