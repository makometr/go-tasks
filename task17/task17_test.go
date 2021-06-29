package task17

func ExampleTask17_chanInt() {
	var v interface{} = make(chan int, 0)
	Task17(v)
	// Output:
	// chan
}

func ExampleTask17_chanSliceString() {
	var v interface{} = make(chan []string, 0)
	Task17(v)
	// Output:
	// chan
}

func ExampleTask17_int64() {
	var v interface{} = int64(9)
	Task17(v)
	// Output:
	// int
}
func ExampleTask17_int() {
	var v interface{} = 10
	Task17(v)
	// Output:
	// int
}
func ExampleTask17_string() {
	var v interface{} = "Moscow"
	Task17(v)
	// Output:
	// string
}
func ExampleTask17_bool() {
	var v interface{} = true
	Task17(v)
	// Output:
	// bool
}

func ExampleTask17_map() {
	var v interface{} = map[int]string{1: "sas", 3: "kek", 2: "Murmansk"}
	Task17(v)
	// Output:
	// undefined
}

func ExampleTask17_slice() {
	var v interface{} = []int{1, 2, 3}
	Task17(v)
	// Output:
	// undefined
}
