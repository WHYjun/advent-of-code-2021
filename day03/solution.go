package day03

import (
	"math"
	"strconv"
	"strings"

	"github.com/WHYjun/advent-of-code-2021/utils"
)

type counter map[int]int

func Part1(input string) int {
	lines := strings.Split(input, "\n")
	gamma, epsilon := 0.0, 0.0
	bitLength := len(lines[0]) - 1
	for i := 0; i <= bitLength; i++ {
		bitCounter := countBit(lines, i)
		if bitCounter[0] >= bitCounter[1] {
			epsilon += math.Pow(2, float64(bitLength-i))
		} else {
			gamma += math.Pow(2, float64(bitLength-i))
		}
	}
	return int(epsilon * gamma)
}

func Part2(input string) int {
	lines := strings.Split(input, "\n")
	o2 := getOxygenRating(lines)
	co2 := getCarbonDioxideRating(lines)
	return o2 * co2
}

func countBit(lines []string, bitIdx int) counter {
	bitCounter := counter{
		0: 0,
		1: 0,
	}
	for _, line := range lines {
		bit, err := strconv.Atoi(string(line[bitIdx]))
		utils.PanicError(err)
		bitCounter[bit]++
	}
	return bitCounter
}

func getOxygenRating(lines []string) int {
	var o2 int64
	bitLength := len(lines[0]) - 1
	for i := 0; i <= bitLength; i++ {
		var tmp []string
		bitCounter := countBit(lines, i)
		for _, line := range lines {
			if bitCounter[0] <= bitCounter[1] {
				bit, err := strconv.Atoi(string(line[i]))
				utils.PanicError(err)
				if bit == 1 {
					tmp = append(tmp, line)
				}
			} else {
				bit, err := strconv.Atoi(string(line[i]))
				utils.PanicError(err)
				if bit == 0 {
					tmp = append(tmp, line)
				}
			}
		}
		lines = tmp
		if len(lines) == 1 {
			o2, _ = strconv.ParseInt(lines[0], 2, 64)
		}
	}
	return int(o2)
}

func getCarbonDioxideRating(lines []string) int {
	var co2 int64
	bitLength := len(lines[0]) - 1
	for i := 0; i <= bitLength; i++ {
		var tmp []string
		bitCounter := countBit(lines, i)
		for _, line := range lines {
			if bitCounter[0] > bitCounter[1] {
				bit, err := strconv.Atoi(string(line[i]))
				utils.PanicError(err)
				if bit == 1 {
					tmp = append(tmp, line)
				}
			} else {
				intBit, err := strconv.Atoi(string(line[i]))
				utils.PanicError(err)
				if intBit == 0 {
					tmp = append(tmp, line)
				}
			}
		}
		lines = tmp
		if len(lines) == 1 {
			co2, _ = strconv.ParseInt(lines[0], 2, 64)
		}
	}
	return int(co2)
}