package psql

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(dsn string) error {
	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	return nil
}

func Close() {
	db, err := DB.DB()
	if err != nil {
		log.Fatal("error getting db", err)
	}

	log.Println("closing db")

	err = db.Close()
	if err != nil {
		log.Fatal("error closing db: ", err)
	}

	log.Println("db closed")
}
