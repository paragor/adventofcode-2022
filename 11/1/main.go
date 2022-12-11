package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	monkeys := []*Monkey{}
	descrs := strings.Split(inputA, "\n\n")
	for _, descr := range descrs {
		monkeys = append(monkeys, &Monkey{
			Descr: descr,
		})
	}
	for _, monkey := range monkeys {
		monkey.Init(monkeys)
	}
	const (
		rounds = 20
	)
	for i := 0; i < rounds; i++ {
		for _, monkey := range monkeys {
			monkey.Turn()
		}
	}

	inspections := []int{}
	for i, monkey := range monkeys {
		fmt.Printf("Monkey %d: %s\n", i, monkey.Print())
		inspections = append(inspections, monkey.Inspections)
	}
	sort.Ints(inspections)
	fmt.Println(inspections[len(inspections)-1] * inspections[len(inspections)-2])
}

type Monkey struct {
	Inspections     int
	Descr           string
	Items           []int
	OperationStr    string
	Operation       func(old int) int
	TestStr         string
	Test            func(item int) bool
	TestTrueMoveTo  *Monkey
	TestFalseMoveTo *Monkey
}

func (m *Monkey) Init(monkeys []*Monkey) {
	for _, str := range strings.Split(m.Descr, "\n")[1:] {
		str = strings.TrimSpace(str)
		if strings.HasPrefix(str, "Starting items: ") {
			items := strings.Split(
				strings.Replace(str, "Starting items: ", "", 1),
				", ",
			)
			for _, item := range items {
				itemI, err := strconv.Atoi(item)
				if err != nil {
					panic(err)
				}
				m.Items = append(m.Items, itemI)
			}

		}
		if strings.HasPrefix(str, "Test: ") {
			m.TestStr = strings.Replace(str, "Test: ", "", 1)
		}
		if strings.HasPrefix(str, "Operation: new = ") {
			m.OperationStr = strings.Replace(str, "Operation: new = ", "", 1)
		}
		if strings.HasPrefix(str, "If true: throw to monkey ") {
			monkeyIdx, err := strconv.Atoi(strings.Replace(str, "If true: throw to monkey ", "", 1))
			if err != nil {
				panic(err)
			}
			m.TestTrueMoveTo = monkeys[monkeyIdx]
		}
		if strings.HasPrefix(str, "If false: throw to monkey ") {
			monkeyIdx, err := strconv.Atoi(strings.Replace(str, "If false: throw to monkey ", "", 1))
			if err != nil {
				panic(err)
			}
			m.TestFalseMoveTo = monkeys[monkeyIdx]
		}
	}

	if strings.HasPrefix(m.TestStr, "divisible by ") {
		divisible, err := strconv.Atoi(strings.Replace(m.TestStr, "divisible by ", "", 1))
		if err != nil {
			panic(err)
		}

		m.Test = func(item int) bool {
			return item%divisible == 0
		}
	}

	operation := strings.Split(m.OperationStr, " ")
	m.Operation = func(old int) int {
		var (
			err error
		)
		first, operator, second := operation[0], operation[1], operation[2]

		res := 0
		reg1 := 0
		reg2 := 0

		if first == "old" {
			reg1 = old
		} else {
			reg1, err = strconv.Atoi(first)
			if err != nil {
				fmt.Printf("%+v", *m)
				panic(err)
			}
		}

		if second == "old" {
			reg2 = old
		} else {
			reg2, err = strconv.Atoi(second)
			if err != nil {
				panic(err)
			}
		}
		if operator == "+" {
			res = reg1 + reg2
		} else if operator == "*" {
			res = reg1 * reg2
		} else {
			panic("unknown operator")
		}
		return res
	}

}

func (m *Monkey) Turn() {
	items := m.Items
	m.Items = []int{}

	var item int
	for {
		if len(items) == 0 {
			break
		}
		m.Inspections++

		item, items = items[0], items[1:]
		newItem := m.Operation(item)
		newItem = newItem / 3
		if m.Test(newItem) {
			m.TestTrueMoveTo.Items = append(m.TestTrueMoveTo.Items, newItem)
		} else {
			m.TestFalseMoveTo.Items = append(m.TestFalseMoveTo.Items, newItem)
		}
	}
}

func (m *Monkey) Print() string {
	return fmt.Sprintf("%v\t| %d", m.Items, m.Inspections)
}

var input = strings.TrimSpace(`
Monkey 0:
  Starting items: 79, 98
  Operation: new = old * 19
  Test: divisible by 23
    If true: throw to monkey 2
    If false: throw to monkey 3

Monkey 1:
  Starting items: 54, 65, 75, 74
  Operation: new = old + 6
  Test: divisible by 19
    If true: throw to monkey 2
    If false: throw to monkey 0

Monkey 2:
  Starting items: 79, 60, 97
  Operation: new = old * old
  Test: divisible by 13
    If true: throw to monkey 1
    If false: throw to monkey 3

Monkey 3:
  Starting items: 74
  Operation: new = old + 3
  Test: divisible by 17
    If true: throw to monkey 0
    If false: throw to monkey 1
`)
var inputA = strings.TrimSpace(`
Monkey 0:
  Starting items: 91, 54, 70, 61, 64, 64, 60, 85
  Operation: new = old * 13
  Test: divisible by 2
    If true: throw to monkey 5
    If false: throw to monkey 2

Monkey 1:
  Starting items: 82
  Operation: new = old + 7
  Test: divisible by 13
    If true: throw to monkey 4
    If false: throw to monkey 3

Monkey 2:
  Starting items: 84, 93, 70
  Operation: new = old + 2
  Test: divisible by 5
    If true: throw to monkey 5
    If false: throw to monkey 1

Monkey 3:
  Starting items: 78, 56, 85, 93
  Operation: new = old * 2
  Test: divisible by 3
    If true: throw to monkey 6
    If false: throw to monkey 7

Monkey 4:
  Starting items: 64, 57, 81, 95, 52, 71, 58
  Operation: new = old * old
  Test: divisible by 11
    If true: throw to monkey 7
    If false: throw to monkey 3

Monkey 5:
  Starting items: 58, 71, 96, 58, 68, 90
  Operation: new = old + 6
  Test: divisible by 17
    If true: throw to monkey 4
    If false: throw to monkey 1

Monkey 6:
  Starting items: 56, 99, 89, 97, 81
  Operation: new = old + 1
  Test: divisible by 7
    If true: throw to monkey 0
    If false: throw to monkey 2

Monkey 7:
  Starting items: 68, 72
  Operation: new = old + 8
  Test: divisible by 19
    If true: throw to monkey 6
    If false: throw to monkey 0
`)
