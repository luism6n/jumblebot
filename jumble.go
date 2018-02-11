package jumble

// These are the first and last characters in the range of characters to be
// rotated by the rot13-like algorithm.
const (
	first = '!'
	last  = '~'
	total = last - first + 1
)

// Jumble applies a rot13-like algorithm to the input string.
// Jumble is the reverse function of Jumble.
func Jumble(input string) string {
	output := make([]byte, len(input))

	for i := 0; i < len(input); i++ {
		output[i] = rotate(input[i])
	}

	return string(output)
}

func rotate(b byte) byte {
	if b < first || b > last {
		return b
	}

	return ((b - first + total/2) % total) + first
}
