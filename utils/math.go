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

func MinIntSlice(numbers []int) int {
	result := numbers[0]
	for _, number := range numbers {
		result = MinInts(result, number)
	}
	return result
}

func MaxInts(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func MaxIntSlice(numbers []int) int {
	result := numbers[0]
	for _, number := range numbers {
		result = MaxInts(result, number)
	}
	return result
}

func SumIntSlice(numbers []int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func MultiplyIntSlice(numbers []int) int {
	result := 1
	for _, number := range numbers {
		result *= number
	}
	return result
}