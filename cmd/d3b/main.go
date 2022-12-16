package main

import (
	"advent-of-code-22/internal/puzzle"
	"fmt"
)

const puzzleDay int = 3
const puzzleInputPart = puzzle.PART_A
const puzzleInputFile = puzzle.FILE_IN

type ElfGroup struct {
	SackOne   string
	SackTwo   string
	SackThree string
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

func NewRuckSackFromStrings(one, two, three string) ElfGroup {
	return ElfGroup{
		SackOne:   one,
		SackTwo:   two,
		SackThree: three,
	}
}

func (r *ElfGroup) sackContentsAsMap(flag int) map[string]bool {
	var sack string
	if flag < 0 {
		sack = r.SackOne
	} else if flag == 0 {
		sack = r.SackTwo
	} else if flag > 0 {
		sack = r.SackThree
	}

	result := make(map[string]bool, len(sack))
	for _, letter := range sack {
		result[fmt.Sprintf("%c", letter)] = true
	}

	return result
}

func (r *ElfGroup) GetCommonItemType() (string, error) {
	a, b, c := r.sackContentsAsMap(-1), r.sackContentsAsMap(0), r.sackContentsAsMap(1)
	ab := make(map[string]bool, 0)

	for k := range b {
		if _, ok := a[k]; ok {
			ab[k] = true
		}
	}

	for k := range ab {
		if _, ok := c[k]; ok {
			return k, nil
		}
	}

	return "", fmt.Errorf("no common items found")
}

func main() {
	total := 0
	lines := puzzle.ReadInputLinesOrPanic(puzzle.GetInputFile(puzzleDay, puzzleInputPart, puzzleInputFile))

	var one, two string
	for i, line := range lines {
		n := (i + 1) % 3
		if n == 1 {
			one = line
		} else if n == 2 {
			two = line
		} else if n == 0 {
			rs := NewRuckSackFromStrings(one, two, line)
			if item, err := rs.GetCommonItemType(); err != nil {
				panic(err)
			} else if priority, err := ItemTypeToPriority(item); err != nil {
				panic(err)
			} else {
				total += priority
			}
		}
	}

	println("total of priorities:", total)
}
