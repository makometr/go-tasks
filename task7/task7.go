package task7

import (
	"sync"
)

// Задание
/*
	Реализовать конкурентную запись в map.
*/

// SafeMap provides concurrent access to map int-int data structure
type SafeMap struct {
	data map[int]int
	mx   sync.RWMutex
}

// NewSafeMap creates and initialize map object
func NewSafeMap() *SafeMap {
	return &SafeMap{data: make(map[int]int)}
}

// Read is concurrent-safe read value by key
func (sm *SafeMap) Read(key int) (int, bool) {
	sm.mx.RLock()
	defer sm.mx.RUnlock()

	v, ok := sm.data[key]
	return v, ok
}

// Write is concurrent-safe write new value by key
func (sm *SafeMap) Write(key, newValue int) {
	sm.mx.Lock()
	defer sm.mx.Unlock()

	sm.data[key] = newValue
}

// Update provides concurrent-safe update of ley value by func f
func (sm *SafeMap) Update(key int, f func(int) int) {
	sm.mx.Lock()
	defer sm.mx.Unlock()

	oldValue := sm.data[key]
	sm.data[key] = f(oldValue)
}

// Task7 with test-func present correcntess of concurrent update by inc n keys in chs number of channels
func Task7(chs int, n int) *SafeMap {
	mapa := NewSafeMap()
	var wg sync.WaitGroup
	wg.Add(chs)

	writer := func() {
		for i := 0; i < n; i++ {
			mapa.Update(i, func(old int) int {
				return old + 1
			})
		}
		wg.Done()
	}

	for i := 0; i < chs; i++ {
		go writer()
	}
	wg.Wait()

	return mapa
}
