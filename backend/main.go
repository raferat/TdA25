package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
	// "github.com/jackc/pgx/v5/pgxpool"
)

func main() {

  db, err := pgx.Connect(context.Background(), "postgres://goserver:server@127.0.0.1:5432/tda25")

  //db, err := pgxpool.New(context.Background(), "postgres://goserver:server@127.0.0.1:5432/tda25")
  if err != nil {
    log.Fatalf("Unable to connect to DB: %+v\n", err)
  } else {
    log.Println("Successful connection to DB.")
  }
  defer db.Close(context.Background());


  http.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "pong")
  })

  log.Fatal(http.ListenAndServe(":3000", nil))
}
