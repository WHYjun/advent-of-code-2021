package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/WHYjun/advent-of-code-2021/utils"
)

func main() {
	day := mustParseDay()
	fmt.Printf("Running day %02d\n", day)

	switch day {
	default:
		utils.PanicError(fmt.Errorf("no such day: %d", day))
	}
}

func mustParseDay() int {
	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		utils.PanicError(fmt.Errorf("should have a vaild integer day (1-25) - %s", err.Error()))
	}
	return day
}
