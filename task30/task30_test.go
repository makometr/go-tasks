package task30

import (
	"reflect"
	"testing"
)

func TestTask30(t *testing.T) {
	type args struct {
		slice []int
		i     int
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{
			name:    "In the middle",
			args:    args{slice: []int{1, 2, 3, 4, 5}, i: 1},
			want:    []int{1, 3, 4, 5},
			wantErr: false,
		},
		{
			name:    "First",
			args:    args{slice: []int{1, 2, 3, 4, 5}, i: 0},
			want:    []int{2, 3, 4, 5},
			wantErr: false,
		},
		{
			name:    "Last",
			args:    args{slice: []int{1, 2, 3, 4, 5}, i: 4},
			want:    []int{1, 2, 3, 4},
			wantErr: false,
		},
		{
			name:    "Out of index",
			args:    args{slice: []int{1, 2, 3, 4, 5}, i: 100},
			want:    []int{1, 2, 3, 4, 5},
			wantErr: true,
		},
		{
			name:    "Empty",
			args:    args{slice: []int{}, i: 10},
			want:    []int{},
			wantErr: true,
		},
		{
			name:    "Nil",
			args:    args{slice: nil, i: 10},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Task30(tt.args.slice, tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("Task30() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Task30() = %v, want %v", got, tt.want)
			}
		})
	}
}
