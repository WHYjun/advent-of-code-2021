package utils

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// MustReadFile reads input file provided by Advent of Code website.
// If it fails to read file, it panics an error message.
func MustReadFile(day int) string {
	filename := fmt.Sprintf("day%02d/input.txt", day)
	file, err := os.Open(filename)
	PanicError(err)
	defer file.Close()

	reader := bufio.NewReader(file)
	contents, err := ioutil.ReadAll(reader)
	PanicError(err)

	return strings.TrimSuffix(string(contents), "\n")
}

// MustConvertLinesToIntegerSlice iterates a list of string to integer slice.
// If it fails to convert string to integer, it panics an error message.
func MustConvertLinesToIntegerSlice(lines []string) []int {
	var numbers []int
	for _, line := range lines {
		number, err := strconv.Atoi(line)
		PanicError(err)
		numbers = append(numbers, number)
	}
	return numbers
}

func ConvertInputToIntSlices(input string) [][]int {
	lines := strings.Split(input, "\n")
	var intSlices [][]int
	for _, line := range lines {
		var intSlice []int
		for _, char := range line {
			val, _ := strconv.Atoi(string(char))
			intSlice = append(intSlice, val)
		}
		intSlices = append(intSlices, intSlice)
	}
	return intSlices
}