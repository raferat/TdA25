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
    <html>
      <head>
        <meta http-equiv="Refresh" content="0; URL=https://www.youtube.com/watch?v=dQw4w9WgXcQ" />
      </head>
      <body>
        Hello TdA
      </body>
    </html>`)
  })

  server.HandleFunc("GET /api", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Content-Type", "text/json")
    w.WriteHeader(http.StatusOK)
    fmt.Fprintln(w, `{"organization": "Student Cyber Games"}`)
  })


  if err := http.ListenAndServe(portStr, server); err != nil {
    log.Fatalln(err)
  }
}
