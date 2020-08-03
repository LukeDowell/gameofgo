package repo

import (
	"dowell.dev/gameofgo/game"
	"errors"
	"github.com/google/uuid"
)

type BoardEntity struct {
	*game.Board

	Id uuid.UUID
}

type GameRepository struct {
	Boards map[uuid.UUID]*game.Board
}

// New is a factory method to create a new GameRepository
func New() *GameRepository {
	return &GameRepository{
		Boards: make(map[uuid.UUID]*game.Board),
	}
}

// Save persists a board to this repository's store
func (r *GameRepository) Save(entity BoardEntity) error {
	if r.Boards[entity.Id] != nil {
		return errors.New("entity already exists with provided id")
	}

	r.Boards[entity.Id] = entity.Board

	return nil
}

// FindById attempts to retrieve a board from this repository. If the
// board isn't found, an error is returned
func (r *GameRepository) FindById(id uuid.UUID) (board *game.Board, err error) {
	if r.Boards[id] == nil {
		return nil, errors.New("board with id " + id.String() + " not found")
	}

	return r.Boards[id], nil
}
