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

	if time.Duration(end.Sub(begin).Seconds()) > time.Duration(secs)+time.Duration(100)*time.Millisecond {
		t.Errorf("Error more then %d sec: %f sec", secs, end.Sub(begin).Seconds())
	}

}
