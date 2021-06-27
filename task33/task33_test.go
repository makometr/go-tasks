package task33

import (
	"testing"
)

func TestTask33Numbers(t *testing.T) {
	type args struct {
		out      chan int
		in       chan int
		numbers  []int
		expected []int
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "Base case",
			args: args{
				numbers:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				expected: []int{0, 2, 4, 6, 8, 10},
			},
		},
		{
			name: "No numbers",
			args: args{
				numbers:  []int{},
				expected: []int{},
			},
		},
		{
			name: "One number even",
			args: args{
				numbers:  []int{100},
				expected: []int{100},
			},
		},
		{
			name: "One number odd",
			args: args{
				numbers:  []int{5},
				expected: []int{},
			},
		},
	}

	for _, tt := range tests {
		tt.args.in = make(chan int, len(tt.args.numbers))
		tt.args.out = make(chan int)

		go func() {
			for _, n := range tt.args.numbers {
				tt.args.in <- n
			}
			close(tt.args.in)
		}()

		filtered := make([]int, 0)
		done := make(chan interface{})
		go func() {
			for v := range tt.args.out {
				filtered = append(filtered, v)
			}
			done <- struct{}{}
		}()

		Task33(tt.args.out, tt.args.in)
		close(tt.args.out)
		<-done

		if equal(tt.args.expected, filtered) != true {
			t.Fatal("Error in Task33 test", tt.name, ". Expected: ", tt.args.expected, "result:", filtered)
		}
	}
}

func TestTask33NilChannelIn(t *testing.T) {
	defer func() { recover() }()

	var in chan int
	out := make(chan int)
	Task33(out, in)

	t.Fatal("Task33 with nil in channel must panics")
}

func TestTask33NilChannelOut(t *testing.T) {
	defer func() { recover() }()

	var in chan int
	out := make(chan int)
	Task33(out, in)

	t.Fatal("Task33 with nil in channel must panics")
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
