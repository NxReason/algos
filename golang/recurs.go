package main

func FactorialRec(n int) int {
	if n == 0 {
		return 1
	}
	return n * FactorialRec(n-1)
}

func Factorial(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func FibRec(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	return FibRec(n-1) + FibRec(n-2)
}

func Fib(n int) int {
	fibs := make([]int, n+2)
	fibs[0] = 0
	fibs[1] = 1
	for i := 2; i <= n; i++ {
		fibs[i] = fibs[i-1] + fibs[i-2]
	}
	return fibs[n]
}