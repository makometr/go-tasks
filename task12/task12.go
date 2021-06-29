package task12

import "fmt"

func update(p *int) {
	b := 2
	p = &b
}

func Task12() {
	var (
		a = 1
		p = &a
	)
	fmt.Println(*p) // 1
	update(p)
	fmt.Println(*p) // 1
}

// Так происходит из-за передачи указателя по значению (копирование).
// - Передавать указатель на указатель
// - Изменять значение с разыменовыванием
