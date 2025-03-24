package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	fmt.Println(Fib(0))
	fmt.Println(Fib(1))
	fmt.Println(Fib(2))
	fmt.Println(Fib(3))
	fmt.Println(Fib(4))
	fmt.Println(Fib(5))
	fmt.Println(Fib(6))
	fmt.Println(Fib(7))
	fmt.Println(Fib(8))
	fmt.Println(Fib(9))
}

func MakeSource(length int) []int {
	source := make([]int, length)
	for i := range source {
		source[i] = rand.IntN(51)
	}
	return source
}

func PrintList(list []int, msg string) {
	fmt.Println(msg)
	fmt.Println(list)
}

func PrintSortingAlgorithm(algoName string, sort func([]int) []int) {
	fmt.Println(algoName)

	source := MakeSource(21)
	PrintList(source, "original list:")

	result := sort(source)
	PrintList(result, "sorted list:")

	fmt.Println()
}