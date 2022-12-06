package digitize_test

import (
	"reflect"
	"testing"
)

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

// testing
func Test_DigitizeMeCaptain(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		args args
		want []int
	}{
		{args{9078562341}, []int{9, 0, 7, 8, 5, 6, 2, 3, 4, 1}},
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
		})
	}
}

// the benchmarks
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
func Benchmark_copyDigitize(b *testing.B) { // winner winner tofu dinner.
	for i := 0; i < b.N; i++ {
		copyDigitize(9078562341)
	}
}
