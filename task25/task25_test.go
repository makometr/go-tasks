package task25

import (
	"sync"
	"testing"
)

func TestSafeCounter(t *testing.T) {
	counter := SafeCounter{}
	chans := 100
	incs := 10

	var wg sync.WaitGroup
	wg.Add(chans)
	for i := 0; i < chans; i++ {
		go func() {
			for i := 0; i < incs; i++ {
				counter.Inc()
			}
			wg.Done()
		}()
	}

	wg.Wait()
	if counter.Get() != int64(chans)*int64(incs) {
		t.Errorf("SafeCounter.Get() = %v, want %v", counter.Get(), chans*incs)
	}
}
