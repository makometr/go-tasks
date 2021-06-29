package task14

import (
	"fmt"
	"reflect"
	"testing"
)

func TestStringSet_Add(t *testing.T) {
	type args struct {
		new string
	}
	tests := []struct {
		name    string
		ss      *StringSet
		args    args
		want    bool
		wantSet *StringSet
	}{
		{
			name:    "Add new elem to empty",
			ss:      newStringSet(),
			args:    args{new: "Strike"},
			want:    false,
			wantSet: newStringSet("Strike"),
		},
		{
			name:    "Add new elem to non-empty",
			ss:      newStringSet("Tree", "Klin"),
			args:    args{new: "Strike"},
			want:    false,
			wantSet: newStringSet("Strike", "Tree", "Klin"),
		},
		{
			name:    "Add new elem to non-empty",
			ss:      newStringSet("Tree", "Klin"),
			args:    args{new: "Tree"},
			want:    true,
			wantSet: newStringSet("Tree", "Klin"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ss.Add(tt.args.new); got != tt.want {
				t.Errorf("StringSet.Add() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(tt.ss, tt.wantSet) {
				t.Errorf("StringSet.Add() = %v, want %v", tt.ss, tt.wantSet)
			}
		})
	}
}

func ExampleStringSet_Add() {
	s := newStringSet("cat", "cat", "dog", "cat", "tree")
	for k := range *s {
		fmt.Println(k)
	}
	// Unordered output:
	// cat
	// dog
	// tree
}
