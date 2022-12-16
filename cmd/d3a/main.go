package main

import (
	"advent-of-code-22/internal/puzzle"
	"fmt"
)

const puzzleDay int = 3
const puzzleInputPart = puzzle.PART_A
const puzzleInputFile = puzzle.FILE_IN

type RuckSack struct {
	CompartmentOne string
	CompartmentTwo string
}

func ItemTypeToPriority(item string) (int, error) {
	if len(item) != 1 {
		return 0, fmt.Errorf("invalid item: %s", item)
	}

	value := []rune(item)[0]

	if value >= 'a' && value <= 'z' {
		println(item, "is", int(value-'a'+1))
		return int(value - 'a' + 1), nil
	}

	println(item, "is", int(value-'A'+27))
	return int(value - 'A' + 27), nil
}

func NewRuckSackFromString(data string) RuckSack {
	return RuckSack{
		CompartmentOne: data[0 : len(data)/2],
		CompartmentTwo: data[len(data)/2:],
	}
}

func (r *RuckSack) compartmentAsMap(one bool) map[string]bool {
	var sack string
	if one {
		sack = r.CompartmentOne
	} else {
		sack = r.CompartmentTwo
	}

	result := make(map[string]bool, len(sack))
	for _, letter := range sack {
		result[fmt.Sprintf("%c", letter)] = true
	}

	return result
}

func (r *RuckSack) GetCommonItemType() (string, error) {
	a, b := r.compartmentAsMap(true), r.compartmentAsMap(false)

	println(fmt.Sprintf("a: %v", r.CompartmentOne), fmt.Sprintf("b: %v", r.CompartmentTwo))

	for k := range b {
		if _, ok := a[k]; ok {
			return k, nil
		}
	}

	return "", fmt.Errorf("no common items found")
}

func main() {
	total := 0
	lines := puzzle.ReadInputLinesOrPanic(puzzle.GetInputFile(puzzleDay, puzzleInputPart, puzzleInputFile))

	for _, line := range lines {
		rs := NewRuckSackFromString(line)
		if item, err := rs.GetCommonItemType(); err != nil {
			panic(err)
		} else if priority, err := ItemTypeToPriority(item); err != nil {
			panic(err)
		} else {
			total += priority
		}
	}

	println("total of priorities:", total)
}
