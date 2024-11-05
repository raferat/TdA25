package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/jackc/pgx/v5/pgxpool"
)

func initServer(router *http.ServeMux) *http.Server {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":4242"
	}

  server := &http.Server{
    Addr: port,
    Handler: router,
  }

  return server
}

func initDBConnection() *pgxpool.Pool {
	db, err := pgxpool.New(context.Background(), "postgres://goserver:server@127.0.0.1:5432/tda25")

  if err != nil {
    log.Fatalf("Unable to create connection pool to DB: %+v\n", err)
  }

  err = db.Ping(context.Background())
	if err != nil {
		log.Fatalf("Unable to connect to DB: %+v\n", err)
	}
  log.Println("Successful connection to DB.")

  return db
}

func main() {
  //gracefull shutdown setup
  termsignals := make(chan os.Signal, 1)
  done := make(chan bool, 1)
  signal.Notify(termsignals, syscall.SIGINT, syscall.SIGTERM)

  //database connection
  db := initDBConnection()
  defer db.Close()

  //server and routes

  router := http.NewServeMux()
  server := initServer(router)
  loadRoutes(router, db)

  //gracefull shutdown
  go func () {
    sig := <- termsignals
    log.Println(sig)

    log.Println(server.Shutdown(context.Background()))
    done <- true
  } ()

  
  //start server and wait till shutdown
	log.Fatal(server.ListenAndServe())
  <- done
}
