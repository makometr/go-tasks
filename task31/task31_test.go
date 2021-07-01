package task31

import (
	"testing"
)

func TestGetDistance(t *testing.T) {
	type args struct {
		a Point2D
		b Point2D
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Example",
			args: args{a: Point2D{1.0, 2.0}, b: Point2D{4.0, 6.0}},
			want: 5.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDistance(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("GetDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
