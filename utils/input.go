package utils

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
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
