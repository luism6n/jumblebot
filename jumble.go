package jumble

// Jumble applies a rot13-like algorithm to the input string.
// Jumble is the reverse function of Jumble.
func Jumble(input string) string {
	output := make([]byte, len(input))

	for i := 0; i < len(input); i++ {
		output[i] = input[i] + 13
	}

	return string(output)
}
