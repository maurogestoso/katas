package katas

import "testing"

func TestCelsiusToFahrenheit(t *testing.T) {
	cases := []struct {
		input, expected int
	}{
		{-30, -22},
		{-10, 14},
		{0, 32},
		{20, 68},
		{30, 86},
	}

	for _, c := range cases {
		result := CelsiusToFahrenheit(c.input)
		if result != c.expected {
			t.Errorf("CelsiusToFahrenheit(%d) == %d, expected %d", c.input, result, c.expected)
		}
	}
}
