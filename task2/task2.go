package task2

import "fmt"

// Задание:
/*
	Написать программу, которая конкурентн	о рассчитает значение квадратов значений взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/

// Комментарий к заданию:
/*
	- Будем считать, то результаты требуется вывести в порядке их появления в слайсе чисел.
*/

// Task2 release
func Task2(nums ...int) (results []string) {
	chans := make([]chan int, len(nums), len(nums))
	for i := 0; i < len(chans); i++ {
		chans[i] = make(chan int)
	}

	for i, num := range nums {
		go func(i, n int) {
			chans[i] <- n * n
		}(i, num)
	}

	for i := 0; i < len(chans); i++ {
		result := <-chans[i]
		fmt.Println(result)
		results = append(results, fmt.Sprint(result))
	}

	fmt.Println("All squares calculated!")
	return results
}
