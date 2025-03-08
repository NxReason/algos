package main

func Merge(source []int) []int {
	result := make([]int, len(source))
	for i := range source {
		result[i] = source[i]
	}

	MergeInner(result, 0, len(result))

	return result
}

func MergeInner(chunk []int, b, e int) {
	if b < e-1 {
		m := (b + e) / 2
		MergeInner(chunk, b, m)
		MergeInner(chunk, m, e)
		MergeParts(chunk, b, m, e)
	}
}

func MergeParts(chunk []int, b, m, e int) {
	n1 := m - b
	n2 := e - m
	left := make([]int, n1)
	right := make([]int, n2)
	for i := 0; i < n1; i++ {
		left[i] = chunk[b+i]
	}
	for i := 0; i < n2; i++ {
		right[i] = chunk[m+i]
	}

	leftIndex, rightIndex := 0, 0
	for i := b; i < e; i++ {
		if leftIndex == n1 {
			chunk[i] = right[rightIndex]
			rightIndex++
			continue
		}
		if rightIndex == n2 {
			chunk[i] = left[leftIndex]
			leftIndex++
			continue
		}

		if left[leftIndex] < right[rightIndex] {
			chunk[i] = left[leftIndex]
			leftIndex++
		} else {
			chunk[i] = right[rightIndex]
			rightIndex++
		}
	}
}