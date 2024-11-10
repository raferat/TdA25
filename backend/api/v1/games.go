package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

  "tdaserver/utils/gameutils"

	"github.com/jackc/pgx/v5/pgtype"
)

type dbscanner interface {
	Scan(dest ...any) error
}

func convertPostgresTo2DArray[T any](pg pgtype.Array[T]) [][]T {
  board_arr := make([][]T, 0, pg.Dimensions()[0].Length)
	for i := int32(0); i < pg.Dimensions()[0].Length*pg.Dimensions()[1].Length; i += pg.Dimensions()[1].Length {
		board_arr = append(board_arr, pg.Elements[i:i+pg.Dimensions()[1].Length])
	}

  return board_arr
}

func scanGameRow[T dbscanner](row T) (*gameutils.GameInfo, error) {
	var (
		id         string
		name       string
		difficulty string
		state      string
		board      pgtype.Array[string]
		created    time.Time
		updated    time.Time
	)

	err := row.Scan(&id, &name, &difficulty, &state, &board, &created, &updated)
	if err != nil {
		return nil, err
	}

  board_arr := convertPostgresTo2DArray(board)

	return &gameutils.GameInfo{
    MandatoryGameInfo: gameutils.MandatoryGameInfo{
      Name: name,
		  Difficulty: difficulty,
		  Board:      board_arr,
    },
		Id:         id,
		State:      state,
		Created:    created,
		Updated:    updated,
	}, nil
}

func convert2DArrayToPostgres[T any](board [][]T) pgtype.Array[T] {
  board_array := pgtype.Array[T]{}
  board_array.SetDimensions([]pgtype.ArrayDimension{
    {Length: 15, LowerBound: 1},
    {Length: 15, LowerBound: 1},
  })

  for i := 0; i < len(board); i++ {
    width := len(board[i])
    for j := 0; j < width; j++ {
      ptr, ok := board_array.ScanIndex(i*width + j).(*T)
      if ok {
        *ptr = board[i][j]
      }
    }
  }
  
  for _, v := range board {
    board_array.Elements = append(board_array.Elements, v...)
  }

  return board_array
}

func (h *Handler) ListGames(w http.ResponseWriter, r *http.Request) *ApiError {
  w.Header().Add("Content-Type", "text/json")

  //sql query
	rows, err := h.DB.Query(r.Context(), "SELECT * FROM games")

  //kdyz je sql spatne
	if err != nil {
		return NewApiError(http.StatusInternalServerError, "", "Could not retrieve games: "+err.Error())
	}
	defer rows.Close()

  //buffer s odpovedi
	buf := bytes.NewBufferString("")

	encoder := json.NewEncoder(buf)

  //zacatek json array
	fmt.Fprint(buf, "[")
	first := true

	for rows.Next() {
		if !first {
			fmt.Fprint(buf, ",") // carky v json, neni na konci (trailing comma)
		} else {
			first = false
		}

		game, err := scanGameRow(rows)
		if err != nil {
			return NewApiError(http.StatusInternalServerError, "", err.Error())
		}

		encoder.Encode(game)
	}

  //konec JSON array
	fmt.Fprintln(buf, "]")
	w.Write(buf.Bytes())
	return nil
}

func (h *Handler) CreateNewGame(w http.ResponseWriter, r *http.Request) *ApiError {
  mgi := &gameutils.MandatoryGameInfo{}
  json.NewDecoder(r.Body).Decode(mgi)

  /*mgi := &gameutils.MandatoryGameInfo{
    Name: "Test from GO",
    Difficulty: "beginner",
    Board: [][]string {
      {"", "X", "O", "X", "O", "X", "O", "", "", "X", "O", "X", "", "O", "X"},
      {"", "X", "O", "X", "O", "X", "O", "", "", "X", "O", "X", "", "O", "X"},
      {"", "X", "O", "X", "O", "X", "O", "", "", "X", "O", "X", "", "O", "X"},
      {"", "X", "O", "X", "O", "X", "O", "", "", "X", "O", "X", "", "O", "X"},
      {"", "X", "O", "X", "O", "X", "O", "", "", "X", "O", "X", "", "O", "X"},
      {"", "X", "O", "X", "O", "X", "O", "", "", "X", "O", "X", "", "O", "X"},
      {"", "X", "O", "X", "O", "X", "O", "", "", "X", "O", "X", "", "O", "X"},
      {"", "X", "O", "X", "O", "X", "O", "", "", "X", "O", "X", "", "O", "X"},
      {"", "X", "O", "X", "O", "X", "O", "", "", "X", "O", "X", "", "O", "X"},
      {"", "X", "O", "X", "O", "X", "O", "", "", "X", "O", "X", "", "O", "X"},
      {"", "X", "O", "X", "O", "X", "O", "", "", "X", "O", "X", "", "O", "X"},
      {"", "X", "O", "X", "O", "X", "O", "", "", "X", "O", "X", "", "O", "X"},
      {"", "X", "O", "X", "O", "X", "O", "", "", "X", "O", "X", "", "O", "X"},
      {"", "X", "O", "X", "O", "X", "O", "", "", "X", "O", "X", "", "O", "X"},
      {"", "X", "O", "X", "O", "X", "O", "", "", "X", "O", "X", "", "O", "X"},
    },
  }*/

  ok, field := mgi.HasEmptyField()
  if !ok {
    return NewApiErrorResponse(http.StatusBadRequest, fmt.Sprintf("Bad request: field `%s` is required", field))
  }
  _, err := h.DB.Exec(r.Context(), "INSERT INTO games (name, game_difficulty, game_state, board) VALUES ($1, $2, $3, $4)",
    mgi.Name, mgi.Difficulty, "unknown", convert2DArrayToPostgres(mgi.Board),
  )

  if err != nil {
    return NewApiError(http.StatusInternalServerError, "", "Error creating new game: " + err.Error())
  }

  return nil
}

func (h *Handler) GetGameByUUID(w http.ResponseWriter, r *http.Request) *ApiError {
  w.Header().Add("Content-Type", "text/json")

	rows, err := h.DB.Query(r.Context(), "SELECT * FROM games WHERE game_id = $1", r.PathValue("uuid"))

  //Query obsahuje syntaktickou chybu
	if err != nil {
		return NewApiError(http.StatusInternalServerError, "", err.Error())
	}
	defer rows.Close()

  //neexistuji hledane hry
	if !rows.Next() {
		return NewApiErrorResponse(http.StatusNotFound, "Resource not found")
	}

  //scanovani z dat z databaze a err
	gamedata, err := scanGameRow(rows)
	if err != nil {
		return NewApiError(http.StatusInternalServerError, "", err.Error())
	}

	json.NewEncoder(w).Encode(gamedata)
	return nil
}

func (h *Handler) UpdateGame(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "update"+r.PathValue("uuid"))
}

func (h *Handler) DeleteGame(w http.ResponseWriter, r *http.Request) *ApiError {
  tag, err := h.DB.Exec(r.Context(), "DELETE FROM games WHERE game_id = $1", r.PathValue("uuid"))
  if err != nil {
    return NewApiError(http.StatusInternalServerError, "", err.Error())
  }

  if tag.RowsAffected() == 0 {
    return NewApiErrorResponse(http.StatusNotFound, "Resource not found")
  }

  w.WriteHeader(http.StatusNoContent)
  return nil
}
