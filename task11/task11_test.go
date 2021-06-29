package task11

import (
	"reflect"
	"testing"
)

func TestTask11(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "One numbe in result",
			args: args{a: []int{1, 2, 3}, b: []int{3, 4, 5}},
			want: []int{3},
		},
		{
			name: "Zero and many numbers",
			args: args{a: []int{0, 0, 1, 1}, b: []int{0, 2, 3}},
			want: []int{0},
		},
		{
			name: "Empty slice",
			args: args{a: []int{0, 1, 3, 3, 4}, b: []int{}},
			want: []int{},
		},
		{
			name: "Nil slice",
			args: args{a: nil, b: nil},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Task11(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Task11() = %v, want %v", got, tt.want)
			}
		})
	}
}
