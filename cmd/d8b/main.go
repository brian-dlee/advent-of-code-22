package main

import (
	"advent-of-code-22/internal/puzzle"
	"fmt"
	"math"
	"strconv"
)

const puzzleDay = 8 // DAY
const puzzleInputPart = puzzle.PART_A
const puzzleInputFile = puzzle.FILE_IN

type Direction uint8

const (
	DirU Direction = 1
	DirD           = 2
	DirL           = 4
	DirR           = 8
)

type TreeRef struct {
	Row         int
	Col         int
	Height      int
	ScenicScore int
}

type TreeGrid struct {
	heights    []float64
	visibility []uint8
	scores     []map[Direction]float64
	RowCount   int
	ColCount   int
}

func NewTreeGrid() *TreeGrid {
	return &TreeGrid{
		heights:    make([]float64, 0),
		visibility: make([]uint8, 0),
		scores:     make([]map[Direction]float64, 0),
		RowCount:   0,
		ColCount:   0,
	}
}

func (g *TreeGrid) AddTree(height int) {
	g.heights = append(g.heights, float64(height))
	g.visibility = append(g.visibility, 0)

	scores := make(map[Direction]float64, 4)
	scores[DirL] = 0
	scores[DirU] = 0
	scores[DirR] = 0
	scores[DirD] = 0

	g.scores = append(g.scores, scores)
}

func (g *TreeGrid) Finalize() error {
	if len(g.heights) != g.RowCount*g.ColCount {
		return fmt.Errorf("invalid tree count: %d != %dx%d", len(g.heights), g.RowCount, g.ColCount)
	}

	row := 0

	vMaxHeight := make([]float64, g.ColCount)
	for i := range vMaxHeight {
		vMaxHeight[i] = -1
	}

	for row < g.RowCount {
		hMaxHeight := float64(-1)
		col := 0
		for col < g.ColCount {
			height := g.GetHeight(row, col)
			if height > vMaxHeight[col] {
				g.addVisibility(row, col, DirU)
			}
			if row == 0 {
				g.setScore(row, col, DirU, 0)
			} else if height > g.GetHeight(row-1, col) {
				g.setScore(row, col, DirU, g.GetScore(row-1, col, DirU)+1)
			} else {
				g.setScore(row, col, DirU, 1)
			}
			if height > hMaxHeight {
				g.addVisibility(row, col, DirL)
			}
			if col == 0 {
				g.setScore(row, col, DirL, 0)
			} else if height > g.GetHeight(row, col-1) {
				g.setScore(row, col, DirL, g.GetScore(row, col-1, DirL)+1)
			} else {
				g.setScore(row, col, DirL, 1)
			}
			vMaxHeight[col] = math.Max(vMaxHeight[col], height)
			hMaxHeight = math.Max(hMaxHeight, height)
			col++
		}
		row++
	}

	row = g.RowCount - 1

	for i := range vMaxHeight {
		vMaxHeight[i] = -1
	}

	for row >= 0 {
		hMaxHeight := float64(-1)
		col := g.ColCount - 1
		for col >= 0 {
			height := g.GetHeight(row, col)
			if height > vMaxHeight[col] {
				g.addVisibility(row, col, DirD)
			}
			if row == g.RowCount-1 {
				g.setScore(row, col, DirD, 0)
			} else if height > g.GetHeight(row+1, col) {
				g.setScore(row, col, DirD, g.GetScore(row+1, col, DirD)+1)
			} else {
				g.setScore(row, col, DirD, 1)
			}
			if height > hMaxHeight {
				g.addVisibility(row, col, DirR)
			}
			if col == g.ColCount-1 {
				g.setScore(row, col, DirR, 0)
			} else if height > g.GetHeight(row, col+1) {
				g.setScore(row, col, DirR, g.GetScore(row, col+1, DirR)+1)
			} else {
				g.setScore(row, col, DirR, 1)
			}
			vMaxHeight[col] = math.Max(vMaxHeight[col], height)
			hMaxHeight = math.Max(hMaxHeight, height)
			col--
		}
		row--
	}

	return nil
}

func (g *TreeGrid) IsVisibleFromDirection(row, col int, direction Direction) bool {
	return g.visibility[row*g.ColCount+col]&uint8(direction) > uint8(0)
}

func (g *TreeGrid) GetHeight(row, col int) float64 {
	if row < 0 || row > g.RowCount-1 {
		return -1
	}

	if col < 0 || col > g.ColCount-1 {
		return -1
	}

	return g.heights[row*g.ColCount+col]
}

