package domain

import "github.com/jinzhu/gorm"

type URL struct {
	gorm.Model
	Hash       string
	URLAdrress string
}

type URLUseCase interface {
	Create(url *URL) map[string]interface{}
	GetByHash(id string) map[string]interface{}
}

type URLRepository interface {
	Store(url *URL) (id uint, err error)
	GetAll() (urlList []URL, err error)
	GetByHash(hash string) (url URL, err error)
	GetByID(id string) (url URL, err error)
}
