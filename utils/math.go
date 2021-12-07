package utils

func AbsInts(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}

func MinInts(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func MaxInts(x, y int) int {
	if x > y {
		return x
	}
	return y
}