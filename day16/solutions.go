package day16

import (
	"strconv"

	"github.com/WHYjun/advent-of-code-2021/utils"
)

var total int
var hex = map[string]string{
	"0": "0000",
	"1": "0001",
	"2": "0010",
	"3": "0011",
	"4": "0100",
	"5": "0101",
	"6": "0110",
	"7": "0111",
	"8": "1000",
	"9": "1001",
	"A": "1010",
	"B": "1011",
	"C": "1100",
	"D": "1101",
	"E": "1110",
	"F": "1111",
}

func Part1(input string) int {
	packets := ""
	for _, bit := range input {
		packets += hex[string(bit)]
	}
	parse(packets)
	return total
}

func Part2(input string) int {
	packets := ""
	for _, bit := range input {
		packets += hex[string(bit)]
	}
	_, val := parse(packets)
	return val
}

func parse(input string) (string, int) {
	version, err := strconv.ParseInt(input[:3], 2, 64)
	utils.PanicError(err)
	total += int(version)
	input = input[3:]

	typeId, err := strconv.ParseInt(input[:3], 2, 64)
	utils.PanicError(err)
	input = input[3:]
	if typeId == 4 {
		t := ""
		for {
			stopBit := input[:1]
			t += input[1:5]
			input = input[5:]
			if stopBit == "0" {
				break
			}
		}
		val, err := strconv.ParseInt(t, 2, 64)
		utils.PanicError(err)
		return input, int(val)
	} else {
		var val int
		lengthTypeId := input[:1]
		input = input[1:]
		subPacketValues := make([]int, 0)
		if lengthTypeId == "0" {
			length := input[:15]
			input = input[15:]
			subPacketLength, err := strconv.ParseInt(length, 2, 64)
			utils.PanicError(err)
			subPackets := input[:int(subPacketLength)]
			for {
				if subPackets == "" {
					break
				} else {
					subPackets, val = parse(subPackets)
					subPacketValues = append(subPacketValues, val)
				}
			}
			input = input[subPacketLength:]
		} else {
			length := input[:11]
			input = input[11:]
			qty, err := strconv.ParseInt(length, 2, 64)
			utils.PanicError(err)
			for i := 0; i < int(qty); i++ {
				input, val = parse(input)
				subPacketValues = append(subPacketValues, val)
			}
		}
		if typeId == 0 {
			return input, utils.SumIntSlice(subPacketValues)
		} else if typeId == 1 {
			return input, utils.MultiplyIntSlice(subPacketValues)
		} else if typeId == 2 {
			return input, utils.MinIntSlice(subPacketValues)
		} else if typeId == 3 {
			return input, utils.MaxIntSlice(subPacketValues)
		} else {
			flag := 0
			if typeId == 5 {
				if subPacketValues[0] > subPacketValues[1] {
					flag = 1
				}
				return input, flag
			} else if typeId == 6 {
				if subPacketValues[0] < subPacketValues[1] {
					flag = 1
				}
				return input, flag
			} else {
				if subPacketValues[0] == subPacketValues[1] {
					flag = 1
				}
				return input, flag
			}
		}
	}
}
