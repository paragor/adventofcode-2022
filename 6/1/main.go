package main

import (
	"fmt"
	"strings"
)

func main() {
	buff := NewBuffer(4, []string{})
	for _, ch := range input {
		result := buff.add(string(ch))
		if result {
			fmt.Println(buff.allCount)
			return
		}
	}
	fmt.Println("hui")
}

type Buffer struct {
	max      int
	arr      []string
	allCount int
}

func NewBuffer(max int, arr []string) *Buffer {
	return &Buffer{max: max, arr: arr}
}

func (b *Buffer) add(s string) bool {
	b.allCount += 1
	b.arr = append(b.arr, s)
	if len(b.arr) > b.max {
		b.arr = b.arr[1:]
	}
	if len(b.arr) < b.max {
		return false
	}
	m := map[string]int{}
	for _, s := range b.arr {
		if m[s] > 0 {
			return false
		}
		m[s] = 1
	}
	return true
}

var input = strings.TrimSpace(`
mjqjpqmgbljsphdztnvjfqwrcgsmlb
`)
