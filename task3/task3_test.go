package task3

import "testing"

func TestTask3(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "Example",
			args:    args{nums: []int{2, 4, 6, 8, 10}},
			want:    2*2 + 4*4 + 6*6 + 8*8 + 10*10,
			wantErr: false,
		},
		{
			name:    "Empty slice data",
			args:    args{nums: []int{}},
			want:    0,
			wantErr: false,
		},
		{
			name:    "Nil slice data",
			args:    args{nums: nil},
			want:    0,
			wantErr: false,
		},
		{
			name:    "With 'error' in data",
			args:    args{nums: []int{1, 3, 4, 0, 5, 6, 7, 8, 9}},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Task3(tt.args.nums)
			if (err != nil) != tt.wantErr {
				t.Errorf("Task3() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Task3() = %v, want %v", got, tt.want)
			}
		})
	}
}
