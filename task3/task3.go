package task3

import (
	"sync"
)

// Задание:
/*
	Дана последовательность чисел  (2,4,6,8,10) найти их сумму квадратов(4+9+16….) с использованием конкурентных вычислений.
*/

// Task3 with wg
func Task3(nums []int) int {
	results := make(chan int, len(nums))

	var wg sync.WaitGroup
	wg.Add(len(nums))
	for i := 0; i < len(nums); i++ {
		go func(index int) {
			results <- nums[index] * nums[index]
			wg.Done()
		}(i)
	}
	wg.Wait()

	var result int
	for i := 0; i < cap(results); i++ {
		result += <-results
	}

	return result
}

// func Task3(nums []int, gn int) int {
// 	todo := make(chan int, len(nums))
// 	results := make(chan int, len(nums))

// 	go func() {
// 		for n := range nums {
// 			todo <- n
// 		}
// 	}()

// 	for i := 0; i < gn; i++ {
// 		go func() {
// 			for value := range todo {
// 				results <- value * value
// 			}
// 		}()
// 	}
// }
