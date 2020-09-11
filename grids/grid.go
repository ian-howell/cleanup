package grids

import "fmt"

type Grid interface {
	NumRows() int
	NumCols() int
	InBounds(row, col int) bool
	IsSet(row, col int) (bool, error)
	Set(row, col int) error
	Unset(row, col int) error
	Toggle(row, col int) error
}

type OutOfBoundsError struct {
	row int
	col int
}

func (err OutOfBoundsError) Error() string {
	return fmt.Sprintf("cell [%d, %d] is out of bounds", err.row, err.col)
}

type grid [][]bool

func New(rows, cols int) grid {
	g := make([][]bool, 0, rows)
	for ; rows > 0; rows-- {
		g = append(g, make([]bool, cols))
	}
	return g
}

func (g grid) NumRows() int {
	return len(g)
}

func (g grid) NumCols() int {
	if len(g) == 0 {
		return 0
	}
	return len(g[0])
}

func (g grid) IsSet(row, col int) (bool, error) {
	if !g.InBounds(row, col) {
		return false, OutOfBoundsError{row: row, col: col}
	}
	return g[row][col], nil
}

func (g grid) Set(row, col int) error {
	if !g.InBounds(row, col) {
		return OutOfBoundsError{row: row, col: col}
	}
	g[row][col] = true
	return nil
}

func (g grid) Unset(row, col int) error {
	if !g.InBounds(row, col) {
		return OutOfBoundsError{row: row, col: col}
	}
	g[row][col] = false
	return nil
}

func (g grid) Toggle(row, col int) error {
	if !g.InBounds(row, col) {
		return OutOfBoundsError{row: row, col: col}
	}
	g[row][col] = !g[row][col]
	return nil
}

func (g grid) InBounds(row, col int) bool {
	return 0 <= row && row < g.NumRows() && 0 <= col && col < g.NumCols()
}
