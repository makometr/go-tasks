package task22_23

import (
	"fmt"
	"sort"
)

const marksNum = 5

type Student struct {
	name  string
	age   int
	marks [marksNum]int
}

// ByMarksSum implements sort.Interface for []Student based on sum of marks
type ByMarksSum []Student

// Len returns number of students in group
func (sg ByMarksSum) Len() int      { return len(sg) }
func (sg ByMarksSum) Swap(i, j int) { sg[i], sg[j] = sg[j], sg[i] }

// Less compares students by marks
func (sg ByMarksSum) Less(i, j int) bool {
	marksSumer := func(s Student) (sum int) {
		for i := 0; i < marksNum; i++ {
			sum += s.marks[i]
		}
		return
	}
	return marksSumer(sg[i]) < marksSumer(sg[j])
}

// Sort22 is
func Sort22(x interface{}) interface{} {
	switch v := x.(type) {
	case []int:
		sort.Ints(v)
		fmt.Println("Ints!	")
		return v
	case []Student:
		sort.Slice(v, func(i, j int) bool {
			return v[i].age < v[j].age
		})
		fmt.Println("Students!")
		return v
	case ByMarksSum:
		sort.Sort(sort.Reverse(v))
		fmt.Println("Group of students by marks!")
		return v
	default:
		fmt.Println("Unknown type")
	}
	return nil
}

func Search23(x []int, v int) int {
	index := sort.SearchInts(x, v)
	if index == len(x) || x[index] != v {
		return len(x)
	}
	return index
}
