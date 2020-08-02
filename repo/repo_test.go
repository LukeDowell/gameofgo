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
		repository = New()
	})

	It("should save a board entity", func() {
		id := uuid.New()
		entity := BoardEntity{
			Board: &game.Board{},
			Id:    id,
		}

		err := repository.Save(entity)

		Expect(err).To(BeNil())
		Expect(repository.Boards[id]).To(Equal(entity.Board))
	})
})
