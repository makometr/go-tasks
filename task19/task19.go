package task19

import "fmt"

var justString string

func someFunc() {
	// Huge string
	v := "sssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssdddddddddddddddddssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssss"
	justString = v[:10]
	// Вместо этого:
	// justString = string([]byte(v[:10]))
	// Копируем нужный нам участок памяти, остальное освобождем
}

func main() {
	someFunc()
	fmt.Println(justString)
}

// Т.к. на созданную большую строку сохраняется ссылка в переменной уровня пакета,
// то память, выделенная для большой строки, gc не будет освбождена (если значение justString не будет далее меняться)
