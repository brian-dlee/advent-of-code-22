package main

import (
	"advent-of-code-22/internal/puzzle"
)

const puzzleDay int = 0
const puzzleInputPart = puzzle.PART_A
const puzzleInputFile = puzzle.FILE_S1

func main() {
	lines := puzzle.ReadInputLinesOrPanic(puzzle.GetInputFile(puzzleDay, puzzleInputPart, puzzleInputFile))

	println("# of lines: %d", len(lines))
}
