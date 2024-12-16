package main

import (
	"fmt"
	"net/http"
)

func router(handler *http.ServeMux) {
	HowDoesItWork(handler)
}

func HowDoesItWork(handler *http.ServeMux) {
	handler.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>HTML5 Boilerplate</title>
</head>

<body>
  <h1>Page Title</h1>
  <p>Hello TdA</p>
</body>

</html>`)
	})

	handler.HandleFunc("GET /api", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{"organization": "Student Cyber Games"}`)
	})
}
