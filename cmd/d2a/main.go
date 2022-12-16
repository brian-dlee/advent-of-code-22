package main

import (
	"advent-of-code-22/internal/puzzle"
	"fmt"
	"strings"
)

type RockPaperScissors struct {
	Rounds []Round
}

func NewRockPaperScissorsFromStrings(matches []string) RockPaperScissors {
	rps := RockPaperScissors{Rounds: make([]Round, len(matches))}

	for i, match := range matches {
		rps.Rounds[i] = NewRoundFromString(strings.TrimSpace(match))
	}

	return rps
}

type Round struct {
	Me       string
	Opponent string
}

func NewRoundFromString(data string) Round {
	parts := strings.Split(data, " ")
	return Round{Opponent: parts[0], Me: parts[1]}
}

func (r *Round) WinScore() int {
	a, b := r.Me, r.Opponent

	aR, bR := a == "X", b == "A"
	aP, bP := a == "Y", b == "B"
	aS, bS := a == "Z", b == "C"

	if aR && bS {
		return 6
	}
	if aP && bR {
		return 6
	}
	if aS && bP {
		return 6
	}

	if aR && bR {
		return 3
	}
	if aP && bP {
		return 3
	}
	if aS && bS {
		return 3
	}

	return 0
}

func (r *Round) ShapeScore() int {
	if r.Me == "X" {
		return 1
	}
	if r.Me == "Y" {
		return 2
	}
	return 3
}

func main() {
	totalScore := 0
	rps := NewRockPaperScissorsFromStrings(
		puzzle.ReadInputLinesOrPanic(puzzle.GetInputFile(2, puzzle.PART_A, puzzle.FILE_IN)),
	)

	for _, round := range rps.Rounds {
		totalScore += round.ShapeScore() + round.WinScore()
	}

	fmt.Printf("Strategy score: %d\n", totalScore)
}
