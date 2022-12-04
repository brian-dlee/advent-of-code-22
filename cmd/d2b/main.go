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
	Outcome string
	Opponent string
}

func NewRoundFromString(data string) Round {
	parts := strings.Split(data, " ")
	return Round{Opponent: parts[0], Outcome: parts[1]}
}

func (r *Round) winScore() int {
	if r.Outcome == "X" { return 0 }
	if r.Outcome == "Y" { return 3 }
	return 6
}

func (r *Round) shapeScore(shape string) int {
	if shape == "X" { return 1 }
	if shape == "Y" { return 2 }
	return 3
}

func (r *Round) Score() int {
	var shape string

	aR := r.Opponent == "A"
	aP := r.Opponent == "B"
	aS := r.Opponent == "C"
	win := r.winScore()

	if win == 6 && aR { shape = "Y" }
	if win == 6 && aP { shape = "Z" }
	if win == 6 && aS { shape = "X" }
	if win == 3 && aR { shape = "X" }
	if win == 3 && aP { shape = "Y" }
	if win == 3 && aS { shape = "Z" }
	if win == 0 && aR { shape = "Z" }
	if win == 0 && aP { shape = "X" }
	if win == 0 && aS { shape = "Y" }

	return win + r.shapeScore(shape)
}

func main() {
	totalScore := 0
	rps := NewRockPaperScissorsFromStrings(
		puzzle.ReadInputLinesOrPanic(puzzle.GetInputFile(2, puzzle.PART_A, puzzle.FILE_IN)),
	)

	for _, round := range rps.Rounds {
		totalScore += round.Score()
	}

	fmt.Printf("Strategy score: %d\n", totalScore)
}
