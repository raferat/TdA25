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

	end := func() {
		server.Shutdown(ctx)
		cancel()
	}

	go func() {
		sig := <-signalChan
		log.Printf("Graceful shutdown: %s\n", sig.String())
		end()
	}()

	return ctx
}

func initServer() *http.Server {
	mux := http.NewServeMux()
	portStr := ":4242"
	if val, ok := os.LookupEnv("PORT"); ok {
		portStr = val
	}

	handler := router(mux)

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
