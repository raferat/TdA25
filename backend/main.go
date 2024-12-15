package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
  server := http.NewServeMux()
  portStr := ":4242"
  if len(os.Args) == 2 {
    portStr = os.Args[1]
  }

  server.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
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

  server.HandleFunc("GET /api", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    fmt.Fprintln(w, `{"organization": "Student Cyber Games"}`)
  })


  if err := http.ListenAndServe(portStr, server); err != nil {
    log.Fatalln(err)
  }
}
