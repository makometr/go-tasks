package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 3, 4, 5, 10, 20, 50, 25, 41, 3}
	sort.Ints(a)
	fmt.Println(a)
	index := sort.SearchInts(sort.IntSlice(a), 50)

	fmt.Println(a[index])
}
