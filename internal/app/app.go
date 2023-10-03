package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/AZRV17/goWEB/internal/config"
	handler "github.com/AZRV17/goWEB/internal/controller/http"
	"github.com/AZRV17/goWEB/internal/repository"
	"github.com/AZRV17/goWEB/internal/service"
	"github.com/AZRV17/goWEB/pkg/db/psql"
)

func Run() {
	config, err := config.NewConfig("internal/config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.Postgres.User, config.Postgres.Password, config.Postgres.Host, config.Postgres.Port, config.Postgres.Db)

	err = psql.Connect(dsn)
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	repository := repository.NewRepository(psql.DB)
	service := service.NewService(repository)
	handler := handler.NewHandler(*service)

	handler.Init(mux)
	srv := &http.Server{
		Addr:    config.Server.Host + ":" + config.Server.Port,
		Handler: mux,
	}

	// listen to OS signals and gracefully shutdown HTTP server
	stopped := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-sigint
		log.Printf("got interruption signal")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Printf("HTTP Server Shutdown Error: %v", err)
		}
		close(stopped)
	}()

	log.Printf("Starting HTTP server on %s", config.Server.Host+":"+config.Server.Port)

	// start HTTP server
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe Error: %v", err)
	}

	<-stopped

	psql.Close()

	log.Printf("server is shutting down")
}
