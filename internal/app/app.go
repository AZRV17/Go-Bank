package app

import (
	"errors"
	"fmt"
	"github.com/AZRV17/goWEB/internal/config"
	deliveryGRPC "github.com/AZRV17/goWEB/internal/delivery/grpc"
	deliveryHTTP "github.com/AZRV17/goWEB/internal/delivery/http"
	"github.com/AZRV17/goWEB/internal/repository"
	grpcServer "github.com/AZRV17/goWEB/internal/server/grpc"
	"github.com/AZRV17/goWEB/internal/server/http"
	"github.com/AZRV17/goWEB/internal/service"
	"github.com/AZRV17/goWEB/pkg/db/psql"
	"github.com/AZRV17/goWEB/pkg/redis"
	"github.com/go-chi/chi/v5"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

func Run() {
	cfg, err := config.NewConfig("internal/config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.Db,
	)

	err = psql.Connect(dsn)
	if err != nil {
		log.Fatal(err)
	}

	defer psql.Close()

	if err := redis.Connect(cfg); err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := redis.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	r := chi.NewRouter()

	// CORS
	r.Use(
		func(next http.Handler) http.Handler {
			return http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					w.Header().Set("Access-Control-Allow-Origin", "*")
					w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
					w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
					w.Header().Set("Access-Control-Allow-Credentials", "true")
					w.Header().Set("Content-Type", "application/json")

					if r.Method == http.MethodOptions {
						w.WriteHeader(http.StatusOK)
						return
					}

					next.ServeHTTP(w, r)
				},
			)
		},
	)

	repo := repository.NewRepository(psql.DB)
	serv := service.NewService(repo)

	httpSrv := httpServer.NewHttpServer(cfg, r)
	grpcSrv := grpcServer.NewGrpcServer(grpc.NewServer(), cfg)

	httpHandler := deliveryHTTP.NewHandler(*serv)
	httpHandler.Init(r)

	grpcHandler := deliveryGRPC.NewHandler(*serv, grpcSrv.GrpcServer)
	grpcHandler.Init()

	// listen to OS signals and gracefully shutdown HTTP srv
	stopped := make(chan struct{})
	go httpSrv.Shutdown(stopped)

	// listen to OS signals and gracefully shutdown GRPC srv
	go grpcSrv.Shutdown()

	log.Printf("Starting HTTP srv on %s\n", cfg.HTTP.Host+":"+cfg.HTTP.Port)
	log.Printf("Starting GRPC srv on %s\n", cfg.GRPC.Host+":"+cfg.GRPC.Port)

	// start HTTP srv
	go func() {
		if err := httpSrv.Run(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP srv ListenAndServe Error: %v", err)
		}
	}()

	// start GRPC srv
	go func() {
		if err := grpcSrv.Run(); !errors.Is(err, grpc.ErrServerStopped) && err != nil {
			log.Fatalf("GRPC srv ListenAndServe Error: %v", err)
		}
	}()

	<-stopped

	log.Printf("srv is shutting down")
}
