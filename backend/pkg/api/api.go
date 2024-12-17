package api

import "net/http"

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
