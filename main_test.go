package digitize_test

import (
	"reflect"
	"testing"
)

// testbenching different ways to digitize a number.

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
	size := howManyDigits(x)
	digit := 0
	digits := make([]int, size)
	for i := size; x >= 1; i-- {
		digit = x % 10
		x /= 10
		copy(digits[i-1:], []int{digit})
	}
	return digits
}

// range fill
func rangeDigitize(x int) []int {
	size := howManyDigits(x)
	digit := 0
	digits := make([]int, size)
	for i := range digits {
		if x < 1 {
			break
		}
		digit = x % 10
		x /= 10
		digits[len(digits)-i-1] = digit
	}
	return digits
}

// gets the number of digits in a number.
func howManyDigits(x int) int {
	digits := 0
	for x >= 1 {
		x /= 10
		digits++
	}
	return digits
}

// range withmanuall set size
func manualSetSizeRangeDigitize(x int) []int {
	digit := 0
	digits := make([]int, 10)
	for i := range digits {
		if x < 1 {
			break
		}
		digit = x % 10
		x /= 10
		digits[10-i-1] = digit
	}
	return digits
}

func Test_howManyDigits(t *testing.T) {
	if got := howManyDigits(1234567890); got != 10 {
		t.Errorf("howManyDigits() = %v, want %v", got, 10)
	} else {
		t.Logf("howManyDigits() = %v, want %v", got, 10)
	}
}

func Test_DigitizeMeCaptain(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		args args
		want []int
	}{
		{args{9078562341}, []int{9, 0, 7, 8, 5, 6, 2, 3, 4, 1}},
		{args{9078562344}, []int{9, 0, 7, 8, 5, 6, 2, 3, 4, 4}},
		{args{9078562345}, []int{9, 0, 7, 8, 5, 6, 2, 3, 4, 5}},
		{args{9078562346}, []int{9, 0, 7, 8, 5, 6, 2, 3, 4, 6}},
		{args{1456489840321549}, []int{1, 4, 5, 6, 4, 8, 9, 8, 4, 0, 3, 2, 1, 5, 4, 9}},
		{args{1234567890}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}},
		{args{1234567891}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1}},
		{args{1234567892}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 2}},
		{args{9223372036854775807}, []int{9, 2, 2, 3, 3, 7, 2, 0, 3, 6, 8, 5, 4, 7, 7, 5, 8, 0, 7}},
	}
	for _, tt := range tests {
		t.Run("digitizing", func(t *testing.T) {
			if got := digitize(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("digitze() = %v, want %v", got, tt.want)
			}
			if got := appendDigitze(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("appendDigitze() = %v, want %v", got, tt.want)
			}
			if got := copyDigitize(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("copyDigitze() = %v, want %v", got, tt.want)
			}
			if got := rangeDigitize(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("digitze() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_manualSetSizeRangeDigitize(t *testing.T) {
	if got := manualSetSizeRangeDigitize(1234567890); !reflect.DeepEqual(got, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}) {
		t.Errorf("manualSetSizeRangeDigitize() = %v, want %v", got, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})
	}
}

// the benchmarks
func Benchmark_howManyDigits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		howManyDigits(9078562341)
	}
}
func Benchmark_digitize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		digitize(9078562341)
	}
}
func Benchmark_appendDigitize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		appendDigitze(9078562341)
	}
}
func Benchmark_copyDigitize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		copyDigitize(9078562341)
	}
}
func Benchmark_rangeDigitize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rangeDigitize(9078562341)
	}
}

func Benchmark_manualSetSizeRangeDigitize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		manualSetSizeRangeDigitize(9078562341)
	}
}
