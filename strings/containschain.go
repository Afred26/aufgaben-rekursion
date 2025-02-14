package strings

// ContainsChain liefert true, falls s eine Kette von count aufeinanderfolgenden
// Vorkommen von symbol enthält.
func ContainsChain(s, symbol string, count int) bool {
	if Length(s) < Length(symbol)*count {
		return false
	}
	if symbol == "" {
		return true
	}
	if count == 0 {
		return true
	}
	if s[:(Length(symbol)*count)] == Chain(symbol, count) {
		return true
	}
	return ContainsChain(s[1:], symbol, count)
}
