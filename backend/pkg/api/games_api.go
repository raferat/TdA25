package api

import (
	"encoding/json"
	"net/http"

	"tdaserver/pkg/db"
	"tdaserver/pkg/utils"
)

func CreateGame(r *http.Request) (int, any) {
	dec := json.NewDecoder(r.Body)
	var base db.GameBase
	err := dec.Decode(&base)
	if err != nil {
		return http.StatusBadRequest, NewBadRequestError(err.Error())
	}

	if ok, msg := utils.CheckIntegrity(base.Board); !ok {
		return http.StatusUnprocessableEntity, NewUnprocessableEntityError(msg)
	}

	game := db.CreateGame(&base)
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

func UpdateGame(r *http.Request) (int, any) {
	uuid := r.PathValue("uuid")
	dec := json.NewDecoder(r.Body)

	var base db.GameBase
	err := dec.Decode(&base)
	if err != nil {
		return http.StatusBadRequest, NewBadRequestError(err.Error())
	}

	if ok, msg := utils.CheckIntegrity(base.Board); !ok {
		return http.StatusUnprocessableEntity, NewUnprocessableEntityError(msg)
	}

	game := db.UpdateGame(uuid, &base)

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
