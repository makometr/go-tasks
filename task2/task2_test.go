package task2

import (
	"reflect"
	"testing"
)

func TestTask2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name        string
		args        args
		wantResults []string
	}{
		{
			name:        "Example",
			args:        args{nums: []int{2, 4, 6, 8}},
			wantResults: []string{"4", "16", "36", "64"},
		},
		{
			name:        "Empty slice",
			args:        args{nums: []int{}},
			wantResults: nil,
		},
		{
			name:        "Nil slice",
			args:        args{nums: []int{}},
			wantResults: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResults := Task2(tt.args.nums...); !reflect.DeepEqual(gotResults, tt.wantResults) {
				t.Errorf("Task2() = %v, want %v", gotResults, tt.wantResults)
			}
		})
	}
}

func ExampleTestTask2() {
	Task2(2, 4, 6)
	// Output:
	// 4
	// 16
	// 36
	// All squares calculated!
}
