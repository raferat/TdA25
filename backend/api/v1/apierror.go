package api

import (
	"encoding/json"
	"net/http"

	"github.com/fatih/color"
)

type ApiError struct {
	Code           int    `json:"code"`
	Reason         string `json:"message"`
	InternalReason string `json:"-"`
	Log            bool   `json:"-"`
}

func NewApiError(code int, msg string, internal string) *ApiError {
	return &ApiError{
		Code:           code,
		Reason:         msg,
		InternalReason: internal,
		Log:            true,
	}
}

func NewApiErrorResponse(code int, msg string) *ApiError {
	return &ApiError{
		Code:   code,
		Reason: msg,
	}
}

type ErrorApiFunc func(w http.ResponseWriter, r *http.Request) *ApiError

func ErrorHandler(errApi ErrorApiFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := errApi(w, r)
		if err != nil {
      if err.Reason == "" {
    		err.Reason = http.StatusText(err.Code)
    	}

			w.WriteHeader(err.Code)
			json.NewEncoder(w).Encode(err)


			logError(err)
		}
	})
}

func logError(err *ApiError) {
	if err.Log {
		color.HiRed("ERROR:\tCode:\t\t%03d\n\tReason:\t\t%s\n\tInternal:\t%s\n", err.Code, err.Reason, err.InternalReason)
	}
}
