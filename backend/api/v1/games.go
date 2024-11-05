package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)



func (h *Handler) ListGames(w http.ResponseWriter, r *http.Request) {
  rows, err := h.DB.Query(r.Context(), "SELECT * FROM games")
  if err != nil {
    log.Println(err)
    return
  }
  defer rows.Close()

  w.WriteHeader(http.StatusOK)
  encoder := json.NewEncoder(w)
  
  fmt.Fprint(w, "[")
  first := true
  
  for rows.Next() {
    var (
      id string
      name string
      difficulty string
      state string
      board pgtype.Array[string]
      created time.Time
      updated time.Time
    )

    if !first {
      fmt.Fprint(w, ",")
    } else {
      first = false
    }

    err = rows.Scan(&id, &name, &difficulty, &state, &board, &created, &updated)

    board_arr := make([][]string, board.Dimensions()[0].Length)
    for i := int32(0); i < board.Dimensions()[0].Length * board.Dimensions()[1].Length; i += board.Dimensions()[1].Length {
      board_arr = append(board_arr, board.Elements[i:i+board.Dimensions()[1].Length])
    }

    fmt.Printf("%+v\n", board_arr)
    

    encoder.Encode(map[string]any{
      "uuid": id,
      "createdAt": created.UTC().Format(time.RFC3339),
      "updatedAt": updated.UTC().Format(time.RFC3339),
      "name": name,
      "difficulty": difficulty,
      "gameState": state,
      "board": board_arr,
    })

    if err != nil {
      log.Println(err)
      continue
    }

    //fmt.Printf("%s\n%s\n%s\n%s\n%+v\n%s\n%s\n", id, name, difficulty, state, board, created, updated)


  }

  fmt.Fprintln(w, "]")
}

func (h *Handler) CreateNewGame(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "Create")
}

func (h *Handler) GetGameByUUID(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "get" + r.PathValue("uuid"))
}

func (h *Handler) UpdateGame(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "update" + r.PathValue("uuid"))
}

func (h *Handler) DeleteGame(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "delete" + r.PathValue("uuid"))
}

