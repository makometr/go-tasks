package task3

import "testing"

func TestTask3(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example",
			args: args{nums: []int{2, 4, 6, 8, 10}},
			want: 2*2 + 4*4 + 6*6 + 8*8 + 10*10,
		},
		{
			name: "Empty slice data",
			args: args{nums: []int{}},
			want: 0,
		},
		{
			name: "Nil slice data",
			args: args{nums: nil},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Task3(tt.args.nums); got != tt.want {
				t.Errorf("Task3() = %v, want %v", got, tt.want)
			}
		})
	}

}
