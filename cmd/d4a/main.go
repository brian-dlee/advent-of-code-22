package main

import (
	"advent-of-code-22/internal/puzzle"
	"fmt"
	"strconv"
	"strings"
)

const puzzleDay = 4
const puzzleInputPart = puzzle.PART_A
const puzzleInputFile = puzzle.FILE_IN

type AssignmentPair struct {
	One Assignment
	Two Assignment
}

func NewAssignmentPairFromString(data string) (*AssignmentPair, error) {
	parts := strings.Split(data, ",")

	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid assignment pair: %s", data)
	}

	a1, err1 := NewAssignmentFromString(parts[0])
	if err1 != nil {
		return nil, err1
	}

	a2, err2 := NewAssignmentFromString(parts[1])
	if err2 != nil {
		return nil, err2
	}

	return &AssignmentPair{One: *a1, Two: *a2}, nil
}

type Assignment struct {
	Start int
	End int
}

func NewAssignmentFromString(data string) (*Assignment, error) {
	parts := strings.Split(data, "-")

	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid assignment: %s", data)
	}

	start, err1 := strconv.Atoi(parts[0])
	if err1 != nil {
		return nil, err1
	}

	end, err2 := strconv.Atoi(parts[1])
	if err2 != nil {
		return nil, err2
	}

	return &Assignment{Start: start, End: end}, nil
}

func (a Assignment) Contains(b Assignment) bool {
	return a.Start <= b.Start && a.End >= b.End
}

func main() {
	lines := puzzle.ReadInputLinesOrPanic(puzzle.GetInputFile(puzzleDay, puzzleInputPart, puzzleInputFile))
	result := 0

	for _, line := range lines {
		p, err := NewAssignmentPairFromString(line)
		if err != nil {
			panic(err)
		}

		if p.One.Contains(p.Two) || p.Two.Contains(p.One) {
			result += 1
		}
	}

	println("Result:", result)
}
