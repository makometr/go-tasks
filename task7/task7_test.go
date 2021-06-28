package task7

import (
	"reflect"
	"testing"
)

func TestTask7(t *testing.T) {

	createModelMap := func(chs, n int) *SafeMap {
		modelMap := NewSafeMap()
		for i := 0; i < n; i++ {
			modelMap.Write(i, chs)
		}
		return modelMap
	}

	type args struct {
		chs int
		n   int
	}
	tests := []struct {
		name string
		args args
		want *SafeMap
	}{
		{
			name: "100 channels do 1000 incs",
			args: args{chs: 100, n: 1000},
			want: createModelMap(100, 1000),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Task7(tt.args.chs, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Task7() = %v, want %v", got, tt.want)
			}
		})
	}
}
