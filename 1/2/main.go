package main

import (
	"fmt"
	"sort"
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

	all := []int{}
	for _, elf := range elfs {
		sum := 0
		for _, foodCalory := range elf {
			c, err := strconv.Atoi(foodCalory)
			if err != nil {
				panic(err)
			}
			sum += c
		}
		all = append(all, sum)
	}
	sort.Ints(all)
	fmt.Println(all[len(all)-3:])
	sum := 0
	for _, x := range all[len(all)-3:] {
		sum += x
	}
	fmt.Println(sum)
}
