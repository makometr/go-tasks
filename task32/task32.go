package task32

import "time"

// Task32 sleep func
func Task32(d time.Duration) {
	timer := time.After(d)
	<-timer
	return
}
