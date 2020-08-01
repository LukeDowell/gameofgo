package api

import (
	"dowell.dev/gameofgo/game"
	"dowell.dev/gameofgo/repo"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func SetupRouter(repo *repo.GameRepository) *gin.Engine {
	router := gin.Default()

	router.POST("/game", func(c *gin.Context) {
		c.JSON(201, GameResponse{
			GameId: uuid.New(),
		})
	})

	return router
}

type GameResponse struct {
	GameId uuid.UUID `json:"gameId"`
	Cells []game.Cell `json:"cells"`
}
