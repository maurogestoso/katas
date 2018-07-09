package katas

import "testing"

func TestConfirmEnding(t *testing.T) {
	type args struct {
		str string
		end string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Returns true if the passed string doesn't end with the passed end",
			args{str: "Hello", end: "lo"},
			true,
		},
		{
			"Returns false if the passed string doesn't end with the passed end",
			args{str: "Hello", end: "la"},
			false,
		},
		{
			"Returns false if the passed end is an empty string",
			args{str: "Hello", end: ""},
			false,
		},
		{
			"Returns false if passed an end longer than the passed string",
			args{str: "abc", end: "bcdef"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConfirmEnding(tt.args.str, tt.args.end); got != tt.want {
				t.Errorf("ConfirmEnding() = %v, want %v", got, tt.want)
			}
		})
	}
}
