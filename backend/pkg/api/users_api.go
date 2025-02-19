package api

import (
	"net/http"

	"tdaserver/pkg/db"
)

func CreateUser(base *db.UserBase, r *http.Request) (int, any) {
	user := db.CreateUser(base)
	if user == nil {
		return http.StatusInternalServerError, NewInternalServerError()
	}

	return http.StatusCreated, user
}

func ListUsers(r *http.Request) (int, any) {
	users := db.ListUsers()

	return http.StatusOK, users
}

func FindUser(r *http.Request) (int, any) {
	uuid := r.PathValue("uuid")
	user := db.GetUserByUUID(uuid)

	if user == nil {
		return http.StatusNotFound, NewNotFoundError()
	}

	return http.StatusOK, user
}

func UpdateUser(base *db.UserBase, r *http.Request) (int, any) {
	uuid := r.PathValue("uuid")

	user := db.UpdateUser(uuid, base)

	if user == nil {
		return http.StatusNotFound, NewNotFoundError()
	}

	return http.StatusOK, user
}

func DeleteUser(r *http.Request) (int, any) {
	uuid := r.PathValue("uuid")

	if db.DeleteUser(uuid) {
		return http.StatusNoContent, nil
	}

	return http.StatusNotFound, NewNotFoundError()
}
