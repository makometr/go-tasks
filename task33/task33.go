package task33

import "fmt"

/* Задание
Даны 2 канала - в первый пишутся рандомные числа после чего они проверяются на четность и отправляются во второй канал.
Результаты работы из второго канала пишутся в stdout.
*/

// Task33 is
func Task33(out chan<- int, in <-chan int) {
	if out == nil {
		panic("Out channel is nil!")
	}
	if in == nil {
		panic("In channel is nil!")
	}

	for v := range in {
		if v%2 == 0 {
			fmt.Println(v)
			out <- v
		}
	}
}
