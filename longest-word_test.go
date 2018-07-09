package katas

import "testing"

func TestLongestWord(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		args args
		want int
	}{
		{args{s: "The quick brown fox jumped over the lazy dog"}, 6},
		{args{s: "May the force be with you"}, 5},
		{args{s: "What is the average airspeed velocity of an unladen swallow"}, 8},
		{args{s: "What if we try a super-long word such as otorhinolaryngology"}, 19},
	}
	for _, tt := range tests {
		if got := LongestWord(tt.args.s); got != tt.want {
			t.Errorf("LongestWord() = %v, want %v", got, tt.want)
		}
	}
}
