package apiserver

import (
	"go-url-shortener/helpers"

	"github.com/jinzhu/gorm"
)

func Start(config *Config) error {
	srv := newServer()
	srv.router.Run()

	db, err := newDB(config.DatabaseURL)
	helpers.ReurnError(err, "Failed on DB initiation")
	defer db.Close()

	return nil
}

func newDB(databaseURL string) (db *gorm.DB, err error) {
	db, err = gorm.Open("postgres", databaseURL)

	db.DB().SetMaxIdleConns(200)
	db.DB().SetMaxOpenConns(20)

	return
}
