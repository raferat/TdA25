package api

import (
	"encoding/json"
	"net/http"

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
