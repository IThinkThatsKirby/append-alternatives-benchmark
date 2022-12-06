package digitize_test

// easy readable append
func digitize(x int) []int {
	digit := 0
	digits := []int{}
	for i := 0; x >= 1; i++ {
		digit = x % 10
		x /= 10
		digits = append([]int{digit}, digits...)
	}
	return digits
}

// make append
func appendDigitze(x int) []int {
	digit := 0
	digits := make([]int, 0)
	for x >= 1 {
		digit = x % 10
		x /= 10
		digits = append([]int{digit}, digits...)
	}
	return digits
}

// make copy
func copyDigitize(x int) []int {
	digit := 0
	digits := make([]int, 10)
	for i := 10; x >= 1; i-- {
		digit = x % 10
		x /= 10
		copy(digits[i-1:], []int{digit})
	}
	return digits
}
