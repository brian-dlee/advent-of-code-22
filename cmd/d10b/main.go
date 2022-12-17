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

const (
	crtWidth  = 40
	crtHeight = 6
)

type Instruction struct {
	Command string
	Value   int
}

func main() {
	var instructions []*Instruction

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

	cycle := 0
	register := 1
	instruction := 0

	screen := make([]rune, crtWidth*crtHeight)
	cursor := 0
	for cursor < len(screen) {
		screen[cursor] = '.'
		cursor++
	}

	for instruction < len(instructions) {
		switch instructions[instruction].Command {
		case "addx":
			applyCycle(&cycle, &screen, register)
			applyCycle(&cycle, &screen, register)
			register += instructions[instruction].Value
			//fmt.Printf("addx %d complete: %d\n", instructions[instruction].Value, register)
		case "noop":
			applyCycle(&cycle, &screen, register)
		}
		instruction++
	}

	applyCycle(&cycle, &screen, register)

	display(screen)
}

func applyCycle(cycle *int, screen *[]rune, register int) {
	//log.Default().Printf("cycle %d: register=%d", *cycle, register)
	pixel := *cycle

	if register-1 <= pixel%40 && pixel%40 <= register+1 {
		(*screen)[pixel] = '#'
	}

	//fmt.Printf("------------ end of cycle %d ------------\n", *cycle)
	//display(*screen)

	*cycle++
}

func display(screen []rune) {
	cursor := 0
	for cursor < crtWidth*crtHeight {
		fmt.Printf("%c", screen[cursor])
		cursor++
		if cursor%crtWidth == 0 {
			fmt.Print("\n")
		}
	}
}
