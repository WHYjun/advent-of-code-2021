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

// MustConvertLinesToIntegerList iterates a list of string to integer list.
// If it fails to convert string to integer, it panics an error message.
func MustConvertLinesToIntegerList(lines []string) []int {
	var numbers []int
	for _, line := range lines {
		number, err := strconv.Atoi(line)
		PanicError(err)
		numbers = append(numbers, number)
	}
	return numbers
}