package main

import (
	"encoding/json"
	"net/http"

	"tdaserver/api/v1"

	"github.com/jackc/pgx/v5/pgxpool"
)

func loadRoutes(router *http.ServeMux, db *pgxpool.Pool) {
	router.HandleFunc("GET /api", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{
			"organization": "Student Cyber Games",
		})
	})

	apiv1handler := api.Handler{
		DB: db,
	}

	router.Handle("GET /api/v1/games", api.ErrorHandler(apiv1handler.ListGames))
	router.HandleFunc("POST /api/v1/games", apiv1handler.CreateNewGame)
	router.Handle("GET /api/v1/games/{uuid}", api.ErrorHandler(apiv1handler.GetGameByUUID))
	router.HandleFunc("PUT /api/v1/games/{uuid}", apiv1handler.UpdateGame)
	router.HandleFunc("DELETE /api/v1/games/{uuid}", apiv1handler.DeleteGame)
}
