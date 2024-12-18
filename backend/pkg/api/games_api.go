package api

import (
	"net/http"

	"tdaserver/pkg/db"
)

func CreateGame(base *db.GameBase, r *http.Request) (int, any) {
	game := db.CreateGame(base)
	if game == nil {
		return http.StatusInternalServerError, NewInternalServerError()
	}

	return http.StatusCreated, game
}

func ListGames(r *http.Request) (int, any) {
	games := db.GameList()

	return http.StatusOK, games
}

func FindGame(r *http.Request) (int, any) {
	uuid := r.PathValue("uuid")
	game := db.FindGame(uuid)

	if game == nil {
		return http.StatusNotFound, NewNotFoundError()
	}

	return http.StatusOK, game
}

func UpdateGame(base *db.GameBase, r *http.Request) (int, any) {
	uuid := r.PathValue("uuid")

	game := db.UpdateGame(uuid, base)

	if game == nil {
		return http.StatusNotFound, NewNotFoundError()
	}

	return http.StatusOK, game
}

func DeleteGame(r *http.Request) (int, any) {
	uuid := r.PathValue("uuid")

	if db.DeleteGame(uuid) {
		return http.StatusNoContent, nil
	}

	return http.StatusNotFound, NewNotFoundError()
}
