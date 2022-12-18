package main

import (
	"advent-of-code-22/internal/puzzle"
	"fmt"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"
)

const puzzleDay = 11 // DAY
const puzzleInputPart = puzzle.PART_A
const puzzleInputFile = puzzle.FILE_IN

type Operand struct {
	Name     string
	IntValue int
}

type Operation struct {
	Target   string
	Operand1 Operand
	Operand2 Operand
	Operator string
}

type Test struct {
	Value   int
	IfTrue  int
	IfFalse int
}

type Monkey struct {
	Number      int
	Inspections int
	Items       []int64
	Operation   Operation
	Test        Test
}

func main() {
	lines := puzzle.ReadInputLinesOrPanic(puzzle.GetInputFile(puzzleDay, puzzleInputPart, puzzleInputFile))
	monkeys, err := readMonkeys(lines)
	if err != nil {
		panic(fmt.Errorf("failed to read monkeys: %s", err))
	}

	divisor := int64(1)
	for _, monkey := range monkeys {
		divisor *= int64(monkey.Test.Value)
	}

	round := 0
	for round < 10_000 {
		//println("======================= round", round, "=======================")
		for _, monkey := range monkeys {
			//println("Monkey", i)
			for _, item := range monkey.Items {
				worry, result := inspect(monkey, item, divisor)
				if result {
					//println("    Item with worry level", worry, "is thrown to monkey", monkey.Test.IfTrue)
					monkeys[monkey.Test.IfTrue].Items = append(monkeys[monkey.Test.IfTrue].Items, worry)
				} else {
					//println("    Item with worry level", worry, "is thrown to monkey", monkey.Test.IfFalse)
					monkeys[monkey.Test.IfFalse].Items = append(monkeys[monkey.Test.IfFalse].Items, worry)
				}
				monkey.Items = []int64{}
			}
		}

		if round+1 == 1 {
			println("== After round", round+1, "==")
			for i, monkey := range monkeys {
				println("Monkey", i, "inspected items", monkey.Inspections, "times.")
			}
		}

		if round+1 == 20 {
			println("== After round", round+1, "==")
			for i, monkey := range monkeys {
				println("Monkey", i, "inspected items", monkey.Inspections, "times.")
			}
		}

		if (round+1)%1000 == 0 {
			println("== After round", round+1, "==")
			for i, monkey := range monkeys {
				println("Monkey", i, "inspected items", monkey.Inspections, "times.")
			}
		}

		round++
	}

	scores := make([]int, len(monkeys))
	for i, monkey := range monkeys {
		scores[i] = monkey.Inspections
	}

	sort.Ints(scores)

	final := int64(scores[len(scores)-1]) * int64(scores[len(scores)-2])

	println("Monkey business:", final)
}

func applyOperation(op Operation, worry int64) int64 {
	v1, v2 := float64(worry), float64(worry)

	if op.Operand1.Name != "old" {
		v1 = float64(op.Operand1.IntValue)
	}

	if op.Operand2.Name != "old" {
		v2 = float64(op.Operand2.IntValue)
	}

	switch op.Operator {
	case "+":
		next := v1 + v2
		//println("    Worry level increased by", v2, "to", next)
		return int64(math.Floor(next))
	case "*":
		next := v1 * v2
		//if op.Operand2.Name == "old" {
		//	println("    Worry level multiplied by itself to", next)
		//} else {
		//	println("    Worry level multiplied by", v2, "to", next)
		//}
		return int64(math.Floor(next))
	default:
		log.Fatalf("unable to process operation: unknown operator %s", op.Operator)
		return 0
	}
}

func inspect(monkey *Monkey, item int64, divisor int64) (int64, bool) {
	//println("  Monkey inspects an item with a worry level of", item)

	worry := applyOperation(monkey.Operation, item)

	if worry > divisor {
		worry %= divisor
	}

	result := worry%int64(monkey.Test.Value) == 0

	//if result {
	//	println("    Current worry level is divisible by", monkey.Test.Value)
	//} else {
	//	println("    Current worry level is not divisible by", monkey.Test.Value)
	//}

	monkey.Inspections++

	return worry, result
}

func readMonkeys(lines []string) ([]*Monkey, error) {
	var monkeys []*Monkey
	var buffer []string

	for _, line := range lines {
		if strings.HasPrefix(line, "Monkey") {
			if buffer != nil {
				if monkey, err := readMonkey(len(monkeys), buffer); err != nil {
					return nil, err
				} else {
					monkeys = append(monkeys, monkey)
				}
			}
			buffer = make([]string, 0)
		}
		buffer = append(buffer, line)
	}

	if monkey, err := readMonkey(len(monkeys), buffer); err != nil {
		return nil, err
	} else {
		monkeys = append(monkeys, monkey)
	}

	return monkeys, nil
}

func readMonkey(number int, buffer []string) (*Monkey, error) {
	items, e := getMonkeyStartingItems(buffer[1])
	if e != nil {
		return nil, e
	}

	operation, e := getMonkeyOperation(buffer[2])
	if e != nil {
		return nil, e
	}

	test, e := getMonkeyTest(buffer[3], buffer[4], buffer[5])
	if e != nil {
		return nil, e
	}

	return &Monkey{
		Number:    number,
		Items:     items,
		Operation: *operation,
		Test:      *test,
	}, nil
}

func getMonkeyStartingItems(text string) ([]int64, error) {
	var items []int64

	parts := strings.Split(strings.TrimPrefix(text, "  Starting items: "), ", ")

	for _, part := range parts {
		if n, err := strconv.Atoi(part); err != nil {
			return nil, err
		} else {
			items = append(items, int64(n))
		}
	}

	return items, nil
}

func getMonkeyOperation(text string) (*Operation, error) {
	var operation Operation

	parts := strings.Split(strings.TrimPrefix(text, "  Operation: "), " ")

	operation.Target = parts[0]
	operation.Operand1 = Operand{Name: parts[2]}
	operation.Operator = parts[3]
	operation.Operand2 = Operand{Name: parts[4]}

	if operation.Operand1.Name != "old" && operation.Operand1.Name != "new" {
		if n, err := strconv.Atoi(operation.Operand1.Name); err != nil {
			return nil, err
		} else {
			operation.Operand1.IntValue = n
		}
	}

	if operation.Operand2.Name != "old" && operation.Operand2.Name != "new" {
		if n, err := strconv.Atoi(operation.Operand2.Name); err != nil {
			return nil, err
		} else {
			operation.Operand2.IntValue = n
		}
	}

	return &operation, nil
}

func getMonkeyTest(l1, l2, l3 string) (*Test, error) {
	var test Test

	v := strings.TrimPrefix(l1, "  Test: divisible by ")
	t := strings.TrimPrefix(l2, "    If true: throw to monkey ")
	f := strings.TrimPrefix(l3, "    If false: throw to monkey ")

	if n, err := strconv.Atoi(v); err != nil {
		return nil, err
	} else {
		test.Value = n
	}

	if n, err := strconv.Atoi(t); err != nil {
		return nil, err
	} else {
		test.IfTrue = n
	}

	if n, err := strconv.Atoi(f); err != nil {
		return nil, err
	} else {
		test.IfFalse = n
	}

	return &test, nil
}
