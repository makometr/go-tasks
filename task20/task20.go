package task20

import "fmt"

func main() {
	slice := []string{"a", "a"}

	func(slice []string) {
		slice = append(slice, "a")
		slice[0] = "b"
		slice[1] = "b"
		fmt.Print(slice)
		// [b b a]
	}(slice)
	fmt.Print(slice)
	// [a a]
}

// Пердеаем по значению ссылочный тип И меняем буфер доя копии слайса перед изменением значений в буфере исходном.
// Если бы append не было, во второй раз вывелось бы [b b]
