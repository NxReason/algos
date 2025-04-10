package main

import (
	"algo/tree"
	"fmt"
	"math/rand/v2"
)

func main() {
	tree.TreeShowcase()
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