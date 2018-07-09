package katas

// Factorialise returns the factorial of the passed int
func Factorialise(n int) int {
	factorial := 1
	for i := 1; i <= n; i++ {
		factorial *= i
	}
	return factorial
}
