package repo_test

import (
	"dowell.dev/gameofgo/game"
	. "dowell.dev/gameofgo/repo"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Repo", func() {
	var repository *GameRepository

	BeforeEach(func() {
		repository = &GameRepository{}
	})

	It("should save a board entity", func() {
		entity := BoardEntity{
			Board: game.Board{},
			Id:    uuid.New(),
		}

		result, err := repository.Save(entity)

		Expect(err).To(BeNil())
		Expect(result).To(Equal(entity))
		Expect(repository.Boards).To(ContainElement(entity))
	})
})
