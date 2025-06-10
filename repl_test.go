package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {

	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
}

	for _, c := range cases {
		actual := cleanInput(c.input)
		expectedLength := len(c.expected)
		actualLength := len(actual)
		if expectedLength != actualLength {
			t.Errorf("expected length does not match actual length")
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("expected: %s, but got %s", expectedWord, word)
			}
		}
	}
}

