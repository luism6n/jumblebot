package jumble

import "testing"

func TestJumble(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"!", "P"},
	}

	for _, c := range testCases {
		output := Jumble(c.input)
		if output != c.expected {
			t.Fatalf("Jumble(%q) should be %q, but was %q", c.input, c.expected, output)
		}
	}
}
