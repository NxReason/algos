package main

func GCD(x, y int) int {
	if y == 0 {
		return x
	}
	if x == 0 {
		return y
	}

	for y != 0 {
		rem := x % y
		x = y
		y = rem
	}

	return x
}