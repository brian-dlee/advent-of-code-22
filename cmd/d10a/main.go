package main

import (
	"advent-of-code-22/internal/puzzle"
	"fmt"
	"strconv"
	"strings"
)

const puzzleDay = 10 // DAY
const puzzleInputPart = puzzle.PART_A
const puzzleInputFile = puzzle.FILE_IN

type Instruction struct {
	Command string
	Value   int
}

func main() {
	var instructions []*Instruction

	tracker := make(map[int]int, 0)
	lines := puzzle.ReadInputLinesOrPanic(puzzle.GetInputFile(puzzleDay, puzzleInputPart, puzzleInputFile))

	for i, line := range lines {
		parts := strings.Split(line, " ")
		instruction := Instruction{
			Command: parts[0],
			Value:   0,
		}

		if len(parts) > 1 {
			if n, err := strconv.Atoi(parts[1]); err != nil {
				panic(fmt.Errorf("failed to read line %d: %s", i, err))
			} else {
				instruction.Value = n
			}
		}

		instructions = append(instructions, &instruction)
	}

	cycle := 1
	register := 1
	cursor := 0

	for cursor < len(instructions) {
		switch instructions[cursor].Command {
		case "addx":
			applyCycle(&cycle, &tracker, register)
			applyCycle(&cycle, &tracker, register)
			register += instructions[cursor].Value
		case "noop":
			applyCycle(&cycle, &tracker, register)
		}
		cursor++
	}

	applyCycle(&cycle, &tracker, register)

	totalSignalStrength := 0
	for _, v := range tracker {
		totalSignalStrength += v
	}

	println("total signal strength:", totalSignalStrength)
}

func applyCycle(cycle *int, tracker *map[int]int, register int) {
	if *cycle-20 == 0 || (*cycle-20)%40 == 0 {
		(*tracker)[*cycle] = register * *cycle
	}
	*cycle++
}
