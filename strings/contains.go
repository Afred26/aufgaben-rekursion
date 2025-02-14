package strings

// Contains prüft, ob der String s die Sequenz seq enthält.
func Contains(s, seq string) bool {
	if Length(s) < Length(seq) {
		return false
	}
	if seq == "" {
		return true
	}
	if s[:Length(seq)] == seq {
		return true
	}
	return Contains(s[1:], seq)
}
