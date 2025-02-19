package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"tdaserver/pkg/db"
	"tdaserver/pkg/utils"
)

type ApiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewBadRequestError(msg string) ApiError {
	return ApiError{
		Code:    http.StatusBadRequest,
		Message: "Bad request: " + msg,
	}
}

func NewNotFoundError() ApiError {
	return ApiError{
		Code:    http.StatusNotFound,
		Message: "Resource not found",
	}
}

func NewUnprocessableEntityError(msg string) ApiError {
	return ApiError{
		Code:    http.StatusUnprocessableEntity,
		Message: "Semantic error: " + msg,
	}
}

func NewInternalServerError() ApiError {
	return ApiError{
		Code:    http.StatusInternalServerError,
		Message: "Internal server error",
	}
}

func ParseGameMiddleware(handler func(*db.GameBase, *http.Request) (int, any)) func(*http.Request) (int, any) {
	return func(r *http.Request) (int, any) {
		dec := json.NewDecoder(r.Body)
		var base db.GameBase
		err := dec.Decode(&base)
		if err != nil {
			return http.StatusBadRequest, NewBadRequestError(err.Error())
		}

		if ok, msg := utils.CheckIntegrity(base.Board); !ok {
			return http.StatusUnprocessableEntity, NewUnprocessableEntityError(msg)
		}

		if ok := utils.CheckDifficulty(base.Difficulty); !ok {
			return http.StatusUnprocessableEntity, NewUnprocessableEntityError("Difficulty value not supported.")
		}

		return handler(&base, r)
	}
}

func ParseUserMiddleware(handler func(*db.UserBase, *http.Request) (int, any)) func(*http.Request) (int, any) {
	return func(r *http.Request) (int, any) {
		dec := json.NewDecoder(r.Body)
		var base db.UserBase
		err := dec.Decode(&base)
		if err != nil {
			return http.StatusBadRequest, NewBadRequestError(err.Error())
		}

		if len(base.Password) < 8 {
			return http.StatusBadRequest, NewBadRequestError("Password must be at least 8 characters in length.")
		}

		if !strings.ContainsAny(base.Password, "abcdefghijklmnopqrstuvwxyz") {
			return http.StatusBadRequest, NewBadRequestError("Password must contain at least 1 lowercase letter.")
		}

		if !strings.ContainsAny(base.Password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
			return http.StatusBadRequest, NewBadRequestError("Password must contain at least 1 uppercase letter.")
		}

		if !strings.ContainsAny(base.Password, "0123456789") {
			return http.StatusBadRequest, NewBadRequestError("Password must contain at least 1 number")
		}

		if !strings.ContainsFunc(base.Password, func(r rune) bool {
			return (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') && (r < '0' || r > '9')
		}) {
			return http.StatusBadRequest, NewBadRequestError("Password must contain at least 1 special character.")
		}

		return handler(&base, r)
	}
}
