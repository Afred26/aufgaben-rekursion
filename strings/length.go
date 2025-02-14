package strings

func Length(s string) int {

	if s == "" {
		return 0
	}
	return 1 + Length(s[1:])
}
