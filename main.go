package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/ian-howell/gocurse/curses"

	"github.com/ian-howell/cleanup/boards"
)

const (
	EOF   = 4
	SPACE = 32
)

func main() {
	screen, err := curses.Initscr()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not initialize curses: %s\n", err.Error())
		os.Exit(1)
	}
	defer curses.Endwin()

	if err := Initialize(screen); err != nil {
		fmt.Fprintf(os.Stderr, "Could not setup common curses settings: %s\n", err.Error())
		os.Exit(1)
	}

	Play(screen)
}

func Initialize(screen *curses.Window) error {
	// Cause key presses to become immediately available
	// Raw is used here to capture all signals
	if err := curses.Raw(); err != nil {
		return err
	}
	// Suppress unnecessary echoing while taking input from the user
	if err := curses.Noecho(); err != nil {
		return err
	}
	// Enables the reading of function keys like F1, F2, arrow keys etc
	if err := screen.Keypad(true); err != nil {
		return err
	}
	// Make the cursor stop blinking
	if err := curses.Curs_set(0); err != nil {
		return err
	}
	if err := curses.Start_color(); err != nil {
		return err
	}
	curses.Init_pair(1, curses.COLOR_BLACK, curses.COLOR_RED)
	curses.Init_pair(2, curses.COLOR_BLACK, curses.COLOR_CYAN)
	return nil
}

func Play(screen *curses.Window) {
	b := boards.New(6, 6)
	Shuffle(&b)
	b.Draw(screen)

	for {
		switch ch := screen.Getch(); ch {
		case SPACE:
			b.Flip()
		case curses.KEY_UP, curses.KEY_DOWN, curses.KEY_LEFT, curses.KEY_RIGHT:
			b.Move(boards.Direction(ch))
		case EOF:
			return
		}
		b.Draw(screen)

		if b.IsSolved() {
			return
		}
	}
}

func Shuffle(b boards.Board) {
	// Just do a whole bunch of random moves
	keyMap := []int{
		curses.KEY_UP,
		curses.KEY_DOWN,
		curses.KEY_LEFT,
		curses.KEY_RIGHT,
	}
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < 5000; i++ {
		b.Move(boards.Direction(keyMap[r.Intn(4)]))
		b.Flip()
	}
}
