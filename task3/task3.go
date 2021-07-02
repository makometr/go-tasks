package task3

import (
	"errors"

	"golang.org/x/sync/errgroup"
)

// Задание:
/*
	Дана последовательность чисел  (2,4,6,8,10) найти их сумму квадратов(4+9+16….) с использованием конкурентных вычислений.
*/

// Task3 with wg
func Task3(nums []int) (int, error) {
	results := make(chan int, len(nums))
	done := make(chan interface{})
	result := 0

	go func() {
		for v := range results {
			result += v
		}
		done <- struct{}{}
	}()

	var g errgroup.Group
	for i := 0; i < len(nums); i++ {
		index := i
		g.Go(func() error {
			if nums[index] == 0 {
				return errors.New("Error: zero")
			}
			results <- nums[index] * nums[index]
			return nil
		})
	}
	err := g.Wait()
	close(results)
	<-done

	if err != nil {
		return 0, err
	}

	return result, nil
}
