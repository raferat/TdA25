package main

import (
	"fmt"
	"net/http"

	"tdaserver/pkg/api"
	"tdaserver/pkg/utils"
)

func router(mux *http.ServeMux) http.Handler {
	HowDoesItWork(mux)

	mux.Handle("GET /", utils.NotFoundMiddleware(http.FileServer(http.Dir("./static"))))

	mux.Handle("POST /api/v1/games", utils.JSONEncodeMiddleware(api.CreateGame))
	mux.Handle("GET /api/v1/games", utils.JSONEncodeMiddleware(api.ListGames))
	mux.Handle("GET /api/v1/games/{uuid}", utils.JSONEncodeMiddleware(api.FindGame))
	mux.Handle("PUT /api/v1/games/{uuid}", utils.JSONEncodeMiddleware(api.UpdateGame))
	mux.Handle("DELETE /api/v1/games/{uuid}", utils.JSONEncodeMiddleware(api.DeleteGame))

	return utils.LoggingMiddleware(mux)
}

func HowDoesItWork(mux *http.ServeMux) {
	mux.HandleFunc("GET /api", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{"organization": "Student Cyber Games"}`)
	})
}
