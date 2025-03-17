package main

import "math"

func Factors(value int) []int {
	factors := make([]int, 0)
	test := 2
	for test < value {
		for value%test == 0 {
			factors = append(factors, test)
			value = value / test
		}
		test++
	}

	if value > 1 {
		factors = append(factors, value)
	}

	return factors
}

func FactorsV2(value int) []int {
	factors := make([]int, 0)

	for value%2 == 0 {
		factors = append(factors, 2)
		value /= 2
	}

	i := 3
	maxFactor := int(math.Sqrt(float64(value)))
	for i <= maxFactor {
		for value % i == 0 {
			factors = append(factors, i)
			value /= i
			maxFactor = int(math.Sqrt(float64(value)))
		}
		i += 2
	}

	if value > 1 { factors = append(factors, value) }

	return factors
}