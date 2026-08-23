package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: "charizard MAYA   	",
			expected: []string{"charizard", "maya"},
		},
		{
			input: "IMONTHELINE",
			expected: []string{"imontheline"},
		},
		{
			input: "I can't wait to Bomb some Dodongos",
			expected: []string{"i", "can't", "wait", "to", "bomb", "some", "dodongos"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Mismatched Data")
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Mismatched Words")
			}
		}
	}
}