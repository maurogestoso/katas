package katas

import "testing"

func TestReverse(t *testing.T) {
	cases := []struct {
		input, expected string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}
	for _, c := range cases {
		result := ReverseString(c.input)
		if result != c.expected {
			t.Errorf("ReverseString(%q) == %q, expected %q", c.input, result, c.expected)
		}
	}
}
