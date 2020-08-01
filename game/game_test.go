package game_test

import (
	. "dowell.dev/gameofgo/game"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Game of life", func() {
	Describe("A live cell", func() {
		It("should die if there are fewer than two neighbors", func() {
			board := boardOf(Cell{X: 0, Y: 0})

			Expect(board.Step().Cells).To(BeEmpty())
		})

		It("should survive if there are two or three neighbors", func() {
			board := boardOf(
				Cell{X: 0, Y: 0},
				Cell{X: 1, Y: 0},
				Cell{X: 2, Y: 0},
			)

			Expect(board.Step().Cells).To(ContainElement(Cell{X: 1, Y: 0}))
		})

		It("should die if there are more than three neighbors", func() {
			board := boardOf(
				Cell{X: 0, Y: 0},
				Cell{X: 1, Y: 0},
				Cell{X: 1, Y: 1},
				Cell{X: 1, Y: -1},
				Cell{X: 2, Y: 0},
			)

			Expect(board.Step().Cells).NotTo(ContainElement(Cell{X: 1, Y: 0}))
		})
	})

	Describe("A dead cell", func() {
		It("should become a live cell if there are exactly three neighbors", func() {
			board := boardOf(
				Cell{X: 0, Y: 0},
				Cell{X: 1, Y: 1},
				Cell{X: 2, Y: 0},
			)

			Expect(board.Step().Cells).To(ContainElement(Cell{X: 1, Y: 0}))
		})
	})

	Describe("The game board", func() {
		It("should be able to determine whether or not a cell exists on the board", func() {
			board := boardOf(Cell{X: 0, Y: 0})

			Expect(board.Contains(Cell{X: 0, Y: 0})).To(BeTrue())
			Expect(board.Contains(Cell{X: 5, Y: 5})).To(BeFalse())
		})

		It("should be able to calculate the number of live neighbors a cell has", func() {
			board := boardOf(
				Cell{X: 0, Y: 0},
				Cell{X: 1, Y: 0},
				Cell{X: 2, Y: 0},
			)

			Expect(len(board.LiveNeighborsFor(Cell{X: 1, Y: 0}))).To(Equal(2))
			Expect(len(board.LiveNeighborsFor(Cell{X: 0, Y: 0}))).To(Equal(1))
			Expect(len(board.LiveNeighborsFor(Cell{X: 5, Y: 5}))).To(Equal(0))
		})
	})
})

func boardOf(cells ...Cell) Board {
	return Board{
		Cells: cells,
	}
}