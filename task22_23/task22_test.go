package task22_23

import (
	"reflect"
	"testing"
)

func TestSort22(t *testing.T) {
	type args struct {
		x interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "Ints",
			args: args{x: []int{1, 5, 3, 4, 2, 6}},
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "Students",
			args: args{x: []Student{
				{name: "Alex", age: 18},
				{name: "Oleg", age: 20},
				{name: "Victor", age: 17}},
			},
			want: []Student{
				{name: "Victor", age: 17},
				{name: "Alex", age: 18},
				{name: "Oleg", age: 20},
			},
		},
		{
			name: "Students by marks",
			args: args{x: ByMarksSum([]Student{
				{name: "Oleg", age: 20, marks: [5]int{4, 4, 4, 4, 4}},
				{name: "Alex", age: 18, marks: [5]int{3, 3, 3, 3, 3}},
				{name: "Victor", age: 17, marks: [5]int{5, 5, 5, 5, 5}},
			}),
			},
			want: ByMarksSum([]Student{
				{name: "Victor", age: 17, marks: [5]int{5, 5, 5, 5, 5}},
				{name: "Oleg", age: 20, marks: [5]int{4, 4, 4, 4, 4}},
				{name: "Alex", age: 18, marks: [5]int{3, 3, 3, 3, 3}},
			}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sort22(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Task22Int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearch23(t *testing.T) {
	type args struct {
		x []int
		v int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Ints non-sorted, Exist",
			args: args{x: []int{1, 3, 4, 5, 10, 20, 50, 25, 41, 3}, v: 10},
			want: 4,
		},
		{
			name: "Ints non-sorted, no-Exist-1",
			args: args{x: []int{1, 3, 4, 5, 10, 20, 50, 25, 41, 3}, v: 45},
			want: 10,
		},
		{
			name: "Ints sorted, Exist",
			args: args{x: []int{1, 3, 3, 4, 5, 10, 20, 25, 41, 50}, v: 10},
			want: 5,
		},
		{
			name: "Ints sorted, no-Exist-1",
			args: args{x: []int{1, 3, 3, 4, 5, 10, 20, 25, 41, 50}, v: 45},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Search23(tt.args.x, tt.args.v); got != tt.want {
				t.Errorf("Search23() = %v, want %v", got, tt.want)
			}
		})
	}
}
