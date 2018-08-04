package ox

import (
	"errors"
	"math"
)

// SIZE of board
const SIZE = 3

// Game structure
type Game struct {
	CBoard Board
}

// Board structure of game
type Board struct {
	Size     int
	Status   [SIZE * SIZE]string
	Row      [SIZE]int
	Column   [SIZE]int
	Diagonal [2]int
}

// NewGame :: Create new game
func NewGame() Game {
	CBoard := Board{Size: SIZE * SIZE}
	return Game{CBoard: CBoard}
}

// CurrentBoard get current status of board
func (g *Game) CurrentBoard() Board {
	return g.CBoard
}

// IsWinner game have winner or not
func (g *Game) IsWinner() bool {
	for i := range g.CBoard.Row {
		if math.Abs(float64(g.CBoard.Row[i])) == SIZE {
			return true
		}
	}
	for i := range g.CBoard.Column {
		if math.Abs(float64(g.CBoard.Column[i])) == SIZE {
			return true
		}
	}
	for i := range g.CBoard.Diagonal {
		if math.Abs(float64(g.CBoard.Diagonal[i])) == SIZE {
			return true
		}
	}
	return false
}

// CurrentStatus of game
func (g *Game) CurrentStatus() string {
	for _, s := range g.CBoard.Status {
		if s == "" {
			return "playing"
		}
	}
	return "finish"
}

// Play with game
func (g *Game) Play(value string, position int) error {
	if g.CBoard.Status[position-1] != "" {
		return errors.New("")
	}
	g.CBoard.Status[position-1] = value

	// Row
	i := (position - 1) / SIZE
	if value == "O" {
		g.CBoard.Row[i]++
	} else {
		g.CBoard.Row[i]--
	}

	// Column
	ic := (position - 1) % SIZE
	if value == "O" {
		g.CBoard.Column[ic]++
	} else {
		g.CBoard.Column[ic]--
	}

	// Diagonal
	if (position-1)%(SIZE+1) == 0 {
		if value == "O" {
			g.CBoard.Diagonal[0]++
		} else {
			g.CBoard.Diagonal[0]--
		}
	}
	if (position-1)%(SIZE-1) == 0 {
		if value == "O" {
			g.CBoard.Diagonal[1]++
		} else {
			g.CBoard.Diagonal[1]--
		}
	}
	return nil
}
