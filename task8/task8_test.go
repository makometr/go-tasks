package task8

import (
	"fmt"
	"testing"
)

func TestTask8(t *testing.T) {
	var n int64 = 15
	Task8(&n, 8, true)
	Task8(&n, 0, false)
	Task8(&n, 1, false)
	Task8(&n, 3, false)

	var expected int64 = 260
	if n != expected {
		t.Error("TestTask8 error. Resultt: ", n, "Expected: ", expected)
		fmt.Printf("%064b\n", n)
		fmt.Printf("%064b\n", expected)
	}
}
