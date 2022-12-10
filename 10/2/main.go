package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	wide   = 40
	height = 6
)

func main() {
	queue := &Queue[*Instruction]{}
	for _, str := range strings.Split(inputA, "\n") {
		str = strings.TrimSpace(str)
		instruction := &Instruction{
			//noop by default
			Add:      0,
			Duration: 1,
		}
		if strings.HasPrefix(str, "addx") {
			add, err := strconv.Atoi(strings.TrimPrefix(str, "addx "))
			if err != nil {
				panic(err)
			}

			instruction.Add = add
			instruction.Duration = 2
		}
		queue.Append(instruction)
	}
	cpu := &CPU{
		ticks:  0,
		memory: 1,
	}
	instr := queue.Next()
	monitor := NewMonitor(1)

	for {
		monitor.Draw(cpu.ticks, cpu.memory)

		instr.Duration--
		cpu.ticks++

		if instr.Duration == 0 {
			cpu.memory += instr.Add
			if queue.Len() == 0 {
				break
			}
			instr = queue.Next()
		}
	}
	for _, row := range monitor.Canvas {
		for _, x := range row {
			fmt.Printf(x)
		}
		fmt.Println()
	}
}

type CPU struct {
	ticks  int
	memory int
}

type Monitor struct {
	Canvas    [][]string
	SpritePos int
}

func NewMonitor(spritePos int) *Monitor {
	rows := make([][]string, height)
	for i := range rows {
		rows[i] = make([]string, wide)
	}
	return &Monitor{SpritePos: spritePos, Canvas: rows}
}

func (s *Monitor) Draw(tick int, spritePos int) {
	vert := (tick) / 40
	horiz := (tick) % 40
	sym := "."
	if horiz >= spritePos-1 && horiz <= spritePos+1 {
		sym = "#"
	}

	s.Canvas[vert][horiz] = sym
}

type Queue[T any] struct {
	items []T
}

func (q *Queue[T]) Len() int {
	return len(q.items)
}

func (q *Queue[T]) Append(instruction T) {
	q.items = append(q.items, instruction)
}
func (q *Queue[T]) Next() T {
	var ret T
	ret, q.items = q.items[0], q.items[1:]
	return ret
}

type Instruction struct {
	Add      int
	Duration int
}

var input = strings.TrimSpace(`
addx 15
addx -11
addx 6
addx -3
addx 5
addx -1
addx -8
addx 13
addx 4
noop
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx -35
addx 1
addx 24
addx -19
addx 1
addx 16
addx -11
noop
noop
addx 21
addx -15
noop
noop
addx -3
addx 9
addx 1
addx -3
addx 8
addx 1
addx 5
noop
noop
noop
noop
noop
addx -36
noop
addx 1
addx 7
noop
noop
noop
addx 2
addx 6
noop
noop
noop
noop
noop
addx 1
noop
noop
addx 7
addx 1
noop
addx -13
addx 13
addx 7
noop
addx 1
addx -33
noop
noop
noop
addx 2
noop
noop
noop
addx 8
noop
addx -1
addx 2
addx 1
noop
addx 17
addx -9
addx 1
addx 1
addx -3
addx 11
noop
noop
addx 1
noop
addx 1
noop
noop
addx -13
addx -19
addx 1
addx 3
addx 26
addx -30
addx 12
addx -1
addx 3
addx 1
noop
noop
noop
addx -9
addx 18
addx 1
addx 2
noop
noop
addx 9
noop
noop
noop
addx -1
addx 2
addx -37
addx 1
addx 3
noop
addx 15
addx -21
addx 22
addx -6
addx 1
noop
addx 2
addx 1
noop
addx -10
noop
noop
addx 20
addx 1
addx 2
addx 2
addx -6
addx -11
noop
noop
noop
`)
var inputA = strings.TrimSpace(`
noop
addx 7
addx -1
addx -1
addx 5
noop
noop
addx 1
addx 3
addx 2
noop
addx 2
addx 5
addx 2
addx 10
addx -9
addx 4
noop
noop
noop
addx 3
addx 5
addx -40
addx 26
addx -23
addx 2
addx 5
addx 26
addx -35
addx 12
addx 2
addx 17
addx -10
addx 3
noop
addx 2
addx 3
noop
addx 2
addx 3
noop
addx 2
addx 2
addx -39
noop
addx 15
addx -12
addx 2
addx 10
noop
addx -1
addx -2
noop
addx 5
noop
addx 5
noop
noop
addx 1
addx 4
addx -25
addx 26
addx 2
addx 5
addx 2
noop
addx -3
addx -32
addx 1
addx 4
addx -2
addx 3
noop
noop
addx 3
noop
addx 6
addx -17
addx 27
addx -7
addx 5
addx 2
addx 3
addx -2
addx 4
noop
noop
addx 5
addx 2
addx -39
noop
noop
addx 2
addx 5
addx 3
addx -2
addx 2
addx 11
addx -4
addx -5
noop
addx 10
addx -18
addx 19
addx 2
addx 5
addx 2
addx 2
addx 3
addx -2
addx 2
addx -37
noop
addx 5
addx 4
addx -1
noop
addx 4
noop
noop
addx 1
addx 4
noop
addx 1
addx 2
noop
addx 3
addx 5
noop
addx -3
addx 5
addx 5
addx 2
addx 3
noop
addx -32
noop
`)
