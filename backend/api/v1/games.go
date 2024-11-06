package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func (h *Handler) ListGames(w http.ResponseWriter, r *http.Request) *ApiError {
	rows, err := h.DB.Query(r.Context(), "SELECT * FROM games")
	if err != nil {
		return NewApiError(http.StatusInternalServerError,
			"Internal server error, try again later.",
			"Could not retrieve games: "+err.Error())
	}
	defer rows.Close()

	buf := bytes.NewBufferString("")

	encoder := json.NewEncoder(buf)

	fmt.Fprint(buf, "[")
	first := true

	for rows.Next() {
		var (
			id         string
			name       string
			difficulty string
			state      string
			board      pgtype.Array[string]
			created    time.Time
			updated    time.Time
		)

		if !first {
			fmt.Fprint(buf, ",")
		} else {
			first = false
		}

		err = rows.Scan(&id, &name, &difficulty, &state, &board, &created, &updated)
		if err != nil {
			return NewApiError(http.StatusInternalServerError,
				"Internal server error - contact administrator",
				err.Error())
		}

		board_arr := make([][]string, 0, board.Dimensions()[0].Length)
		for i := int32(0); i < board.Dimensions()[0].Length*board.Dimensions()[1].Length; i += board.Dimensions()[1].Length {
			board_arr = append(board_arr, board.Elements[i:i+board.Dimensions()[1].Length])
		}
		encoder.Encode(map[string]any{
			"uuid":       id,
			"createdAt":  created.UTC().Format(time.RFC3339),
			"updatedAt":  updated.UTC().Format(time.RFC3339),
			"name":       name,
			"difficulty": difficulty,
			"gameState":  state,
			"board":      board_arr,
		})
	}
	fmt.Fprintln(buf, "]")
	w.Write(buf.Bytes())
	return nil
}

func (h *Handler) CreateNewGame(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Create")
}

func (h *Handler) GetGameByUUID(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "get"+r.PathValue("uuid"))
}

func (h *Handler) UpdateGame(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "update"+r.PathValue("uuid"))
}

func (h *Handler) DeleteGame(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "delete"+r.PathValue("uuid"))
}
