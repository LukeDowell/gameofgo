package game

type Cell struct {
	X int
	Y int
}

type Board struct {
	Cells []Cell
}

// Step moves the game forward one "step" in time. A new board
// is returned representing the next logical state of the game.
func (b Board) Step() Board {
	var next []Cell
	for _, c := range b.Cells {
		neighbors := b.LiveNeighborsFor(c)

		if len(neighbors) < 4 && len(neighbors) > 1 {
			next = append(next, c)
		}

		allNeighbors := []Cell{
			{c.X - 1, c.Y - 1},
			{c.X, c.Y - 1},
			{c.X + 1, c.Y - 1},
			{c.X - 1, c.Y},
			{c.X + 1, c.Y},
			{c.X - 1, c.Y + 1},
			{c.X, c.Y + 1},
			{c.X + 1, c.Y + 1},
		}

		for _, n := range allNeighbors {
			if !b.Contains(n) && len(b.LiveNeighborsFor(n)) == 3 {
				next = append(next, n)
			}
		}
	}

	return Board{Cells: next}
}

// LiveNeighborsFor calculates the number of live neighbors
// that exist for the provided cell on this board
func (b Board) LiveNeighborsFor(c Cell) []Cell {
	var neighbors []Cell
	for _, v := range b.Cells {
		if v.X == c.X && v.Y == c.Y {
			continue
		}

		if v.X <= c.X + 1 && v.X >= c.X -1 && v.Y <= c.Y + 1 && v.Y >= c.Y - 1 {
			neighbors = append(neighbors, v)
		}
	}
	return neighbors
}

// Contains returns true if the provided cell exists on this board
func (b Board) Contains(c Cell) bool {
	for _, v := range b.Cells {
		if v.X == c.X && v.Y == c.Y {
			return true
		}
	}
	return false
}