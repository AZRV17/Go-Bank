package app

import (
	"log"
	"net/http"

	"github.com/AZRV17/goWEB/pkg/db/psql"
)

func Run() {
	dsn := "postgres://root:sa@localhost:5431/gowebdb"

	err := psql.Connect(dsn)
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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("We're live !!!"))
	})

	http.ListenAndServe(":8080", nil)
}
