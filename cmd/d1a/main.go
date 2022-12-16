package main

import (
	"advent-of-code-22/internal/optional"
	"advent-of-code-22/internal/puzzle"
	"fmt"
)

type ElfBackpack struct {
	Total  int
	Number int
}

func main() {
	current := &ElfBackpack{Number: 0, Total: 0}
	highest := current
	elves := make([]*ElfBackpack, 0)
	values := optional.MapStringsToOptionalInts(
		puzzle.ReadInputLinesOrPanic(puzzle.GetInputFile(1, puzzle.PART_A, puzzle.FILE_IN)),
	)

	for _, value := range values {
		if value == nil {
			if current.Total > highest.Total {
				highest = current
			}

			elves = append(elves, current)
			current = &ElfBackpack{Number: current.Number + 1, Total: 0}
		} else {
			current.Total += *value
		}
	}

	fmt.Printf("Highest: %d with %d Calories\n", highest.Number, highest.Total)
}
