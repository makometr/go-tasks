package task16

import "fmt"

func main() {
	n := 0
	if true {
		n := 1
		n++
	}
	fmt.Println(n) // 0
}

// Т.к. := создает новую переменную в текущей области видимости, которая затеняет предыдущую
