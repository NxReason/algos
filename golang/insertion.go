package main

func Insertion(source []int) []int {
	result := make([]int, len(source))
	for i := range source {
		result[i] = source[i]
	}

	for i := 1; i < len(result); i++ {
		key := result[i]
		j := i - 1
		for j >= 0 && result[j] > key {
			result[j+1] = result[j]
			j = j - 1
		}
		result[j+1] = key
	}

	return result
}
