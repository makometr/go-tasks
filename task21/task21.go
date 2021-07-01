package task21

import (
	"fmt"
	"sync"
)

// 21.	Написать программу, которая в конкурентном виде читает элементы из массива в stdout.

// Другой вариант с cихнронизацией и множественным чтением реализован во втором задании

// Sync "Конкурентно" ~синхронно
func Sync(data []int) {
	ch := make(chan int, 10)

	go func() {
		defer close(ch)
		for i := 0; i < len(data); i++ {
			ch <- data[i]
		}
	}()

	for v := range ch {
		fmt.Println(v)
	}
}

// Async Более конкурентно, но без синхронизации
func Async(workersNum int, data []int) {
	elemsPerWorker := len(data) / workersNum
	tailSize := len(data) % workersNum

	var wg sync.WaitGroup
	wg.Add(workersNum)
	for wi := 0; wi < workersNum; wi++ {
		beginIndex := elemsPerWorker * wi
		size := elemsPerWorker
		if wi == workersNum-1 {
			size += tailSize
		}
		go func(index, size, id int) {
			for i := index; i < index+size; i++ {
				// fmt.Println(id, ":", data[i])
				fmt.Println(data[i])
			}
			wg.Done()
		}(beginIndex, size, wi)

	}

	wg.Wait()
}
