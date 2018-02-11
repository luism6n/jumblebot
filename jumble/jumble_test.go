package jumble

import "testing"

const (
	asciiPrintableLowerHalf = "!\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~"
	asciiPrintableUpperHalf = "PQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~!\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNO"
	asciiNonPrintable       = "\x00\x01\x02\x03\x04\x05\x06\x07\x08\x09\x0a\x0b\x0c\x0d\x0e\x0f\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1c\x1d\x1e\x1f\x20\x7f"
	weirdUTF8               = "€®ŧ←↓→øþæßðđŋħł»©“”µẃéŕýúíóṕáśǵḱĺźçǘńḿẁèỳùìòàǜǹẽỹũĩõãṽñ"
)

func TestJumble(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{asciiPrintableLowerHalf, asciiPrintableUpperHalf},
		{asciiPrintableUpperHalf, asciiPrintableLowerHalf},
		{asciiNonPrintable, asciiNonPrintable},
		{weirdUTF8, weirdUTF8},
	}

	for _, c := range testCases {
		output := Jumble(c.input)
		if output != c.expected {
			t.Fatalf("Jumble(%q) should be %q, but was %q", c.input, c.expected, output)
		}
	}
}
