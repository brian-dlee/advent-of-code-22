package main

import (
	"advent-of-code-22/internal/puzzle"
)

const puzzleDay = 0 // DAY
const puzzleInputPart = puzzle.PART_A
const puzzleInputFile = puzzle.FILE_S1

func main() {
	lines := puzzle.ReadInputLinesOrPanic(puzzle.GetInputFile(puzzleDay, puzzleInputPart, puzzleInputFile))

	println("# of lines:", len(lines))
}
