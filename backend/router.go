package main

import (
	"net/http"
	"os"
	"path/filepath"

	"tdaserver/pkg/api"
	"tdaserver/pkg/utils"
)

func router(mux *http.ServeMux) http.Handler {
	mux.Handle("GET /api", utils.JSONEncodeMiddleware(func(r *http.Request) (int, any) {
		return http.StatusOK, map[string]string{"organization": "Student Cyber Games"}
	}))
	mux.Handle("/api/", http.StripPrefix("/api", registerApiRoutes()))
	mux.Handle("/", fileRouter("./website", "fallback.html"))

	return utils.LoggingMiddleware(mux)
}

func registerApiRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /", utils.JSONEncodeMiddleware(func(r *http.Request) (int, any) {
		return http.StatusOK, map[string]string{"organization": "Student Cyber Games"}
	}))

	mux.Handle("POST /v1/games",
		utils.JSONEncodeMiddleware(
			api.ParseGameMiddleware(api.CreateGame),
		))

	mux.Handle("GET /v1/games", utils.JSONEncodeMiddleware(api.ListGames))
	mux.Handle("GET /v1/games/{uuid}", utils.JSONEncodeMiddleware(api.FindGame))

	mux.Handle("PUT /v1/games/{uuid}",
		utils.JSONEncodeMiddleware(
			api.ParseGameMiddleware(api.UpdateGame),
		))

	mux.Handle("DELETE /v1/games/{uuid}", utils.JSONEncodeMiddleware(api.DeleteGame))

	return mux
}

func fileRouter(static string, fallback string) http.Handler {
	fileServer := http.FileServer(http.Dir(static))
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		filePath := filepath.Join(static, r.URL.Path)

		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			http.ServeFile(w, r, filepath.Join(static, fallback))
			return
		}

		fileServer.ServeHTTP(w, r)
	})
}
