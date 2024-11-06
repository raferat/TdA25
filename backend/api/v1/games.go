package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type gameinfo struct {
	Id         string     `json:"uuid"`
	Name       string     `json:"name"`
	Difficulty string     `json:"difficulty"`
	State      string     `json:"gameState"`
	Board      [][]string `json:"board"`
	Created    time.Time  `json:"createdAt"`
	Updated    time.Time  `json:"updatedAt"`
}

type dbscanner interface {
	Scan(dest ...any) error
}

func scanGameRow[T dbscanner](row T) (*gameinfo, error) {
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

	board_arr := make([][]string, 0, board.Dimensions()[0].Length)
	for i := int32(0); i < board.Dimensions()[0].Length*board.Dimensions()[1].Length; i += board.Dimensions()[1].Length {
		board_arr = append(board_arr, board.Elements[i:i+board.Dimensions()[1].Length])
	}

	return &gameinfo{
		Id:         id,
		Name:       name,
		Difficulty: difficulty,
		State:      state,
		Board:      board_arr,
		Created:    created,
		Updated:    updated,
	}, nil
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

func (h *Handler) CreateNewGame(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Create")
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

func (h *Handler) DeleteGame(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "delete"+r.PathValue("uuid"))
}
