package katas

import "testing"

func TestFactorialise(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		args args
		want int
	}{
		{args{n: 5}, 120},
		{args{n: 10}, 3628800},
		{args{n: 20}, 2432902008176640000},
	}
	for _, tt := range tests {
		if got := Factorialise(tt.args.n); got != tt.want {
			t.Errorf("Factorialise() = %v, want %v", got, tt.want)
		}
	}
}
