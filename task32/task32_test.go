package task32

import (
	"testing"
	"time"
)

func TestTask32(t *testing.T) {
	beginTime := time.Now()
	Task32(time.Second * 2)
	endTime := time.Now()

	diff := endTime.Sub(beginTime)
	if diff < time.Second*2 {
		t.Error("error")
	}
}
