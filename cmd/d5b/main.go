package main

import (
	"advent-of-code-22/internal/puzzle"
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
)

const puzzleDay int = 5
const puzzleInputPart = puzzle.PART_A
const puzzleInputFile = puzzle.FILE_IN
const verbose = false

var stackLabelRegex = regexp.MustCompile(`^\s*\d+\s+(?:\s+\d+\s+)*$`)
var moveRegex = regexp.MustCompile("^move (\\d+) from (\\d+) to (\\d+)$")

type Stack struct {
	Containers []string
}

func (s *Stack) minimize() {
	i := len(s.Containers)

	for i > 0 {
		if s.Containers[i-1] != "" {
			break
		}

		i -= 1
	}

	if i == len(s.Containers) {
		return
	}

	s.Containers = s.Containers[0:i]
}

func (s *Stack) PushMany(containers []string) {
	for _, container := range containers {
		s.Containers = append(s.Containers, container)
	}
}

func (s *Stack) PeekMany(count int) []string {
	result := make([]string, count)

	i := 1
	for i <= count {
		if len(s.Containers)-i >= 0 {
			result[i-1] = s.Containers[len(s.Containers)-i]
		} else {
			result[i-1] = ""
		}

		i += 1
	}

	return result
}

func (s *Stack) PopMany(count int) []string {
	if len(s.Containers) >= count {
		result := s.Containers[len(s.Containers)-count:len(s.Containers)]
		s.Containers = s.Containers[0:len(s.Containers)-count]
		return result
	}

	panic(fmt.Errorf("cannot Stack.PopMany() - size %d, count %d", len(s.Containers), count))
}

type Move struct {
	Count int
	Source int
	Target int
}

type CraneMap struct {
	Stacks []*Stack
	Moves []*Move
}

func NewCraneMapFromLines(stackCount int, stackRows []string, moveRows []string) (*CraneMap, error) {
	craneMap := &CraneMap{
		Stacks: make([]*Stack, stackCount),
		Moves:  make([]*Move, 0),
	}

	log.Default().Printf("creating CraneMap: stackCount=%d, stackLineCount=%d, moveLineCount=%d", stackCount, len(stackRows), len(moveRows))

	i := 0
	for i < stackCount {
		craneMap.Stacks[i] = &Stack{Containers: make([]string, len(stackRows))}
		i += 1
	}

	i = 0
	for i < len(stackRows) {
		stack := 0
		cursor := 0
		line := stackRows[len(stackRows) - 1 - i]

		log.Default().Printf("building CraneMap: parsing containers from line - stackLineNumber=%d, stackLine=`%s`", i+1, line)

		for cursor < len(line) {
			log.Default().Printf("building CraneMap: parsing container - stackLineNumber=%d, stackNumber=%d", i+1, stack+1)

			start, end := stack*4, math.Min(float64(stack*4+4), float64(len(line)))

			content := strings.Trim(strings.TrimSpace(line[start:int(end)]), "[]")

			log.Default().Printf("building CraneMap: extracted container name - stackLineNumber=%d, stackNumber=%d, content=`%s`", i+1, stack+1, content)

			craneMap.Stacks[stack].Containers[i] = content

			stack += 1
			cursor = int(end)
		}

		i += 1
	}

	for _, stack := range craneMap.Stacks {
		stack.minimize()
	}

	for i, move := range moveRows {
		craneMap.Moves = append(craneMap.Moves, &Move{})

		result := moveRegex.FindAllStringSubmatch(move, -1)


		if len(result) != 1 {
			return nil, fmt.Errorf("failed to create CraneMap: move instruction %d is invalid in '%s' - match count should be 1, got %d", i+1, move, len(result))
		}

		for j, subMatch := range result[0][1:] {
			if n, err := strconv.Atoi(strings.TrimSpace(subMatch)); err != nil {
				return nil, fmt.Errorf("failed to create CraneMap: move isntruction %d is invalid in '%s' - invalid numeric value in position %d", i+1, move, j+1)
			} else if j == 0 {
				craneMap.Moves[i].Count = n
			} else if j == 1 {
				craneMap.Moves[i].Source = n - 1
			} else if j == 2 {
				craneMap.Moves[i].Target = n - 1
			}
		}
	}

	return craneMap, nil
}

