package db

import (
	"log"
	"time"

	"tdaserver/pkg/utils"

	"github.com/google/uuid"
)

func CreateGame(base *GameBase) *Game {
	now := time.Now()
	result := &Game{
		GameBase:  *base,
		Uuid:      uuid.NewString(),
		CreatedAt: now,
		UpdatedAt: now,
		GameState: utils.CalculateGameState(base.Board),
	}

	_, err := createGame.Exec(result.Uuid, result.CreatedAt, result.UpdatedAt, result.Name, result.Difficulty, result.GameState, utils.SerializeBoard(result.Board))
	if err != nil {
		log.Printf("Error creating game: %#v\n", err)
		return nil
	}

	return result
}
