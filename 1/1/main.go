package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	elfs := [][]string{}
	currentElf := []string{}
	for _, str := range strings.Split(input, "\n") {
		if str == "" {
			elfs = append(elfs, currentElf)
			currentElf = []string{}
			continue
		}
		currentElf = append(currentElf, strings.TrimSpace(str))
	}
	if len(currentElf) != 0 {
		elfs = append(elfs, currentElf)
	}

	max := 0
	for _, elf := range elfs {
		sum := 0
		for _, foodCalory := range elf {
			c, err := strconv.Atoi(foodCalory)
			if err != nil {
				panic(err)
			}
			sum += c
		}
		if sum > max {
			max = sum
		}
	}
	fmt.Println(max)
}
