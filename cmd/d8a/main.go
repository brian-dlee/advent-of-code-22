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
	Row    int
	Col    int
	Height int
}

type TreeGrid struct {
	heights    []int
	visibility []uint8
	RowCount   int
	ColCount   int
}

func NewTreeGrid() *TreeGrid {
	return &TreeGrid{
		heights:    make([]int, 0),
		visibility: make([]uint8, 0),
		RowCount:   0,
		ColCount:   0,
	}
}

func (g *TreeGrid) AddTree(height int) {
	g.heights = append(g.heights, height)
	g.visibility = append(g.visibility, 0)
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
			height := float64(g.GetHeight(row, col))
			if height > vMaxHeight[col] {
				g.visibility[row*g.ColCount+col] |= uint8(DirU)
			}
			if height > hMaxHeight {
				g.visibility[row*g.ColCount+col] |= uint8(DirL)
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
			height := float64(g.GetHeight(row, col))
			if height > vMaxHeight[col] {
				g.visibility[row*g.ColCount+col] |= uint8(DirD)
			}
			if height > hMaxHeight {
				g.visibility[row*g.ColCount+col] |= uint8(DirR)
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

func (g *TreeGrid) GetHeight(row, col int) int {
	if row < 0 || row > g.RowCount-1 {
		return -1
	}

	if col < 0 || col > g.ColCount-1 {
		return -1
	}

	return g.heights[row*g.ColCount+col]
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
					Height: g.GetHeight(row, col),
				})
			}
			col++
		}
		row++
	}

	return visible
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
}
