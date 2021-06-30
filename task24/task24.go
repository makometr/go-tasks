package task24

import "fmt"

func Task24() {
	slice1 := make([]int, 100)     // слайс на 100 нулей
	slice2 := make([]int, 10, 100) // слайс на 10 нулей, но размер (капасити) такой же, что и у первого
	fmt.Println(slice1, slice2)
}
