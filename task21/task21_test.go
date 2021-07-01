package task21

func ExampleSync() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	Sync(data)
	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
	// 6
	// 7
	// 8
	// 9
	// 10
}

func ExampleAsync_workersLessSize() {
	Async(5, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	// Unordered output:
	// 0
	// 1
	// 2
	// 3
	// 4
	// 5
	// 6
	// 7
	// 8
	// 9
	// 10
}

func ExampleAsync_workersMoreSize() {
	Async(10, []int{0, 1, 2, 3, 4, 5})
	// Unordered output:
	// 0
	// 1
	// 2
	// 3
	// 4
	// 5
}
