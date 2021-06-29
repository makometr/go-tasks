package task15

import "testing"

func TestSwap1(t *testing.T) {
	var a, b int = 10, 100

	Swap1(&a, &b)
	if a != 100 || b != 10 {
		t.Error("Error in swap. Result:", a, b, "Expected:", 100, 10)
	}
}

func TestSwap2(t *testing.T) {
	var a, b int = 10, 100

	Swap1(&a, &b)
	if a != 100 || b != 10 {
		t.Error("Error in swap. Result:", a, b, "Expected:", 100, 10)
	}
}
