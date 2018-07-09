package katas

import "testing"

func TestFactorialise(t *testing.T) {
	cases := []struct {
		input, expected int
	}{
		{5, 120},
		{10, 3628800},
		{20, 2432902008176640000},
	}
	for _, c := range cases {
		result := Factorialise(c.input)
		if result != c.expected {
			t.Errorf("Factorialise(%d) == %d. Expected %d", c.input, result, c.expected)
		}
	}
}
