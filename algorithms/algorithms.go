package algorithms

func FibonacciRecursive(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return FibonacciRecursive(n-1) + FibonacciRecursive(n-2)
	}
}

func Fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	prev := 0
	current := 1
	for i := 2; i <= n; i++ {
		prev, current = current, current+prev
	}
	return current
}
