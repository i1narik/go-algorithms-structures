package algorithms

import "math"

func FibonacciRecursive(n uint64) uint64 {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return FibonacciRecursive(n-1) + FibonacciRecursive(n-2)
	}
}

func Fibonacci(n uint64) uint64 {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	var prev uint64 = 0
	var current uint64 = 1
	for i := uint64(2); i <= n; i++ {
		prev, current = current, current+prev
	}
	return current
}

func Fibonacci1(n uint64) uint64 {
	sqrt5 := math.Sqrt(float64(5))
	phi := (sqrt5 + 1) / 2
	return uint64(math.Pow(phi, float64(n)) / (sqrt5 + 0.5))
}
