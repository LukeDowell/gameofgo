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

func New() *GameRepository {
	return &GameRepository{
		Boards: make(map[uuid.UUID]*game.Board),
	}
}

// Save persists a board to some form of database
func (r *GameRepository) Save(entity BoardEntity) error {
	if r.Boards[entity.Id] != nil {
		return errors.New("entity already exists with provided id")
	}

	r.Boards[entity.Id] = entity.Board

	return nil
}
