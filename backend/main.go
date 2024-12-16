package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"tdaserver/pkg/db"
)

func gracefulShutdown(server *http.Server) context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	signalChan := make(chan os.Signal, 5)
	signal.Notify(signalChan, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGSEGV, syscall.SIGTERM)

	go func() {
		sig := <-signalChan
		log.Printf("Graceful shutdown: %s\n", sig.String())
		server.Shutdown(ctx)
		cancel()
	}()

	return ctx
}

func initServer() *http.Server {
	handler := http.NewServeMux()
	portStr := ":4242"
	if val, ok := os.LookupEnv("PORT"); ok {
		portStr = val
	}

	router(handler)

	return &http.Server{
		Addr:    portStr,
		Handler: handler,
	}
}

func main() {
	server := initServer()
	ctx := gracefulShutdown(server)
	_, end := db.InitDB(ctx)

	log.Printf("Server running on: %s\n", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
	<-end.Done()
}