func (g *TreeGrid) GetMaximumScenicScore() *TreeRef {
	var t *TreeRef

	row := 0
	for row < g.RowCount {
		col := 0
		for col < g.ColCount {
			s1 := g.GetScore(row, col, DirL)
			s2 := g.GetScore(row, col, DirU)
			s3 := g.GetScore(row, col, DirR)
			s4 := g.GetScore(row, col, DirD)
			sT := s1 * s2 * s3 * s4

			if t == nil || sT > float64(t.ScenicScore) {
				t = &TreeRef{
					Row:         row,
					Col:         col,
					Height:      int(g.GetHeight(row, col)),
					ScenicScore: int(sT),
				}
			}
			col++
		}
		row++
	}

	return t
}

func (g *TreeGrid) GetScore(row, col int, direction Direction) float64 {
	if row < 0 || row > g.RowCount-1 {
		return -1
	}

	if col < 0 || col > g.ColCount-1 {
		return -1
	}

	return g.scores[row*g.ColCount+col][direction]
}

func (g *TreeGrid) GetVisibleTrees() []TreeRef {
	row, visible := 0, make([]TreeRef, 0)

	for row < g.RowCount {
		col := 0
		for col < g.ColCount {
			isVisible := g.IsVisibleFromDirection(row, col, DirL) ||
				g.IsVisibleFromDirection(row, col, DirU) ||
				g.IsVisibleFromDirection(row, col, DirR) ||
				g.IsVisibleFromDirection(row, col, DirD)

			if isVisible {
				visible = append(visible, TreeRef{
					Row:    row,
					Col:    col,
					Height: int(g.GetHeight(row, col)),
				})
			}
			col++
		}
		row++
	}

	return visible
}

func (g *TreeGrid) GetVisibility(row, col int) uint8 {
	if row < 0 || row > g.RowCount-1 {
		return 0
	}

	if col < 0 || col > g.ColCount-1 {
		return 0
	}

	return g.visibility[row*g.ColCount+col]
}

func (g *TreeGrid) addVisibility(row, col int, visibility Direction) {
	if row < 0 || row > g.RowCount-1 {
		return
	}

	if col < 0 || col > g.ColCount-1 {
		return
	}

	g.visibility[row*g.ColCount+col] |= uint8(visibility)
}

func (g *TreeGrid) setScore(row, col int, direction Direction, score float64) {
	if row < 0 || row > g.RowCount-1 {
		return
	}

	if col < 0 || col > g.ColCount-1 {
		return
	}

	g.scores[row*g.ColCount+col][direction] = score
}

func main() {
	lines := puzzle.ReadInputLinesOrPanic(puzzle.GetInputFile(puzzleDay, puzzleInputPart, puzzleInputFile))
	grid := NewTreeGrid()
	total := 0

	for i, line := range lines {
		for j, char := range line {
			if n, err := strconv.Atoi(fmt.Sprintf("%c", char)); err != nil {
				panic(fmt.Errorf("failed to read tree height: %d,%d - %s", i, j, char))
			} else {
				grid.AddTree(n)
				total++
			}
		}
		grid.RowCount++
	}

	grid.ColCount = total / grid.RowCount

	if err := grid.Finalize(); err != nil {
		panic(err)
	}

	println("Visible tree count:", len(grid.GetVisibleTrees()))

	var max float64 = -1
	y := 0
	for y < grid.RowCount {
		x := 0
		for x < grid.ColCount {
			max = math.Max(getScenicScoreTheSlowWay(grid, y, x), max)
			x++
		}
		y++
	}

	println("Maximum scenic score:", grid.GetMaximumScenicScore().ScenicScore)
	println("Maximum scenic score (the slow way):", int(max))
}

func getScenicScoreTheSlowWay(grid *TreeGrid, row, col int) float64 {
	h := grid.GetHeight(row, col)
	distance := 0
	l, u, r, d := 0, 0, 0, 0
	x, y := 0, 0

	x = col - 1
	distance = 0
	if col > 0 {
		for x >= 0 {
			distance += 1
			if grid.GetHeight(row, x) >= h {
				break
			}
			x--
		}
	}
	l = distance

	y = row - 1
	distance = 0
	if row > 0 {
		for y >= 0 {
			distance += 1
			if grid.GetHeight(y, col) >= h {
				break
			}
			y--
		}
	}
	u = distance

	x = col + 1
	distance = 0
	if col < grid.ColCount-1 {
		for x <= grid.ColCount-1 {
			distance += 1
			if grid.GetHeight(row, x) >= h {
				break
			}
			x++
		}
	}
	r = distance

	y = row + 1
	distance = 0
	if row < grid.RowCount-1 {
		for y <= grid.RowCount-1 {
			distance += 1
			if grid.GetHeight(y, col) >= h {
				break
			}
			y++
		}
	}
	d = distance

	return float64(l * u * r * d)
}
