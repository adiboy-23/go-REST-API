package main

import (
	"context"
	//"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/adiboy-23/go-REST-API/internal/config"
)

func main() {
	//load config
	cfg := config.MustLoad()

	//database setup

	//setup router
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) { // ther is a request pointer not object
		w.Write([]byte("Welcome to Student API"))
	})

	//setup http server
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	slog.Info("Starting server at %s", slog.String("address", cfg.Addr))

	//now use channels for sync
	done := make(chan os.Signal, 1) //buffer size=1
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	//graceful shutdown
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Fail to start server")
		}
	}()

	<-done // till any ctrl-C : signal interupt is not pressed , we will not pass this statement execution

	slog.Info("Shutting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		slog.Error("Failed to shutdown server", slog.String("error", err.Error()))
	}

	slog.Info("Server shutdown successfully")

}
