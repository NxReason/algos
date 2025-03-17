package main

import (
	"fmt"
	"math/rand/v2"
)

func LinearCongruentialGen(seed int, a, b, m int) []int {
	result := make([]int, 50)
	result[0] = seed

	for i := 1; i < 50; i++ {
		result[i] = (a*result[i-1] + b) % m
	}

	return result
}

func ShowRandomizeList() {
	source := make([]int, 20)
	for i := range source {
		source[i] = i
	}
	fmt.Println(source)
	RandomizeList(source)
	fmt.Println(source)
}

func RandomizeList(arr []int) {
	length := len(arr)
	for i := 0; i < length; i++ {
		j := rand.IntN(length - i) + i
		arr[i], arr[j] = arr[j], arr[i]
	}
}