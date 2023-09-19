package app

import (
	"fmt"
	"log"
	"net/http"

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

	defer func() {
		db, err := psql.DB.DB()
		if err != nil {
			log.Fatal(err)
		}

		log.Fatal(db.Close())
	}()

	mux := http.NewServeMux()

	repositories := repository.NewRepositories(psql.DB)
	service := service.NewService(repositories)
	handler := handler.NewHandler(*service)

	handler.Init(mux)

	log.Println("Server started on port " + config.Server.Port)
	log.Fatal(http.ListenAndServe(config.Server.Host+":"+config.Server.Port, mux))
}
