package katas

import "testing"

func TestReverseString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		args args
		want string
	}{
		{args{s: "Hello, world"}, "dlrow ,olleH"},
		{args{s: "Hello, 世界"}, "界世 ,olleH"},
		{args{s: ""}, ""},
	}
	for _, tt := range tests {
		if got := ReverseString(tt.args.s); got != tt.want {
			t.Errorf("ReverseString() = %v, want %v", got, tt.want)
		}
	}
}
