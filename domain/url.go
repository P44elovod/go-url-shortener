package domain

import "github.com/jinzhu/gorm"

type URL struct {
	gorm.Model
	Hash       string
	URLAdrress string
}

type URLUseCase interface{}

type URLRepository interface {
	Store(url *URL) (id uint, err error)
	GetAll() (urlList []URL)
	GetByHash(hash string) (url URL)
	GetByID(id string) (url URL)
}
