package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"tdaserver/pkg/api"
	"tdaserver/pkg/utils"
)

func router(mux *http.ServeMux) http.Handler {
	mux.Handle("/realtime/freeplay", http.HandlerFunc(api.HandleFreeplayGame))
	mux.Handle("GET /api", utils.LoggingMiddleware(utils.CompressionMiddleware(utils.JSONEncodeMiddleware(func(r *http.Request) (int, any) {
		return http.StatusOK, map[string]string{"organization": "Student Cyber Games"}
	}))))
	mux.Handle("/api/", utils.LoggingMiddleware(utils.CompressionMiddleware(http.StripPrefix("/api", registerApiRoutes()))))
	mux.Handle("/", utils.LoggingMiddleware(utils.CompressionMiddleware(fileRouter("./website", "fallback.html"))))

	return mux
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

	//================== USERS ===================

	mux.Handle("POST /v1/users", utils.JSONEncodeMiddleware(api.ParseUserMiddleware(api.CreateUser)))
	mux.Handle("GET /v1/users", utils.JSONEncodeMiddleware(api.ListUsers))
	mux.Handle("GET /v1/users/{uuid}", utils.JSONEncodeMiddleware(api.FindUser))
	mux.Handle("PUT /v1/users/{uuid}", utils.JSONEncodeMiddleware(api.ParseUserMiddleware(api.UpdateUser)))
	mux.Handle("DELETE /v1/users/{uuid}", utils.JSONEncodeMiddleware(api.DeleteUser))

	//================== REALTIME ================

	mux.Handle("POST /realtime/login", http.HandlerFunc(api.Login))
	// mux.Handle("GET /realtime/freeplay", http.HandlerFunc(api.HandleFreeplayGame))

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

		if time := shouldBeCached(filePath); time != 0 {
			w.Header().Add("Cache-Control", fmt.Sprintf("max-age=%d", time))
		}

		fileServer.ServeHTTP(w, r)
	})
}

func shouldBeCached(path string) int {
	if strings.HasSuffix(path, ".ttf") || strings.HasSuffix(path, ".svg") || strings.HasSuffix(path, ".png") {
		return 31536000
	}

	if strings.HasSuffix(path, ".js") {
		return 86400
	}

	return 3600
}
