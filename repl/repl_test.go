package repl

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello world  ",
			expected: []string{"hello", "world"},
		}, {
			input:    "  hey alvin ",
			expected: []string{"hey", "alvin"},
		}, {
			input:    "  my ghost are in my room  ",
			expected: []string{"my", "ghost", "are", "in", "my", "room"},
		}, {
			input:    "  hey alvin     how  ",
			expected: []string{"hey", "alvin", "how"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		expectedLength := len(c.expected)
		gotLength := len(actual)
		if gotLength != expectedLength {
			t.Errorf("got: %v, expected: %v", gotLength, expectedLength)
		}
		for i := range actual {
			gotWord := actual[i]
			expectedWord := c.expected[i]
			if gotWord != expectedWord {
				t.Errorf("got: %s, expected: %s", gotWord, expectedWord)
			}
		}

	}
}
