package task18

import "fmt"

func someAction(v []int8, b int8) {
	fmt.Println("Before:")
	fmt.Printf("	Address of slice in func         : %p\n", &v)
	fmt.Printf("	Address of slice's buffer in func: %p\n\n", &v[0])

	v[0] = 100
	v = append(v, b)

	fmt.Println("After:")
	fmt.Printf("	Address of slice in func         : %p\n", &v)
	fmt.Printf("	Address of slice's buffer in func: %p\n\n", &v[0])
}

func main() {
	var a = []int8{1, 2, 3, 4, 5}
	fmt.Printf("Address of slice in main         : %p\n", &a)
	fmt.Printf("Address of slice's buffer in main: %p\n\n", &a[0])

	someAction(a, 6)
	fmt.Printf("Address of slice in main         : %p\n", &a)
	fmt.Printf("Address of slice's buffer in main: %p\n\n", &a[0])
	fmt.Println(a)
	// [100 2 3 4 5]
}

// Ответ: первый элемент изменится, но новый не добавится.
// Снова из-за копирования по значению. А слайс - ссылочный тип.
