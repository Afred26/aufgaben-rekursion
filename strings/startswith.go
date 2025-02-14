package strings

// StartsWith liefert true, falls der string s mit der Sequenz seq beginnt.
func StartsWith(s, seq string) bool {
	if seq == "" {
		return true
	}
	if s == "" {
		return false
	}
	if s[:Length(seq)] == seq {
		return true
	}
	return false
}
