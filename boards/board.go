package boards

import (
	"github.com/ian-howell/gocurse/curses"

	"github.com/ian-howell/cleanup/grids"
)

type Direction int

var directionMap = map[Direction][]int{
	curses.KEY_DOWN:  []int{+1, +0},
	curses.KEY_RIGHT: []int{+0, +1},
	curses.KEY_UP:    []int{-1, -0},
	curses.KEY_LEFT:  []int{-0, -1},
}

type Board interface {
	Draw(win *curses.Window) error
	IsSolved() bool
	Move(dir Direction)
	Flip()
}

var _ Board = &board{}

type board struct {
	grids.Grid
	currentRow int
	currentCol int
}

func New(rows, cols int) board {
	return board{Grid: grids.New(rows, cols)}
}

func (b board) Draw(win *curses.Window) error {
	for r := 0; r < b.NumRows(); r++ {
		for c := 0; c < b.NumCols(); c++ {
			b.drawTile(win, r, c)
		}
	}
	win.Redraw()
	return nil
}

func (b *board) Flip() {
	for _, m := range directionMap {
		// Safe to ignore errors here
		b.Toggle(b.currentRow+m[0], b.currentCol+m[1])
	}
}

func (b *board) Move(dir Direction) {
	delta := directionMap[dir]
	newRow, newCol := b.currentRow+delta[0], b.currentCol+delta[1]
	if b.InBounds(newRow, newCol) {
		b.currentRow, b.currentCol = newRow, newCol
	}
}

func (b board) IsSolved() bool {
	for r := 0; r < b.NumRows(); r++ {
		for c := 0; c < b.NumCols(); c++ {
			if isSet, _ := b.IsSet(r, c); !isSet {
				return false
			}
		}
	}
	return true
}

func (b board) drawTile(win *curses.Window, row, col int) error {
	isSet, err := b.IsSet(row, col)
	if err != nil {
		return err
	}

	color := curses.Color_pair(1)
	if isSet {
		color = curses.Color_pair(2)
	}

	startRow, endRow := 3*row, 3*(row+1)
	startCol, endCol := 3*col, 3*(col+1)
	for r := startRow; r < endRow; r++ {
		for c := startCol; c < endCol; c++ {
			win.Addch(c, r, ' ', color)
		}
	}

	if row == b.currentRow && col == b.currentCol {
		win.Addch(startCol+1, startRow+1, 'X', color)
	}

	return nil
}
