package repo

import (
	"dowell.dev/gameofgo/game"
	"github.com/google/uuid"
)

type BoardEntity struct {
	game.Board

	Id uuid.UUID
}

type GameRepository struct {
	Boards *[]BoardEntity
}

// Save persists a board to some form of database
func (r GameRepository) Save(entity BoardEntity) (result BoardEntity, err error) {
	return BoardEntity{}, nil
}
