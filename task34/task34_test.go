package task34

import "testing"

func TestTask34(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Unique ascii string",
			args: args{"qwertyuiop[] asdfghjkl;'zx cvb nm,./"},
			want: true,
		},
		{
			name: "NON Unique ascii string",
			args: args{"qwer tyuiop[]asq ghjkl;'zxcv bng,./"},
			want: false,
		},
		{
			name: "Unique unicode string",
			args: args{"Ͳ Θ π Ω Ć"},
			want: true,
		},
		{
			name: "NON Unique unicode string",
			args: args{"Ͳ Θ π Ω Ć   Ͳ Θ π Ω Ć"},
			want: false,
		},
		{
			name: "Empty string",
			args: args{""},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Task34(tt.args.s); got != tt.want {
				t.Errorf("Task34() = %v, want %v", got, tt.want)
			}
		})
	}
}