func (m *CraneMap) ExecuteMovesAndReturnTopString() string {
	stacks := make([]*Stack, len(m.Stacks))

	log.Default().Printf("executing moves - moveCount=%d", len(m.Moves))

	for i, stack := range m.Stacks {
		stacks[i] = &Stack{Containers: make([]string, len(stack.Containers))}

		for j, container := range stack.Containers {
			stacks[i].Containers[j] = container
		}
	}

	for i, move := range m.Moves {
		log.Default().Printf("executing move %d - Count=%d, Source=%d, Target=%d", i, move.Count, move.Source, move.Target)

		log.Default().Printf("executing move %d - Source=%d, SourceCount=%d, Target=%d, TargetCount=%d", i, move.Source, len(stacks[move.Source].Containers), move.Target, len(stacks[move.Target].Containers))

		if verbose && i == 0 {
			dump(fmt.Sprintf("before %d - Moving %s from %d to %d", i, stacks[move.Source].PeekMany(move.Count), move.Source, move.Target), stacks)
		}

		pop := stacks[move.Source].PopMany(move.Count)
		stacks[move.Target].PushMany(pop)

		log.Default().Printf("move complete %d - Source=%d, SourceCount=%d, Target=%d, TargetCount=%d, Container=%s", i, move.Source, len(stacks[move.Source].Containers), move.Target, len(stacks[move.Source].Containers), pop)

		if verbose {
			dump(fmt.Sprintf("after %d - Moved %s from %d to %d", i, pop, move.Source, move.Target), stacks)
		}
	}

	result := make([]string, len(m.Stacks))

	for i, stack := range stacks {
		result[i] = stack.PeekMany(1)[0]
	}

	output := strings.Join(result, "")

	if verbose {
		dump(fmt.Sprintf("done - %s", output), stacks)
	}

	return output
}

func main() {
	var craneMap *CraneMap

	lines := puzzle.ReadInputLinesOrPanic(puzzle.GetInputFile(puzzleDay, puzzleInputPart, puzzleInputFile))

	for i, line := range lines {
		log.Default().Printf("looking for stack labels: lineNumber=%d, line=`%s`", i+1, line)

		if stackLabelRegex.MatchString(line) {
			log.Default().Printf("found stack labels: lineNumber=%d", i+1)

			count := determineStackCountByLabelsLine(line)

			log.Default().Printf("determined stack count by labels: stackCount=%d", count)

			log.Default().Printf("splitting input into stack and moves: splitIndex=%d", i)

			for j, preview := range lines[i+2:] {
				log.Default().Printf("preview after split: index=%d, content=`%s`", j, preview)
			}

			if result, err := NewCraneMapFromLines(count, lines[0:i], lines[i+2:]); err != nil {
				panic(err)
			} else {
				craneMap = result
			}
			break
		}
	}

	if craneMap == nil {
		panic(fmt.Errorf("failed to read input: no stack label line detected"))
	}

	println("stack count:", len(craneMap.Stacks), "move count:", len(craneMap.Moves))

	result := craneMap.ExecuteMovesAndReturnTopString()

	println("top after execution:", result)
}

func determineStackCountByLabelsLine(line string) int {
	charactersInEachLabel := 4.0
	lengthWithCompleteFinalLabel := len(line)+1

	return int(math.Floor(float64(lengthWithCompleteFinalLabel) / charactersInEachLabel))
}

func dump(title string, stacks []*Stack) {
	fmt.Println("========================================================================================")
	fmt.Println(title)

	max := 0
	for _, stack := range stacks {
		max = int(math.Max(float64(len(stack.Containers)), float64(max)))
	}

	row := max - 1
	for row >= 0 {
		for i, stack := range stacks {
			if row >= len(stack.Containers) {
				fmt.Printf(" %s ", " ")
			} else {
				fmt.Printf(" %s ", stack.Containers[row])
			}

			if i == len(stacks) - 1 {
				fmt.Printf("\n")
			} else {
				fmt.Printf(" ")
			}
		}

		row -= 1
	}

	fmt.Println("========================================================================================")
}
