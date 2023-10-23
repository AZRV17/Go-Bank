package app

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/AZRV17/goWEB/internal/config"
	delivery "github.com/AZRV17/goWEB/internal/delivery/http"
	"github.com/AZRV17/goWEB/internal/repository"
	"github.com/AZRV17/goWEB/internal/server"
	"github.com/AZRV17/goWEB/internal/service"
	"github.com/AZRV17/goWEB/pkg/db/psql"
	"github.com/go-chi/chi/v5"
)

func Run() {
	cfg, err := config.NewConfig("internal/config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.Db)

	err = psql.Connect(dsn)
	if err != nil {
		log.Fatal(err)
	}

	defer psql.Close()

	r := chi.NewRouter()

	repo := repository.NewRepository(psql.DB)
	service := service.NewService(repo)
	handler := delivery.NewHandler(*service)

	handler.Init(r)

	srv := server.NewHttpServer(cfg, r)

	// listen to OS signals and gracefully shutdown HTTP srv
	stopped := make(chan struct{})

	go srv.Shutdown(stopped)

	log.Printf("Starting HTTP srv on %s", ":"+cfg.Server.Port)

	// start HTTP srv
	if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("HTTP srv ListenAndServe Error: %v", err)
	}

	<-stopped

	log.Printf("srv is shutting down")
}
