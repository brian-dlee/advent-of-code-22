package main

import (
	"advent-of-code-22/internal/optional"
	"advent-of-code-22/internal/puzzle"
	"fmt"
	"sort"
)

type ElfBackpack struct {
	Total  int
	Number int
}

type RankedElves struct {
	Elves []*ElfBackpack
	Size  int
}

func NewRankedElves(size int) RankedElves {
	return RankedElves{
		Elves: make([]*ElfBackpack, 0),
		Size:  size,
	}
}

func (r *RankedElves) Total() int {
	total := 0

	for _, e := range r.Elves {
		total += e.Total
	}

	return total
}

func (r *RankedElves) Len() int {
	return len(r.Elves)
}

func (r *RankedElves) Less(i, j int) bool {
	if r.Elves[i] == nil {
		return true
	}

	if r.Elves[j] == nil {
		return false
	}

	return r.Elves[i].Total < r.Elves[j].Total
}

func (r *RankedElves) Swap(i, j int) {
	tmp := r.Elves[i]
	r.Elves[i] = r.Elves[j]
	r.Elves[j] = tmp
}

func (r *RankedElves) Submit(elf *ElfBackpack) {
	r.Elves = append(r.Elves, elf)

	sort.Sort(sort.Reverse(r))

	if len(r.Elves) > r.Size {
		r.Elves = r.Elves[0:r.Size]
	}
}

func main() {
	current := &ElfBackpack{Number: 0, Total: 0}
	ranking := NewRankedElves(3)
	values := optional.MapStringsToOptionalInts(
		puzzle.ReadInputLinesOrPanic(puzzle.GetInputFile(1, puzzle.PART_A, puzzle.FILE_IN)),
	)

	for _, value := range values {
		if value == nil {
			ranking.Submit(current)
			current = &ElfBackpack{Number: current.Number + 1, Total: 0}
		} else {
			current.Total += *value
		}
	}

	fmt.Printf("Total of ranking: Top %d have %d calories total\n", ranking.Size, ranking.Total())
}
