package apiserver

import (
	"go-url-shortener/helpers"

	_urlMain "go-url-shortener/url"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Start(config *Config) error {

	db := newDB(config.DatabaseURL)

	defer db.Close()

	srv := newServer()
	_urlMain.InitURL(db, srv.router)

	srv.router.Run()
	return nil
}

func newDB(databaseURL string) *gorm.DB {
	db, err := gorm.Open("postgres", databaseURL)
	helpers.ReurnError(err, "Failed on DB initiation")
	db.DB().SetMaxIdleConns(200)
	db.DB().SetMaxOpenConns(20)

	return db
}
