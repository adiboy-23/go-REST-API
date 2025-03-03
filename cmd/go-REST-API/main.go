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
	"github.com/adiboy-23/go-REST-API/internal/http/handlers/student"
	"github.com/adiboy-23/go-REST-API/internal/storage/sqlite"
)

func main() {
	//load config
	cfg := config.MustLoad()

	//database setup
	storage, err := sqlite.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	slog.Info("Storage initalized", slog.String("env", cfg.Env), slog.String("version", "1.0.0"))

	//setup router
	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.New(storage))

	router.HandleFunc("Get /api/students/{id}", student.GetById(storage))

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

	err = server.Shutdown(ctx)

	if err != nil {
		slog.Error("Failed to shutdown server", slog.String("error", err.Error()))
	}

	slog.Info("Server shutdown successfully")

}
