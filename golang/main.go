package main

import (
	"algo/list"
	"fmt"
	"math/rand/v2"
)

func main() {
	// c1 := list.IntegerCell { Value: 42, Next: nil }
	// c2 := list.IntegerCell { Value: 7, Next: &c1 }
	// c3 := list.IntegerCell { Value: 24, Next: &c2 }
	// top := &c3
	// for top != nil {
	// 	fmt.Println(top.Value)
	// 	top = top.Next
	// }
	intList := list.IntegerList { }
	intList.Append(42)
	intList.Append(7)
	intList.Append(24)
	intList.Prepend(5)
	intList.Append(1)
	fmt.Println(intList.ToString())
	fmt.Println(intList.Contains(24))
	fmt.Println(intList.Contains(253254))

	intList.InsertAt(15, 1)
	intList.InsertAt(-5, 0)
	intList.InsertAt(-2, 7)
	fmt.Println(intList.ToString())
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