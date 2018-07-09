package katas

import "testing"

func TestCelsiusToFahrenheit(t *testing.T) {
	type args struct {
		celsius int
	}
	tests := []struct {
		args args
		want int
	}{
		{args{celsius: -30}, -22},
		{args{celsius: -10}, 14},
		{args{celsius: 0}, 32},
		{args{celsius: 20}, 68},
		{args{celsius: 30}, 86},
	}
	for _, tt := range tests {
		if got := CelsiusToFahrenheit(tt.args.celsius); got != tt.want {
			t.Errorf("CelsiusToFahrenheit() = %v, want %v", got, tt.want)
		}
	}
}
