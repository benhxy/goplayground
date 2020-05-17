package goleetcode

// Learning: Writing unit tests.
func numJewelsInStones(J string, S string) int {
	// Add all jewels to set.
	jewels := make(map[rune]bool)
	for _, j := range []rune(J) {
		jewels[j] = true
	}

	// Increment count if stone is jewel.
	jewelCount := 0
	for _, c := range []rune(S) {
		if jewels[c] {
			jewelCount++
		}
	}

	return jewelCount
}
