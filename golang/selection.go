package main

func Selection(source []int) []int {
	result := make([]int, len(source))
	for i := range source {
		result[i] = source[i]
	}

	for i := 0; i < len(result)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(result); j++ {
			if result[j] < result[minIndex] {
				minIndex = j
			}
		}
		result[i], result[minIndex] = result[minIndex], result[i]
	}

	return result
}