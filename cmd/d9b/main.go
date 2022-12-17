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

func dump(maxX, maxY int, knots []*Coord, visited map[string]bool) {
	y := maxY - 1
	fmt.Println(strings.Repeat("=", maxX))
	for y >= 0 {
		x := 0
		for x < maxX {
			space := ""
			for k, knot := range knots {
				if knot.x == x && knot.y == y {
					if k == 0 {
						space = "H"
					} else {
						space = fmt.Sprintf("%d", k)
					}
				}
				if space != "" {
					break
				}
			}

			if space != "" {
				fmt.Print(space)
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
	knots := make([]*Coord, 10)
	visited := map[string]bool{ZeroZero().ToString(): true}

	for i := range knots {
		knots[i] = ZeroZero()
	}

	dump(20, 20, knots, visited)

	for i, line := range lines {
		m, err := ToMovement(line)
		if err != nil {
			panic(fmt.Errorf("failed to read line: %s", err))
		}

		log.Default().Printf("movement %d: %v", i, m)

		j := 0

		for j < m.Value {
			for k, knot := range knots {
				if k == 0 {
					switch m.Direction {
					case UP:
						knot.y++
					case DOWN:
						knot.y--
					case RIGHT:
						knot.x++
					case LEFT:
						knot.x--
					}
					continue
				}

				head := knots[k-1]
				tail := knot
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
			}

			dump(10, 10, knots, visited)

			//log.Default().Printf("positions %d,%d: h: %v, t: %v", i, j, head, tail)

			visited[knots[len(knots)-1].ToString()] = true

			j++
		}

		log.Default().Printf("visited count %d: %d", i, len(visited))
	}

	println("# of visits:", len(visited))
}
