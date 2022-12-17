package main

import (
	"advent-of-code-22/internal/puzzle"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

const puzzleDay = 9 // DAY
const puzzleInputPart = puzzle.PART_A
const puzzleInputFile = puzzle.FILE_IN

type Direction string

const (
	UP    Direction = "U"
	DOWN            = "D"
	LEFT            = "L"
	RIGHT           = "R"
)

type Movement struct {
	Direction Direction
	Value     int
}

func ToMovement(line string) (*Movement, error) {
	parts := strings.Split(line, " ")

	if n, err := strconv.Atoi(parts[1]); err != nil {
		return nil, err
	} else {
		return &Movement{Direction: Direction(parts[0]), Value: n}, nil
	}
}

type Coord struct {
	x int
	y int
}

func ZeroZero() *Coord {
	return &Coord{
		x: 0,
		y: 0,
	}
}

func (c *Coord) Distance(o *Coord) float64 {
	return math.Abs(float64(c.x)-float64(o.x)) + math.Abs(float64(c.y)-float64(o.y))
}

func (c *Coord) ToString() string {
	return fmt.Sprintf("%d,%d", c.x, c.y)
}

func dump(maxX, maxY int, head, tail *Coord, visited map[string]bool) {
	y := maxY - 1
	fmt.Println(strings.Repeat("=", maxX))
	for y >= 0 {
		x := 0
		for x < maxX {
			if head.x == x && head.y == y {
				fmt.Print("H")
			} else if tail.x == x && tail.y == y {
				fmt.Print("T")
			} else if x == 0 && y == 0 {
				fmt.Print("s")
			} else if _, ok := visited[fmt.Sprintf("%d,%d", x, y)]; ok {
				fmt.Print("*")
			} else {
				fmt.Print(".")
			}
			x++
		}
		fmt.Print("\n")
		y--
	}
	fmt.Println(strings.Repeat("=", maxX))
	fmt.Print("\n")
}

func main() {
	lines := puzzle.ReadInputLinesOrPanic(puzzle.GetInputFile(puzzleDay, puzzleInputPart, puzzleInputFile))
	head := ZeroZero()
	tail := ZeroZero()
	visited := map[string]bool{head.ToString(): true}

	dump(10, 10, head, tail, visited)

	for i, line := range lines {
		m, err := ToMovement(line)
		if err != nil {
			panic(fmt.Errorf("failed to read line: %s", err))
		}

		log.Default().Printf("movement %d: %v", i, m)

		j := 0

		for j < m.Value {
			switch m.Direction {
			case UP:
				head.y++
			case DOWN:
				head.y--
			case RIGHT:
				head.x++
			case LEFT:
				head.x--
			}

			row := head.y == tail.y
			col := head.x == tail.x
			distance := head.Distance(tail)

			if !row && !col && distance > 2 {
				if head.y > tail.y {
					tail.y++
				} else {
					tail.y--
				}
				if head.x > tail.x {
					tail.x++
				} else {
					tail.x--
				}
			} else if row && !col && distance > 1 {
				if head.x > tail.x {
					tail.x++
				} else {
					tail.x--
				}
			} else if !row && col && distance > 1 {
				if head.y > tail.y {
					tail.y++
				} else {
					tail.y--
				}
			}

			dump(10, 10, head, tail, visited)

			log.Default().Printf("positions %d,%d: h: %v, t: %v", i, j, head, tail)

			visited[tail.ToString()] = true

			j++
		}

		log.Default().Printf("visited count %d: %d", i, len(visited))
	}

	println("# of visits:", len(visited))
}
