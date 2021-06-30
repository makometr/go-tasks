package task25

import (
	"sync/atomic"
)

// SafeCounter is
type SafeCounter struct {
	data int64
}

// Inc incremets counter value atomically
func (sc *SafeCounter) Inc() int64 {
	return atomic.AddInt64(&sc.data, 1)
}

// Get current counter value atomically
func (sc *SafeCounter) Get() int64 {
	return atomic.LoadInt64(&sc.data)
}
