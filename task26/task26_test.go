package task26

import "testing"

func TestTask26(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example",
			args: args{s: "bròwn"},
			want: "nwòrb",
		},
		{
			name: "More unicode!",
			args: args{s: "Hel 世界"},
			want: "界世 leH",
		},
		{
			name: "Empty string",
			args: args{s: ""},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Task26(tt.args.s); got != tt.want {
				t.Errorf("Task26() = %v, want %v", got, tt.want)
			}
		})
	}
}
