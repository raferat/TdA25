package db

import (
	"database/sql"
	"log"
	"time"

	"tdaserver/pkg/utils"

	"github.com/google/uuid"
)

func scanGameFromRow[T interface{ Scan(...any) error }](r T) (*Game, error) {
	elem := &Game{}
	var board []byte
	err := r.Scan(&elem.Uuid,
		&elem.CreatedAt.Time,
		&elem.UpdatedAt.Time,
		&elem.Name,
		&elem.Difficulty,
		&elem.GameState,
		&board,
	)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Printf("Error scanning game values: %#v\n", err)
		}
		return nil, err
	}

	elem.Board, err = utils.DeserializeBoard(board)
	if err != nil {
		log.Printf("Error deserialize board: %#v\n", err)
		return nil, err
	}

	return elem, nil
}

//================================================================= POST /api/v1/games

func CreateGame(base *GameBase) *Game {
	now := time.Now()
	result := &Game{
		GameBase:  *base,
		Uuid:      uuid.NewString(),
		CreatedAt: MyTime{Time: now},
		UpdatedAt: MyTime{Time: now},
		GameState: utils.CalculateGameState(base.Board),
	}

	_, err := createGame.Exec(result.Uuid, result.CreatedAt.Time, result.UpdatedAt.Time, result.Name, result.Difficulty, result.GameState, utils.SerializeBoard(result.Board))
	if err != nil {
		log.Printf("Error creating game: %#v\n", err)
		return nil
	}

	return result
}

//================================================================= GET /api/v1/games

func GameList() (result []*Game) {
	rows, err := getGames.Query()
	if err != nil {
		log.Printf("Error get game list: %#v\n", err)
		return result
	}
	defer rows.Close()

	for rows.Next() {
		elem, err := scanGameFromRow(rows)
		if err != nil {
			return result
		}

		result = append(result, elem)
	}

	return result
}

//================================================================= GET /api/v1/games/{uuid}

func FindGame(uuid string) *Game {
	row := findGame.QueryRow(uuid)
	result, err := scanGameFromRow(row)
	if err != nil {
		return nil
	}

	return result
}

//================================================================= PUT /api/v1/games/{uuid}

func UpdateGame(uuid string, base *GameBase) *Game {
	result := &Game{
		GameBase:  *base,
		UpdatedAt: MyTime{time.Now()},
		GameState: utils.CalculateGameState(base.Board),
		Uuid:      uuid,
	}

	row := editGame.QueryRow(result.UpdatedAt.Time,
		result.Name,
		result.Difficulty,
		result.GameState,
		utils.SerializeBoard(result.Board),
		result.Uuid,
	)

	err := row.Scan(&result.CreatedAt.Time)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Printf("Error scanning time after update: %#v\n", err)
		}
		return nil
	}

	return result
}

//================================================================= DELETE /api/v1/games/{uuid}

// return true if successful
func DeleteGame(uuid string) bool {
	res, err := deleteGame.Exec(uuid)
	if err != nil {
		log.Printf("Error deleting: %#v\n", err)
	}

	num, err := res.RowsAffected()
	if err != nil {
		log.Println("Toto nesmi nastat. Pokud toto nastalo, znamena to konec civilizace. Buh nam vsem pomahej.")
		panic(err)
	}

	return num == 1
}
