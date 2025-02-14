package calc

// BinomialCoefficient erwartet zwei Zahlen n und k und liefert den
// Binomialkoeffizienten "n über k".
func BinomialCoefficientFactorial(n, k int) int {

	return Factorial(n) / (Factorial(k) * Factorial(n-k))
}

func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}

// BinomialCoefficient erwartet zwei Zahlen n und k und liefert den
// Binomialkoeffizienten "n über k".
func BinomialCoefficientPascal(n, k int) int {
	if k > n {
		return 0
	}
	if n == 0 || k == 0 {
		return 1
	}
	return BinomialCoefficientPascal(n-1, k) + BinomialCoefficientPascal(n-1, k-1)
}
