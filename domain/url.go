package domain

import "github.com/jinzhu/gorm"

type URL struct {
	gorm.Model
	Hash       string
	URLAdrress string
}

type URLUseCase interface{}
type URLRepository interface{}
