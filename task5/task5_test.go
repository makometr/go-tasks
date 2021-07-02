package task5

import (
	"testing"
	"time"
)

func TestTask5(t *testing.T) {
	secs := 2
	begin := time.Now()
	Task5(secs)
	end := time.Now()

	if end.Sub(begin) > time.Duration(secs)+time.Duration(150)*time.Millisecond {
		t.Errorf("Error more then %d sec: %d", secs, end.Sub(begin))
	}

}
