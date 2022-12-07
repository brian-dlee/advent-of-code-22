package main

import (
	"advent-of-code-22/internal/puzzle"
	"fmt"
	"log"
)

const puzzleDay = 6 // DAY
const puzzleInputPart = puzzle.PART_A
const puzzleInputFile = puzzle.FILE_IN

func main() {
	lines := puzzle.ReadInputLinesOrPanic(puzzle.GetInputFile(puzzleDay, puzzleInputPart, puzzleInputFile))

	if len(lines) > 1 {
		panic(fmt.Errorf("only one line is expected"))
	}

	buffer := make([]rune, 4)
	cursor := 0
	passed := 0

	for i, c := range lines[0] {
		log.Default().Printf("checking character %d: %c", i, c)

		if cursor > 3 && isBufferUniqueWithCursor(buffer, cursor) {
			break
		}

		buffer[cursor%len(buffer)] = c

		passed += 1
		cursor += 1
	}

	println("result:", passed)
}

func isBufferUniqueWithCursor(chars []rune, startPosition int) bool {
	i := 0
	store := make(map[rune]bool, 0)

	log.Default().Printf("checking buffer: %s", string(chars))

	for i < len(chars) {
		c := chars[(startPosition+i)%len(chars)]

		if _, ok := store[c]; ok {
			log.Default().Printf("buffer check: %s - found duplicate: %c", string(chars), c)
			return false
		} else {
			log.Default().Printf("buffer check %s - found unique: %c", string(chars), c)
		}

		store[c] = true
		i++
	}

	return true
}
