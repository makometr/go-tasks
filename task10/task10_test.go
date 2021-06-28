package task10

import (
	"reflect"
	"testing"
)

func TestTask10(t *testing.T) {
	type args struct {
		ts []int
	}
	tests := []struct {
		name string
		args args
		want map[int][]int
	}{
		{
			name: "Positive values",
			args: args{
				[]int{10, 11, 15, 6, 22},
			},
			want: map[int][]int{10: {10, 11, 15}, 0: {6}, 20: {22}},
		},
		{
			name: "Negative values",
			args: args{
				[]int{-10, -11, -15, -6, -22},
			},
			want: map[int][]int{-10: {-10, -11, -15}, -20: {-22}, 0: {-6}},
		},
		{
			name: "question?",
			args: args{
				[]int{1, 2, 3, -1, -2, -3},
			},
			want: map[int][]int{0: {1, 2, 3, -1, -2, -3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Task10(tt.args.ts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Task10() = %v, want %v", got, tt.want)
			}
		})
	}
}
