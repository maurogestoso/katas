package katas

import "testing"

func TestLongestWord(t *testing.T) {
	cases := []struct {
		input    string
		expected int
	}{
		{"The quick brown fox jumped over the lazy dog", 6},
		{"May the force be with you", 5},
		{"What is the average airspeed velocity of an unladen swallow", 8},
		{"What if we try a super-long word such as otorhinolaryngology", 19},
	}
	for _, c := range cases {
		result := LongestWord(c.input)
		if result != c.expected {
			t.Errorf("LongestWord(%q) == %d. Expected %d", c.input, result, c.expected)
		}
	}
}
